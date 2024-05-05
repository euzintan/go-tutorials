package store

import (
	"example/QuizGame/model"
)

var Questions = []model.Question{
	{ID: "1", QuestionDescription: "What has to be broken before you can use it?", Options: []string{"A Clock", "An Egg", "An Apple"}, AnswerIndex: 1},
	{ID: "2", QuestionDescription: "I’m tall when I’m young, and I’m short when I’m old. What am I?", Options: []string{"A Tree", "A Pencil", "A Candle"}, AnswerIndex: 2},
	{ID: "3", QuestionDescription: "What has a head, a tail, is brown, and has no legs?", Options: []string{"A Penny", "A Snake", "A Drum"}, AnswerIndex: 0},
	{ID: "4", QuestionDescription: "The more you take, the more you leave behind. What am I?", Options: []string{"Footsteps", "Breath", "Thoughts"}, AnswerIndex: 0},
	{ID: "5", QuestionDescription: "What has keys but can’t open locks?", Options: []string{"A Piano", "A Map", "A Secret"}, AnswerIndex: 1},
}
