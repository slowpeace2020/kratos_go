package biz

import (
	"context"

	v1 "realworld/api/realworld/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// RealWorld is a RealWorld model.
type RealWorld struct {
	Hello string
}

type ArticleRepo interface {
}

type CommentRepo interface {
}

type TagRepo interface {
}

type SocialUsecase struct {
	arpo ArticleRepo
	copo CommentRepo
	tapo TagRepo

	log *log.Helper
}

// RealWorldUsecase is a RealWorld usecase.

func NewSocialUsecase(arpo ArticleRepo, copo CommentRepo, tapo TagRepo, logger log.Logger) *SocialUsecase {
	return &SocialUsecase{arpo: arpo, copo: copo, tapo: tapo, log: log.NewHelper(logger)}
}

// CreateArticle creates a Article, and returns the new Article.
func (uc *SocialUsecase) CreateArticle(ctx context.Context) error {
	return nil
}
