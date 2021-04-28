package main

import (
	"fmt"

	"github.com/bikram/aspiration_mapper/mapper"
)

func NewSkipString(i int, s string) mapper.SkipString {
	return mapper.SkipString{
		Data:    []rune(s),
		Skipper: i,
	}
}

func main() {
	s := NewSkipString(2, "Aspiration.com")
	mapper.MapString(&s)
	fmt.Println(s)
}
