package model

import (
	"fmt"

	"github.com/RafaelFleitas/API-Golang/src/configuration/logger"
	"github.com/RafaelFleitas/API-Golang/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *UserDomain) CreateUser() *rest_err.RestErr {

	logger.Info("Init CreateUser model",
		zap.String("journey", "createUser"),
	)

	ud.EncryptPassword()
	fmt.Println(ud.Password)

	return nil
}
