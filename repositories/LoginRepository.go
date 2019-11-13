package repositories

import (
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/djamboe/mtools-login-service/interfaces"
	"github.com/djamboe/mtools-login-service/models"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type LoginRepositoryWithCircuitBreaker struct {
	LoginRepository interfaces.ILoginRepository
}

func (repository *LoginRepositoryWithCircuitBreaker) GetUserByEmailAndPassword(username string, password string) (models.UserModel, error) {
	output := make(chan models.UserModel, 1)
	hystrix.ConfigureCommand("get_user_by_username_and_password", hystrix.CommandConfig{Timeout: 1000})
	errors := hystrix.Go("get_user_by_username_and_password", func() error {
		user, _ := repository.LoginRepository.GetUserByEmailAndPassword(username, password)
		output <- user
		return nil
	}, nil)

	select {
	case out := <-output:
		return out, nil
	case err := <-errors:
		println(err)
		return models.UserModel{}, err
	}
}

type LoginRepository struct {
	interfaces.IDbHandler
}

func (repository *LoginRepository) GetUserByEmailAndPassword(username string, password string) (models.UserModel, error) {
	row, err := repository.Query(fmt.Sprintf("SELECT * from user_model where username = '%s' AND password '%s'", username, password))
	if err != nil {
		return models.UserModel{}, err
	}

	var user models.UserModel
	row.Next()
	row.Scan(&user.Id, &user.Name)

	return user, nil

}
