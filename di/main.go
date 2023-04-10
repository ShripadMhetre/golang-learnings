package di

import (
	"fmt"
	"github.com/shripadmhetre/golang-learnings/di/client"
	"github.com/shripadmhetre/golang-learnings/di/service"
	"os"
)

func Main() {
	fl, err := os.OpenFile("di/logs/data.txt", os.O_RDWR, 0755)

	if err != nil {
		panic(err)
	}
	defer fl.Close()

	logger := &client.FileLogger{File: fl}
	httpClient := client.NewHttpClient(logger)
	service := service.NewConcatService(logger, httpClient)

	// service := CreateConcatService()

	result := service.GetAll(
		"http://example.com",
		"https://lorem-ipsum.in",
	)

	fmt.Println("Concatenated response: ", result)
}
