package service

import (
	"github.com/RafaelFleitas/API-Golang/src/configuration/logger"
	"github.com/RafaelFleitas/API-Golang/src/configuration/rest_err"
	"github.com/RafaelFleitas/API-Golang/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUserService(userDomain model.UserDomainInterface) (model.UserDomainInterface, *rest_err.RestErr) {

	logger.Info("Init CreateUser model",
		zap.String("journey", "createUser"),
	)

	// Criptografa a senha antes de passar para o repositório salvar no banco
	userDomain.EncryptPassword()

	// Repassa para o repositório que vai executar o INSERT no Oracle
	return ud.userRepository.CreateUser(userDomain)
}
