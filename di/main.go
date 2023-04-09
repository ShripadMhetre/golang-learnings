package di

import (
	"bytes"
	"fmt"
)

type Logger struct{}

func (logger *Logger) Log(message string) {
	fmt.Println(message)
}

type HttpClient struct {
	logger *Logger
}

func (client *HttpClient) Get(url string) string {
	client.logger.Log("Getting " + url)

	// make an HTTP request
	return "Response from - " + url
}

func NewHttpClient(logger *Logger) *HttpClient {
	return &HttpClient{logger}
}

type ConcatService struct {
	logger *Logger
	client *HttpClient
}

func (service *ConcatService) GetAll(urls ...string) string {
	service.logger.Log("Running GetAll :-")

	var result bytes.Buffer

	for _, url := range urls {
		result.WriteString(service.client.Get(url) + " ")
	}

	return result.String()
}

func NewConcatService(logger *Logger, client *HttpClient) *ConcatService {
	return &ConcatService{logger, client}
}

func Main() {
	logger := &Logger{}
	client := NewHttpClient(logger)
	service := NewConcatService(logger, client)

	// service := CreateConcatService()

	result := service.GetAll(
		"http://example.com",
		"https://lorem-ipsum.in",
	)

	fmt.Println("Concatenated response: ", result)
}
