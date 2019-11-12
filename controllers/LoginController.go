package controllers

import (
	"net/http"

	"github.com/djamboe/mtools-login-service/interfaces"
)

type LoginController struct {
	interfaces.ILoginService
}

func (controller *LoginController) LoginProcess(res http.ResponseWriter, req *http.Request) {

}
