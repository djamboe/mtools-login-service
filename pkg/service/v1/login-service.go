package v1

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

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
	if err := s.checkAPI(req.Api); err != nil {
		return nil, err
	}

	loginController := ServiceContainer().InjectLoginController()
	response, err := loginController.LoginProcess(req.Username, req.Password)

	if err != nil {
		panic(err)
	}
	return &v1.LoginResponse{
		Api:     apiVersion,
		Message: response.Username,
		Error:   false,
		User:    nil,
	}, nil
}
