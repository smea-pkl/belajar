package main

import (
	"log/slog"
	"net/http"
	"pkl/backend/configs"
	"pkl/backend/database"
	"pkl/backend/pkgs/san_config"
	"pkl/backend/services/user_service"

	connectcors "connectrpc.com/cors"
	"github.com/rs/cors"
)

func main() {

	var cfg configs.AppConfig

	err := san_config.NewConfiguration(&cfg, san_config.NewYamlSecretProvider("./config.yaml"))
	if err != nil {
		panic(err)
	}

	// inisiasi database
	db, err := database.ConnectDB(cfg)
	if err != nil {
		panic(err)
	}

	// 1. buat object server -> mux
	muxServer := http.NewServeMux()

	// 2. buat api url /hello, terus ketika dibuka, nanti respon "hello"
	muxServer.HandleFunc("GET /hello", handleHello)

	// register user service
	user_service.UserServiceRegister(muxServer, db)

	// 3. running server mux

	slog.Info("running server", "listen", cfg.ListenAddress) // kasih log di cli

	// set protocol
	protocols := new(http.Protocols)
	protocols.SetHTTP1(true)
	protocols.SetUnencryptedHTTP2(true)

	// instance server baru
	appServer := &http.Server{
		Addr:      cfg.ListenAddress,
		Handler:   withCORS(muxServer),
		Protocols: protocols,
	}

	// running server
	err = appServer.ListenAndServe()
	if err != nil {
		panic(err)
	}

}

func handleHello(w http.ResponseWriter, r *http.Request) {
	content := "hello"

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(content))
}

func withCORS(handler http.Handler) http.Handler {
	// connectcors.AllowedHeaders() covers only the CONNECT PROTOCOL headers
	// (Content-Type, Connect-Protocol-Version, timeouts, …). It does NOT include Authorization
	// — that is ours to add.
	//
	// Omitting it fails in a nasty, asymmetric way: Login carries no Authorization header, so it
	// sails through, and EVERY authenticated call afterwards is killed by the browser at
	// preflight — before the request ever reaches the server. Nothing appears in the server log.
	// From the UI it looks like the API returned nothing.
	allowedHeaders := append(connectcors.AllowedHeaders(), "Authorization")

	// PUT is not a Connect method; it is added for the local document file endpoint, which clients
	// upload to directly with a cross-origin PUT.
	allowedMethods := append(connectcors.AllowedMethods(), http.MethodPut)

	middleware := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: allowedMethods,
		AllowedHeaders: allowedHeaders,
		ExposedHeaders: connectcors.ExposedHeaders(),
	})

	app := middleware.Handler(handler)

	// An AGENT's requests are not our app's — ChatGPT, Claude — so they bypass the app's origin list and answer their own
	// CORS. The bearer, or PKCE, is what authorizes them, never an origin.
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		app.ServeHTTP(w, r)
	})
}
