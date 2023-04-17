package singleton

import (
	"fmt"
)

func Main() {

	for i := 0; i < 5; i++ {
		go getInstance()
	}

	fmt.Scanln()
}
