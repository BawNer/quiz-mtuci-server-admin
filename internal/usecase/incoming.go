package usecase

import (
	"context"
	"quiz-mtuci-server/internal/entity"
	"quiz-mtuci-server/pkg/logger"
	"time"
)

type ServiceUseCase struct {
	logger *logger.Logger
	jwt    JWT
	repo   QuizRepo
	auth   AuthRepo
}

func New(logger *logger.Logger, j JWT, r QuizRepo, a AuthRepo) *ServiceUseCase {
	return &ServiceUseCase{
		logger: logger,
		jwt:    j,
		repo:   r,
		auth:   a,
	}
}

func (s *ServiceUseCase) GetAllQuiz(ctx context.Context) ([]*entity.QuizUI, error) {
	quizzes, err := s.repo.GetAllQuiz(ctx)
	if err != nil {
		return nil, err
	}

	for _, quiz := range quizzes {
		user, err := s.auth.GetUserByID(ctx, quiz.AuthorID)
		if err != nil {
			return nil, err
		}

		quiz.Author = user
	}

	return quizzes, nil
}

func (s *ServiceUseCase) GetQuizById(ctx context.Context, quizId int) (*entity.QuizUI, error) {
	return s.repo.GetQuizById(ctx, quizId)
}

func (s *ServiceUseCase) GetQuizByHash(ctx context.Context, quizHash string) (*entity.QuizUI, error) {
	return s.repo.GetQuizByHash(ctx, quizHash)
}

func (s *ServiceUseCase) SaveQuiz(ctx context.Context, quiz *entity.QuizUI) (*entity.QuizUI, error) {
	quiz.AuthorID = s.jwt.Parse(ctx.Value("token").(map[string]interface{})).ID
	return s.repo.SaveQuiz(ctx, quiz)
}

func (s *ServiceUseCase) GetUserByLoginWithPassword(ctx context.Context, user entity.UserLogin) (*entity.User, error) {
	foundedUser, err := s.auth.GetUserByLoginWithPassword(ctx, user)
	if err != nil {
		return nil, err
	}
	token, err := s.jwt.Create(time.Hour*24, foundedUser)
	if err != nil {
		return nil, err
	}
	response := &entity.User{
		ID:         foundedUser.ID,
		Email:      foundedUser.Email,
		Name:       foundedUser.Name,
		PassText:   foundedUser.PassText,
		NumberZach: foundedUser.NumberZach,
		IsStudent:  foundedUser.IsStudent,
		Token:      token,
	}

	return response, nil
}

func (s *ServiceUseCase) SaveReviewers(ctx context.Context, reviewer *entity.Reviewers) error {
	reviewer.UserID = s.jwt.Parse(ctx.Value("token").(map[string]interface{})).ID
	return s.repo.SaveReviewers(ctx, reviewer)
}
