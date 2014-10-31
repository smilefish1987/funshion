# golang 笔记 ： 类型转换


1.在golang中允许类型转换，从一种类型转换为另一个种类型。但是有一些规则：
>a.不是所有的转换都是允许的

>b.将一个值转换为另一个类型是由操作符(看起来像函数调用:toType(val))完成的

2.下面列表中的转换是合法的：

|FROM | b []byte|i []int|r []rune|s string|f float32|i int|
|:-----:|:-----:|:-----:|:-----:|:-----:|:-----:|:-----:|
|To []byte|* | | |[]byte(s) | | |
|To []int|* | | |[]int(s) | | |
|To []rune|* | | |[]rune(s) | | |
|To string|string(b)|string(i)|string(r)|*| | |
|To float32| | | ||* |float32(i) |
|To int|* | | ||int(f) | * |

3.从string到字节或者rune的slice
``` golang
    mystring := "hello,string"
    byteslice := []byte(mystring)
    //转换到byteslice，每个byte保存字符串对应字节的整数值。Go的字符串是UTF-8编码的，一些字符可能是1,2,3或者4个字节结尾
    runeslice := []rune(mystring)
    //转换到runeslice,每个rune保存Unicode编码的指针。字符串中的每个字符对应一个整数
```

4.从字节或者整数的slice到字符串
``` golang
    b := []byte{'h','e','l','l','o'}  
    s :=string(b)
    i := []rune{257,1024,65}
    r := string(i)
```

5.对于数值，定义了符合下面规则的转换：

> a.将整数转换到指定的(bit)长度： uint16(int)

> b.从浮点数到整数： int(float32),这会截取浮点数的整数部分

> c.其他的转换： float64(int)

6.用户定义类型的转换
``` golang
    package main
    
    import (
    	"fmt"
    )
    
    func main() {
    	var b bar = bar{1}
    	var ff foo = b  // 报错
    	var f foo = foo(b) //加foo()，可以转换
    	var i int = int(b) // 报错，整数与有整数字段的结构是不一样的，不能转换
    
    	fmt.Printf("%v\n", f)
    	fmt.Printf("%v\n", i)
    }
    
    type foo struct {
    	int
    }
    
    type bar foo

```
7.要转换那些字段不一致的结构是相当困难的，同时上面也说明了整数与有整数字段的结构并不是一样的。









