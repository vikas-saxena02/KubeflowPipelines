// Code generated by go-swagger; DO NOT EDIT.

package experiment_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// APIListExperimentsResponse api list experiments response
//
// swagger:model apiListExperimentsResponse
type APIListExperimentsResponse struct {

	// A list of experiments returned.
	Experiments []*APIExperiment `json:"experiments"`

	// The token to list the next page of experiments.
	NextPageToken string `json:"next_page_token,omitempty"`

	// The total number of experiments for the given query.
	TotalSize int32 `json:"total_size,omitempty"`
}

// Validate validates this api list experiments response
func (m *APIListExperimentsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExperiments(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIListExperimentsResponse) validateExperiments(formats strfmt.Registry) error {
	if swag.IsZero(m.Experiments) { // not required
		return nil
	}

	for i := 0; i < len(m.Experiments); i++ {
		if swag.IsZero(m.Experiments[i]) { // not required
			continue
		}

		if m.Experiments[i] != nil {
			if err := m.Experiments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("experiments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("experiments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this api list experiments response based on the context it is used
func (m *APIListExperimentsResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateExperiments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIListExperimentsResponse) contextValidateExperiments(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Experiments); i++ {

		if m.Experiments[i] != nil {

			if swag.IsZero(m.Experiments[i]) { // not required
				return nil
			}

			if err := m.Experiments[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("experiments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("experiments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *APIListExperimentsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIListExperimentsResponse) UnmarshalBinary(b []byte) error {
	var res APIListExperimentsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
