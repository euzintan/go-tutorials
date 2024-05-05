package main

import (
	"example/QuizGame/api"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/questions", api.GetAllQuestions)
	router.GET("/questions/:id", api.GetQuestionById)
	router.POST("/questions", api.CreateAQuestion)
	router.POST("/answer/:id", api.AnswerAQuestion)
	router.Run("localhost:8080")
}
