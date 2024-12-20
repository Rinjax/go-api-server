package api

import (
	"jax/api/endpoints"
	"net/http"
)

func routes(endpoint *endpoints.Endpoint) *http.ServeMux {
	mux := &http.ServeMux{}
	mux.HandleFunc("POST /auth/register", endpoint.Register)
	mux.HandleFunc("POST /auth/login", endpoint.Login)
	mux.HandleFunc("POST /auth/logout", endpoint.Logout)

	return mux
}