package service

import (
	v1 "realworld/api/realworld/v1"
	"realworld/internal/biz"
)

// RealWorldService is a greeter service.
type RealWorldService struct {
	v1.UnimplementedRealWorldServer

	uc *biz.RealWorldUsecase
}

// NewRealWorldService new a greeter service.
func NewRealWorldService(uc *biz.RealWorldUsecase) *RealWorldService {
	return &RealWorldService{uc: uc}
}

// SayHello implements realworld.RealWorldServer.
//func (s *RealWorldService) SayHello(ctx context.Context, in *v1.HelloRequest) (*v1.HelloReply, error) {
//	g, err := s.uc.CreateRealWorld(ctx, &biz.RealWorld{Hello: in.Name})
//	if err != nil {
//		return nil, err
//	}
//	return &v1.HelloReply{Message: "Hello " + g.Hello}, nil
//}
