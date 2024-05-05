package model

type Question struct {
	ID                  string   `json: "id"`
	QuestionDescription string   `json: "question_des"`
	Options             []string `json: "options"`
	AnswerIndex         int      `json: "answer_index"`
}
