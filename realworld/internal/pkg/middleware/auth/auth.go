package auth

import (
	"context"
	"github.com/davecgh/go-spew/spew"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
)

func JWTAuth() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				authJWT := tr.RequestHeader().Get("Authorization")
				spew.Dump(authJWT)
			}
			return handler(ctx, req)

		}
	}

}
