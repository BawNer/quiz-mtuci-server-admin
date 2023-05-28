package entity

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type JSONB []interface{}

func (j JSONB) Value() (driver.Value, error) {
	valueString, err := json.MarshalIndent(j, "", "\t")
	return string(valueString), err
}

func (j *JSONB) Scan(value interface{}) error {
	var data []byte
	switch v := value.(type) {
	case string:
		data = []byte(v)
	case []byte:
		data = v
	}
	return json.Unmarshal(data, &j)
}

type Reviewers struct {
	ID       int       `json:"id"`
	UserID   int       `json:"userID"`
	QuizID   int       `json:"quizID"`
	Answers  JSONB     `json:"answers"`
	ClosedAt time.Time `json:"closedAt"`
}
