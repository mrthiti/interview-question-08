package handlers

import (
	"net/http"
	"strconv"
	"thaibev-assignment/backend/database"
	"thaibev-assignment/backend/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetQuestions(c *gin.Context) {
	var questions []models.Question
	database.DB.Order("order_number asc").Find(&questions)
	c.JSON(http.StatusOK, questions)
}

func CreateQuestion(c *gin.Context) {
	var req models.CreateQuestionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var count int64
	database.DB.Model(&models.Question{}).Count(&count)

	question := models.Question{
		OrderNumber: int(count) + 1,
		Question:    req.Question,
		Choice1:     req.Choice1,
		Choice2:     req.Choice2,
		Choice3:     req.Choice3,
		Choice4:     req.Choice4,
	}

	database.DB.Create(&question)
	c.JSON(http.StatusCreated, question)
}

func DeleteQuestion(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var question models.Question
	if result := database.DB.First(&question, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "question not found"})
		return
	}

	deletedOrder := question.OrderNumber
	database.DB.Delete(&question)

	// re-number remaining questions
	database.DB.Model(&models.Question{}).
		Where("order_number > ?", deletedOrder).
		UpdateColumn("order_number", gorm.Expr("order_number - 1"))

	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
