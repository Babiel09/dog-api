package models

import "time"

type Perguntas struct {
	Id              uint      `json:"id"`
	CreatedAt       time.Time `json:"created"`
	TitleOfQuestion string    `json:"title"`
	Question        string    `json:"question"`
}
