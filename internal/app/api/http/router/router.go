package router

import (
	"net/http"

	v1 "github.com/folder-app/internal/app/api/http/v1"
)

func SetupV1(api *v1.API, httpServer *http.ServeMux) {
	httpServer.HandleFunc("POST /api/v1/user/create", api.CreateUser)
}
