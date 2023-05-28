package grpc

import (
	"context"
	"fmt"
	"net"
	"quiz-mtuci-server/internal/usecase"
	"quiz-mtuci-server/pkg/logger"
	"quiz-mtuci-server/pkg/metrics"
	"quiz-mtuci-server/proto/example"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

type Server struct {
	example.UnimplementedExampleServer
	logger         *logger.Logger
	exampleUseCase *usecase.ServiceUseCase
}

func NewServer(logger *logger.Logger, exampleUseCase *usecase.ServiceUseCase) *Server {
	return &Server{
		logger:         logger,
		exampleUseCase: exampleUseCase,
	}
}

func (s *Server) Start(port string) error {
	addr := fmt.Sprintf(":%s", port)

	listener, err := net.Listen("tcp", addr)

	if err != nil {
		s.logger.Error().Err(err)
		return err
	}

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)

	example.RegisterExampleServer(grpcServer, s)

	grpc_health_v1.RegisterHealthServer(grpcServer, health.NewServer())

	s.logger.Info().Msgf("Start serve GRPC at :%s", port)

	go grpcServer.Serve(listener)

	return nil
}

func (s *Server) GetTasks(c context.Context, req *example.Request) (*example.Response, error) {

	beginTime := time.Now()

	defer func() {
		metrics.SetRequestTime("GetTasksGRPC", float64(time.Since(beginTime).Milliseconds()))
		s.logger.Info().Msgf("GetTasksGRPC time %d", int(time.Since(beginTime).Milliseconds()))
	}()

	s.logger.Info().Msgf("got grpc query GetTasks")

	tasks, err := s.exampleUseCase.GetAllTasks(c)

	if err != nil {
		return nil, err
	}

	response := example.Response{Tasks: []*example.Task{}}

	for _, task := range tasks {
		response.Tasks = append(response.Tasks, &example.Task{Name: task.Name})
	}

	return &response, nil
}
