package service

import (
	"context"
)

// Implement yor service methods methods.
// e.x: Foo(ctx context.Context,s string)(rs string, err error)
type UcService interface {
	GDeviceID(ctx context.Context, imei, imsi, mac string) (did string, err error)
}

type stubUcService struct{}

// Get a new instance of the service.
// If you want to add service middleware this is the place to put them.
func New() (s *stubUcService) {
	s = &stubUcService{}
	return s
}

// Implement the business logic of GDeviceID
func (uc *stubUcService) GDeviceID(ctx context.Context, imei string, imsi string, mac string) (did string, err error) {
	return did, err
}
