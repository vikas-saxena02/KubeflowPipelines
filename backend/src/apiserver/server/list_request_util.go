// Copyright 2018 The Kubeflow Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	apiv1beta1 "github.com/kubeflow/pipelines/backend/api/v1beta1/go_client"
	apiv2beta1 "github.com/kubeflow/pipelines/backend/api/v2beta1/go_client"
	"github.com/kubeflow/pipelines/backend/src/apiserver/common"
	"github.com/kubeflow/pipelines/backend/src/apiserver/filter"
	"github.com/kubeflow/pipelines/backend/src/apiserver/list"
	"github.com/kubeflow/pipelines/backend/src/apiserver/model"
	"github.com/kubeflow/pipelines/backend/src/common/util"
	"google.golang.org/protobuf/encoding/protojson"
)

const (
	defaultPageSize = 20
	maxPageSize     = 200
)

func validateFilterV1(referenceKey *apiv1beta1.ResourceKey) (*model.FilterContext, error) {
	filterContext := &model.FilterContext{}
	if referenceKey != nil {
		refType, err := toModelResourceTypeV1(referenceKey.Type)
		if err != nil {
			return nil, util.Wrap(err, "Unrecognized resource reference type")
		}
		filterContext.ReferenceKey = &model.ReferenceKey{Type: refType, ID: referenceKey.Id}
	}
	return filterContext, nil
}

func validatePagination(pageToken string, pageSize int, keyFieldName string, queryString string,
	modelFieldByApiFieldMapping map[string]string,
) (*common.PaginationContext, error) {
	sortByFieldName, isDesc, err := parseSortByQueryString(queryString, modelFieldByApiFieldMapping)
	if err != nil {
		return nil, util.Wrap(err, "Invalid query string")
	}
	if pageSize < 0 {
		return nil, util.NewInvalidInputError("The page size should be greater than 0. Got %v", strconv.Itoa(pageSize))
	}
	if pageSize == 0 {
		// Use default page size if not provided.
		pageSize = defaultPageSize
	}
	if pageSize > maxPageSize {
		pageSize = maxPageSize
	}
	if sortByFieldName == "" {
		// By default, sort by key field.
		sortByFieldName = keyFieldName
	}
	token, err := deserializePageToken(pageToken)
	if err != nil {
		return nil, util.Wrap(err, "Invalid page token")
	}
	return &common.PaginationContext{
		PageSize:        pageSize,
		SortByFieldName: sortByFieldName,
		KeyFieldName:    keyFieldName,
		IsDesc:          isDesc,
		Token:           token,
	}, nil
}

func parseSortByQueryString(queryString string, modelFieldByApiFieldMapping map[string]string) (string, bool, error) {
	// ignore the case of the letter. Split query string by space
	queryList := strings.Fields(strings.ToLower(queryString))
	// Check the query string format.
	if len(queryList) > 2 || (len(queryList) == 2 && queryList[1] != "desc" && queryList[1] != "asc") {
		return "", false, util.NewInvalidInputError(
			"Received invalid sort by format '%v'. Supported format: \"field_name\", \"field_name desc\", or \"field_name asc\"", queryString)
	}
	isDesc := false
	if len(queryList) == 2 && queryList[1] == "desc" {
		isDesc = true
	}
	sortByApiField := ""
	if len(queryList) > 0 {
		sortByApiField = queryList[0]
	}
	// Check if the field can be sorted.
	sortByFieldName, ok := modelFieldByApiFieldMapping[sortByApiField]
	if !ok {
		return "", false, util.NewInvalidInputError("Cannot sort on field %v. Supported fields %v",
			sortByApiField, keysString(modelFieldByApiFieldMapping))
	}
	return sortByFieldName, isDesc, nil
}

func keysString(modelFieldByApiFieldMapping map[string]string) string {
	keys := make([]string, 0, len(modelFieldByApiFieldMapping))
	for k := range modelFieldByApiFieldMapping {
		if k != "" {
			keys = append(keys, k)
		}
	}
	return "[" + strings.Join(keys, ", ") + "]"
}

