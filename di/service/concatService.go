package service

import (
	"bytes"
	"fmt"

	"github.com/shripadmhetre/golang-learnings/di/client"
)

type ConcatService struct {
	logger client.Logger
	client *client.HttpClient
}

func (service *ConcatService) GetAll(urls ...string) string {
	err := service.logger.Log("Running GetAll:\n")

	if err != nil {
		fmt.Println(err)
	}

	var result bytes.Buffer

	for _, url := range urls {
		result.WriteString(service.client.Get(url) + " ")
	}

	resultStr := result.String()

	err = service.logger.Log(resultStr)
	if err != nil {
		fmt.Println(err)
	}

	return resultStr
}

func NewConcatService(logger client.Logger, client *client.HttpClient) *ConcatService {
	return &ConcatService{logger, client}
}
