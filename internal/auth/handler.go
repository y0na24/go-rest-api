package auth

import (
	"go-rest-api/configs"
	"go-rest-api/pkg/requestUtils"
	"go-rest-api/pkg/responseUtils"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type AuthHandlerDeps struct {
	Config *configs.Config
}

type AuthHandler struct {
	AuthHandlerDeps
}

func NewAuthHandler(mux *http.ServeMux, deps AuthHandlerDeps) {
	handler := AuthHandler{
		AuthHandlerDeps: deps,
	}
	mux.HandleFunc("POST /auth/login", handler.login())
	mux.HandleFunc("POST /auth/register", handler.register())
}

func (handler *AuthHandler) register() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		payload, err := requestUtils.HandleBody[RegisterRequest](w, req)

		if err != nil {
			responseUtils.JsonResponse(w, err.Error(), 402)
			return
		}

		validate := validator.New()
		err = validate.Struct(payload)

		if err != nil {
			responseUtils.JsonResponse(w, err.Error(), 402)
			return
		}

		res := RegisterResponse{
			Token: "123",
		}
		responseUtils.JsonResponse(w, res, 200)
	}
}

func (handler *AuthHandler) login() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		payload, err := requestUtils.HandleBody[LoginRequest](w, req)

		if err != nil {
			responseUtils.JsonResponse(w, err.Error(), 402)
			return
		}

		validate := validator.New()
		err = validate.Struct(payload)

		if err != nil {
			responseUtils.JsonResponse(w, err.Error(), 402)
			return
		}

		res := LoginResponse{
			Token: "123",
		}
		responseUtils.JsonResponse(w, res, 200)
	}
}
