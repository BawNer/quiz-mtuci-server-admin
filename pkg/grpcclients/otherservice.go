package grpcclients

import (
	"context"
	"quiz-mtuci-server/pkg/logger"
	"quiz-mtuci-server/pkg/metrics"
	"quiz-mtuci-server/proto/otherService"
	"strconv"
	"time"

	"google.golang.org/grpc"
)

type OtherServiceClient struct {
	Client     otherService.OtherServiceClient
	connection *grpc.ClientConn
	ip         string
	port       int
	logger     *logger.Logger
}

func NewOtherServiceClient(ip string, port int, logger *logger.Logger) (*OtherServiceClient, error) {
	address := ip + ":" + strconv.Itoa(port)

	connection, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client := otherService.NewOtherServiceClient(connection)

	return &OtherServiceClient{
		Client:     client,
		connection: connection,
		ip:         ip,
		port:       port,
		logger:     logger,
	}, nil
}

func (p *OtherServiceClient) Process(name string) error {
	beginTime := time.Now()

	defer func() {
		metrics.SetRequestTime("ProcessGRPC", float64(time.Since(beginTime).Milliseconds()))
		p.logger.Info().Msgf("ProcessGRPC time %d", int(time.Since(beginTime).Milliseconds()))
	}()

	_, err := p.Client.Process(context.Background(), &otherService.ProcessRequest{
		Name: name,
	})

	if err != nil {
		return err
	}

	return nil
}
