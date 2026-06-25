package repository

import (
	"context"

	"github.com/RafaelFleitas/API-Golang/src/configuration/logger"
	"github.com/RafaelFleitas/API-Golang/src/configuration/rest_err"
	"github.com/RafaelFleitas/API-Golang/src/model"
)

func (ur *userRepository) CreateUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_err.RestErr) {

	logger.Info("Init create user repository")

	// ExecContext executa o INSERT no Oracle. Os :1, :2, :3, :4 são os valores passados na ordem abaixo.
	// O ID é gerado automaticamente pelo banco (GENERATED ALWAYS AS IDENTITY), não precisa ser informado.
	_, err := ur.databaseConnection.ExecContext(
		context.Background(),
		"INSERT INTO users (name, email, password, age) VALUES (:1, :2, :3, :4)",
		userDomain.GetName(),
		userDomain.GetEmail(),
		userDomain.GetPassword(),
		userDomain.GetAge(),
	)

	if err != nil {
		return nil, rest_err.NewInternalServerError(err.Error())
	}

	return userDomain, nil

}
