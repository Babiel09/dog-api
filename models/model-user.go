package models

import (
	"gopkg.in/validator.v2"
)

type User struct {
	Id    uint   `json:"id"`
	Nome  string `json:"nome" validate:"nonzero"`
	Cargo string `json:"cargo" validate:"nonzero"`
	Senha string `json:"senha" validate:"nonzero"`
	Idade int    `json:"idade" validate:"min=8"`
}

// Método para validação
func UserValidation(usuario *User) error {
	if err := validator.Validate(usuario); err != nil {
		return err
	}
	return nil
}
