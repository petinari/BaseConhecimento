package models

import (
	"errors"
	"strings"

	"github.com/badoux/checkmail"
)

type User struct {
	ID       uint64 `json:"id,omitempty"`
	Name     string `json:"nome,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Admin    bool   `json:"admin,omitempty"`
}

func (u *User) Preparar(etapa string) error {
	if erro := u.validar(etapa); erro != nil {
		return erro
	}
	if erro := u.formatar(etapa); erro != nil {
		return erro
	}
	return nil
}

func (usuario *User) validar(etapa string) error {
	if usuario.Name == "" {
		return errors.New("o nome é obrigatório, não pode estar em branco")
	}

	if usuario.Email == "" {
		return errors.New("o email é obrigatório, não pode estar em branco")
	}
	if erro := checkmail.ValidateFormat((usuario.Email)); erro != nil {
		return errors.New("o email inserio é invalido")
	}
	if etapa == "cadastro" && usuario.Password == "" {
		return errors.New("a senha é obrigatória, não pode estar em branco")
	}
	return nil
}

func (usuario *User) formatar(etapa string) error {
	usuario.Name = strings.TrimSpace(usuario.Name)

	usuario.Email = strings.TrimSpace(usuario.Email)
	// if etapa == "cadastro" {
	// 	senhaComHash, erro := seguranca.Hash(usuario.Password)
	// 	if erro != nil {
	// 		return erro
	// 	}
	// 	usuario.Password = string(senhaComHash)
	// }
	return nil

}
