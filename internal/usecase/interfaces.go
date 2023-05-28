package usecase

import (
	"context"
	"quiz-mtuci-server/internal/entity"
)

type AnalyzeTaskIDType string

type (
	QuizRepo interface {
		GetAllQuiz(ctx context.Context) ([]*entity.QuizUI, error)
		GetQuizById(ctx context.Context, quizId int) (*entity.QuizUI, error)
		GetQuizByHash(ctx context.Context, quizHash string) (*entity.QuizUI, error)
		SaveQuiz(ctx context.Context, quiz *entity.QuizUI) (*entity.QuizUI, error)
		SaveReviewers(ctx context.Context, reviewer *entity.Reviewers) error
	}
	AuthRepo interface {
		GetUserByLoginWithPassword(ctx context.Context, login entity.UserLogin) (*entity.User, error)
		GetUserByID(ctx context.Context, id int) (*entity.User, error)
	}
	UseCase interface {
		GetAllQuiz(ctx context.Context) ([]*entity.QuizUI, error)
		GetQuizById(ctx context.Context, quizId int) (*entity.QuizUI, error)
		GetQuizByHash(ctx context.Context, quizHash string) (*entity.QuizUI, error)
		SaveQuiz(ctx context.Context, quiz *entity.QuizUI) (*entity.QuizUI, error)
		GetUserByLoginWithPassword(ctx context.Context, login entity.UserLogin) (*entity.User, error)
		SaveReviewers(ctx context.Context, reviewer *entity.Reviewers) error
	}
)
