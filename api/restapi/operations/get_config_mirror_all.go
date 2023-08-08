// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/flomesh-io/fsmxlb/api/models"
)

// GetConfigMirrorAllHandlerFunc turns a function with the right signature into a get config mirror all handler
type GetConfigMirrorAllHandlerFunc func(GetConfigMirrorAllParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetConfigMirrorAllHandlerFunc) Handle(params GetConfigMirrorAllParams) middleware.Responder {
	return fn(params)
}

// GetConfigMirrorAllHandler interface for that can handle valid get config mirror all params
type GetConfigMirrorAllHandler interface {
	Handle(GetConfigMirrorAllParams) middleware.Responder
}

// NewGetConfigMirrorAll creates a new http.Handler for the get config mirror all operation
func NewGetConfigMirrorAll(ctx *middleware.Context, handler GetConfigMirrorAllHandler) *GetConfigMirrorAll {
	return &GetConfigMirrorAll{Context: ctx, Handler: handler}
}

/*
	GetConfigMirrorAll swagger:route GET /config/mirror/all getConfigMirrorAll

# Get

Get
*/
type GetConfigMirrorAll struct {
	Context *middleware.Context
	Handler GetConfigMirrorAllHandler
}

func (o *GetConfigMirrorAll) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetConfigMirrorAllParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetConfigMirrorAllOKBody get config mirror all o k body
//
// swagger:model GetConfigMirrorAllOKBody
type GetConfigMirrorAllOKBody struct {

	// mirr attr
	MirrAttr []*models.MirrorGetEntry `json:"mirrAttr"`
}

// Validate validates this get config mirror all o k body
func (o *GetConfigMirrorAllOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMirrAttr(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetConfigMirrorAllOKBody) validateMirrAttr(formats strfmt.Registry) error {
	if swag.IsZero(o.MirrAttr) { // not required
		return nil
	}

	for i := 0; i < len(o.MirrAttr); i++ {
		if swag.IsZero(o.MirrAttr[i]) { // not required
			continue
		}

		if o.MirrAttr[i] != nil {
			if err := o.MirrAttr[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getConfigMirrorAllOK" + "." + "mirrAttr" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getConfigMirrorAllOK" + "." + "mirrAttr" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get config mirror all o k body based on the context it is used
func (o *GetConfigMirrorAllOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateMirrAttr(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetConfigMirrorAllOKBody) contextValidateMirrAttr(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.MirrAttr); i++ {

		if o.MirrAttr[i] != nil {

			if swag.IsZero(o.MirrAttr[i]) { // not required
				return nil
			}

			if err := o.MirrAttr[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getConfigMirrorAllOK" + "." + "mirrAttr" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getConfigMirrorAllOK" + "." + "mirrAttr" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetConfigMirrorAllOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetConfigMirrorAllOKBody) UnmarshalBinary(b []byte) error {
	var res GetConfigMirrorAllOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
