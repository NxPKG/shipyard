// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package recorder

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/khulnasoft/shipyard/api/v1/models"
)

// GetRecorderMasksOKCode is the HTTP code returned for type GetRecorderMasksOK
const GetRecorderMasksOKCode int = 200

/*
GetRecorderMasksOK Success

swagger:response getRecorderMasksOK
*/
type GetRecorderMasksOK struct {

	/*
	  In: Body
	*/
	Payload []*models.RecorderMask `json:"body,omitempty"`
}

// NewGetRecorderMasksOK creates GetRecorderMasksOK with default headers values
func NewGetRecorderMasksOK() *GetRecorderMasksOK {

	return &GetRecorderMasksOK{}
}

// WithPayload adds the payload to the get recorder masks o k response
func (o *GetRecorderMasksOK) WithPayload(payload []*models.RecorderMask) *GetRecorderMasksOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get recorder masks o k response
func (o *GetRecorderMasksOK) SetPayload(payload []*models.RecorderMask) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRecorderMasksOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.RecorderMask, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
