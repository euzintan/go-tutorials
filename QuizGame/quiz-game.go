package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type question struct {
	ID                  string   `json: "id"`
	QuestionDescription string   `json: "question_des"`
	Options             []string `json: "options"`
	AnswerIndex         int      `json: "answer_index"`
}

type userAnswer struct {
	ID     string `json: "id"`
	Answer int    `json: "answer"`
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
	question, err := getQuestionById(c.Param("id"))
	var answer userAnswer

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	if err := c.BindJSON(&answer); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Poorly formed answer request"})
		return
	}

	if question.AnswerIndex == answer.Answer {
		c.IndentedJSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Correct, the answer is %v, %v", answer.ID, question.Options[answer.Answer])})
	} else if answer.Answer >= len(question.Options) {
		c.IndentedJSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Wrong, %v is not even a valid option", answer.Answer)})
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Wrong, the answer is not %v", question.Options[answer.Answer])})
	}
}

func main() {
	router := gin.Default()
	router.GET("/questions", GetAllQuestions)
	router.GET("/questions/:id", GetQuestionById)
	router.POST("/questions/:id", AnswerAQuestion)
	router.Run("localhost:8080")
}

// var name string
// var age uint
// var livesLeft uint = 3

// CheckIfGameOver := func() bool {
// 	if livesLeft <= 0 {
// 		fmt.Println("Womp Womp Game Over")
// 	} else {
// 		fmt.Printf("You just lost 1 life, you have %v lives left\n", livesLeft)
// 	}
// 	return (livesLeft <= 0)
// }

// fmt.Print("Enter your name: ")
// fmt.Scan(&name)
// fmt.Printf("Welcome to my game, %v\nEnter your age: ", name)
// fmt.Scan(&age)

// if age >= 10 {
// 	fmt.Printf("%v?! You're old, but wise\n", age)
// } else {
// 	fmt.Printf("%v?! too young!\nCome back in %v years\n", age, 10-age)
// 	return
// }

// var answer string
// var ended bool = false

// for !ended {
// 	fmt.Println("What do you hear first, Lightning or Thunder?")

// 	fmt.Scan(&answer)

// 	if strings.ToUpper(answer) == "LIGHTNING" {
// 		fmt.Println("Wrong, you can't hear lightning")
// 		livesLeft--
// 		ended = CheckIfGameOver()
// 	} else if strings.ToUpper(answer) == "THUNDER" {
// 		fmt.Println("Correct!\nCongratulations I'm out of questions for you!")
// 		return
// 	} else {
// 		fmt.Println("That's not one of the options")
// 		livesLeft--
// 		ended = CheckIfGameOver()
// 	}
// }
