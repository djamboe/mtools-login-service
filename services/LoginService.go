package services

import (
	"github.com/djamboe/mtools-login-service/interfaces"
	"github.com/djamboe/mtools-login-service/models"
)

type LoginService struct {
	interfaces.ILoginRepository
}

func (service *LoginService) DoLogin(username string, password string) (models.UserModel, error) {
	var user models.UserModel
	user, err := service.GetUserByEmailAndPassword(username, password)
	if err != nil {
		return models.UserModel{}, nil
	}
	return user, nil
}
