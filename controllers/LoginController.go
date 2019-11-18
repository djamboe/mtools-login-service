package controllers

import (
	"encoding/json"
	"github.com/djamboe/mtools-login-service/models"
	"github.com/djamboe/mtools-login-service/viewmodels"
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
	var response viewmodels.LoginVM

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

	if login.Id != 0 {
		response.Error = false
		response.Id = login.Id
		response.Username = login.Username
	}else{
		response.Error = true
		response.Id = login.Id
		response.Username = login.Username
	}

	render.JSON(res, req,response)

}
