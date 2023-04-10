package beforedi

import (
	"bytes"
	"fmt"
	"os"
)

type Logger interface {
	Log(message string) error
}

type FileLogger struct {
	file *os.File
}

func (l *FileLogger) Log(message string) error {
	_, err := l.file.Write([]byte(message))
	return err
}

type StdoutLogger struct{}

func (l *StdoutLogger) Log(message string) error {
	_, err := fmt.Println(message)
	return err
}

type HttpClient struct {
	logger Logger
}

func (client *HttpClient) Get(url string) string {
	err := client.logger.Log("Getting " + url)

	if err != nil {
		fmt.Println(err)
	}

	// make an HTTP request
	return "Response from - " + url
}

func NewHttpClient() *HttpClient {
	logger := &StdoutLogger{}
	return &HttpClient{logger}
}

type ConcatService struct {
	logger Logger
	client *HttpClient
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

	return result.String()
}

func NewConcatService() *ConcatService {
	fl, err := os.OpenFile("beforedi/data.txt", os.O_RDWR, 0755)

	if err != nil {
		panic(err)
	}

	//defer fl.Close()

	logger := &FileLogger{file: fl}
	//logger := &StdoutLogger{}
	client := NewHttpClient()

	return &ConcatService{logger, client}
}

func Main() {
	service := NewConcatService()

	result := service.GetAll(
		"http://example.com",
		"https://lorem-ipsum.in",
	)

	service.logger.Log(result)
}
