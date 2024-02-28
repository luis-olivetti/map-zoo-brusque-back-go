package service

import (
	"github.com/luis-olivetti/map-zoo-brusque-back-go/internal/model"
	"github.com/luis-olivetti/map-zoo-brusque-back-go/pkg/helper"
	"github.com/spf13/viper"
)

type UserService interface {
	Authenticate(user, password string) (bool, error)
	GenerateJWT(user string) (*model.UserJWT, error)
}

type userService struct {
	*Service
	config *viper.Viper
}

func NewUserService(service *Service, config *viper.Viper) UserService {
	return &userService{
		Service: service,
		config:  config,
	}
}

func (s *userService) GenerateJWT(user string) (*model.UserJWT, error) {
	secret := s.config.GetString("jwt.secret")

	token, err := helper.GenerateJWT(user, secret)

	return &model.UserJWT{
		Token: token,
	}, err
}

func (s *userService) Authenticate(user, password string) (bool, error) {
	hashedUser := helper.GenerateSHA256(user)
	hashedPassword := helper.GenerateSHA256(password)

	storedUsername := s.config.GetString("login.username")
	storedPassword := s.config.GetString("login.password")

	return hashedUser == storedUsername && hashedPassword == storedPassword, nil
}
