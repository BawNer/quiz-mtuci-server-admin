package repo

import (
	"context"
	"fmt"
	"quiz-mtuci-server/internal/entity"
	"quiz-mtuci-server/pkg/logger"
	"quiz-mtuci-server/pkg/postgres"

	"github.com/google/uuid"
)

type QuizRepo struct {
	*postgres.Postgres
	l *logger.Logger
}

func New(pg *postgres.Postgres, l *logger.Logger) *QuizRepo {
	return &QuizRepo{pg, l}
}

func (r *QuizRepo) GetAllQuiz(ctx context.Context) ([]*entity.QuizUI, error) {
	var (
		response []*entity.QuizUI
		quizzes  []entity.Quiz
	)
	result := r.DB.Table("quizzes").Find(&quizzes)
	if result.Error != nil {
		return nil, fmt.Errorf("quiz repo err %v", result.Error)
	}

	for _, quiz := range quizzes {
		var (
			questionsUI []entity.QuestionsUI
			questions   []entity.Question
			answers     []entity.AnswersOption
		)
		if err := r.DB.Table("questions").Where("quiz_id = ?", quiz.ID).Find(&questions); err.Error != nil {
			return nil, err.Error
		}

		for _, question := range questions {

			if err := r.DB.Table("answers_options").Where("question_id = ?", question.ID).Find(&answers); err.Error != nil {
				return nil, err.Error
			}

			questionsUI = append(questionsUI, entity.QuestionsUI{
				ID:             question.ID,
				Label:          question.Label,
				Description:    question.Description,
				AnswersOptions: answers,
			})
		}
		response = append(response, &entity.QuizUI{
			ID:        quiz.ID,
			AuthorID:  quiz.AuthorID,
			Type:      quiz.Type,
			QuizHash:  quiz.QuizHash,
			Title:     quiz.Title,
			Questions: questionsUI,
			Active:    quiz.Active,
			CreatedAt: quiz.CreatedAt,
			UpdatedAt: quiz.UpdatedAt,
		})
	}

	return response, nil
}

func (r *QuizRepo) GetQuizById(ctx context.Context, quizId int) (*entity.QuizUI, error) {
	var (
		questionsUI []entity.QuestionsUI
		quiz        entity.Quiz
		questions   []entity.Question
	)
	result := r.DB.Table("quizzes").First(&quiz, quizId)
	if result.Error != nil {
		return nil, fmt.Errorf("quiz repo err %v", result.Error)
	}

	if err := r.DB.Table("questions").Where("quiz_id = ?", quiz.ID).Find(&questions); err.Error != nil {
		return nil, err.Error
	}

	for _, question := range questions {
		var answers []entity.AnswersOption

		if err := r.DB.Table("answers_options").Where("question_id = ?", question.ID).Find(&answers); err.Error != nil {
			return nil, err.Error
		}

		questionsUI = append(questionsUI, entity.QuestionsUI{
			ID:             question.ID,
			Label:          question.Label,
			Description:    question.Description,
			AnswersOptions: answers,
		})
	}

	response := &entity.QuizUI{
		ID:        quiz.ID,
		AuthorID:  quiz.AuthorID,
		Type:      quiz.Type,
		QuizHash:  quiz.QuizHash,
		Title:     quiz.Title,
		Questions: questionsUI,
		Active:    quiz.Active,
		CreatedAt: quiz.CreatedAt,
		UpdatedAt: quiz.UpdatedAt,
	}

	return response, nil
}

func (r *QuizRepo) GetQuizByHash(ctx context.Context, quizHash string) (*entity.QuizUI, error) {
	var (
		questionsUI []entity.QuestionsUI
		quiz        entity.Quiz
		questions   []entity.Question
	)
	result := r.DB.Table("quizzes").Where("quiz_hash = ?", quizHash).First(&quiz)
	if result.Error != nil {
		return nil, fmt.Errorf("quiz repo err %v", result.Error)
	}

	if err := r.DB.Table("questions").Where("quiz_id = ?", quiz.ID).Find(&questions); err.Error != nil {
		return nil, err.Error
	}

	for _, question := range questions {
		var answers []entity.AnswersOption

		if err := r.DB.Table("answers_options").Where("question_id = ?", question.ID).Find(&answers); err.Error != nil {
			return nil, err.Error
		}

		questionsUI = append(questionsUI, entity.QuestionsUI{
			ID:             question.ID,
			Label:          question.Label,
			Description:    question.Description,
			AnswersOptions: answers,
		})
	}

	response := &entity.QuizUI{
		ID:        quiz.ID,
		AuthorID:  quiz.AuthorID,
		Type:      quiz.Type,
		QuizHash:  quiz.QuizHash,
		Title:     quiz.Title,
		Questions: questionsUI,
		Active:    quiz.Active,
		CreatedAt: quiz.CreatedAt,
		UpdatedAt: quiz.UpdatedAt,
	}

	return response, nil
}

func (r *QuizRepo) SaveQuiz(ctx context.Context, quiz *entity.QuizUI) (*entity.QuizUI, error) {
	var (
		questions []entity.QuestionsUI
		answers   []entity.AnswersOption
	)

	newQuiz := entity.Quiz{
		AuthorID: quiz.AuthorID,
		Type:     quiz.Type,
		QuizHash: uuid.New().String(),
		Title:    quiz.Title,
		Active:   quiz.Active,
	}
	if createQuiz := r.DB.Table("quizzes").Create(&newQuiz); createQuiz.Error != nil {
		return nil, createQuiz.Error
	}

	// добавляем вопросы к квизу
	for _, question := range quiz.Questions {
		newQuestions := entity.Question{
			QuizID:      newQuiz.ID,
			Label:       question.Label,
			Description: question.Description,
		}
		if createQuestion := r.DB.Table("questions").Create(&newQuestions); createQuestion.Error != nil {
			return nil, createQuestion.Error
		}
		// добавляем варианты ответа
		for _, answer := range question.AnswersOptions {
			newAnswerOption := entity.AnswersOption{
				QuestionID:  newQuestions.ID,
				Label:       answer.Label,
				Description: answer.Description,
			}
			if createAnswerOption := r.DB.Table("answers_options").Create(&newAnswerOption); createAnswerOption.Error != nil {
				return nil, createAnswerOption.Error
			}

			answers = append(answers, entity.AnswersOption{
				ID:          newAnswerOption.ID,
				QuestionID:  newQuestions.ID,
				Label:       newAnswerOption.Label,
				Description: newAnswerOption.Description,
			})
		}

		questions = append(questions, entity.QuestionsUI{
			ID:             newQuestions.ID,
			Label:          newQuestions.Label,
			Description:    newQuestions.Description,
			AnswersOptions: answers,
		})
	}

	createdQuiz := &entity.QuizUI{
		ID:        newQuiz.ID,
		AuthorID:  newQuiz.AuthorID,
		Type:      newQuiz.Type,
		QuizHash:  newQuiz.QuizHash,
		Title:     newQuiz.Title,
		Questions: questions,
		Active:    newQuiz.Active,
		CreatedAt: newQuiz.CreatedAt,
		UpdatedAt: newQuiz.UpdatedAt,
	}

	return createdQuiz, nil
}

func (r *QuizRepo) SaveReviewers(ctx context.Context, reviewer *entity.Reviewers) error {
	err := r.DB.Table("reviewers").Save(&reviewer)
	if err.Error != nil {
		return err.Error
	}

	return nil
}