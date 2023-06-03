package usecase

import (
	"context"
	"quiz-mtuci-server/internal/entity"
	"quiz-mtuci-server/pkg/logger"
	"strconv"
	"strings"
	"time"
)

type ServiceUseCase struct {
	logger *logger.Logger
	jwt    JWT
	repo   QuizRepo
	mtuci  MtuciRepo
}

func New(logger *logger.Logger, j JWT, r QuizRepo, m MtuciRepo) *ServiceUseCase {
	return &ServiceUseCase{
		logger: logger,
		jwt:    j,
		repo:   r,
		mtuci:  m,
	}
}

func (s *ServiceUseCase) GetAllQuiz(ctx context.Context) ([]*entity.QuizUI, error) {
	var response []*entity.QuizUI

	quizzes, err := s.repo.GetAllQuiz(ctx)
	if err != nil {
		return nil, err
	}

	for _, quiz := range quizzes {
		user, err := s.mtuci.GetUserByID(ctx, quiz.AuthorID)
		if err != nil {
			return nil, err
		}
		quiz.Author = user

		var groups []*entity.Group
		quiz.AccessFor = strings.ReplaceAll(quiz.AccessFor, " ", "")
		if quiz.AccessFor != "*" {
			groupsIDs := strings.Split(quiz.AccessFor, ",")
			for _, v := range groupsIDs {
				groupID, err := strconv.Atoi(v)
				if err != nil {
					return nil, err
				}
				group, err := s.mtuci.GetGroupById(ctx, groupID)
				if err != nil {
					return nil, err
				}
				groups = append(groups, group)
			}
		} else {
			groups, err = s.mtuci.GetAllGroups(ctx)
			if err != nil {
				return nil, err
			}
		}
		response = append(response, &entity.QuizUI{
			ID:        quiz.ID,
			AuthorID:  quiz.AuthorID,
			Author:    quiz.Author,
			AccessFor: groups,
			QuizHash:  quiz.QuizHash,
			Title:     quiz.Title,
			Questions: quiz.Questions,
			Active:    quiz.Active,
			CreatedAt: quiz.CreatedAt,
			UpdatedAt: quiz.UpdatedAt,
		})
	}

	return response, nil
}

func (s *ServiceUseCase) GetQuizById(ctx context.Context, quizId int) (*entity.QuizUI, error) {
	quiz, err := s.repo.GetQuizById(ctx, quizId)
	if err != nil {
		return nil, err
	}
	author, err := s.mtuci.GetUserByID(ctx, quiz.AuthorID)
	if err != nil {
		return nil, err
	}
	quiz.Author = author
	var groups []*entity.Group
	quiz.AccessFor = strings.ReplaceAll(quiz.AccessFor, " ", "")
	if quiz.AccessFor != "*" {
		groupsIDs := strings.Split(quiz.AccessFor, ",")
		for _, v := range groupsIDs {
			groupID, err := strconv.Atoi(v)
			if err != nil {
				return nil, err
			}
			group, err := s.mtuci.GetGroupById(ctx, groupID)
			if err != nil {
				return nil, err
			}
			groups = append(groups, group)
		}
	} else {
		groups, err = s.mtuci.GetAllGroups(ctx)
		if err != nil {
			return nil, err
		}
	}

	return &entity.QuizUI{
		ID:        quiz.ID,
		AuthorID:  quiz.AuthorID,
		Author:    quiz.Author,
		AccessFor: groups,
		QuizHash:  quiz.QuizHash,
		Title:     quiz.Title,
		Questions: quiz.Questions,
		Active:    quiz.Active,
		CreatedAt: quiz.CreatedAt,
		UpdatedAt: quiz.UpdatedAt,
	}, nil
}

func (s *ServiceUseCase) SaveQuiz(ctx context.Context, quiz *entity.QuizUISaveRequest) (*entity.QuizUI, error) {
	quiz.AuthorID = s.jwt.Parse(ctx.Value("token").(map[string]interface{})).ID
	savedQuiz, err := s.repo.SaveQuiz(ctx, quiz)
	if err != nil {
		return nil, err
	}
	author, err := s.mtuci.GetUserByID(ctx, savedQuiz.AuthorID)
	if err != nil {
		return nil, err
	}
	savedQuiz.Author = author

	var groups []*entity.Group
	savedQuiz.AccessFor = strings.ReplaceAll(savedQuiz.AccessFor, " ", "")
	if savedQuiz.AccessFor != "*" {
		groupsIDs := strings.Split(savedQuiz.AccessFor, ",")
		for _, v := range groupsIDs {
			groupID, err := strconv.Atoi(v)
			if err != nil {
				return nil, err
			}
			group, err := s.mtuci.GetGroupById(ctx, groupID)
			if err != nil {
				return nil, err
			}
			groups = append(groups, group)
		}
	} else {
		groups, err = s.mtuci.GetAllGroups(ctx)
		if err != nil {
			return nil, err
		}
	}
	response := &entity.QuizUI{
		ID:        savedQuiz.ID,
		AuthorID:  savedQuiz.AuthorID,
		Author:    savedQuiz.Author,
		AccessFor: groups,
		QuizHash:  savedQuiz.QuizHash,
		Title:     savedQuiz.Title,
		Questions: savedQuiz.Questions,
		Active:    savedQuiz.Active,
		CreatedAt: savedQuiz.CreatedAt,
		UpdatedAt: savedQuiz.UpdatedAt,
	}
	return response, nil
}

func (s *ServiceUseCase) GetUserByLoginWithPassword(ctx context.Context, user entity.UserLogin) (*entity.User, error) {
	foundedUser, err := s.mtuci.GetUserByLoginWithPassword(ctx, user)
	if err != nil {
		return nil, err
	}
	token, err := s.jwt.Create(time.Hour*24, foundedUser)
	if err != nil {
		return nil, err
	}
	foundedUser.Token = token

	return foundedUser, nil
}

func (s *ServiceUseCase) DeleteQuiz(ctx context.Context, quizID int) error {
	return s.repo.DeleteQuiz(ctx, quizID)
}

func (s *ServiceUseCase) GetAllGroups(ctx context.Context) ([]*entity.Group, error) {
	return s.mtuci.GetAllGroups(ctx)
}
