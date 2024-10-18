package domain

import "modulo/src/exeption"

func (ud *UserDomain) CreateUser() *exeption.Erro {
	ud.EcriptyPassword()
	return nil
}
