package main

import (
	"fmt"
	"time"
)

func ChangeValue(s *string) {
	*s = "asd"
}
func ChangeValue2(s string) {
	s = "changed!"
}

func main() {
	var str string = "hello world"
	ChangeValue2(str)
	ChangeValue(&str)
	fmt.Println(str)
	fmt.Println(time.Now())
}
