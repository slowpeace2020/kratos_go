package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"realworld/internal/biz"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func (r *userRepo) GetUserByEmail(ctx context.Context, email string) (*biz.User, error) {
	//TODO implement me
	return nil, nil
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepo) CreateUser(ctx context.Context, u *biz.User) error {
	return nil
}

type profileRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo .
func NewProfileRepo(data *Data, logger log.Logger) biz.ProfileRepo {
	return &profileRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

//func (r *realworldRepo) Save(ctx context.Context, g *biz.RealWorld) (*biz.RealWorld, error) {
//	return g, nil
//}
//
//func (r *realworldRepo) Update(ctx context.Context, g *biz.RealWorld) (*biz.RealWorld, error) {
//	return g, nil
//}
//
//func (r *realworldRepo) FindByID(context.Context, int64) (*biz.RealWorld, error) {
//	return nil, nil
//}
//
//func (r *realworldRepo) ListByHello(context.Context, string) ([]*biz.RealWorld, error) {
//	return nil, nil
//}
//
//func (r *realworldRepo) ListAll(context.Context) ([]*biz.RealWorld, error) {
//	return nil, nil
//}
