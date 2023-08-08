// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/flomesh-io/fsmxlb/api/models"
)

// PostConfigVlanVlanIDMemberNoContentCode is the HTTP code returned for type PostConfigVlanVlanIDMemberNoContent
const PostConfigVlanVlanIDMemberNoContentCode int = 204

/*
PostConfigVlanVlanIDMemberNoContent OK

swagger:response postConfigVlanVlanIdMemberNoContent
*/
type PostConfigVlanVlanIDMemberNoContent struct {
}

// NewPostConfigVlanVlanIDMemberNoContent creates PostConfigVlanVlanIDMemberNoContent with default headers values
func NewPostConfigVlanVlanIDMemberNoContent() *PostConfigVlanVlanIDMemberNoContent {

	return &PostConfigVlanVlanIDMemberNoContent{}
}

// WriteResponse to the client
func (o *PostConfigVlanVlanIDMemberNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// PostConfigVlanVlanIDMemberBadRequestCode is the HTTP code returned for type PostConfigVlanVlanIDMemberBadRequest
const PostConfigVlanVlanIDMemberBadRequestCode int = 400

/*
PostConfigVlanVlanIDMemberBadRequest Malformed arguments for API call

swagger:response postConfigVlanVlanIdMemberBadRequest
*/
type PostConfigVlanVlanIDMemberBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostConfigVlanVlanIDMemberBadRequest creates PostConfigVlanVlanIDMemberBadRequest with default headers values
func NewPostConfigVlanVlanIDMemberBadRequest() *PostConfigVlanVlanIDMemberBadRequest {

	return &PostConfigVlanVlanIDMemberBadRequest{}
}

// WithPayload adds the payload to the post config vlan vlan Id member bad request response
func (o *PostConfigVlanVlanIDMemberBadRequest) WithPayload(payload *models.Error) *PostConfigVlanVlanIDMemberBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post config vlan vlan Id member bad request response
func (o *PostConfigVlanVlanIDMemberBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConfigVlanVlanIDMemberBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostConfigVlanVlanIDMemberUnauthorizedCode is the HTTP code returned for type PostConfigVlanVlanIDMemberUnauthorized
const PostConfigVlanVlanIDMemberUnauthorizedCode int = 401

/*
PostConfigVlanVlanIDMemberUnauthorized Invalid authentication credentials

swagger:response postConfigVlanVlanIdMemberUnauthorized
*/
type PostConfigVlanVlanIDMemberUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostConfigVlanVlanIDMemberUnauthorized creates PostConfigVlanVlanIDMemberUnauthorized with default headers values
func NewPostConfigVlanVlanIDMemberUnauthorized() *PostConfigVlanVlanIDMemberUnauthorized {

	return &PostConfigVlanVlanIDMemberUnauthorized{}
}

// WithPayload adds the payload to the post config vlan vlan Id member unauthorized response
func (o *PostConfigVlanVlanIDMemberUnauthorized) WithPayload(payload *models.Error) *PostConfigVlanVlanIDMemberUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post config vlan vlan Id member unauthorized response
func (o *PostConfigVlanVlanIDMemberUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConfigVlanVlanIDMemberUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostConfigVlanVlanIDMemberForbiddenCode is the HTTP code returned for type PostConfigVlanVlanIDMemberForbidden
const PostConfigVlanVlanIDMemberForbiddenCode int = 403

/*
PostConfigVlanVlanIDMemberForbidden Capacity insufficient

swagger:response postConfigVlanVlanIdMemberForbidden
*/
type PostConfigVlanVlanIDMemberForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostConfigVlanVlanIDMemberForbidden creates PostConfigVlanVlanIDMemberForbidden with default headers values
func NewPostConfigVlanVlanIDMemberForbidden() *PostConfigVlanVlanIDMemberForbidden {

	return &PostConfigVlanVlanIDMemberForbidden{}
}

// WithPayload adds the payload to the post config vlan vlan Id member forbidden response
func (o *PostConfigVlanVlanIDMemberForbidden) WithPayload(payload *models.Error) *PostConfigVlanVlanIDMemberForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post config vlan vlan Id member forbidden response
func (o *PostConfigVlanVlanIDMemberForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConfigVlanVlanIDMemberForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostConfigVlanVlanIDMemberNotFoundCode is the HTTP code returned for type PostConfigVlanVlanIDMemberNotFound
const PostConfigVlanVlanIDMemberNotFoundCode int = 404

/*
PostConfigVlanVlanIDMemberNotFound Vlan interface is not defined

swagger:response postConfigVlanVlanIdMemberNotFound
*/
type PostConfigVlanVlanIDMemberNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostConfigVlanVlanIDMemberNotFound creates PostConfigVlanVlanIDMemberNotFound with default headers values
func NewPostConfigVlanVlanIDMemberNotFound() *PostConfigVlanVlanIDMemberNotFound {

	return &PostConfigVlanVlanIDMemberNotFound{}
}

// WithPayload adds the payload to the post config vlan vlan Id member not found response
func (o *PostConfigVlanVlanIDMemberNotFound) WithPayload(payload *models.Error) *PostConfigVlanVlanIDMemberNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post config vlan vlan Id member not found response
func (o *PostConfigVlanVlanIDMemberNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConfigVlanVlanIDMemberNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostConfigVlanVlanIDMemberConflictCode is the HTTP code returned for type PostConfigVlanVlanIDMemberConflict
const PostConfigVlanVlanIDMemberConflictCode int = 409

/*
PostConfigVlanVlanIDMemberConflict Resource Conflict. VLAN member already exists on this VLAN interface OR Vlan member is being added to 2nd Vlan inteface as an untagged member.

swagger:response postConfigVlanVlanIdMemberConflict
*/
type PostConfigVlanVlanIDMemberConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostConfigVlanVlanIDMemberConflict creates PostConfigVlanVlanIDMemberConflict with default headers values
func NewPostConfigVlanVlanIDMemberConflict() *PostConfigVlanVlanIDMemberConflict {

	return &PostConfigVlanVlanIDMemberConflict{}
}

// WithPayload adds the payload to the post config vlan vlan Id member conflict response
func (o *PostConfigVlanVlanIDMemberConflict) WithPayload(payload *models.Error) *PostConfigVlanVlanIDMemberConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post config vlan vlan Id member conflict response
func (o *PostConfigVlanVlanIDMemberConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConfigVlanVlanIDMemberConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostConfigVlanVlanIDMemberInternalServerErrorCode is the HTTP code returned for type PostConfigVlanVlanIDMemberInternalServerError
const PostConfigVlanVlanIDMemberInternalServerErrorCode int = 500

/*
PostConfigVlanVlanIDMemberInternalServerError Internal service error

swagger:response postConfigVlanVlanIdMemberInternalServerError
*/
type PostConfigVlanVlanIDMemberInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostConfigVlanVlanIDMemberInternalServerError creates PostConfigVlanVlanIDMemberInternalServerError with default headers values
func NewPostConfigVlanVlanIDMemberInternalServerError() *PostConfigVlanVlanIDMemberInternalServerError {

	return &PostConfigVlanVlanIDMemberInternalServerError{}
}

// WithPayload adds the payload to the post config vlan vlan Id member internal server error response
func (o *PostConfigVlanVlanIDMemberInternalServerError) WithPayload(payload *models.Error) *PostConfigVlanVlanIDMemberInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post config vlan vlan Id member internal server error response
func (o *PostConfigVlanVlanIDMemberInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConfigVlanVlanIDMemberInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostConfigVlanVlanIDMemberServiceUnavailableCode is the HTTP code returned for type PostConfigVlanVlanIDMemberServiceUnavailable
const PostConfigVlanVlanIDMemberServiceUnavailableCode int = 503

/*
PostConfigVlanVlanIDMemberServiceUnavailable Maintanence mode

swagger:response postConfigVlanVlanIdMemberServiceUnavailable
*/
type PostConfigVlanVlanIDMemberServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostConfigVlanVlanIDMemberServiceUnavailable creates PostConfigVlanVlanIDMemberServiceUnavailable with default headers values
func NewPostConfigVlanVlanIDMemberServiceUnavailable() *PostConfigVlanVlanIDMemberServiceUnavailable {

	return &PostConfigVlanVlanIDMemberServiceUnavailable{}
}

// WithPayload adds the payload to the post config vlan vlan Id member service unavailable response
func (o *PostConfigVlanVlanIDMemberServiceUnavailable) WithPayload(payload *models.Error) *PostConfigVlanVlanIDMemberServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post config vlan vlan Id member service unavailable response
func (o *PostConfigVlanVlanIDMemberServiceUnavailable) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConfigVlanVlanIDMemberServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
