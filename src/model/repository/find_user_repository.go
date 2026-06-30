package repository

import (
	"context"
	"database/sql"

	"github.com/RafaelFleitas/API-Golang/src/configuration/logger"
	"github.com/RafaelFleitas/API-Golang/src/configuration/rest_err"
	"github.com/RafaelFleitas/API-Golang/src/model"
	"github.com/RafaelFleitas/API-Golang/src/model/repository/entity"
	"github.com/RafaelFleitas/API-Golang/src/model/repository/entity/converter"
	"go.uber.org/zap"
)

func (ur *userRepository) FindUserByEmailRepository(email string) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindByEmail user repository")

	row := ur.databaseConnection.QueryRowContext(
		context.Background(),
		"SELECT id, name, email, password, age FROM users WHERE email = :1",
		email,
	)

	userEntity := &entity.UserEntity{}

	err := row.Scan(&userEntity.ID, &userEntity.Name, &userEntity.Email, &userEntity.Password, &userEntity.Age)

	if err != nil {
		logger.Error("Error trying to find user by email", err)
		if err == sql.ErrNoRows {
			return nil, rest_err.NewNotFoundError("user not found")
		}
		return nil, rest_err.NewInternalServerError("Error trying to find user by email")
	}

	logger.Info("FindUserByEmail repository executed successfully",
		zap.String("journey", "FindUserByEmail"),
		zap.String("email: ", email))

	return converter.ConvertEntityToDomain(userEntity), nil

}

func (ur *userRepository) FindUserByIdRepository(id int64) (model.UserDomainInterface, *rest_err.RestErr) {
	logger.Info("Init FindById user repository")

	row := ur.databaseConnection.QueryRowContext(
		context.Background(),
		"SELECT id, name, email, password, age FROM users WHERE ID = :1",
		id,
	)

	userEntity := &entity.UserEntity{}

	err := row.Scan(&userEntity.ID, &userEntity.Name, &userEntity.Email, &userEntity.Password, &userEntity.Age)

	if err != nil {
		logger.Error("Error trying to find user by ID", err)
		if err == sql.ErrNoRows {
			return nil, rest_err.NewNotFoundError("user not found")
		}
		return nil, rest_err.NewInternalServerError("Error trying to find user by ID")
	}

	logger.Info("FindUserByEmail repository executed successfully",
		zap.String("journey", "FindUserById"),
		zap.Int64("ID: ", id))

	return converter.ConvertEntityToDomain(userEntity), nil

}
