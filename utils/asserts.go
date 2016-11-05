package utils

import "fmt"

func MustTrue(val bool) {
	if !val {
		fmt.Println("failed")
	} else {
		fmt.Println("passed")
	}
}


