package model

import (
	"golang.org/x/crypto/bcrypt"
)

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string //
	GetName() string
	GetAge() int8

	EncryptPassword()
}

// Cria um struct para receber os dados brutos
// e retorna em formato UserDomainInterface

func NewUserDomain(email, password, name string, age int8) UserDomainInterface {
	return &userDomain{
		email, password, name, age,
	}
}

// Estrutura que armazena as informações do usuário
type userDomain struct {
	email    string
	password string
	name     string
	age      int8
}

// Os métodos Getters servem como uma forma controlada de acessar os valores fora do pacote
func (ud *userDomain) GetEmail() string {
	return ud.email
}
func (ud *userDomain) GetPassword() string {
	return ud.password
}

func (ud *userDomain) GetName() string {
	return ud.name
}
func (ud *userDomain) GetAge() int8 {
	return ud.age
}

// Substitui a senha do usuário pelo hash gerado no bcrypt
func (ud *userDomain) EncryptPassword() {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(ud.password), bcrypt.DefaultCost)
	ud.password = string(hashedPassword)
}
