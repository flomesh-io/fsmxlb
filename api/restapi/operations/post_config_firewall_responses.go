// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/flomesh-io/fnlb/api/models"
)

// PostConfigFirewallNoContentCode is the HTTP code returned for type PostConfigFirewallNoContent
const PostConfigFirewallNoContentCode int = 204

/*
PostConfigFirewallNoContent OK

swagger:response postConfigFirewallNoContent
*/
type PostConfigFirewallNoContent struct {
}

// NewPostConfigFirewallNoContent creates PostConfigFirewallNoContent with default headers values
func NewPostConfigFirewallNoContent() *PostConfigFirewallNoContent {

	return &PostConfigFirewallNoContent{}
}

// WriteResponse to the client
func (o *PostConfigFirewallNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// PostConfigFirewallBadRequestCode is the HTTP code returned for type PostConfigFirewallBadRequest
const PostConfigFirewallBadRequestCode int = 400

/*
PostConfigFirewallBadRequest Malformed arguments for API call

swagger:response postConfigFirewallBadRequest
*/
type PostConfigFirewallBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostConfigFirewallBadRequest creates PostConfigFirewallBadRequest with default headers values
func NewPostConfigFirewallBadRequest() *PostConfigFirewallBadRequest {

	return &PostConfigFirewallBadRequest{}
}

// WithPayload adds the payload to the post config firewall bad request response
func (o *PostConfigFirewallBadRequest) WithPayload(payload *models.Error) *PostConfigFirewallBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post config firewall bad request response
func (o *PostConfigFirewallBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConfigFirewallBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostConfigFirewallUnauthorizedCode is the HTTP code returned for type PostConfigFirewallUnauthorized
const PostConfigFirewallUnauthorizedCode int = 401

/*
PostConfigFirewallUnauthorized Invalid authentication credentials

swagger:response postConfigFirewallUnauthorized
*/
type PostConfigFirewallUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostConfigFirewallUnauthorized creates PostConfigFirewallUnauthorized with default headers values
func NewPostConfigFirewallUnauthorized() *PostConfigFirewallUnauthorized {

	return &PostConfigFirewallUnauthorized{}
}

// WithPayload adds the payload to the post config firewall unauthorized response
func (o *PostConfigFirewallUnauthorized) WithPayload(payload *models.Error) *PostConfigFirewallUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post config firewall unauthorized response
func (o *PostConfigFirewallUnauthorized) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConfigFirewallUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostConfigFirewallForbiddenCode is the HTTP code returned for type PostConfigFirewallForbidden
const PostConfigFirewallForbiddenCode int = 403

/*
PostConfigFirewallForbidden Capacity insufficient

swagger:response postConfigFirewallForbidden
*/
type PostConfigFirewallForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostConfigFirewallForbidden creates PostConfigFirewallForbidden with default headers values
func NewPostConfigFirewallForbidden() *PostConfigFirewallForbidden {

	return &PostConfigFirewallForbidden{}
}

// WithPayload adds the payload to the post config firewall forbidden response
func (o *PostConfigFirewallForbidden) WithPayload(payload *models.Error) *PostConfigFirewallForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post config firewall forbidden response
func (o *PostConfigFirewallForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConfigFirewallForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostConfigFirewallNotFoundCode is the HTTP code returned for type PostConfigFirewallNotFound
const PostConfigFirewallNotFoundCode int = 404

/*
PostConfigFirewallNotFound Resource not found

swagger:response postConfigFirewallNotFound
*/
type PostConfigFirewallNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostConfigFirewallNotFound creates PostConfigFirewallNotFound with default headers values
func NewPostConfigFirewallNotFound() *PostConfigFirewallNotFound {

	return &PostConfigFirewallNotFound{}
}

// WithPayload adds the payload to the post config firewall not found response
func (o *PostConfigFirewallNotFound) WithPayload(payload *models.Error) *PostConfigFirewallNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post config firewall not found response
func (o *PostConfigFirewallNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConfigFirewallNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostConfigFirewallConflictCode is the HTTP code returned for type PostConfigFirewallConflict
const PostConfigFirewallConflictCode int = 409

/*
PostConfigFirewallConflict Resource Conflict. VLAN already exists OR dependency VRF/VNET not found

swagger:response postConfigFirewallConflict
*/
type PostConfigFirewallConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostConfigFirewallConflict creates PostConfigFirewallConflict with default headers values
func NewPostConfigFirewallConflict() *PostConfigFirewallConflict {

	return &PostConfigFirewallConflict{}
}

// WithPayload adds the payload to the post config firewall conflict response
func (o *PostConfigFirewallConflict) WithPayload(payload *models.Error) *PostConfigFirewallConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post config firewall conflict response
func (o *PostConfigFirewallConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConfigFirewallConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostConfigFirewallInternalServerErrorCode is the HTTP code returned for type PostConfigFirewallInternalServerError
const PostConfigFirewallInternalServerErrorCode int = 500

/*
PostConfigFirewallInternalServerError Internal service error

swagger:response postConfigFirewallInternalServerError
*/
type PostConfigFirewallInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostConfigFirewallInternalServerError creates PostConfigFirewallInternalServerError with default headers values
func NewPostConfigFirewallInternalServerError() *PostConfigFirewallInternalServerError {

	return &PostConfigFirewallInternalServerError{}
}

// WithPayload adds the payload to the post config firewall internal server error response
func (o *PostConfigFirewallInternalServerError) WithPayload(payload *models.Error) *PostConfigFirewallInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post config firewall internal server error response
func (o *PostConfigFirewallInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConfigFirewallInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostConfigFirewallServiceUnavailableCode is the HTTP code returned for type PostConfigFirewallServiceUnavailable
const PostConfigFirewallServiceUnavailableCode int = 503

/*
PostConfigFirewallServiceUnavailable Maintanence mode

swagger:response postConfigFirewallServiceUnavailable
*/
type PostConfigFirewallServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostConfigFirewallServiceUnavailable creates PostConfigFirewallServiceUnavailable with default headers values
func NewPostConfigFirewallServiceUnavailable() *PostConfigFirewallServiceUnavailable {

	return &PostConfigFirewallServiceUnavailable{}
}

// WithPayload adds the payload to the post config firewall service unavailable response
func (o *PostConfigFirewallServiceUnavailable) WithPayload(payload *models.Error) *PostConfigFirewallServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post config firewall service unavailable response
func (o *PostConfigFirewallServiceUnavailable) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostConfigFirewallServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
