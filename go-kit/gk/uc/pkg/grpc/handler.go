package grpc

import (
	"context"
	"errors"

	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/yiv/go-lab/go-kit/gk/uc/pkg/endpoints"
	"github.com/yiv/go-lab/go-kit/gk/uc/pkg/grpc/pb"
	oldcontext "golang.org/x/net/context"
)

type grpcServer struct {
	gDeviceID grpctransport.Handler
}

// MakeGRPCServer makes a set of endpoints available as a gRPC server.
func MakeGRPCServer(endpoints endpoints.Endpoints) (req pb.UcServer) {
	req = &grpcServer{
		gDeviceID: grpctransport.NewServer(
			endpoints.GDeviceIDEndpoint,
			DecodeGRPCGDeviceIDRequest,
			EncodeGRPCGDeviceIDResponse,
		),
	}
	return req
}

// DecodeGRPCGDeviceIDRequest is a transport/grpc.DecodeRequestFunc that converts a
// gRPC request to a user-domain request. Primarily useful in a server.
// TODO: Do not forget to implement the decoder, you can find an example here :
// https://github.com/go-kit/kit/blob/master/examples/addsvc/transport_grpc.go#L62-L65
func DecodeGRPCGDeviceIDRequest(_ context.Context, grpcReq interface{}) (req interface{}, err error) {
	err = errors.New("'GDeviceID' Decoder is not impelement")
	return req, err
}

// EncodeGRPCGDeviceIDResponse is a transport/grpc.EncodeResponseFunc that converts a
// user-domain response to a gRPC reply. Primarily useful in a server.
// TODO: Do not forget to implement the encoder, you can find an example here :
// https://github.com/go-kit/kit/blob/master/examples/addsvc/transport_grpc.go#L62-L65
func EncodeGRPCGDeviceIDResponse(_ context.Context, grpcReply interface{}) (res interface{}, err error) {
	err = errors.New("'GDeviceID' Encoder is not impelement")
	return res, err
}

func (s *grpcServer) GDeviceID(ctx oldcontext.Context, req *pb.GDeviceIDRequest) (rep *pb.GDeviceIDReply, err error) {
	_, rp, err := s.gDeviceID.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	rep = rp.(*pb.GDeviceIDReply)
	return rep, err
}
