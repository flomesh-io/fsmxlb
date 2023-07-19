// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteConfigTunnelVxlanVxlanIDPeerPeerIPHandlerFunc turns a function with the right signature into a delete config tunnel vxlan vxlan ID peer peer IP handler
type DeleteConfigTunnelVxlanVxlanIDPeerPeerIPHandlerFunc func(DeleteConfigTunnelVxlanVxlanIDPeerPeerIPParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteConfigTunnelVxlanVxlanIDPeerPeerIPHandlerFunc) Handle(params DeleteConfigTunnelVxlanVxlanIDPeerPeerIPParams) middleware.Responder {
	return fn(params)
}

// DeleteConfigTunnelVxlanVxlanIDPeerPeerIPHandler interface for that can handle valid delete config tunnel vxlan vxlan ID peer peer IP params
type DeleteConfigTunnelVxlanVxlanIDPeerPeerIPHandler interface {
	Handle(DeleteConfigTunnelVxlanVxlanIDPeerPeerIPParams) middleware.Responder
}

// NewDeleteConfigTunnelVxlanVxlanIDPeerPeerIP creates a new http.Handler for the delete config tunnel vxlan vxlan ID peer peer IP operation
func NewDeleteConfigTunnelVxlanVxlanIDPeerPeerIP(ctx *middleware.Context, handler DeleteConfigTunnelVxlanVxlanIDPeerPeerIPHandler) *DeleteConfigTunnelVxlanVxlanIDPeerPeerIP {
	return &DeleteConfigTunnelVxlanVxlanIDPeerPeerIP{Context: ctx, Handler: handler}
}

/*
	DeleteConfigTunnelVxlanVxlanIDPeerPeerIP swagger:route DELETE /config/tunnel/vxlan/{vxlanID}/peer/{PeerIP} deleteConfigTunnelVxlanVxlanIdPeerPeerIp

# Remove a one of vxlan remote(peer) ip address configuration

Return a list of existing tunnels of a type. If there're no tunnels to return, empty list will be returned.
*/
type DeleteConfigTunnelVxlanVxlanIDPeerPeerIP struct {
	Context *middleware.Context
	Handler DeleteConfigTunnelVxlanVxlanIDPeerPeerIPHandler
}

func (o *DeleteConfigTunnelVxlanVxlanIDPeerPeerIP) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteConfigTunnelVxlanVxlanIDPeerPeerIPParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}