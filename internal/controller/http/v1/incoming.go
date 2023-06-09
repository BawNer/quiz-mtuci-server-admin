package v1

import (
	"fmt"
	"net/http"
	"quiz-mtuci-server/internal/entity"
	"strconv"

	"github.com/gin-gonic/gin"

	"quiz-mtuci-server/internal/usecase"
	"quiz-mtuci-server/pkg/logger"
)

type serviceRoutes struct {
	t usecase.UseCase
	l logger.Interface
	m *usecase.MiddlewareStruct
}

func newQuizRoutes(handler *gin.RouterGroup, t usecase.UseCase, l logger.Interface, m *usecase.MiddlewareStruct) {
	r := &serviceRoutes{t, l, m}

	h := handler.Group("/quiz")
	h.Use(m.AuthGuard())
	h.Use(m.IsAdmin())
	{
		h.GET("/", r.GetAllQuiz)
		h.GET("/:id", r.GetQuizById)
		h.POST("/", r.SaveQuiz)
		h.DELETE("/:id", r.DeleteQuiz)
	}
}

func newUserRoutes(handler *gin.RouterGroup, t usecase.UseCase, l logger.Interface, m *usecase.MiddlewareStruct) {
	r := &serviceRoutes{t, l, m}

	h := handler.Group("/users")
	{
		h.POST("/login", r.GetUserByLoginWithPassword)
		h.GET("/groups", r.GetAllGroups)
	}
}

func (s *serviceRoutes) GetAllQuiz(c *gin.Context) {
	quizzes, err := s.t.GetAllQuiz(c)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if quizzes == nil {
		quizzes = make([]*entity.QuizUI, 0)
	}

	c.JSON(http.StatusOK, entity.QuizzesResponseUI{
		Success:     true,
		Description: "",
		Quizzes:     quizzes,
	})
}

func (s *serviceRoutes) GetQuizById(c *gin.Context) {
	quizID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, "problem when get id params")
		return
	}

	quiz, err := s.t.GetQuizById(c, quizID)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, entity.QuizResponseUI{
		Success:     true,
		Description: "",
		Quiz:        quiz,
	})
}

func (s *serviceRoutes) SaveQuiz(c *gin.Context) {
	var request entity.QuizUISaveRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		errorResponse(c, http.StatusInternalServerError, fmt.Sprintf("Error parse request json, %s", err))
		return
	}

	quiz, err := s.t.SaveQuiz(c, &request)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, entity.QuizResponseUI{
		Success:     true,
		Description: "Опрос создан!",
		Quiz:        quiz,
	})
}

func (s *serviceRoutes) GetUserByLoginWithPassword(c *gin.Context) {
	var userLogin entity.UserLogin
	if err := c.ShouldBindJSON(&userLogin); err != nil {
		errorResponse(c, http.StatusInternalServerError, fmt.Sprintf("Error parse body, %v", err))
	}

	user, err := s.t.GetUserByLoginWithPassword(c, userLogin)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Header("Authorization", fmt.Sprintf("Bearer %s", user.Token))
	c.JSON(http.StatusOK, entity.ResponseUserLogin{
		Success:     true,
		Description: "Login success",
		User:        user,
	})
}

func (s *serviceRoutes) DeleteQuiz(c *gin.Context) {
	quizID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, "problem when get id params")
		return
	}

	if err := s.t.DeleteQuiz(c, quizID); err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, entity.PositiveResponseUI{
		Success:     true,
		Description: "Quiz has been removed!",
	})
}

func (s *serviceRoutes) GetAllGroups(c *gin.Context) {
	groups, err := s.t.GetAllGroups(c)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, entity.GroupsResponse{
		Success:     true,
		Description: "All groups found",
		Groups:      groups,
	})
}
