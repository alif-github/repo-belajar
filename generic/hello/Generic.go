package hello

import (
	"fmt"
)

func Generic[T any](param T) T {
	fmt.Println(param)
	return param
}