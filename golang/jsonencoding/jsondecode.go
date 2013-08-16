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

	var book Book

	//调用encoding/json包中的Unmarshal对数据进行jsondecode
	errdecode := json.Unmarshal(bstr, &book)

	/**
	 * 假设一个JSON对象有一个名为"Foo"的索引，要将"Foo"所对应的值填充到目标结构体的目标字段上，
	 * json.Unmarshal()将会遵循如下顺序进行查找匹配：
	 * (1) 一个包含Foo标签的字段
	 * (2) 一个名为Foo的字段
	 * (3) 一个名为Foo或者FOO或者除了首字母其他字母不区分大小写的名为Foo的字段
	 * 这些字段在类型声明中必须都是以大写字母开头，可被导出的字段
	 * 如果JSON中的字段在Go目标类型中不存在，json.Unmarshal()函数在解析过程中会丢弃该字段
	 */

	if errdecode != nil {
		fmt.Println("jsondecode error")
	}

	/**
	 * JSON中的布尔值将会转换为Go中的bool类型
	 * 数值会被转换为Go中的float64类型
	 * 字符串转换后还是string类型
	 * JSON数组会转换为[]interface{}类型
	 * JSON对象会转换为map[string] interface{}类型
	 * null值会转换为nil
	 * 在Go的标准库encoding/json包中，允许使用map[string] interface{}和 []interface{}类型的值
	 * 来分别存放未知结构的JSON对象或数组
	 */

	fmt.Println(book.Title)
	fmt.Println(book.Authors)
	fmt.Println(book.Publisher)
	fmt.Println(book.Price)
	fmt.Println(book.IsPublished)

	/*
		for k, v := range book {
			fmt.Println("the property %s of book is  %s ", k, v)
		}*/

}
