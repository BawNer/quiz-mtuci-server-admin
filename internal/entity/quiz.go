package entity

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type JSONBQuestions []Question

func (j JSONBQuestions) Value() (driver.Value, error) {
	valueString, err := json.MarshalIndent(j, "", "\t")
	return string(valueString), err
}

func (j *JSONBQuestions) Scan(value interface{}) error {
	var data []byte
	switch v := value.(type) {
	case string:
		data = []byte(v)
	case []byte:
		data = v
	}
	return json.Unmarshal(data, &j)
}

type AnswersOption struct {
	ID          int    `json:"id"`
	QuestionID  int    `json:"-"`
	Label       string `json:"label"`
	Description string `json:"description,omitempty"`
}

type Question struct {
	ID          int    `json:"id"`
	QuizID      int    `json:"-"`
	Label       string `json:"label"`
	Description string `json:"description"`
}

type Quiz struct {
	ID        int       `json:"id"`
	AuthorID  int       `json:"authorId"`
	AccessFor string    `json:"accessFor"`
	QuizHash  string    `json:"quizHash"`
	Title     string    `json:"title"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type QuizFromDB struct {
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