// Decode page token. If page token is empty, we assume listing the first page and return a nil Token.
func deserializePageToken(pageToken string) (*common.Token, error) {
	if pageToken == "" {
		return nil, nil
	}
	tokenBytes, err := base64.StdEncoding.DecodeString(pageToken)
	if err != nil {
		return nil, util.NewInvalidInputErrorWithDetails(err, "Invalid package token")
	}
	var token common.Token
	err = json.Unmarshal(tokenBytes, &token)
	if err != nil {
		return nil, util.NewInvalidInputErrorWithDetails(err, "Invalid package token")
	}
	return &token, nil
}

// parseAPIFilter attempts to decode a url-encoded JSON-stringified api
// filter object. An empty string is considered valid input, and equivalent to
// the nil filter, which trivially does nothing.
func parseAPIFilter(encoded string, apiVersion string) (interface{}, error) {
	if encoded == "" {
		return nil, nil
	}
	decoded, err := url.QueryUnescape(encoded)
	if err != nil {
		return nil, util.NewInvalidInputError("failed to parse valid filter from %q: %v", encoded, err)
	}

	transformedJSON, err := transformJSONForBackwardCompatibility(decoded)
	if err != nil {
		fmt.Printf("Failed to transform JSON: %v\n", err)
		return nil, err
	}

	switch apiVersion {
	case "v2beta1":
		f := &apiv2beta1.Filter{}
		if err := protojson.Unmarshal([]byte(transformedJSON), f); err != nil {
			return nil, util.NewInvalidInputError("failed to parse valid filter from %q: %v", encoded, err)
		}
		return f, nil
	case "v1beta1":
		f := &apiv1beta1.Filter{}
		if err := protojson.Unmarshal([]byte(transformedJSON), f); err != nil {
			return nil, util.NewInvalidInputError("failed to parse valid filter from %q: %v", encoded, err)
		}
		return f, nil
	default:
		return nil, util.NewUnknownApiVersionError("filter "+apiVersion, encoded)
	}
}

// Validates list options for a given resource and listing parameters.
// apiVersion cat be set to "v1beta1" or "v2beta1". Depending on the value,
// the corresponding API filter message will be used when parsing filterSpec.
func validatedListOptions(listable list.Listable, pageToken string, pageSize int, sortBy string, filterSpec string, apiVersion string) (*list.Options, error) {
	defaultOpts := func() (*list.Options, error) {
		if listable == nil {
			return nil, util.NewInvalidInputError("Please specify a valid type to list. E.g., list runs or list jobs")
		}
		f, err := parseAPIFilter(filterSpec, apiVersion)
		if err != nil {
			return nil, err
		}
		newFilter, err := filter.New(f)
		if err != nil {
			return nil, err
		}

		return list.NewOptions(listable, pageSize, sortBy, newFilter)
	}

	if pageToken == "" {
		return defaultOpts()
	}

	opts, err := list.NewOptionsFromToken(pageToken, pageSize)
	if err != nil {
		return nil, err
	}

	if sortBy != "" || filterSpec != "" {
		// Sanity check that these match the page token.
		do, err := defaultOpts()
		if err != nil {
			return nil, err
		}

		if !opts.Matches(do) {
			return nil, util.NewInvalidInputError("page token does not match the supplied sort by and/or filtering criteria. Either specify the same criteria or leave the latter empty if page token is specified")
		}
	}

	return opts, nil
}

// transformJSONForBackwardCompatibility replaces specific JSON key names to maintain
// backward compatibility with older APIs. Previously, KFP deviated from the typical
// snake_case naming convention for protobuf field names for Filter predicate values.
// This function replaces specific JSON key names to maintain backward compatibility
// with older APIs.
// See Predicate.value in backend/api/v2beta1/filter.proto for these values.
// Previously
func transformJSONForBackwardCompatibility(jsonStr string) (string, error) {
	replacer := strings.NewReplacer(
		`"intValue":`, `"int_value":`,
		`"longValue":`, `"long_value":`,
		`"stringValue":`, `"string_value":`,
		`"timestampValue":`, `"timestamp_value":`,
		`"intValues":`, `"int_values":`,
		`"longValues":`, `"long_values":`,
		`"stringValues":`, `"string_values":`,
	)
	return replacer.Replace(jsonStr), nil
}
