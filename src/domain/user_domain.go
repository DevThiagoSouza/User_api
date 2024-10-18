package domain

import (
	"crypto/md5"
	"encoding/hex"
	"modulo/src/exeption"
)

func NewUserDomain(
	email, password, name, sobrenome string,
) *UserDomain {
	return &UserDomain{
		email, password, name, sobrenome,
	}
}

type UserDomain struct {
	Email     string
	Password  string
	Name      string
	SobreNome string
}

func (ud *UserDomain) EcriptyPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.Password))
	ud.Password = hex.EncodeToString(hash.Sum(nil))
}

type UserDomainInterface interface {
	CreateUser() *exeption.Erro
	UpdateUser(string) *exeption.Erro
	UserId(string) (*UserDomain, *exeption.Erro)
	UserEmail(string) (*UserDomain, *exeption.Erro)
	DeleteUser(string) *exeption.Erro
}
