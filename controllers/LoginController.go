package controllers

import (
	"net/http"

	"github.com/djamboe/mtools-login-service/interfaces"
	"github.com/djamboe/mtools-login-service/models"
)

type LoginController struct {
	interfaces.ILoginService
}

func (controller *LoginController) LoginProcess(res http.ResponseWriter, req *http.Request) {
	var param models.UserLoginParamModel
}
