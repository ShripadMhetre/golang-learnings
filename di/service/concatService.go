package service

import (
	"bytes"

	"github.com/shripadmhetre/golang-learnings/di/client"
)

type ConcatService struct {
	logger *client.Logger
	client *client.HttpClient
}

func (service *ConcatService) GetAll(urls ...string) string {
	service.logger.Log("Running GetAll :-")

	var result bytes.Buffer

	for _, url := range urls {
		result.WriteString(service.client.Get(url) + " ")
	}

	return result.String()
}

func NewConcatService(logger *client.Logger, client *client.HttpClient) *ConcatService {
	return &ConcatService{logger, client}
}
