package main

import (
	"database/sql"
	"github.com/djamboe/mtools-login-service/controllers"
	"github.com/djamboe/mtools-login-service/infrastructures"
	"github.com/djamboe/mtools-login-service/repositories"
	"github.com/djamboe/mtools-login-service/services"
	"sync"
)

type IserviceContainer interface {
	InjectLoginController() controllers.LoginController
}

type kernel struct{}

func (k *kernel) InjectLoginController() controllers.LoginController {
	mysqlConn, _ := sql.Open("mysql", "root:@tcp(localhost:3306)/marketing-tools?charset=utf8")
	mysqlHandler := &infrastructures.DBHandler{}
	mysqlHandler.Conn = mysqlConn

	//sqlConn, _ := sql.Open("sqlite3", "user.db")
	//sqliteHandler := &infrastructures.SQLiteHandler{}
	//sqliteHandler.Conn = sqlConn

	loginRepository := &repositories.LoginRepository{mysqlHandler}
	loginService := &services.LoginService{&repositories.LoginRepositoryWithCircuitBreaker{loginRepository}}
	loginController := controllers.LoginController{loginService}

	return loginController
}

var (
	k             *kernel
	containerOnce sync.Once
)

func ServiceContainer() IserviceContainer {
	if k == nil {
		containerOnce.Do(func() {
			k = &kernel{}
		})
	}
	return k
}
