// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package endpoint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/khulnasoft/shipyard/api/v1/models"
)

// PatchEndpointIDOKCode is the HTTP code returned for type PatchEndpointIDOK
const PatchEndpointIDOKCode int = 200

/*
PatchEndpointIDOK Success

swagger:response patchEndpointIdOK
*/
type PatchEndpointIDOK struct {
}

// NewPatchEndpointIDOK creates PatchEndpointIDOK with default headers values
func NewPatchEndpointIDOK() *PatchEndpointIDOK {

	return &PatchEndpointIDOK{}
}

// WriteResponse to the client
func (o *PatchEndpointIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// PatchEndpointIDInvalidCode is the HTTP code returned for type PatchEndpointIDInvalid
const PatchEndpointIDInvalidCode int = 400

/*
PatchEndpointIDInvalid Invalid modify endpoint request

swagger:response patchEndpointIdInvalid
*/
type PatchEndpointIDInvalid struct {

	/*
	  In: Body
	*/
	Payload models.Error `json:"body,omitempty"`
}

// NewPatchEndpointIDInvalid creates PatchEndpointIDInvalid with default headers values
func NewPatchEndpointIDInvalid() *PatchEndpointIDInvalid {

	return &PatchEndpointIDInvalid{}
}

// WithPayload adds the payload to the patch endpoint Id invalid response
func (o *PatchEndpointIDInvalid) WithPayload(payload models.Error) *PatchEndpointIDInvalid {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch endpoint Id invalid response
func (o *PatchEndpointIDInvalid) SetPayload(payload models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchEndpointIDInvalid) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// PatchEndpointIDForbiddenCode is the HTTP code returned for type PatchEndpointIDForbidden
const PatchEndpointIDForbiddenCode int = 403

/*
PatchEndpointIDForbidden Forbidden

swagger:response patchEndpointIdForbidden
*/
type PatchEndpointIDForbidden struct {
}

// NewPatchEndpointIDForbidden creates PatchEndpointIDForbidden with default headers values
func NewPatchEndpointIDForbidden() *PatchEndpointIDForbidden {

	return &PatchEndpointIDForbidden{}
}

// WriteResponse to the client
func (o *PatchEndpointIDForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(403)
}

// PatchEndpointIDNotFoundCode is the HTTP code returned for type PatchEndpointIDNotFound
const PatchEndpointIDNotFoundCode int = 404

/*
PatchEndpointIDNotFound Endpoint does not exist

swagger:response patchEndpointIdNotFound
*/
type PatchEndpointIDNotFound struct {
}

// NewPatchEndpointIDNotFound creates PatchEndpointIDNotFound with default headers values
func NewPatchEndpointIDNotFound() *PatchEndpointIDNotFound {

	return &PatchEndpointIDNotFound{}
}

// WriteResponse to the client
func (o *PatchEndpointIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// PatchEndpointIDTooManyRequestsCode is the HTTP code returned for type PatchEndpointIDTooManyRequests
const PatchEndpointIDTooManyRequestsCode int = 429

/*
PatchEndpointIDTooManyRequests Rate-limiting too many requests in the given time frame

swagger:response patchEndpointIdTooManyRequests
*/
type PatchEndpointIDTooManyRequests struct {
}

// NewPatchEndpointIDTooManyRequests creates PatchEndpointIDTooManyRequests with default headers values
func NewPatchEndpointIDTooManyRequests() *PatchEndpointIDTooManyRequests {

	return &PatchEndpointIDTooManyRequests{}
}

// WriteResponse to the client
func (o *PatchEndpointIDTooManyRequests) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(429)
}

// PatchEndpointIDFailedCode is the HTTP code returned for type PatchEndpointIDFailed
const PatchEndpointIDFailedCode int = 500

/*
PatchEndpointIDFailed Endpoint update failed

swagger:response patchEndpointIdFailed
*/
type PatchEndpointIDFailed struct {

	/*
	  In: Body
	*/
	Payload models.Error `json:"body,omitempty"`
}

// NewPatchEndpointIDFailed creates PatchEndpointIDFailed with default headers values
func NewPatchEndpointIDFailed() *PatchEndpointIDFailed {

	return &PatchEndpointIDFailed{}
}

// WithPayload adds the payload to the patch endpoint Id failed response
func (o *PatchEndpointIDFailed) WithPayload(payload models.Error) *PatchEndpointIDFailed {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the patch endpoint Id failed response
func (o *PatchEndpointIDFailed) SetPayload(payload models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PatchEndpointIDFailed) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
