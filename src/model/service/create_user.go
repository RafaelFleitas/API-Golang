package service

import (
	"fmt"

	"github.com/RafaelFleitas/API-Golang/src/configuration/logger"
	"github.com/RafaelFleitas/API-Golang/src/configuration/rest_err"
	"github.com/RafaelFleitas/API-Golang/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(userDomain model.UserDomainInterface) *rest_err.RestErr {

	logger.Info("Init CreateUser model",
		zap.String("journey", "createUser"),
	)

	userDomain.EncryptPassword()
	fmt.Println(userDomain.GetPassword())

	return nil
}
