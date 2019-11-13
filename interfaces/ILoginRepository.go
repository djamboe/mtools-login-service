package interfaces

import (
	"github.com/djamboe/mtools-login-service/models"
)

type ILoginRepository interface {
	GetUserByEmailAndPassword(username string, password string) (models.UserModel, error)
}
