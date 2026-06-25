package model

import (
	"golang.org/x/crypto/bcrypt"
)

// UserDomainInterface define o que um usuário precisa ter na aplicação.
// Toda a aplicação usa essa interface — nunca a struct diretamente.
// Isso protege os campos privados e facilita trocar a implementação no futuro.
type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetName() string
	GetAge() int8
	GetID() int64
	SetID(id int64)

	EncryptPassword()
}

// userDomain é a struct privada que guarda os dados do usuário.
// Os campos são privados para que só sejam acessados pelos métodos abaixo.
type userDomain struct {
	id       int64
	email    string
	password string
	name     string
	age      int8
}

// Getters — única forma de ler os dados do usuário fora desse pacote
func (ud *userDomain) GetEmail() string    { return ud.email }
func (ud *userDomain) GetPassword() string { return ud.password }
func (ud *userDomain) GetName() string     { return ud.name }
func (ud *userDomain) GetAge() int8        { return ud.age }
func (ud *userDomain) GetID() int64        { return ud.id }
func (ud *userDomain) SetID(id int64)      { ud.id = id }

// NewUserDomain é o construtor do usuário. Recebe os dados da requisição e devolve a interface.
func NewUserDomain(email, password, name string, age int8) UserDomainInterface {
	return &userDomain{
		email:    email,
		password: password,
		name:     name,
		age:      age,
	}
}

// EncryptPassword substitui a senha em texto puro pelo hash bcrypt antes de salvar no banco
func (ud *userDomain) EncryptPassword() {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(ud.password), bcrypt.DefaultCost)
	ud.password = string(hashedPassword)
}
