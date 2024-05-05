package api

import (
	"example/QuizGame/model"
	"example/QuizGame/store"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getQuestionById(id string) (*model.Question, error) {
	for i, q := range store.Questions {
		if q.ID == id {
			return &store.Questions[i], nil
		}
	}
	return nil, fmt.Errorf("Book with ID: %v cannot be found.", id)
}

func GetAllQuestions(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, store.Questions)
}

func GetQuestionById(c *gin.Context) {
	question, err := getQuestionById(c.Param("id"))

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
	}

	c.IndentedJSON(http.StatusOK, question)
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

	newQuestion := model.Question{
		QuestionDescription: submittedQuestion.QuestionDescription,
		Options:             submittedQuestion.Options,
		AnswerIndex:         submittedQuestion.AnswerIndex,
		ID:                  strconv.Itoa(len(store.Questions) + 1),
	}

	store.Questions = append(store.Questions, newQuestion)

	c.IndentedJSON(http.StatusOK, newQuestion)
}
