package di

import (
	"fmt"
	"github.com/shripadmhetre/golang-learnings/di/httpclient"
	"github.com/shripadmhetre/golang-learnings/di/service"
	"os"
)

func Main() {
	fl, err := os.OpenFile("di/logs/data.txt", os.O_RDWR, 0755)

	if err != nil {
		panic(err)
	}
	defer fl.Close()

	logger := &httpclient.FileLogger{File: fl}
	logger2 := &httpclient.StdoutLogger{}
	httpClient := httpclient.NewHttpClient(logger2)
	service := service.NewConcatService(logger, httpClient)

	result := service.GetAll(
		"http://example.com",
		"https://lorem-ipsum.in",
	)

	fmt.Println("Concatenated response: ", result)
}
