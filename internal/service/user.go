package service

import (
	"github.com/luis-olivetti/map-zoo-brusque-back-go/internal/model"
	"github.com/luis-olivetti/map-zoo-brusque-back-go/pkg/helper"
	"github.com/spf13/viper"
	"go.uber.org/zap/zapcore"
)

type UserService interface {
	GenerateJWT(user, password string) (*model.UserJWT, error)
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

func (s *userService) GenerateJWT(user, password string) (*model.UserJWT, error) {
	secret := s.config.GetString("jwt.secret")

	s.logger.Info("generating jwt", zapcore.Field{Key: "secret", Type: zapcore.StringType, String: secret})

	token, err := helper.GenerateJWT(user, secret)

	return &model.UserJWT{
		Token: token,
	}, err
}
