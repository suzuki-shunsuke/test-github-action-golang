package hello

import "fmt"

func Hello() string {
	hello := make([]string, 1, 1)
	fmt.Println(hello)
	return "hello"
}
