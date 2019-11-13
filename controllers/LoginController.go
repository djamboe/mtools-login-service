package controllers

import (
	"encoding/json"
	"github.com/djamboe/mtools-login-service/models"
	"github.com/go-chi/render"
	"io/ioutil"
	"net/http"

	"github.com/djamboe/mtools-login-service/interfaces"
)

type LoginController struct {
	interfaces.ILoginService
}

func (controller *LoginController) LoginProcess(res http.ResponseWriter, req *http.Request) {
	var user models.UserLoginParamModel
	s, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(s, &user)
	if err != nil {
		panic(err)
	}

	username := user.Username
	password := user.Password

	login, err := controller.DoLogin(username, password)
	if err != nil {
		panic(err)
	}
	render.JSON(res, req, login.Username)

}
