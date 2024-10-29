package models

type Perguntas struct {
	Id              uint   `json:"id"`
	TitleOfQuestion string `json:"title"`
	Question        string `json:"question"`
}
