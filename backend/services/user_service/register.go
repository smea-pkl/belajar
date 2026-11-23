package user_service

import (
	"context"
	"net/http"
	"pkl/backend/gen/pkl/users/v1/usersv1connect"

	"connectrpc.com/connect"
	"gorm.io/gorm"
)

type CheckLoginInterceptor struct {
}

// WrapStreamingClient implements [connect.Interceptor].
func (c *CheckLoginInterceptor) WrapStreamingClient(connect.StreamingClientFunc) connect.StreamingClientFunc {
	panic("unimplemented")
}

// WrapStreamingHandler implements [connect.Interceptor].
func (c *CheckLoginInterceptor) WrapStreamingHandler(connect.StreamingHandlerFunc) connect.StreamingHandlerFunc {
	panic("unimplemented")
}

// WrapUnary implements [connect.Interceptor].
func (c *CheckLoginInterceptor) WrapUnary(next connect.UnaryFunc) connect.UnaryFunc {
	return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {

		req.Header().Get("Authorization")

		return next(ctx, req)
	}
}

func UserServiceRegister(muxServer *http.ServeMux, db *gorm.DB) error {

	interceptor := connect.WithInterceptors(
		&CheckLoginInterceptor{},
	)

	path, handler := usersv1connect.NewUserServiceHandler(&userServiceImpl{db: db}, interceptor)
	muxServer.Handle(path, handler)

	return nil
}
