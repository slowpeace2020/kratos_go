package biz

import (
	"context"
	"fmt"
	"golang.org/x/crypto/bcrypt"

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
	Email        string
	Token        string
	Username     string
	Bio          string
	Image        string
	PasswordHash string
}

type UserLogin struct {
	Email    string
	Token    string
	Username string
	Bio      string
	Image    string
}

func hashPassword(password string) string {
	b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v", b)
	return string(b)
}

func verifyPassword(hashed, input string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(input)); err != nil {
		return false
	}

	return true
}

type UserRepo interface {
	CreateUser(ctx context.Context, user *User) error
	GetUserByEmail(ctx context.Context, email string) (*User, error)
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

func (uc *UserUsecase) Register(ctx context.Context, username, email, password string) (*UserLogin, error) {
	user := &User{
		Email:        email,
		Username:     username,
		PasswordHash: hashPassword(password),
	}
	if err := uc.us.CreateUser(ctx, user); err != nil {
		return nil, err
	}
	return &UserLogin{Email: email, Username: username, Token: "xxx"}, nil
}

func (uc *UserUsecase) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	return nil, nil
}

func (uc *UserUsecase) Login(ctx context.Context, email, password string) (*UserLogin, error) {

	u, err := uc.us.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if !verifyPassword(u.PasswordHash, password) {
		return nil, errors.New(1, "Login failed", "")
	}
	return &UserLogin{Email: u.Email, Username: u.Username, Bio: u.Bio, Image: u.Image, Token: "abx"}, nil
}
