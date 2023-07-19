// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/cybwan/fsmxlb/api/models"
)

// DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedNoContentCode is the HTTP code returned for type DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedNoContent
const DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedNoContentCode int = 204

/*
DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedNoContent OK

swagger:response deleteConfigVlanVlanIdMemberIfNameTaggedTaggedNoContent
*/
type DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedNoContent struct {
}

// NewDeleteConfigVlanVlanIDMemberIfNameTaggedTaggedNoContent creates DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedNoContent with default headers values
func NewDeleteConfigVlanVlanIDMemberIfNameTaggedTaggedNoContent() *DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedNoContent {

	return &DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedNoContent{}
}

// WriteResponse to the client
func (o *DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedBadRequestCode is the HTTP code returned for type DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedBadRequest
const DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedBadRequestCode int = 400

/*
DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedBadRequest Malformed arguments for API call

swagger:response deleteConfigVlanVlanIdMemberIfNameTaggedTaggedBadRequest
*/
type DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteConfigVlanVlanIDMemberIfNameTaggedTaggedBadRequest creates DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedBadRequest with default headers values
func NewDeleteConfigVlanVlanIDMemberIfNameTaggedTaggedBadRequest() *DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedBadRequest {

	return &DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedBadRequest{}
}

// WithPayload adds the payload to the delete config vlan vlan Id member if name tagged tagged bad request response
func (o *DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedBadRequest) WithPayload(payload *models.Error) *DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete config vlan vlan Id member if name tagged tagged bad request response
func (o *DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedUnauthorizedCode is the HTTP code returned for type DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedUnauthorized
const DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedUnauthorizedCode int = 401

/*
DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedUnauthorized Invalid authentication credentials

swagger:response deleteConfigVlanVlanIdMemberIfNameTaggedTaggedUnauthorized
*/
type DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteConfigVlanVlanIDMemberIfNameTaggedTaggedUnauthorized creates DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedUnauthorized with default headers values
func NewDeleteConfigVlanVlanIDMemberIfNameTaggedTaggedUnauthorized() *DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedUnauthorized {

	return &DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedUnauthorized{}
}

// WithPayload adds the payload to the delete config vlan vlan Id member if name tagged tagged unauthorized response
func (o *DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedUnauthorized) WithPayload(payload *models.Error) *DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete config vlan vlan Id member if name tagged tagged unauthorized response
func (o *DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedNotFoundCode is the HTTP code returned for type DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedNotFound
const DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedNotFoundCode int = 404

/*
DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedNotFound Vlan interface is not defined/Vlan member is not found on this Vlan interface

swagger:response deleteConfigVlanVlanIdMemberIfNameTaggedTaggedNotFound
*/
type DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteConfigVlanVlanIDMemberIfNameTaggedTaggedNotFound creates DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedNotFound with default headers values
func NewDeleteConfigVlanVlanIDMemberIfNameTaggedTaggedNotFound() *DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedNotFound {

	return &DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedNotFound{}
}

// WithPayload adds the payload to the delete config vlan vlan Id member if name tagged tagged not found response
func (o *DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedNotFound) WithPayload(payload *models.Error) *DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete config vlan vlan Id member if name tagged tagged not found response
func (o *DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedInternalServerErrorCode is the HTTP code returned for type DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedInternalServerError
const DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedInternalServerErrorCode int = 500

/*
DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedInternalServerError Internal service error

swagger:response deleteConfigVlanVlanIdMemberIfNameTaggedTaggedInternalServerError
*/
type DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteConfigVlanVlanIDMemberIfNameTaggedTaggedInternalServerError creates DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedInternalServerError with default headers values
func NewDeleteConfigVlanVlanIDMemberIfNameTaggedTaggedInternalServerError() *DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedInternalServerError {

	return &DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedInternalServerError{}
}

// WithPayload adds the payload to the delete config vlan vlan Id member if name tagged tagged internal server error response
func (o *DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedInternalServerError) WithPayload(payload *models.Error) *DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete config vlan vlan Id member if name tagged tagged internal server error response
func (o *DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedServiceUnavailableCode is the HTTP code returned for type DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedServiceUnavailable
const DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedServiceUnavailableCode int = 503

/*
DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedServiceUnavailable Maintanence mode

swagger:response deleteConfigVlanVlanIdMemberIfNameTaggedTaggedServiceUnavailable
*/
type DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteConfigVlanVlanIDMemberIfNameTaggedTaggedServiceUnavailable creates DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedServiceUnavailable with default headers values
func NewDeleteConfigVlanVlanIDMemberIfNameTaggedTaggedServiceUnavailable() *DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedServiceUnavailable {

	return &DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedServiceUnavailable{}
}

// WithPayload adds the payload to the delete config vlan vlan Id member if name tagged tagged service unavailable response
func (o *DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedServiceUnavailable) WithPayload(payload *models.Error) *DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete config vlan vlan Id member if name tagged tagged service unavailable response
func (o *DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedServiceUnavailable) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteConfigVlanVlanIDMemberIfNameTaggedTaggedServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}