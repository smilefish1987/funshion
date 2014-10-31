# golang笔记 变量和常量


**变量和常量**

1.golang的源文件以 **.go** 为后缀
2.golang的语法类C,但是通常语句后面的**;**号不是必须的，只有在一行写多条语句的情况下，每条语句需要用**;**号分隔
3.golang通过关键字**var**声明变量，并且**变量的类型在变量名之后**，以空格隔开：

``` golang
    var b int
    b = 10
```
    
4.当定义一个变量，变量的默认值为其类型的null值，int为0，string为"",bool为false,float32为0.0
5.对于变量的声明和赋初始值，可以通过一种类型推断的简单方式实现：

``` golang
    b := 10
    //等同于
    var b int
    b = 10
```
    
6.在golang中声明的变量没有使用会报编译错误
7.平行声明和推断多个变量：

``` golang
    var a,b int
    c,d := 12,34
```
    
8.通过分组的方式声明多个变量：
``` golang
    var (
		m int
		n bool
	)

	m, n = 1, false
```

9.有一个特殊的变量 **_**，任何分配给它的值都会被忽略，主要用于函数或者遍历时忽略某个变量返回的值

``` golang
    _, ok := monthdays["Feb"]  //用于忽略Feb键对应的值
```

10.常量在golang中，通过关键字const声明并初始化。它们在编译时被创建，**只能是整型，布尔型和字符串类型**，也可以通过const分组的方式声明多个常量,也可以通过**iota**来定义常量

``` golang

    const a  = 45
    const a int = 45
    
    const (
		c3 int    = 12     //c3=12
		c4 string = "9"    //c4=9
		c5        = iota   //c5=2
		c6                 //c6=3
		c7 = iota          //c7=4
	)

	const (
		cc1 = iota   //cc1=0   可以把iota理解成组的索引？
		cc2          //cc2=1
		cc3 = 7      //cc3=7
		cc4 = iota   //cc4=3
		cc5          //cc5=4
	)
``` 