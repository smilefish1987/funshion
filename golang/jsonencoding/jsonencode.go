package main

import (
	"encoding/json"
	"fmt"
)

type Book struct {
	Title       string
	Authors     []string
	Publisher   string
	Price       float32
	IsPublished bool
}

func main() {
	authros := []string{"Smilefish", "Benben", "hello", "world", "mm", "gg"}

	gobook := Book{
		"golang programming",
		authros,
		"china",
		9.98,
		true,
	}

	//调用encoding/json包中的Marshal对数据进行jsonencode
	bstr, err := json.Marshal(gobook)

	if err != nil {
		fmt.Println("jsonencode error")
	}

	fmt.Println(string(bstr))
}
