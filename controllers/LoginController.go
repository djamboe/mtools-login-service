package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/djamboe/mtools-login-service/interfaces"
	"github.com/djamboe/mtools-login-service/models"
	"github.com/go-chi/render"
)

type LoginController struct {
	interfaces.ILoginService
}

func (controller *LoginController) LoginProcess(res http.ResponseWriter, req *http.Request) {
	var params models.UserLoginParamModel
	s, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(s, &params)
	if err != nil {
		panic(err)
	}
	render.JSON(res, req, params)
}
