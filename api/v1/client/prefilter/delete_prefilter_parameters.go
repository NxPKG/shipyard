// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package prefilter

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

	"github.com/khulnasoft/shipyard/api/v1/models"
)

// NewDeletePrefilterParams creates a new DeletePrefilterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeletePrefilterParams() *DeletePrefilterParams {
	return &DeletePrefilterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeletePrefilterParamsWithTimeout creates a new DeletePrefilterParams object
// with the ability to set a timeout on a request.
func NewDeletePrefilterParamsWithTimeout(timeout time.Duration) *DeletePrefilterParams {
	return &DeletePrefilterParams{
		timeout: timeout,
	}
}

// NewDeletePrefilterParamsWithContext creates a new DeletePrefilterParams object
// with the ability to set a context for a request.
func NewDeletePrefilterParamsWithContext(ctx context.Context) *DeletePrefilterParams {
	return &DeletePrefilterParams{
		Context: ctx,
	}
}

// NewDeletePrefilterParamsWithHTTPClient creates a new DeletePrefilterParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeletePrefilterParamsWithHTTPClient(client *http.Client) *DeletePrefilterParams {
	return &DeletePrefilterParams{
		HTTPClient: client,
	}
}

/*
DeletePrefilterParams contains all the parameters to send to the API endpoint

	for the delete prefilter operation.

	Typically these are written to a http.Request.
*/
type DeletePrefilterParams struct {

	/* PrefilterSpec.

	   List of CIDR ranges for filter table
	*/
	PrefilterSpec *models.PrefilterSpec

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete prefilter params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeletePrefilterParams) WithDefaults() *DeletePrefilterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete prefilter params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeletePrefilterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete prefilter params
func (o *DeletePrefilterParams) WithTimeout(timeout time.Duration) *DeletePrefilterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete prefilter params
func (o *DeletePrefilterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete prefilter params
func (o *DeletePrefilterParams) WithContext(ctx context.Context) *DeletePrefilterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete prefilter params
func (o *DeletePrefilterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete prefilter params
func (o *DeletePrefilterParams) WithHTTPClient(client *http.Client) *DeletePrefilterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete prefilter params
func (o *DeletePrefilterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPrefilterSpec adds the prefilterSpec to the delete prefilter params
func (o *DeletePrefilterParams) WithPrefilterSpec(prefilterSpec *models.PrefilterSpec) *DeletePrefilterParams {
	o.SetPrefilterSpec(prefilterSpec)
	return o
}

// SetPrefilterSpec adds the prefilterSpec to the delete prefilter params
func (o *DeletePrefilterParams) SetPrefilterSpec(prefilterSpec *models.PrefilterSpec) {
	o.PrefilterSpec = prefilterSpec
}

// WriteToRequest writes these params to a swagger request
func (o *DeletePrefilterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.PrefilterSpec != nil {
		if err := r.SetBodyParam(o.PrefilterSpec); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
