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

func (s *RealWorldService) Register(ctx context.Context, req *v1.RegistrationRequest) (reply *v1.UserReply, err error) {
	return &v1.UserReply{}, nil
}

func (s *RealWorldService) GetCurrentUser(ctx context.Context, req *v1.GetCurrentUserRequest) (reply *v1.UserReply, err error) {
	return &v1.UserReply{}, nil
}
func (s *RealWorldService) UpdateUser(ctx context.Context, req *v1.UpdateUserRequest) (reply *v1.UserReply, err error) {
	return &v1.UserReply{}, nil
}
func (s *RealWorldService) GetProfile(ctx context.Context, req *v1.GetProfileRequest) (reply *v1.ProfileReply, err error) {
	return &v1.ProfileReply{}, nil
}
func (s *RealWorldService) FollowUser(ctx context.Context, req *v1.FollowUserRequest) (reply *v1.ProfileReply, err error) {
	return &v1.ProfileReply{}, nil
}
func (s *RealWorldService) UnfollowUser(ctx context.Context, req *v1.UnfollowUserRequest) (reply *v1.ProfileReply, err error) {
	return &v1.ProfileReply{}, nil
}
