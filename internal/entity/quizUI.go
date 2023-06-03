package entity

import (
	"time"
)

type QuizUI struct {
	ID        int           `json:"id"`
	AuthorID  int           `json:"-"`
	Author    *User         `json:"author"`
	AccessFor []*Group      `json:"accessFor"`
	QuizHash  string        `json:"quizHash"`
	Title     string        `json:"title"`
	Questions []QuestionsUI `json:"questions"`
	Active    bool          `json:"active"`
	CreatedAt time.Time     `json:"createdAt"`
	UpdatedAt time.Time     `json:"updatedAt"`
}

type QuizUISaveResponse struct {
	ID        int           `json:"id"`
	AuthorID  int           `json:"-"`
	Author    *User         `json:"author"`
	AccessFor []*Group      `json:"accessFor"`
	QuizHash  string        `json:"quizHash"`
	Title     string        `json:"title"`
	Questions []QuestionsUI `json:"questions"`
	Active    bool          `json:"active"`
	CreatedAt time.Time     `json:"createdAt"`
	UpdatedAt time.Time     `json:"updatedAt"`
}

type QuizUISaveRequest struct {
	ID        int           `json:"id"`
	AuthorID  int           `json:"-"`
	Author    *User         `json:"author"`
	AccessFor string        `json:"accessFor"`
	QuizHash  string        `json:"quizHash"`
	Title     string        `json:"title"`
	Questions []QuestionsUI `json:"questions"`
	Active    bool          `json:"active"`
	CreatedAt time.Time     `json:"createdAt"`
	UpdatedAt time.Time     `json:"updatedAt"`
}

type QuestionsUI struct {
	ID             int             `json:"id"`
	Label          string          `json:"label"`
	Description    string          `json:"description"`
	AnswersOptions []AnswersOption `json:"answersOptions"`
}

type QuizResponseUI struct {
	Success     bool    `json:"success"`
	Description string  `json:"description"`
	Quiz        *QuizUI `json:"quiz"`
}

type QuizzesResponseUI struct {
	Success     bool      `json:"success"`
	Description string    `json:"description"`
	Quizzes     []*QuizUI `json:"quizzes"`
}
