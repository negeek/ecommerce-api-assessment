package controllers

import (
	"net/http"

	"github.com/negeek/ecommerce-api-assessment/repositories"
	"github.com/negeek/ecommerce-api-assessment/services"
	"github.com/negeek/ecommerce-api-assessment/utils"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{userService: userService}
}

func (uc *UserController) Register(w http.ResponseWriter, r *http.Request) {
	var user repositories.User
	err := utils.Unmarshall(r.Body, &user)
	if err != nil {
		utils.JsonResponse(w, false, http.StatusBadRequest, "error reading request data", nil)
		return
	}

	err = uc.userService.Register(&user)
	if err != nil {
		utils.JsonResponse(w, false, http.StatusBadRequest, err.Error(), nil)
		return
	}

	utils.JsonResponse(w, true, http.StatusCreated, "User registered successfully", nil)
}

func (uc *UserController) Login(w http.ResponseWriter, r *http.Request) {
	var user repositories.User
	err := utils.Unmarshall(r.Body, &user)
	if err != nil {
		utils.JsonResponse(w, false, http.StatusBadRequest, "error reading request data", nil)
		return
	}

	token, err := uc.userService.Login(&user)
	if err != nil {
		utils.JsonResponse(w, false, http.StatusUnauthorized, err.Error(), nil)
		return
	}

	utils.JsonResponse(w, true, http.StatusOK, "Login successful", map[string]string{"token": token})
}
