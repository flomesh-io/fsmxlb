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

// GetConfigPortAllHandlerFunc turns a function with the right signature into a get config port all handler
type GetConfigPortAllHandlerFunc func(GetConfigPortAllParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetConfigPortAllHandlerFunc) Handle(params GetConfigPortAllParams) middleware.Responder {
	return fn(params)
}

// GetConfigPortAllHandler interface for that can handle valid get config port all params
type GetConfigPortAllHandler interface {
	Handle(GetConfigPortAllParams) middleware.Responder
}

// NewGetConfigPortAll creates a new http.Handler for the get config port all operation
func NewGetConfigPortAll(ctx *middleware.Context, handler GetConfigPortAllHandler) *GetConfigPortAll {
	return &GetConfigPortAll{Context: ctx, Handler: handler}
}

/*
	GetConfigPortAll swagger:route GET /config/port/all getConfigPortAll

# Get all of the port interfaces

Get all of the port interfaces.
*/
type GetConfigPortAll struct {
	Context *middleware.Context
	Handler GetConfigPortAllHandler
}

func (o *GetConfigPortAll) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetConfigPortAllParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// GetConfigPortAllOKBody get config port all o k body
//
// swagger:model GetConfigPortAllOKBody
type GetConfigPortAllOKBody struct {

	// port attr
	PortAttr []*models.PortEntry `json:"portAttr"`
}

// Validate validates this get config port all o k body
func (o *GetConfigPortAllOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePortAttr(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetConfigPortAllOKBody) validatePortAttr(formats strfmt.Registry) error {
	if swag.IsZero(o.PortAttr) { // not required
		return nil
	}

	for i := 0; i < len(o.PortAttr); i++ {
		if swag.IsZero(o.PortAttr[i]) { // not required
			continue
		}

		if o.PortAttr[i] != nil {
			if err := o.PortAttr[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getConfigPortAllOK" + "." + "portAttr" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getConfigPortAllOK" + "." + "portAttr" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get config port all o k body based on the context it is used
func (o *GetConfigPortAllOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidatePortAttr(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetConfigPortAllOKBody) contextValidatePortAttr(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.PortAttr); i++ {

		if o.PortAttr[i] != nil {

			if swag.IsZero(o.PortAttr[i]) { // not required
				return nil
			}

			if err := o.PortAttr[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getConfigPortAllOK" + "." + "portAttr" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getConfigPortAllOK" + "." + "portAttr" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetConfigPortAllOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetConfigPortAllOKBody) UnmarshalBinary(b []byte) error {
	var res GetConfigPortAllOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
