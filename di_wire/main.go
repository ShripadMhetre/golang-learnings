package di_wire

import "fmt"

func Main() {
	service := CreateConcatService()

	result := service.GetAll(
		"http://example.com",
		"https://drewolson.org",
	)

	fmt.Println(result)
}
