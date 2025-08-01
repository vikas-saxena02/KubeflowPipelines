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
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	apiv1beta1 "github.com/kubeflow/pipelines/backend/api/v1beta1/go_client"
	apiv2beta1 "github.com/kubeflow/pipelines/backend/api/v2beta1/go_client"
	"github.com/kubeflow/pipelines/backend/src/apiserver/common"
	"github.com/kubeflow/pipelines/backend/src/apiserver/list"
	"github.com/kubeflow/pipelines/backend/src/apiserver/model"
	"github.com/kubeflow/pipelines/backend/src/common/util"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/testing/protocmp"
)

var fakeModelFieldsBySortableAPIFields = map[string]string{
	"":            "Name",
	"name":        "Name",
	"author":      "Author",
	"description": "Description",
}

func getFakeModelToken() string {
	token := common.Token{
		SortByFieldValue: "bar",
		KeyFieldValue:    "foo",
	}
	expectedJson, _ := json.Marshal(token)
	return base64.StdEncoding.EncodeToString(expectedJson)
}

func TestValidateFilterV1(t *testing.T) {
	referenceKey := &apiv1beta1.ResourceKey{Type: apiv1beta1.ResourceType_EXPERIMENT, Id: "123"}
	ctx, err := validateFilterV1(referenceKey)
	assert.Nil(t, err)
	assert.Equal(t, &model.FilterContext{ReferenceKey: &model.ReferenceKey{Type: model.ExperimentResourceType, ID: "123"}}, ctx)
}

func TestValidateFilterV1_ToModelResourceTypeFailed(t *testing.T) {
	referenceKey := &apiv1beta1.ResourceKey{Type: apiv1beta1.ResourceType_UNKNOWN_RESOURCE_TYPE, Id: "123"}
	_, err := validateFilterV1(referenceKey)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "Unrecognized resource reference type")
}

func TestValidatePagination(t *testing.T) {
	token := getFakeModelToken()
	context, err := validatePagination(token, 3, "Name",
		"", fakeModelFieldsBySortableAPIFields)
	assert.Nil(t, err)
	expected := &common.PaginationContext{
		PageSize:        3,
		SortByFieldName: "Name",
		KeyFieldName:    "Name",
		Token:           &common.Token{SortByFieldValue: "bar", KeyFieldValue: "foo"},
	}
	assert.Equal(t, expected, context)
}

func TestValidatePagination_NegativePageSizeError(t *testing.T) {
	token := getFakeModelToken()
	_, err := validatePagination(token, -1, "Name",
		"", fakeModelFieldsBySortableAPIFields)
	assert.Equal(t, codes.InvalidArgument, err.(*util.UserError).ExternalStatusCode())
}

func TestValidatePagination_DefaultPageSize(t *testing.T) {
	token := getFakeModelToken()
	context, err := validatePagination(token, 0, "Name",
		"", fakeModelFieldsBySortableAPIFields)
	expected := &common.PaginationContext{
		PageSize:        defaultPageSize,
		SortByFieldName: "Name",
		KeyFieldName:    "Name",
		Token:           &common.Token{SortByFieldValue: "bar", KeyFieldValue: "foo"},
	}
	assert.Nil(t, err)
	assert.Equal(t, expected, context)
}

func TestValidatePagination_DefaultSorting(t *testing.T) {
	token := getFakeModelToken()
	context, err := validatePagination(token, 0, "Name",
		"", fakeModelFieldsBySortableAPIFields)
	expected := &common.PaginationContext{
		PageSize:        defaultPageSize,
		SortByFieldName: "Name",
		KeyFieldName:    "Name",
		Token:           &common.Token{SortByFieldValue: "bar", KeyFieldValue: "foo"},
	}
	assert.Nil(t, err)
	assert.Equal(t, expected, context)
}

func TestValidatePagination_InvalidToken(t *testing.T) {
	_, err := validatePagination("invalid token", 0, "",
		"", fakeModelFieldsBySortableAPIFields)
	assert.Equal(t, codes.InvalidArgument, err.(*util.UserError).ExternalStatusCode())
}

