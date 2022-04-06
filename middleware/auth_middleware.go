package middleware

import (
	"golang_restful_api/helper"
	"golang_restful_api/model/web"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.Header.Get("X-API-Key") == "RAHASIA" {
		// ok
		middleware.Handler.ServeHTTP(writer, request)
	} else {
		// error
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)
		webResponse := web.CategoryWebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}
		helper.WriteToResponseBody(writer, webResponse)
	}
}
