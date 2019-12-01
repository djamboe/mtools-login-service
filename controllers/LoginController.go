package controllers

import (
	"github.com/djamboe/mtools-login-service/interfaces"
	"github.com/djamboe/mtools-login-service/viewmodels"
)

type LoginController struct {
	interfaces.ILoginService
}

func (controller *LoginController) LoginProcess(username string, password string) (viewmodels.LoginVM, error) {
	var response viewmodels.LoginVM

	login, err := controller.DoLogin(username, password)
	if err != nil {
		panic(err)
	}

	if login.Id != 0 {
		response.Error = false
		response.Id = login.Id
		response.Username = login.Username
	} else {
		response.Error = true
		response.Id = login.Id
		response.Username = login.Username
	}

	return response, nil

}
