package controllers

import (
	"github.com/djamboe/mtools-login-service/interfaces"
	"github.com/djamboe/mtools-login-service/models"
)

type LoginController struct {
	interfaces.ILoginService
}

func (controller *LoginController) LoginProcess(username string, password string) (models.UserModel, error) {
	login, err := controller.DoLogin(username, password)
	if err != nil {
		panic(err)
	}
	return login, nil
}
