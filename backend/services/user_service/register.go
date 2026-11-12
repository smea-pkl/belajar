package user_service

import (
	"net/http"
	"pkl/backend/gen/pkl/users/v1/usersv1connect"

	"gorm.io/gorm"
)

func UserServiceRegister(muxServer *http.ServeMux, db *gorm.DB) error {
	path, handler := usersv1connect.NewUserServiceHandler(&userServiceImpl{db: db})
	muxServer.Handle(path, handler)

	return nil
}