func TestDeserializePageToken(t *testing.T) {
	token := common.Token{
		SortByFieldValue: "bar",
		KeyFieldValue:    "foo",
	}
	expectedJson, _ := json.Marshal(token)
	tokenString := base64.StdEncoding.EncodeToString(expectedJson)
	actualToken, err := deserializePageToken(tokenString)
	assert.Nil(t, err)
	assert.Equal(t, token, *actualToken)
}
func TestTransformJSONForBackwardCompatibility(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		expected  string
		expectErr bool
	}{
		{
			name:      "Standard case with multiple replacements",
			input:     `{"intValue": 1, "stringValue": "test", "longValue": 1234567890}`,
			expected:  `{"int_value": 1, "string_value": "test", "long_value": 1234567890}`,
			expectErr: false,
		},
		{
			name:      "No replacements needed",
			input:     `{"int_value": 1, "string_value": "test", "long_value": 1234567890}`,
			expected:  `{"int_value": 1, "string_value": "test", "long_value": 1234567890}`,
			expectErr: false,
		},
		{
			name:      "Empty input",
			input:     ``,
			expected:  ``,
			expectErr: false,
		},
		{
			name:      "Nested JSON structure",
			input:     `{"data": {"intValue": 5, "stringValue": "nested"}}`,
			expected:  `{"data": {"int_value": 5, "string_value": "nested"}}`,
			expectErr: false,
		},
		{
			name:      "Input with no recognized fields",
			input:     `{"otherValue": "value"}`,
			expected:  `{"otherValue": "value"}`,
			expectErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := transformJSONForBackwardCompatibility(test.input)

			if test.expectErr {
				assert.NotNil(t, err, "Expected an error but got none")
				return
			}

			assert.Nil(t, err, "Expected no error but got: %+v", err)
			assert.Equal(t, test.expected, result, "Unexpected result for input: %v", test.input)
		})
	}
}
func TestDeserializePageToken_InvalidEncodingStringError(t *testing.T) {
	_, err := deserializePageToken("this is a invalid token")
	assert.Equal(t, codes.InvalidArgument, err.(*util.UserError).ExternalStatusCode())
}

func TestDeserializePageToken_UnmarshalError(t *testing.T) {
	_, err := deserializePageToken(base64.StdEncoding.EncodeToString([]byte("invalid token")))
	assert.Equal(t, codes.InvalidArgument, err.(*util.UserError).ExternalStatusCode())
}

func TestParseSortByQueryString_EmptyString(t *testing.T) {
	modelField, isDesc, err := parseSortByQueryString("", fakeModelFieldsBySortableAPIFields)
	assert.Nil(t, err)
	assert.Equal(t, "Name", modelField)
	assert.False(t, isDesc)
}

func TestParseSortByQueryString_FieldNameOnly(t *testing.T) {
	modelField, isDesc, err := parseSortByQueryString("Name", fakeModelFieldsBySortableAPIFields)
	assert.Nil(t, err)
	assert.Equal(t, "Name", modelField)
	assert.False(t, isDesc)
}

func TestParseSortByQueryString_FieldNameWithDescFlag(t *testing.T) {
	modelField, isDesc, err := parseSortByQueryString("Name desc", fakeModelFieldsBySortableAPIFields)
	assert.Nil(t, err)
	assert.Equal(t, "Name", modelField)
	assert.True(t, isDesc)
}

func TestParseSortByQueryString_FieldNameWithAscFlag(t *testing.T) {
	modelField, isDesc, err := parseSortByQueryString("Name asc", fakeModelFieldsBySortableAPIFields)
	assert.Nil(t, err)
	assert.Equal(t, "Name", modelField)
	assert.False(t, isDesc)
}

func TestParseSortByQueryString_NotSortableFieldName(t *testing.T) {
	_, _, err := parseSortByQueryString("foobar", fakeModelFieldsBySortableAPIFields)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "Cannot sort on field foobar")
}

func TestParseSortByQueryString_IncorrectDescFlag(t *testing.T) {
	_, _, err := parseSortByQueryString("id foobar", fakeModelFieldsBySortableAPIFields)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "Received invalid sort by format 'id foobar'")
}

func TestParseSortByQueryString_StringTooLong(t *testing.T) {
	_, _, err := parseSortByQueryString("Name desc foo", fakeModelFieldsBySortableAPIFields)
	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), "Received invalid sort by format 'Name desc foo'")
}

func TestParseAPIFilter_EmptyStringYieldsNilFilter(t *testing.T) {
	f, err := parseAPIFilter("", "v1beta1")
	assert.Nil(t, err)
	assert.Nil(t, f)

	f, err = parseAPIFilter("", "v2beta1")
	assert.Nil(t, err)
	assert.Nil(t, f)
}

func TestParseAPIFilter_InvalidStringYieldsError(t *testing.T) {
	f, err := parseAPIFilter("lkjlkjlkj", "v1beta1")
	assert.NotNil(t, err)
	assert.Nil(t, f)

	f, err = parseAPIFilter("lkjlkjlkj", "v2beta1")
	assert.NotNil(t, err)
	assert.Nil(t, f)
}

