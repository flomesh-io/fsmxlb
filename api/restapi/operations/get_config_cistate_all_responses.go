// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/flomesh-io/fnlb/api/models"
)

// GetConfigCistateAllOKCode is the HTTP code returned for type GetConfigCistateAllOK
const GetConfigCistateAllOKCode int = 200

/*
GetConfigCistateAllOK OK

swagger:response getConfigCistateAllOK
*/
type GetConfigCistateAllOK struct {

	/*
	  In: Body
	*/
	Payload *GetConfigCistateAllOKBody `json:"body,omitempty"`
}

// NewGetConfigCistateAllOK creates GetConfigCistateAllOK with default headers values
func NewGetConfigCistateAllOK() *GetConfigCistateAllOK {

	return &GetConfigCistateAllOK{}
}

// WithPayload adds the payload to the get config cistate all o k response
func (o *GetConfigCistateAllOK) WithPayload(payload *GetConfigCistateAllOKBody) *GetConfigCistateAllOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get config cistate all o k response
func (o *GetConfigCistateAllOK) SetPayload(payload *GetConfigCistateAllOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigCistateAllOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetConfigCistateAllUnauthorizedCode is the HTTP code returned for type GetConfigCistateAllUnauthorized
const GetConfigCistateAllUnauthorizedCode int = 401

/*
GetConfigCistateAllUnauthorized Invalid authentication credentials

swagger:response getConfigCistateAllUnauthorized
*/
type GetConfigCistateAllUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetConfigCistateAllUnauthorized creates GetConfigCistateAllUnauthorized with default headers values
func NewGetConfigCistateAllUnauthorized() *GetConfigCistateAllUnauthorized {

	return &GetConfigCistateAllUnauthorized{}
}

// WithPayload adds the payload to the get config cistate all unauthorized response
func (o *GetConfigCistateAllUnauthorized) WithPayload(payload *models.Error) *GetConfigCistateAllUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get config cistate all unauthorized response
func (o *GetConfigCistateAllUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigCistateAllUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetConfigCistateAllInternalServerErrorCode is the HTTP code returned for type GetConfigCistateAllInternalServerError
const GetConfigCistateAllInternalServerErrorCode int = 500

/*
GetConfigCistateAllInternalServerError Internal service error

swagger:response getConfigCistateAllInternalServerError
*/
type GetConfigCistateAllInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetConfigCistateAllInternalServerError creates GetConfigCistateAllInternalServerError with default headers values
func NewGetConfigCistateAllInternalServerError() *GetConfigCistateAllInternalServerError {

	return &GetConfigCistateAllInternalServerError{}
}

// WithPayload adds the payload to the get config cistate all internal server error response
func (o *GetConfigCistateAllInternalServerError) WithPayload(payload *models.Error) *GetConfigCistateAllInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get config cistate all internal server error response
func (o *GetConfigCistateAllInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigCistateAllInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetConfigCistateAllServiceUnavailableCode is the HTTP code returned for type GetConfigCistateAllServiceUnavailable
const GetConfigCistateAllServiceUnavailableCode int = 503

/*
GetConfigCistateAllServiceUnavailable Maintanence mode

swagger:response getConfigCistateAllServiceUnavailable
*/
type GetConfigCistateAllServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetConfigCistateAllServiceUnavailable creates GetConfigCistateAllServiceUnavailable with default headers values
func NewGetConfigCistateAllServiceUnavailable() *GetConfigCistateAllServiceUnavailable {

	return &GetConfigCistateAllServiceUnavailable{}
}

// WithPayload adds the payload to the get config cistate all service unavailable response
func (o *GetConfigCistateAllServiceUnavailable) WithPayload(payload *models.Error) *GetConfigCistateAllServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get config cistate all service unavailable response
func (o *GetConfigCistateAllServiceUnavailable) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetConfigCistateAllServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
