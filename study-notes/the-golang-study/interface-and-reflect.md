# golang 笔记 接口和反射

1.在Golang中,通过interface关键字来定义接口。接口是方法的集合。每个类型都有接口，意味着对那个类型定义了方法集合。

``` golang
package main

//定义了结构S
type S struct {
	num int
}

//给机构绑定了两个方法
func (p *S) Get() int {
	return p.num
}

func (p *S) Set(val int) {
	p.num = val
}

//定义了接口类型，仅仅是方法的集合，对于接口I，S是合法的实现，因为S定义了I所需要的两个方法。注意，即便没有明确定义S实现了I,这也是正确的(非侵入式接口)
type I interface {
	Get() int
	Set(int)
}

//接口值，下面这个函数参数p保存了接口类型的值
func funcI(p I) {
	println(p.Get())
	p.Set(2)
	println(p.Get())
}

func main() {
	s := S{10}
	funcI(&s) //S实现了I,可以调用f向其传递S类型的值的指针，获取s的地址，而不是S的值的原因，是因为在s的指针上定义了方法。可以让方法接受值类型，但是这样Set方法就不会期望的那样工作
}
```

2.无需明确一个类型是否实现了一个接口。这就意味着Golang实现了叫做duck typing的模式。但是这又不是纯粹的duck typing,因为如果可能的话Golang编译器将对类型是否实现接口进行静态检查。Golang也有纯粹动态的方面，可将一个接口类型转换成另一个。通常情况下，转换的检查是在运行时进行的，当在已有接口值中存储的类型值不匹配将要转换到的接口（非法转换时），程序会抛出运行时错误。

3.Golang中的联合了接口值，静态类型检查，运行时动态转换，以及无需明确定义类型适配一个接口，这些给Golang带来了强大，灵活，高效和容易编写的特定。

4.当有多个struct都实现了同一个interface,但是在实际中又需要知道实现interface的是哪个struct类型时，可以使用两种方式type switch和 comma-ok来判断一个接口类型是否实现了某个特定的接口：
``` golang
package main

import (
	"fmt"
)

func main() {
	s := S{10}
	r := R{30}
	checkTypeSwitch(&s)
	checkTypeSwitch(&r)
	checkTypecommaOk(&s)
	checkTypecommaOk(&r)
}

type S struct {
	num int
}

func (p *S) Get() int {
	return p.num
}

func (p *S) Set(val int) {
	p.num = val
}

type R struct {
	i int
}

func (r *R) Get() int {
	return r.i
}

func (r *R) Set(val int) {
	r.i = val
}

type I interface {
	Get() int
	Set(int)
}

//通过comma-ok的方式判断
//Go语言里面有一个语法，可以直接判断是否是该类型的变量： value, ok = element.(T)，这里value就是变量的值，ok是一个bool类型，element是interface变量，T是断言的类型。如果element里面确实存储了T类型的数值，那么ok返回true，否则返回false

func checkTypecommaOk(p I) {
	if t, ok := p.(*R); ok {
		println("*R yes", t)
	} else {
		println("*R no", t)
	}

}

//通过switch-type的方式判断
// interfacevalue.(type)语法不能在switch外的任何逻辑里面使用，如果你要在switch外面判断一个类型就使用comma-ok
func checkTypeSwitch(p I) {

	switch t := p.(type) {
	case *S:
		fmt.Printf("%v\t %v\n", t, "*S")
	case *R:
		fmt.Printf("%v\t %v\n", t, "*R")
	default:
		println("type unknown", t)
	}
}
```

