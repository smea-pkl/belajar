package user_service

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	usersv1 "pkl/backend/gen/pkl/users/v1"
	"pkl/backend/models"

	"connectrpc.com/connect"
	"gorm.io/gorm"
)

type userServiceImpl struct {
	db *gorm.DB
}

// Register implements [usersv1connect.UserServiceHandler].
func (u *userServiceImpl) Register(ctx context.Context, req *connect.Request[usersv1.RegisterRequest]) (*connect.Response[usersv1.RegisterResponse], error) {
	pay := req.Msg

	passHash := md5.Sum([]byte(pay.Password))
	passString := hex.EncodeToString(passHash[:])

	user := models.User{
		Name:     pay.Name,
		Username: pay.Username,
		Password: passString,
	}
	err := u.db.Save(&user).Error
	if err != nil {
		res := usersv1.RegisterResponse{
			Name:     "",
			Username: "",
		}

		return connect.NewResponse(&res), err
	}

	res := usersv1.RegisterResponse{
		Name:     pay.Name,
		Username: pay.Username,
	}
	return connect.NewResponse(&res), nil
}

// Login implements [usersv1connect.UserServiceHandler].
func (u *userServiceImpl) Login(
	ctx context.Context,
	req *connect.Request[usersv1.LoginRequest],
) (*connect.Response[usersv1.LoginResponse], error) {

	pay := req.Msg

	if pay.Username == "heri" && pay.Password == "1234" {
		res := usersv1.LoginResponse{
			Msg: "login berhasil",
		}

		return connect.NewResponse(&res), nil
	}

	return &connect.Response[usersv1.LoginResponse]{}, errors.New("Login Gagal")
}
