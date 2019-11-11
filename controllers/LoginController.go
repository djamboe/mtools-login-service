package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/djamboe/mtools-login-service/interfaces"

	"github.com/djamboe/mtools-login-service/viewmodels"
	"github.com/go-chi/chi"
)

type LoginController struct {
	interfaces.ILoginService
}

func (controller *LoginController) LoginProcess(res http.ResponseWriter, req *http.Request) {

	player1Name := chi.URLParam(req, "player1")
	player2Name := chi.URLParam(req, "player2")

	scores, err := controller.GetScores(player1Name, player2Name)
	if err != nil {
		//Handle error
	}

	json.NewEncoder(res).Encode(viewmodels.ScoresVM{scores})
}
