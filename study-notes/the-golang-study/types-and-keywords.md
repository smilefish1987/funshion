# golang笔记 类型和关键字


**类型**

一.golang的内置简单类型：

1.bool 取值只能是预定义常量true和false

2.数字类型：
int (有符号，根据硬件平台决定长度，32位就是32,64位就是64位)
uint (无符号，其他的同int)
int8,int16,int32,int64,byte,uint8,uint16,uint32,uint64
float32/float64 (浮点数)

3.字符串
string (**以""双引号包裹的字符序列，''单引号表示字符，不表示字符串**)
[字符串一点被赋值，不能被修改，也就是说在golang中，字符串是不可变的]
可以通过下面的方式修改字符串:

``` golang
    var str string = "hello"
	strslice := []rune(str)
	strslice[1] = 'm'
	str2 := string(strslice)
	fmt.Println(str2)  //hello
	fmt.Println(str)   //hmllo
```
	
多行字符串的拼接：

``` golang
    str := "hello" +  //+号放在行尾
    "world"
``` 
    
原始字符串(以反引号**`**包裹的字符串),原始字符串标识的值，在引号中是不转义的

``` golang
    str := `hello
        world`
```    
4.rune (int32的别名，用UTF-8编码。在遍历字符串中的字符时，需要用到rune,用以获取实际的字符)

5.复数类型 (golang原生支持复数，形式：re+im *i* ,re表示实部，im表示虚部)

>  complex64  32位re,32位im

>  complex128  64位re,64位im 
    
6.error类型 默认值为nil

**Golang关键字**

    break       default        func        interface       select
    case        defer          go          map             struct
    chan        else           goto        package         switch
    const       fallthrough    if          range           type
    continue    for            import      return          var
