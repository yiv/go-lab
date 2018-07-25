package endpoints

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/yiv/go-lab/go-kit/gk/uc/pkg/service"
)

// Endpoints collects all of the endpoints that compose an add service. It's
// meant to be used as a helper struct, to collect all of the endpoints into a
// single parameter.

type Endpoints struct {
	GDeviceIDEndpoint endpoint.Endpoint
}
type GDeviceIDRequest struct {
	Imei string
	Imsi string
	Mac  string
}
type GDeviceIDResponse struct {
	Did string
	Err error
}

func New(svc service.UcService) (ep Endpoints) {
	ep.GDeviceIDEndpoint = MakeGDeviceIDEndpoint(svc)
	return ep
}

// MakeGDeviceIDEndpoint returns an endpoint that invokes GDeviceID on the service.
// Primarily useful in a server.
func MakeGDeviceIDEndpoint(svc service.UcService) (ep endpoint.Endpoint) {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GDeviceIDRequest)
		did, err := svc.GDeviceID(ctx, req.Imei, req.Imsi, req.Mac)
		return GDeviceIDResponse{Did: did, Err: err}, nil
	}
}
