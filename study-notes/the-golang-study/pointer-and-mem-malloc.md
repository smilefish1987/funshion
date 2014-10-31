# golang笔记 指针 和 内存分配


**指针**

1.Golang由指针，但是没有指针运算
2.指针的出现，让调用函数的参数的值传递效率的提高和修改参数的值在函数外体现提供了可能。
3.指针的定义和使用：

``` golang
package main

import (
	"fmt"
)

func main() {
	
	var p *int   
	fmt.Printf("%v\n", p) // nil

	var i int = 10
	p = &i
	fmt.Printf("%v\n", p) // 0xc208000158
	fmt.Printf("%v\n", *p) // 10

	*p = 100
	fmt.Printf("%v\n", *p)  //100
	fmt.Printf("%v\n", i)   //100

}
```
4.一个新定义的或者没有任何指向的指针，其值为 nil
5.没有指针运算，所有*p++,它表示(*p)++,先获取值，然后对值进行加1

**内存分配**

1.Golang提供了两个内存分配的原语，new和make,他们用于不同的类型的内存分配。Golang是有垃圾回收机制的(GC)，所有无需担心内存的分配和回收。

2.用new分配内存
``` golang
func new(Type) *Type

package main

func main() {

	p := new(int)
	println(*p)
	*p = 10
	println(*p)
}
```
new(T)分配了零值填充的T类型的内存空间，并且返回其地址，一个*T类型的值。用Golang的术语说，它返回一个指针，指向新分配的类型T的零值。也就是所new分配的数据结构的实例是已经初始化过的，可以直接使用。

3.使用make分配内存
``` golang 
func make(Type, size IntegerType) Type
/*The make built-in function allocates and initializes an object of type slice, map, or chan (only). Like new, the first argument is a type, not a value. Unlike new, make's return type is the same as the type of its argument, not a pointer to it. The specification of the result depends on the type: Slice: The size specifies the length. The capacity of the slice is
equal to its length. A second integer argument may be provided to
specify a different capacity; it must be no smaller than the
length, so make([]int, 0, 10) allocates a slice of length 0 and
capacity 10.
Map: An initial allocation is made according to the size but the
resulting map has length 0. The size may be omitted, in which case
a small starting size is allocated.
Channel: The channel's buffer is initialized with the specified
buffer capacity. If zero, or the size is omitted, the channel is unbuffered.*/
```
3.1 make只能创建slice,map和channel,并且返回一个有初始值（非零）的T类型，而不是*T 。本质来讲，导致这三个类型由所不同的原有是指向数据结构的引用在使用前必须被初始化。

``` golang
    var p *[]int = new([]int)  //new分配slice结构内存，不推荐使用
    var v []int = make([]int,100)  //v指向一个新分配的有100个整数的数组
    
    v := make([]int,100)
    mp := make(map[string]int, 10)
	mp["mm"] = 12
	println(mp["mm"])
	// make仅适用于map,slice和channel，并且返回的不是指针，应该使用new获取特定的指针
```
4.new(T) 返回 *T 指向一个零值T; make(T) 返回初始化后的T
5.有些时候零值不能满足需求，必须有一个用于初始化的构造函数,当然也可以使用复合声明去掉初始化代码
``` golang
func NewFile(fd int,name string) *File{
    if fd < 0{
        return nil
    }
    f := new(File)
    f.fd = fd
    f.name = name
    f.dirinfo = nil
    f.nepipe = 0
    return f
}

//使用复合声明
func NewFile(fd int,name string) *File{
    if fd<0 {
        return nil
    }
    
    f := File{fd,name,nil,0} //复合声明一个对象
    return &f
}

//也可以更简洁
func NewFile(fd int,name string) *File{
    if fd<0 {
        return nil
    }
    
    return &File{fd,name,nil,0} //复合声明一个对象
}
```
6.复合声明所有的项目(字段)都必须按顺序全部写上，然而，通过对元素用字段和值成对的标识，初始化内容可以按任意顺序出现，并且可以省略初始化为零值的字段.
``` golang
    return &File{fd,name,nil,0} //所有项目按顺序全部写上
    return &File{fd:fd,name:name} //可以任意顺序，并且省略初始化为零值的字段
```

7.在特定的情况下，如果复合声明不包含任何字段，它创建特定类型的零值。也就是下面的表达式是等价的
``` golang
    new(File) 等价于 &File{}
```

8.复合声明同样可以用于创建array,slice和map，通过制定适当的索引和map键值来标识字段。
``` golang
    const Enone, Einval = 0, 1
	ar := [...]string{Enone: "no error", Einval: "invalid argument"}
	sl := []string{Enone: "no error", Einval: "invalid argument"}
	ma := map[int]string{Enone: "no error", Einval: "invalid argument"}
	println(ar[0])
	println(sl[1])
	println(ma[Enone])
```






