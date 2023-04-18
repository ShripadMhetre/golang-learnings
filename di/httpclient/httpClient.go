package httpclient

import (
	"fmt"
	"os"
)

type Logger interface {
	Log(message string) error
}

type FileLogger struct {
	File *os.File
}

func (l *FileLogger) Log(message string) error {
	_, err := l.File.Write([]byte(message))
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

func NewHttpClient(logger Logger) *HttpClient {
	return &HttpClient{logger}
}
