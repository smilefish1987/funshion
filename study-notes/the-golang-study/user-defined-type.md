# golang 笔记  自定义类型


1.Golang允许定义新的类型，通过关键字type实现：
``` golang
    type foo int
    //创建了一个新的类型 foo 作用和int一样
```

2.通过struct关键字创建更加复杂的类型
``` golang
package main

import (
	"fmt"
)

func main() {

	p1 := new(Person)
	p1.name = "smile"
	p1.age = 23

	fmt.Printf("%v\n", p1) // &{smile 23}

	p2 := Person{"fish", 25}
	fmt.Printf("%v\n", p2) // {fish 25}

	p3 := Person{age: 24, name: "benben"}
	fmt.Printf("%v\n", p3)   // {benben 24}
	println(p3.age, p3.name) //24 benben

	p4 := Null{}   //没有字段的struct
	fmt.Printf("%v\n", p4)

	p5 := NoName{"hello", 45}  //匿名字段的struct
	fmt.Printf("%v\n", p5)
}

type Null struct {
}

type NoName struct {
	string
	int
}

type Person struct {
	name string
	age  int
}
```

3.在struct中首字母大写的字段可以被导出，也就是在其他的包中可以进行读写；字段名以小写字母开头是当前包私有的，其他的包中不能进行读写。

4.给新创建的类型，添加方法或者通过将新创建的类型作为参数传递给函数，以便操作和处理：
``` golang
package main

func main() {

	p := Person{"fish", 13}
	p1 := &Person{"smile", 15}
	sayHello(&p)   // hello fish
	(*p1).sayHello()  // hello smile
	
	p1.sayHello()    // hello smile
	// 如果x可以获取地址，并且&x的方法包含了m，x.m()是(&x).m()更短的写法
	//这里Golang编译器会查找Person类型的变量p的,如果没有找到就会查找*Person类型的方法列表，并且将其转化为(&p).sayHello()
	p.sayHello()  // hello fish
}

type Person struct {
	name string
	age  int
}

//通过函数操作新定义的类型，把类型作为参数
func sayHello(p *Person) {
	println("hello " + p.name)
}

//通过方法操作新定义的类型，将方法绑定到类型
func (p *Person) sayHello() {
	println("hello " + (*p).name)
}
```

5.到底是使用函数还是方法完全由程序员说了算，但是如果需要满足接口就必须使用方法。

6.通过两种不同的方式基于现有的struct创建新的类型：不带struct创建的新类型Tperson没有任何Person的方法；而带struct创建的新类型Tperson却从Person集成了所有的方法，被绑定到其匿名字段Person
``` golang
package main

import (
	"fmt"
)

func main() {
	p := Person{"fish", 24}
	fmt.Printf("%v\n", p)  //{fish 24}
	p.growth()     // growth
	fmt.Printf("%v\n", p.age) // 25
	p.sayHello()       //hello fish

	pt := Tperson{"smile", 23}
	fmt.Printf("%v\n", pt)    //{smile 23}
	//pt.sayHello() //报错 pt没有这个方法
	//pt.growth()   //报错 pt没有这个方法

	ps := Sperson{Person{"benben", 25}} 
	fmt.Printf("%v\n", ps)  //{{benben 25}}
	ps.sayHello()   // hello benben
	ps.growth()    //growth
}

type Tperson Person //定义新类型

type Sperson struct { //定义新类型
	Person
}

type Person struct {
	name string
	age  int
}

func (p *Person) sayHello() {
	println("hello " + p.name)
}

func (p *Person) growth() {
	p.age++
	println("growth")
}
```

7.Golang没有传统意义上的继承，但是有时候也需要从已经实现的类型中“继承”并修改一些方法。在Golang中我们可以通过嵌入一个类型这种组合的方式来实现。





