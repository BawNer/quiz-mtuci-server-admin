package usecase

import (
	"context"
	"quiz-mtuci-server/internal/entity"
)

type AnalyzeTaskIDType string

type (
	QuizRepo interface {
		GetAllQuiz(ctx context.Context) ([]*entity.QuizFromDB, error)
		GetQuizById(ctx context.Context, quizId int) (*entity.QuizUI, error)
		SaveQuiz(ctx context.Context, quiz *entity.QuizUISaveRequest) (*entity.QuizUISaveRequest, error)
		DeleteQuiz(ctx context.Context, quizID int) error
	}
	MtuciRepo interface {
		GetUserByLoginWithPassword(ctx context.Context, login entity.UserLogin) (*entity.User, error)
		GetUserByID(ctx context.Context, id int) (*entity.User, error)
		GetAllGroups(ctx context.Context) ([]*entity.Group, error)
		GetGroupById(ctx context.Context, groupId int) (*entity.Group, error)
	}
	UseCase interface {
		GetAllQuiz(ctx context.Context) ([]*entity.QuizUI, error)
		GetQuizById(ctx context.Context, quizId int) (*entity.QuizUI, error)
		SaveQuiz(ctx context.Context, quiz *entity.QuizUISaveRequest) (*entity.QuizUI, error)
		GetUserByLoginWithPassword(ctx context.Context, login entity.UserLogin) (*entity.User, error)
		DeleteQuiz(ctx context.Context, quizID int) error
		GetAllGroups(ctx context.Context) ([]*entity.Group, error)
	}
)
