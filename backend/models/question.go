package models

import "gorm.io/gorm"

type Question struct {
	gorm.Model
	OrderNumber int    `json:"order_number" gorm:"not null"`
	Question    string `json:"question" gorm:"not null"`
	Choice1     string `json:"choice_1"`
	Choice2     string `json:"choice_2"`
	Choice3     string `json:"choice_3"`
	Choice4     string `json:"choice_4"`
}

type CreateQuestionRequest struct {
	Question string `json:"question" binding:"required"`
	Choice1  string `json:"choice_1"`
	Choice2  string `json:"choice_2"`
	Choice3  string `json:"choice_3"`
	Choice4  string `json:"choice_4"`
}
