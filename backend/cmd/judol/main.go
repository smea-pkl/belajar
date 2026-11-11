package main

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	usersv1 "pkl/backend/gen/pkl/users/v1"
	"pkl/backend/gen/pkl/users/v1/usersv1connect"

	"connectrpc.com/connect"
)

type userServiceImpl struct{}

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

func main() {
	// 1. buat object server -> mux
	muxServer := http.NewServeMux()

	path, handler := usersv1connect.NewUserServiceHandler(&userServiceImpl{})

	// 2. buat api url /hello, terus ketika dibuka, nanti respon "hello"
	muxServer.HandleFunc("GET /hello", handleHello)
	muxServer.Handle(path, handler)

	// 3. running server mux

	slog.Info("running server at localhost:8000") // kasih log di cli

	// set protocol
	protocols := new(http.Protocols)
	protocols.SetHTTP1(true)
	protocols.SetUnencryptedHTTP2(true)

	// instance server baru
	appServer := &http.Server{
		Addr:      "localhost:8000",
		Handler:   muxServer,
		Protocols: protocols,
	}

	// running server
	err := appServer.ListenAndServe()
	if err != nil {
		panic(err)
	}

}

type User struct {
	Name     string
	Username string
	Image    string
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	// 1. data
	user := User{
		Name:     "Bambang",
		Username: "bambang123",
		Image:    "",
	}

	// 2. ubah data struct ke format json
	content, err := json.MarshalIndent(user, "", "  ")
	if err != nil {

		// jika terjadi error

		w.WriteHeader(http.StatusInternalServerError)
		pesan := "gagal encoding json"
		w.Write([]byte(pesan))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(content)
}
