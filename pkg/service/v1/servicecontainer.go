package v1

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/djamboe/mtools-login-service/controllers"
	"github.com/djamboe/mtools-login-service/infrastructures"
	"github.com/djamboe/mtools-login-service/repositories"
	"github.com/djamboe/mtools-login-service/services"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
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
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mongoDBConn, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		fmt.Errorf("todo: couldn't connect to mongo: %v", err)
	}
	err = mongoDBConn.Connect(ctx)
	if err != nil {
		fmt.Errorf("todo: mongo client couldn't connect with background context: %v", err)
	}

	if err != nil {
		log.Fatal(err)
	}

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
