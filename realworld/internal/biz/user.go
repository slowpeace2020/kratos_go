package biz

import (
	"context"

	v1 "realworld/api/realworld/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrGreeterUnspecified is Greeter Unspecified.
	ErrGreeterUnspecified = errors.NotFound(v1.ErrorReason_GREETER_UNSPECIFIED.String(), "Greeter Unspecified")
)

// RealWorld is a RealWorld model.

type User struct {
	Username string
}
type UserRepo interface {
	CreateUser(ctx context.Context, user *User) error
}

type ProfileRepo interface {
}

type UserUsecase struct {
	us UserRepo
	pr ProfileRepo

	log *log.Helper
}

// NewUserUsecase is a user usecase.

func NewUserUsecase(us UserRepo, pr ProfileRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{us: us, pr: pr, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) Register(ctx context.Context, user *User) error {
	if err := uc.us.CreateUser(ctx, user); err != nil {

	}
	return nil
}
