// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/flomesh-io/fsmxlb/api/models"
)

// PostConfigVlanNoContentCode is the HTTP code returned for type PostConfigVlanNoContent
const PostConfigVlanNoContentCode int = 204

/*
PostConfigVlanNoContent OK

swagger:response postConfigVlanNoContent
*/
type PostConfigVlanNoContent struct {
}

// NewPostConfigVlanNoContent creates PostConfigVlanNoContent with default headers values
func NewPostConfigVlanNoContent() *PostConfigVlanNoContent {

	return &PostConfigVlanNoContent{}
}

// WriteResponse to the client
func (o *PostConfigVlanNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// PostConfigVlanBadRequestCode is the HTTP code returned for type PostConfigVlanBadRequest
const PostConfigVlanBadRequestCode int = 400

/*
PostConfigVlanBadRequest Malformed arguments for API call

swagger:response postConfigVlanBadRequest
*/
type PostConfigVlanBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostConfigVlanBadRequest creates PostConfigVlanBadRequest with default headers values
func NewPostConfigVlanBadRequest() *PostConfigVlanBadRequest {

	return &PostConfigVlanBadRequest{}
}

// WithPayload adds the payload to the post config vlan bad request response
func (o *PostConfigVlanBadRequest) WithPayload(payload *models.Error) *PostConfigVlanBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post config vlan bad request response
func (o *PostConfigVlanBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConfigVlanBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostConfigVlanUnauthorizedCode is the HTTP code returned for type PostConfigVlanUnauthorized
const PostConfigVlanUnauthorizedCode int = 401

/*
PostConfigVlanUnauthorized Invalid authentication credentials

swagger:response postConfigVlanUnauthorized
*/
type PostConfigVlanUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostConfigVlanUnauthorized creates PostConfigVlanUnauthorized with default headers values
func NewPostConfigVlanUnauthorized() *PostConfigVlanUnauthorized {

	return &PostConfigVlanUnauthorized{}
}

// WithPayload adds the payload to the post config vlan unauthorized response
func (o *PostConfigVlanUnauthorized) WithPayload(payload *models.Error) *PostConfigVlanUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post config vlan unauthorized response
func (o *PostConfigVlanUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConfigVlanUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostConfigVlanForbiddenCode is the HTTP code returned for type PostConfigVlanForbidden
const PostConfigVlanForbiddenCode int = 403

/*
PostConfigVlanForbidden Capacity insufficient

swagger:response postConfigVlanForbidden
*/
type PostConfigVlanForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostConfigVlanForbidden creates PostConfigVlanForbidden with default headers values
func NewPostConfigVlanForbidden() *PostConfigVlanForbidden {

	return &PostConfigVlanForbidden{}
}

// WithPayload adds the payload to the post config vlan forbidden response
func (o *PostConfigVlanForbidden) WithPayload(payload *models.Error) *PostConfigVlanForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post config vlan forbidden response
func (o *PostConfigVlanForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConfigVlanForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostConfigVlanNotFoundCode is the HTTP code returned for type PostConfigVlanNotFound
const PostConfigVlanNotFoundCode int = 404

/*
PostConfigVlanNotFound Resource not found

swagger:response postConfigVlanNotFound
*/
type PostConfigVlanNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostConfigVlanNotFound creates PostConfigVlanNotFound with default headers values
func NewPostConfigVlanNotFound() *PostConfigVlanNotFound {

	return &PostConfigVlanNotFound{}
}

// WithPayload adds the payload to the post config vlan not found response
func (o *PostConfigVlanNotFound) WithPayload(payload *models.Error) *PostConfigVlanNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post config vlan not found response
func (o *PostConfigVlanNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConfigVlanNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostConfigVlanConflictCode is the HTTP code returned for type PostConfigVlanConflict
const PostConfigVlanConflictCode int = 409

/*
PostConfigVlanConflict Resource Conflict. VLAN already exists OR dependency VRF/VNET not found

swagger:response postConfigVlanConflict
*/
type PostConfigVlanConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostConfigVlanConflict creates PostConfigVlanConflict with default headers values
func NewPostConfigVlanConflict() *PostConfigVlanConflict {

	return &PostConfigVlanConflict{}
}

// WithPayload adds the payload to the post config vlan conflict response
func (o *PostConfigVlanConflict) WithPayload(payload *models.Error) *PostConfigVlanConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post config vlan conflict response
func (o *PostConfigVlanConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConfigVlanConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostConfigVlanInternalServerErrorCode is the HTTP code returned for type PostConfigVlanInternalServerError
const PostConfigVlanInternalServerErrorCode int = 500

/*
PostConfigVlanInternalServerError Internal service error

swagger:response postConfigVlanInternalServerError
*/
type PostConfigVlanInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostConfigVlanInternalServerError creates PostConfigVlanInternalServerError with default headers values
func NewPostConfigVlanInternalServerError() *PostConfigVlanInternalServerError {

	return &PostConfigVlanInternalServerError{}
}

// WithPayload adds the payload to the post config vlan internal server error response
func (o *PostConfigVlanInternalServerError) WithPayload(payload *models.Error) *PostConfigVlanInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post config vlan internal server error response
func (o *PostConfigVlanInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConfigVlanInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostConfigVlanServiceUnavailableCode is the HTTP code returned for type PostConfigVlanServiceUnavailable
const PostConfigVlanServiceUnavailableCode int = 503

/*
PostConfigVlanServiceUnavailable Maintanence mode

swagger:response postConfigVlanServiceUnavailable
*/
type PostConfigVlanServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostConfigVlanServiceUnavailable creates PostConfigVlanServiceUnavailable with default headers values
func NewPostConfigVlanServiceUnavailable() *PostConfigVlanServiceUnavailable {

	return &PostConfigVlanServiceUnavailable{}
}

// WithPayload adds the payload to the post config vlan service unavailable response
func (o *PostConfigVlanServiceUnavailable) WithPayload(payload *models.Error) *PostConfigVlanServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post config vlan service unavailable response
func (o *PostConfigVlanServiceUnavailable) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConfigVlanServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
