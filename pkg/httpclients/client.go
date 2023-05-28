package httpclients

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"quiz-mtuci-server/config"
	"quiz-mtuci-server/pkg/logger"
	"quiz-mtuci-server/pkg/metrics"
	"strings"
	"time"
)

type ExampleClientSettingsResp struct {
	Priority int `json:"priority"`
}

type ExampleClientSettings struct {
	httpClient *http.Client
	config     *config.ExampleClientSettings
	logger     *logger.Logger
}

func NewExampleClientSettings(config *config.ExampleClientSettings, logger *logger.Logger) (*ExampleClientSettings, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	httpClient := &http.Client{Transport: tr}

	return &ExampleClientSettings{
		httpClient: httpClient,
		config:     config,
		logger:     logger,
	}, nil
}

func (p *ExampleClientSettings) GetExampleClientSettings() (*ExampleClientSettingsResp, error) {
	p.logger.Info().Msg("call GetExampleClientSettings")

	beginTime := time.Now()

	defer func() {
		metrics.SetRequestTime("GetExampleClientSettings", float64(time.Since(beginTime).Milliseconds()))
		p.logger.Info().Msgf("GetExampleClientSettings time %d", int(time.Since(beginTime).Milliseconds()))
	}()

	url := p.config.Url

	if !strings.HasSuffix(url, "/") {
		url += "/"
	}

	req, err := http.NewRequestWithContext(context.TODO(), "GET", url, http.NoBody)

	if err != nil {
		return nil, err
	}

	resp, err := p.httpClient.Do(req)

	if err != nil {
		p.logger.Error().Err(err)

		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		p.logger.Error().Err(err)

		return nil, err
	}

	response := ExampleClientSettingsResp{}

	err = json.Unmarshal(body, &response)

	if err != nil {
		p.logger.Error().Err(err)

		return nil, err
	}

	p.logger.Info().Msg(fmt.Sprintf("got response ExampleClientSettings: %v", response))

	return &response, nil
}
