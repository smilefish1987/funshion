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

	/*
	 *  jsonencode之后的数据类型的映射关系：
	 *  布尔值转换为JSON后还是布尔类型
	 *  浮点数和整数会被转化为JSON里边的常规数字
	 *  字符串将以UTF-8编码转换输出为Unicode字符集的字符串，特殊字符将会被转义，比如: <:\u003c
	 *  数组和切片会转换为JSON里边的数组，但[]byte类型的值将会被转化为Base64编码后的字符串，
	 *  slice类型的零值会被转化为null
	 *  结构体会转化为JSON对象，并且只有结构体里边以大写字母开头的可被导出的字段才会被转化输出，
	 *  而这些可导出的字段会作为JSON对象的字符串索引
	 *  转化一个map类型的数据结构时，该数据的类型必须是map[string] T (T可以是encoding/json包支持的任意数据类型)
	 */

	fmt.Println(string(bstr))

	// the output is : {"Title":"golang programming","Authors":["Smilefish","Benben","hello","world","mm","gg"],"Publisher":"china","Price":9.98,"IsPublished":true}
}
