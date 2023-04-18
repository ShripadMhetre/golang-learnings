package service

import (
	"bytes"
	"fmt"

	"github.com/shripadmhetre/golang-learnings/di/httpclient"
)

type ConcatService struct {
	logger httpclient.Logger
	client *httpclient.HttpClient
}

func (service *ConcatService) GetAll(urls ...string) string {
	err := service.logger.Log("Running GetAll:\n")

	if err != nil {
		fmt.Println(err)
	}

	var result bytes.Buffer

	for _, url := range urls {
		result.WriteString(service.client.Get(url) + ", ")
	}

	resultStr := result.String()

	err = service.logger.Log(resultStr)
	if err != nil {
		fmt.Println(err)
	}

	return resultStr
}

func NewConcatService(logger httpclient.Logger, client *httpclient.HttpClient) *ConcatService {
	return &ConcatService{logger, client}
}
