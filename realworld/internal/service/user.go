package service

import (
	"context"
	v1 "realworld/api/realworld/v1"
)

func (s *RealWorldService) Login(ctx context.Context, request *v1.LoginRequest) (reply *v1.UserReply, err error) {
	return &v1.UserReply{
		User: &v1.UserReply_Data{Email: "test@gmail", Token: "text", Bio: "ll", Image: "dbg"},
	}, nil
}
