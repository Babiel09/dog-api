package models

type Caes struct {
	Id           uint    `json:"id"`
	Nome         string  `json:"nome"`
	Idade        int     `json:"idade"`
	Peso         float64 `json:"peso"`
	Temperamento bool    `json:"temperamento"`
}
