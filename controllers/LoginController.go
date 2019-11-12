package controllers

import (
	"fmt"
	"net/http"

	"github.com/djamboe/mtools-login-service/interfaces"
)

type LoginController struct {
	interfaces.ILoginService
}

func (controller *LoginController) LoginProcess(res http.ResponseWriter, req *http.Request) {
	//user := models.UserLoginParamModel{}
	//err := json.NewDecoder(req.Body).Decode(&user)
	//if err != nil {
	//	panic(err)
	//}
	//userJson, err := json.Marshal(user)
	//if err != nil {
	//	panic(err)
	//}
	//
	//res.Header().Set("Content-Type", "application/json")
	//res.WriteHeader(http.StatusOK)
	//res.Write(userJson)
	fmt.Fprintf(res, "Hello World!")
}