5.空接口(interface{})不包含任何的method，正因为如此，所有的类型都实现了空interface。空interface对于描述起不到任何的作用(因为它不包含任何的method），但是空interface在我们需要存储任意类型的数值的时候相当有用，因为它可以存储任意类型的数值。它有点类似于C语言的void*类型。
``` golang
package main
func main() {
	s := S{10}
	println(efunc(&s))  //10
	println(efunc(s))   //会报错，因为s没有Get方法，*s有
	i := 5
	println(efunc(i))  //会报错，因为int没有Get()方法
}

type S struct {
	num int
}

func (p *S) Get() int {
	return p.num
}

func (p *S) Set(val int) {
	p.num = val
}

type I interface {
	Get() int
	Set(int)
}

//接收任何类型值的函数
func efunc(val interface{}) int {
	return val.(I).Get()
}
```

6.方法就是有接收者的函数。可以在任意类型上定义方法(除了非本地类型[定义在其他包中的类型]，包括内建类型：int类型不能有方法)

``` golang
    type Foo int
    func (self Foo) Emit(){
        fmt.Printf("%v\n",self)
    }
    
    //在内建类型int上不能定义新的方法
    func (i int) Emit(){
        //do-something
    }
    
    //在非本地类型 net.AddrError上不能定义新的方法
    func (a *net.AddrError) Emit(){
        //do-something
    }
```

7.接口定义了一个方法的集合。方法包含实际的代码。一个接口就定义，而方法就是实现。因此方法的接收者不能定义为接口类型，这样子会引起invalid receiver type..的编译器错误。【接收者类型必须是T或者*T，这里的T是类型名。T叫做接收者基础类型或简称基础类型。基础类型一定不能是指针或接口类型，并且定义在与方法相同的包中】

8.在Golang中创建指向接口的指针是无意义的，实际上创建接口值的指针也是非法的【语言的改变是使用指针指向接口值而不再自动反引用指针。指向接口值的指针通常是低级的错误，而不是正确的代码】

9.根据规则单方法接口命名为方法名加上-er后缀：Reader,Writer,Formatter等。

10.接口嵌入接口
```golang
type Interface interface {
    sort.Interface //嵌入字段sort.Interface
    Push(x interface{}) //a Push method to push elements into the heap
    Pop() interface{} //a Pop elements that pops elements from the heap
}
//如果一个interface1作为interface2的一个嵌入字段，那么interface2隐式的包含了interface1里面的method。

// io.ReadWriter接口就包含了Reader和Writer两个接口
// io.ReadWriter
type ReadWriter interface {
    Reader
    Writer
}
```

11.Go语言实现了反射，所谓反射就是能检查程序在运行时的状态。我们一般用到的包是reflect包。
``` golang
package main

import (
	"fmt"
	"reflect"
)

type T struct {
	A int
	B string
}

func main() {
	var x float64 = 3.4
	fmt.Println("Type: ", reflect.TypeOf(x)) //Type:float64

	v := reflect.ValueOf(x)
	fmt.Println("type: ", v.Type()) //type :float64
	fmt.Println("kind is float64: ", v.Kind() == reflect.Float64) // kind is float64 : true
	fmt.Println("value: ", v.Float())  //value: 3.4

	p := reflect.ValueOf(&x)
	fmt.Println("type of p: ", p.Type()) // type of p: *float64
	fmt.Println("settability of p :", p.CanSet()) //settablility of p : false

	v = p.Elem()
	fmt.Println("settability of v: ", v.CanSet())
 // settablility of v : true
	v.SetFloat(7.2)
	fmt.Println(v.Interface())  //7.2
	fmt.Println(x)  //7.2

	t := T{345, "smile123"}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()

	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface()) 
		//0: A int = 345
		//1: B string = smile123
	}
}
```

12.使用reflect一般分成三步，要去反射是一个类型的值(这些值都实现了空interface)

>12.1 需要把它转化成reflect对象(reflect.Type或者reflect.Value，根据不同的情况调用不同的函数)

>``` golang
t := reflect.TypeOf(i)    //得到类型的元数据,通过t我们能获取类型定义里面的所有元素
v := reflect.ValueOf(i)   //得到实际的值，通过v我们获取存储在里面的值，还可以去改变值
>```

>12.2 转化为reflect对象之后我们就可以进行一些操作了

>>a.就是将reflect对象转化成相应的值
>``` golang
tag := t.Elem().Field(0).Tag  //获取定义在struct里面的标签
name := v.Elem().Field(0).String()  //获取存储在第一个字段里面的值
>```
>>b.获取反射值相应的类型和数值
>``` golang
var x float64 = 3.4
v := reflect.ValueOf(x)
fmt.Println("type:", v.Type())
fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
fmt.Println("value:", v.Float())
>```

>12.3 反射的话，那么反射的字段必须是可修改的，我们前面学习过传值和传引用，这个里面也是一样的道理。反射的字段必须是可读写的意思是,如果下面这样写，那么会发生错误
>``` golang
var x float64 = 3.4
v := reflect.ValueOf(x)
v.SetFloat(7.1)
>```
>如果要修改相应的值，必须这样写
>``` golang
var x float64 = 3.4
p := reflect.ValueOf(&x)
v := p.Elem()
v.SetFloat(7.1)
>```




