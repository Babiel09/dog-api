package models

import (
	"gopkg.in/validator.v2"
)

type Caes struct {
	Id      uint   `json:"id"`
	Title   string `json:"title" validate:"nonzero"`
	Legenda string `json:"legenda" validate:"nonzero"`
}

//Criando o método de validação das informações

func ValidarInformations(cao *Caes) error {
	//Verifica se não há um erro na vlidação da minha API
	if err := validator.Validate(cao); err != nil {
		return err
	}
	//Caso não tenha retornara como nulo
	return nil
}
