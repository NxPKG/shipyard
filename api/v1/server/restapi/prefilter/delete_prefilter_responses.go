// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package prefilter

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/khulnasoft/shipyard/api/v1/models"
)

// DeletePrefilterOKCode is the HTTP code returned for type DeletePrefilterOK
const DeletePrefilterOKCode int = 200

/*
DeletePrefilterOK Deleted

swagger:response deletePrefilterOK
*/
type DeletePrefilterOK struct {

	/*
	  In: Body
	*/
	Payload *models.Prefilter `json:"body,omitempty"`
}

// NewDeletePrefilterOK creates DeletePrefilterOK with default headers values
func NewDeletePrefilterOK() *DeletePrefilterOK {

	return &DeletePrefilterOK{}
}

// WithPayload adds the payload to the delete prefilter o k response
func (o *DeletePrefilterOK) WithPayload(payload *models.Prefilter) *DeletePrefilterOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete prefilter o k response
func (o *DeletePrefilterOK) SetPayload(payload *models.Prefilter) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeletePrefilterOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeletePrefilterForbiddenCode is the HTTP code returned for type DeletePrefilterForbidden
const DeletePrefilterForbiddenCode int = 403

/*
DeletePrefilterForbidden Forbidden

swagger:response deletePrefilterForbidden
*/
type DeletePrefilterForbidden struct {
}

// NewDeletePrefilterForbidden creates DeletePrefilterForbidden with default headers values
func NewDeletePrefilterForbidden() *DeletePrefilterForbidden {

	return &DeletePrefilterForbidden{}
}

// WriteResponse to the client
func (o *DeletePrefilterForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// DeletePrefilterInvalidCIDRCode is the HTTP code returned for type DeletePrefilterInvalidCIDR
const DeletePrefilterInvalidCIDRCode int = 461

/*
DeletePrefilterInvalidCIDR Invalid CIDR prefix

swagger:response deletePrefilterInvalidCIdR
*/
type DeletePrefilterInvalidCIDR struct {

	/*
	  In: Body
	*/
	Payload models.Error `json:"body,omitempty"`
}

// NewDeletePrefilterInvalidCIDR creates DeletePrefilterInvalidCIDR with default headers values
func NewDeletePrefilterInvalidCIDR() *DeletePrefilterInvalidCIDR {

	return &DeletePrefilterInvalidCIDR{}
}

// WithPayload adds the payload to the delete prefilter invalid c Id r response
func (o *DeletePrefilterInvalidCIDR) WithPayload(payload models.Error) *DeletePrefilterInvalidCIDR {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete prefilter invalid c Id r response
func (o *DeletePrefilterInvalidCIDR) SetPayload(payload models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeletePrefilterInvalidCIDR) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(461)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// DeletePrefilterFailureCode is the HTTP code returned for type DeletePrefilterFailure
const DeletePrefilterFailureCode int = 500

/*
DeletePrefilterFailure Prefilter delete failed

swagger:response deletePrefilterFailure
*/
type DeletePrefilterFailure struct {

	/*
	  In: Body
	*/
	Payload models.Error `json:"body,omitempty"`
}

// NewDeletePrefilterFailure creates DeletePrefilterFailure with default headers values
func NewDeletePrefilterFailure() *DeletePrefilterFailure {

	return &DeletePrefilterFailure{}
}

// WithPayload adds the payload to the delete prefilter failure response
func (o *DeletePrefilterFailure) WithPayload(payload models.Error) *DeletePrefilterFailure {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete prefilter failure response
func (o *DeletePrefilterFailure) SetPayload(payload models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeletePrefilterFailure) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
