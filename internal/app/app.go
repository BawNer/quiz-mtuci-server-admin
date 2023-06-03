package app

import (
	"fmt"
	"os"
	"os/signal"
	"quiz-mtuci-server/pkg/mysql"
	"syscall"

	"github.com/gin-gonic/gin"

	"quiz-mtuci-server/config"
	v1 "quiz-mtuci-server/internal/controller/http/v1"
	"quiz-mtuci-server/internal/usecase"
	"quiz-mtuci-server/internal/usecase/repo"
	"quiz-mtuci-server/pkg/httpserver"
	"quiz-mtuci-server/pkg/logger"
	"quiz-mtuci-server/pkg/postgres"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	pg, err := postgres.New(cfg.Postgres)

	if err != nil {
		l.Fatal().Err(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()

	msq, err := mysql.New(cfg.MySQL)

	if err != nil {
		l.Fatal().Err(fmt.Errorf("app - Run - mysql.New: %w", err))
	}
	defer msq.Close()

	jwtManager := usecase.NewJWT([]byte(cfg.JWT.Secret))

	quizUseCase := usecase.New(
		l,
		jwtManager,
		repo.New(pg, l),
		repo.NewMtuciRepo(msq, l),
	)

	handler := gin.New()
	middlewares := usecase.NewMiddleware(jwtManager, cfg)
	v1.NewRouter(handler, l, quizUseCase, middlewares)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info().Msgf("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error().Err(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	err = httpServer.Shutdown()
	if err != nil {
		l.Error().Err(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
