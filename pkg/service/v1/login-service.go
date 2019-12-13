package v1

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	jwt "github.com/dgrijalva/jwt-go"
	v1 "github.com/djamboe/mtools-login-service/pkg/api/v1"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
)

// toDoServiceServer is implementation of v1.ToDoServiceServer proto interface
type loginServiceServer struct {
}

// NewToDoServiceServer creates ToDo service
func NewLoginServiceServer() v1.LoginServiceServer {
	return &loginServiceServer{}
}

// checkAPI checks if the API version requested by client is supported by server
func (s *loginServiceServer) checkAPI(api string) error {
	// API version is "" means use current version of the service
	if len(api) > 0 {
		if apiVersion != api {
			return status.Errorf(codes.Unimplemented,
				"unsupported API version: service implements API version '%s', but asked for '%s'", apiVersion, api)
		}
	} else {
		return status.Errorf(codes.Unimplemented,
			"please input your api version")
	}
	return nil
}

// Create new todo task
func (s *loginServiceServer) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginResponse, error) {
	// check if the API version requested by client is supported by server
	message := "Successfully login !"
	errorStatus := false
	if err := s.checkAPI(req.Api); err != nil {
		message = "Unsupported api version !"
		errorStatus = true
	}
	loginController := ServiceContainer().InjectLoginController()
	response, err := loginController.LoginProcess(req.Username, req.Password)

	if err != nil {
		message = "Login failed, an error occured"
		errorStatus = true
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uId":   response.Id,
		"uDbId": response.DbId,
		"uLvl":  response.Level,
		"uMid":  response.MemberId,
		"uP":    response.Parent,
		"uE":    response.UserEmail,
		"uS":    response.Status,
	})
	tokenString, err := token.SignedString([]byte("1A-kYaSQJISPt6zI5ZvBD1g9cNy9SqBr0WTmbZuzUHLpfNt28r0rSxImBwO1rjl_TQp44pafqJ9Y4GQzogiZM8qH6vBByu5_AMtLs0AOcaq2jU0vwkJ6OgrnrWkaJQ1cyQAmjp4Kr5ZfOO_riN8xbdO2C8BT15Ks4OOL_4SBdRT8fEYiHnIutZ29oG17Q0pCN53MTII7Dv-eM5QzrrVojmvrAJK3KoC4bi6Uh_7P6t892c4IWiZnzOpKdK7ZhgW2fQFTurHlrmAgU8WYOE8Eui0FU5WVZtHRBrcbRCGxIbjKeonCbJfJ8BDwaI0WsfTjYclEetAoKTNUZ8sG_mspyQ"))

	if err != nil {
		fmt.Println(err)
	}

	if response.Id == 0 {
		message = "Invalid credentials !"
		errorStatus = true
		tokenString = ""
	}
	return &v1.LoginResponse{
		Api:     apiVersion,
		Message: message,
		Error:   errorStatus,
		Token:   tokenString,
	}, nil
}
