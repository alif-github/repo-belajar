package main

import (
	"fmt"
	"genericlib/hello"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println(hello.PanggilHello())
	var angka int64
	angka = 1
	jawaban := Generic[int64](angka)
	jawaban2 := hello.Generic[int64](angka)
	fmt.Println("Ini jawaban 1: ", jawaban)
	fmt.Println("Ini jawaban 2: ", jawaban2)
}

func Generic[T any](param T) T {
	fmt.Println("Ini di dalem generic nya: ", param)
	return param
}
