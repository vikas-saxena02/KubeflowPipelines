// Code generated by go-swagger; DO NOT EDIT.

package experiment_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewExperimentServiceGetExperimentV1Params creates a new ExperimentServiceGetExperimentV1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExperimentServiceGetExperimentV1Params() *ExperimentServiceGetExperimentV1Params {
	return &ExperimentServiceGetExperimentV1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewExperimentServiceGetExperimentV1ParamsWithTimeout creates a new ExperimentServiceGetExperimentV1Params object
// with the ability to set a timeout on a request.
func NewExperimentServiceGetExperimentV1ParamsWithTimeout(timeout time.Duration) *ExperimentServiceGetExperimentV1Params {
	return &ExperimentServiceGetExperimentV1Params{
		timeout: timeout,
	}
}

// NewExperimentServiceGetExperimentV1ParamsWithContext creates a new ExperimentServiceGetExperimentV1Params object
// with the ability to set a context for a request.
func NewExperimentServiceGetExperimentV1ParamsWithContext(ctx context.Context) *ExperimentServiceGetExperimentV1Params {
	return &ExperimentServiceGetExperimentV1Params{
		Context: ctx,
	}
}

// NewExperimentServiceGetExperimentV1ParamsWithHTTPClient creates a new ExperimentServiceGetExperimentV1Params object
// with the ability to set a custom HTTPClient for a request.
func NewExperimentServiceGetExperimentV1ParamsWithHTTPClient(client *http.Client) *ExperimentServiceGetExperimentV1Params {
	return &ExperimentServiceGetExperimentV1Params{
		HTTPClient: client,
	}
}

/*
ExperimentServiceGetExperimentV1Params contains all the parameters to send to the API endpoint

	for the experiment service get experiment v1 operation.

	Typically these are written to a http.Request.
*/
type ExperimentServiceGetExperimentV1Params struct {

	/* ID.

	   The ID of the experiment to be retrieved.
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the experiment service get experiment v1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExperimentServiceGetExperimentV1Params) WithDefaults() *ExperimentServiceGetExperimentV1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the experiment service get experiment v1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExperimentServiceGetExperimentV1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the experiment service get experiment v1 params
func (o *ExperimentServiceGetExperimentV1Params) WithTimeout(timeout time.Duration) *ExperimentServiceGetExperimentV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the experiment service get experiment v1 params
func (o *ExperimentServiceGetExperimentV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the experiment service get experiment v1 params
func (o *ExperimentServiceGetExperimentV1Params) WithContext(ctx context.Context) *ExperimentServiceGetExperimentV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the experiment service get experiment v1 params
func (o *ExperimentServiceGetExperimentV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the experiment service get experiment v1 params
func (o *ExperimentServiceGetExperimentV1Params) WithHTTPClient(client *http.Client) *ExperimentServiceGetExperimentV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the experiment service get experiment v1 params
func (o *ExperimentServiceGetExperimentV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the experiment service get experiment v1 params
func (o *ExperimentServiceGetExperimentV1Params) WithID(id string) *ExperimentServiceGetExperimentV1Params {
	o.SetID(id)
	return o
}

// SetID adds the id to the experiment service get experiment v1 params
func (o *ExperimentServiceGetExperimentV1Params) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExperimentServiceGetExperimentV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
