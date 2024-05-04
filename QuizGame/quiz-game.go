package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type question struct {
	ID                  string   `json: "id"`
	QuestionDescription string   `json: "question_des"`
	Options             []string `json: "options"`
	AnswerIndex         int      `json: "answer_index"`
}

var questions = []question{
	{ID: "1", QuestionDescription: "What has to be broken before you can use it?", Options: []string{"A Clock", "An Egg", "An Apple"}, AnswerIndex: 1},
	{ID: "2", QuestionDescription: "I’m tall when I’m young, and I’m short when I’m old. What am I?", Options: []string{"A Tree", "A Pencil", "A Candle"}, AnswerIndex: 2},
	{ID: "3", QuestionDescription: "What has a head, a tail, is brown, and has no legs?", Options: []string{"A Penny", "A Snake", "A Drum"}, AnswerIndex: 0},
	{ID: "4", QuestionDescription: "The more you take, the more you leave behind. What am I?", Options: []string{"Footsteps", "Breath", "Thoughts"}, AnswerIndex: 0},
	{ID: "5", QuestionDescription: "What has keys but can’t open locks?", Options: []string{"A Piano", "A Map", "A Secret"}, AnswerIndex: 1},
}

func GetAllQuestions(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, questions)
}

func GetQuestionById(c *gin.Context) {
	question, err := getQuestionById(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
	}

	c.IndentedJSON(http.StatusOK, question)
}

func getQuestionById(id string) (*question, error) {
	for i, q := range questions {
		if q.ID == id {
			return &questions[i], nil
		}
	}
	return nil, fmt.Errorf("Book with ID: %v cannot be found.", id)
}

func AnswerAQuestion(c *gin.Context) {
	questionId := c.Param("id")
	question, err := getQuestionById(questionId)

	var answer struct {
		AnswerIndex int `json:"answer_idx"`
	}

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	if err := c.BindJSON(&answer); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Poorly formed answer request"})
		return
	}

	if question.AnswerIndex == answer.AnswerIndex {
		c.IndentedJSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Correct, the answer is %v, %v", questionId, question.Options[answer.AnswerIndex])})
	} else if answer.AnswerIndex >= len(question.Options) {
		c.IndentedJSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Wrong, %v is not even a valid option", answer.AnswerIndex)})
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Wrong, the answer is not %v", question.Options[answer.AnswerIndex])})
	}
}

func CreateAQuestion(c *gin.Context) {

	var submittedQuestion struct {
		QuestionDescription string   `json:"question_des"`
		Options             []string `json:"options"`
		AnswerIndex         int      `json:"answer_index"`
	}

	if err := c.BindJSON(&submittedQuestion); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Poorly formed question"})
		return
	}

	if submittedQuestion.AnswerIndex >= len(submittedQuestion.Options) {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Poorly formed question"})
		return
	}

	newQuestion := question{
		QuestionDescription: submittedQuestion.QuestionDescription,
		Options:             submittedQuestion.Options,
		AnswerIndex:         submittedQuestion.AnswerIndex,
		ID:                  strconv.Itoa(len(questions) + 1),
	}

	questions = append(questions, newQuestion)

	c.IndentedJSON(http.StatusOK, newQuestion)
}

func main() {
	router := gin.Default()
	router.GET("/questions", GetAllQuestions)
	router.GET("/questions/:id", GetQuestionById)
	router.POST("/questions", CreateAQuestion)
	router.POST("/answer/:id", AnswerAQuestion)
	router.Run("localhost:8080")
}
