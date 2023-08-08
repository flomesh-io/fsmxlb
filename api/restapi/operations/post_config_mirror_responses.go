// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/flomesh-io/fsmxlb/api/models"
)

// PostConfigMirrorNoContentCode is the HTTP code returned for type PostConfigMirrorNoContent
const PostConfigMirrorNoContentCode int = 204

/*
PostConfigMirrorNoContent OK

swagger:response postConfigMirrorNoContent
*/
type PostConfigMirrorNoContent struct {
}

// NewPostConfigMirrorNoContent creates PostConfigMirrorNoContent with default headers values
func NewPostConfigMirrorNoContent() *PostConfigMirrorNoContent {

	return &PostConfigMirrorNoContent{}
}

// WriteResponse to the client
func (o *PostConfigMirrorNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// PostConfigMirrorBadRequestCode is the HTTP code returned for type PostConfigMirrorBadRequest
const PostConfigMirrorBadRequestCode int = 400

/*
PostConfigMirrorBadRequest Malformed arguments for API call

swagger:response postConfigMirrorBadRequest
*/
type PostConfigMirrorBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostConfigMirrorBadRequest creates PostConfigMirrorBadRequest with default headers values
func NewPostConfigMirrorBadRequest() *PostConfigMirrorBadRequest {

	return &PostConfigMirrorBadRequest{}
}

// WithPayload adds the payload to the post config mirror bad request response
func (o *PostConfigMirrorBadRequest) WithPayload(payload *models.Error) *PostConfigMirrorBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post config mirror bad request response
func (o *PostConfigMirrorBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConfigMirrorBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostConfigMirrorUnauthorizedCode is the HTTP code returned for type PostConfigMirrorUnauthorized
const PostConfigMirrorUnauthorizedCode int = 401

/*
PostConfigMirrorUnauthorized Invalid authentication credentials

swagger:response postConfigMirrorUnauthorized
*/
type PostConfigMirrorUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostConfigMirrorUnauthorized creates PostConfigMirrorUnauthorized with default headers values
func NewPostConfigMirrorUnauthorized() *PostConfigMirrorUnauthorized {

	return &PostConfigMirrorUnauthorized{}
}

// WithPayload adds the payload to the post config mirror unauthorized response
func (o *PostConfigMirrorUnauthorized) WithPayload(payload *models.Error) *PostConfigMirrorUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post config mirror unauthorized response
func (o *PostConfigMirrorUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConfigMirrorUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostConfigMirrorForbiddenCode is the HTTP code returned for type PostConfigMirrorForbidden
const PostConfigMirrorForbiddenCode int = 403

/*
PostConfigMirrorForbidden Capacity insufficient

swagger:response postConfigMirrorForbidden
*/
type PostConfigMirrorForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostConfigMirrorForbidden creates PostConfigMirrorForbidden with default headers values
func NewPostConfigMirrorForbidden() *PostConfigMirrorForbidden {

	return &PostConfigMirrorForbidden{}
}

// WithPayload adds the payload to the post config mirror forbidden response
func (o *PostConfigMirrorForbidden) WithPayload(payload *models.Error) *PostConfigMirrorForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post config mirror forbidden response
func (o *PostConfigMirrorForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConfigMirrorForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostConfigMirrorNotFoundCode is the HTTP code returned for type PostConfigMirrorNotFound
const PostConfigMirrorNotFoundCode int = 404

/*
PostConfigMirrorNotFound Resource not found

swagger:response postConfigMirrorNotFound
*/
type PostConfigMirrorNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostConfigMirrorNotFound creates PostConfigMirrorNotFound with default headers values
func NewPostConfigMirrorNotFound() *PostConfigMirrorNotFound {

	return &PostConfigMirrorNotFound{}
}

// WithPayload adds the payload to the post config mirror not found response
func (o *PostConfigMirrorNotFound) WithPayload(payload *models.Error) *PostConfigMirrorNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post config mirror not found response
func (o *PostConfigMirrorNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConfigMirrorNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostConfigMirrorConflictCode is the HTTP code returned for type PostConfigMirrorConflict
const PostConfigMirrorConflictCode int = 409

/*
PostConfigMirrorConflict Resource Conflict. VLAN already exists OR dependency VRF/VNET not found

swagger:response postConfigMirrorConflict
*/
type PostConfigMirrorConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostConfigMirrorConflict creates PostConfigMirrorConflict with default headers values
func NewPostConfigMirrorConflict() *PostConfigMirrorConflict {

	return &PostConfigMirrorConflict{}
}

// WithPayload adds the payload to the post config mirror conflict response
func (o *PostConfigMirrorConflict) WithPayload(payload *models.Error) *PostConfigMirrorConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post config mirror conflict response
func (o *PostConfigMirrorConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConfigMirrorConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostConfigMirrorInternalServerErrorCode is the HTTP code returned for type PostConfigMirrorInternalServerError
const PostConfigMirrorInternalServerErrorCode int = 500

/*
PostConfigMirrorInternalServerError Internal service error

swagger:response postConfigMirrorInternalServerError
*/
type PostConfigMirrorInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostConfigMirrorInternalServerError creates PostConfigMirrorInternalServerError with default headers values
func NewPostConfigMirrorInternalServerError() *PostConfigMirrorInternalServerError {

	return &PostConfigMirrorInternalServerError{}
}

// WithPayload adds the payload to the post config mirror internal server error response
func (o *PostConfigMirrorInternalServerError) WithPayload(payload *models.Error) *PostConfigMirrorInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post config mirror internal server error response
func (o *PostConfigMirrorInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConfigMirrorInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostConfigMirrorServiceUnavailableCode is the HTTP code returned for type PostConfigMirrorServiceUnavailable
const PostConfigMirrorServiceUnavailableCode int = 503

/*
PostConfigMirrorServiceUnavailable Maintanence mode

swagger:response postConfigMirrorServiceUnavailable
*/
type PostConfigMirrorServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostConfigMirrorServiceUnavailable creates PostConfigMirrorServiceUnavailable with default headers values
func NewPostConfigMirrorServiceUnavailable() *PostConfigMirrorServiceUnavailable {

	return &PostConfigMirrorServiceUnavailable{}
}

// WithPayload adds the payload to the post config mirror service unavailable response
func (o *PostConfigMirrorServiceUnavailable) WithPayload(payload *models.Error) *PostConfigMirrorServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post config mirror service unavailable response
func (o *PostConfigMirrorServiceUnavailable) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConfigMirrorServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