func TestParseAPIFilter_DecodesEncodedStringV1(t *testing.T) {
	// in was generated by calling JSON.stringify followed by encodeURIComponent in
	// the browser on the following JSON string:
	//   {"predicates":[{"op":"EQUALS","key":"testkey","stringValue":"testvalue"}]}

	in := "%7B%22predicates%22%3A%5B%7B%22op%22%3A%22EQUALS%22%2C%22key%22%3A%22testkey%22%2C%22stringValue%22%3A%22testvalue%22%7D%5D%7D"

	// The above should correspond the following filter:
	want := &apiv1beta1.Filter{
		Predicates: []*apiv1beta1.Predicate{
			{
				Key: "testkey", Op: apiv1beta1.Predicate_EQUALS,
				Value: &apiv1beta1.Predicate_StringValue{StringValue: "testvalue"},
			},
		},
	}

	got, err := parseAPIFilter(in, "v1beta1")
	if !cmp.Equal(got, want, cmpopts.EquateEmpty(), protocmp.Transform()) || err != nil {
		t.Errorf("parseAPIString(%q) =\nGot %+v, %v\n Want %+v, <nil>\nDiff: %s",
			in, got, err, want, cmp.Diff(want, got))
	}
}

func TestParseAPIFilter_DecodesEncodedString(t *testing.T) {
	// in was generated by calling JSON.stringify followed by encodeURIComponent in
	// the browser on the following JSON string:
	//   {"predicates":[{"operation":"EQUALS","key":"testkey","stringValue":"testvalue"}]}

	in := "%7B%22predicates%22%3A%5B%7B%22operation%22%3A%22EQUALS%22%2C%22key%22%3A%22testkey%22%2C%22stringValue%22%3A%22testvalue%22%7D%5D%7D"

	// The above should correspond the following filter:
	want := &apiv2beta1.Filter{
		Predicates: []*apiv2beta1.Predicate{
			{
				Key: "testkey", Operation: apiv2beta1.Predicate_EQUALS,
				Value: &apiv2beta1.Predicate_StringValue{StringValue: "testvalue"},
			},
		},
	}

	got, err := parseAPIFilter(in, "v2beta1")
	if !cmp.Equal(got, want, cmpopts.EquateEmpty(), protocmp.Transform()) || err != nil {
		t.Errorf("parseAPIString(%q) =\nGot %+v, %v\n Want %+v, <nil>\nDiff: %s",
			in, got, err, want, cmp.Diff(want, got))
	}
}

type fakeListable struct {
	PrimaryKey       string
	FakeName         string
	CreatedTimestamp int64
}

func (f *fakeListable) PrimaryKeyColumnName() string {
	return "PrimaryKey"
}

func (f *fakeListable) DefaultSortField() string {
	return "CreatedTimestamp"
}

var fakeAPIToModelMap = map[string]string{
	"timestamp": "CreatedTimestamp",
	"name":      "FakeName",
	"id":        "PrimaryKey",
}

func (f *fakeListable) APIToModelFieldMap() map[string]string {
	return fakeAPIToModelMap
}

func (f *fakeListable) GetModelName() string {
	return ""
}

func (f *fakeListable) GetField(name string) (string, bool) {
	if field, ok := fakeAPIToModelMap[name]; ok {
		return field, true
	} else {
		return "", false
	}
}

func (f *fakeListable) GetFieldValue(name string) interface{} {
	switch name {
	case "CreatedTimestamp":
		return f.CreatedTimestamp
	case "FakeName":
		return f.FakeName
	case "PrimaryKey":
		return f.PrimaryKey
	}
	return nil
}

func (f *fakeListable) GetSortByFieldPrefix(name string) string {
	return ""
}

func (f *fakeListable) GetKeyFieldPrefix() string {
	return ""
}

func TestValidatedListOptions_Errors(t *testing.T) {
	opts, err := list.NewOptions(&fakeListable{}, 10, "name asc", nil)
	if err != nil {
		t.Fatalf("list.NewOptions() = _, %+v; Want nil error", err)
	}

	npt, err := opts.NextPageToken(&fakeListable{})
	if err != nil {
		t.Fatalf("opt.NextPageToken() = _, %+v; Want nil error", err)
	}

	_, err = validatedListOptions(&fakeListable{}, npt, 10, "name asc", "", "v1beta1")
	if err != nil {
		t.Fatalf("validatedListOptions(fakeListable, 10, \"name asc\") = _, %+v; Want nil error", err)
	}

	_, err = validatedListOptions(&fakeListable{}, npt, 10, "name asc", "", "v2beta1")
	if err != nil {
		t.Fatalf("validatedListOptions(fakeListable, 10, \"name asc\") = _, %+v; Want nil error", err)
	}

	_, err = validatedListOptions(&fakeListable{}, npt, 10, "name desc", "", "v1beta1")
	if err == nil {
		t.Fatalf("validatedListOptions(fakeListable, 10, \"name desc\") = _, %+v; Want error", err)
	}

	_, err = validatedListOptions(&fakeListable{}, npt, 10, "name desc", "", "v2beta1")
	if err == nil {
		t.Fatalf("validatedListOptions(fakeListable, 10, \"name desc\") = _, %+v; Want error", err)
	}
}
