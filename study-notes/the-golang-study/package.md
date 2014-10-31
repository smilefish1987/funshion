# golang 笔记  包

1.golang中包就是函数和数据的集合。一个golang包可以包含多个文件，只需要在文件名的开始都通过package关键字定义这些文件属于同一个包名就可以。在实际开发中golang的包名也对应$GOPATH环境变量中src下的目录名，通过go build命令编译包，go install 安装包。

2.包名定义的语法： package <package-name>

3.包名约定为全小写

4.包实践
4.1新建一个名为cal.go的文件，将其放在$GOPATH/src/cal文件夹，内容如下：

``` golang
/*
    包的文档用于说明包是用来做什么用的，出现在包定义语句package cal的前面，这段话会出现在go doc生成的关于包的页面上
*/
package cal

// 用于将函数的两个参数a,b相加并返回结果
func Add(a int, b int) int {
	return a + b
}
//用于将函数的两个参数a,b想乘并返回结果
func Multi(a int, b int) int {
	return a * b
}
//这是一个私有的函数，包外面不能使用
func divdy(a int, b int) int {
	return a % b
}
```
4.2 执行下面的编译和安装命令
``` shell
$ go build
$ go install
```

4.3 新建一个名为test.go的包的测试文件，将其放在$GOPATH/src文件，内容如下：
``` golang

package main

import (
	"cal"
)

func main() {

	println(cal.Add(2, 3))   //5
	println(cal.Multi(4, 5)) //20
}
```

4.4.通过下面的命令编译和运行，可以得到相应的结果
``` shell
    $ go run test.go
```

4.5 通过import语句导入包，包名是默认名称。当然也可以通过在导入语句中制定其他的名称来覆盖默认的名称：
``` golang
    import  bar "cal"  //现在可以通过bar访问cal包中的内容
    bar.Add(1,2)
```

5.在Golang中，当函数的首字母大写的时候，该函数会被从包中导出(在包外部是可见的，也就是公有的)，总结golang的可见性规则：
5.1 公有函数/属性的名字以大写字母开头
5.2 私有函数/属性的名字以小写字母开头

6.需要编译成可执行文件的go源文件，必须有一个main包，作为项目的入口，main包中必须定义一个如下所示的main函数作为程序执行的入口。
``` golang
func main(){
    // do-something
}
```

7.包的单元测试，包的测试文件也在包的目录中，被命名为：packagename_test.go,每个测试函数都有相同的标识，形式如下： func TestXxx(t *testing.T){}, 其中Xxx为包中能被导出的需要被测试的函数名。通过在$GOPATH/src下包名的文件夹中执行 go test 命令来执行测试。

8.包的测试实践
8.1 在$GOPATH/src/cal下面，新建一个cal_test.go的文件，内容如下：

``` golang

package cal

import (
	"testing"
)

//每一个函数对应一个Test开头的函数，且只能对应一个
func TestAdd(t *testing.T) {
	if 5 != Add(2, 3) {
		t.Log("2+3 should 5")
		t.Fail()
	}

	// if 6 != Add(2, 3) {
	// 	t.Log("2+3 is not 6")
	// 	t.FailNow()
	// }
}

func TestMulti(t *testing.T) {
	if 10 != Multi(2, 5) {
		t.Log("2 * 5 should 10")
		t.FailNow()
	}
}

```

8.2 到文件夹$GOPATH/src/cal 执行如下命令:
``` golang
    $ go test
```
可以看到结果：

``` shell
    PASS
    ok     cal  0.004s
```

9.testing包中的函数说明：
``` golang
    
    func (t *T) Fail() //标记测试函数失败，但仍然继续执行
    
    func (t *T) FailNow() //标记测试函数失败，并且中断其执行。当前文件中的其余的测试将被跳过，然后执行下一个文件中的测试
    
    func (t *T) Log(args ...interface{})  //用默认格式对其参数进行格式化，并记录文件到错误日志
    
    func (t *T) Fatal(args ...interface{})  //等价于Log() 后跟随FailNow()
    
```

10.常用包的说明：

| 包名   |  用途与说明                       |
| -------| :----------------------------------: |
|fmt  |fmt包中实现了格式化的I/O函数，格式化短语：%v  默认格式的值，当打印结构时，加(%+v)会增加字符名; %#v Golang样式的值表达式; %T  带有类型的Golang样式的值表达|
|io   |io包提供了原始的I/O操作界面。它主要的任务是对os包这样的原始I/O进行封装，增加一些其他相关，使其具有抽象功能用在公告的接口上|
|bufio | bufio包实现了带缓存的I/O,它封装于io.Reader和io.Writer对象，创建了另一个对象(Reader和Writer)在提供缓冲的同时实现了一些文本I/O的功能 |
| sort | sort包提供了对数组和用户定义集合的原始的排序功能|
| strconv | strconv包提供了将字符串转换成基本数据类型，或者从基本数据类型转换为字符串的功能|
| os | os包提供了与平台无关的操作系统功能接口，其设计是Unix形式的 |
| sync | sync包提供了基本的同步原语，比如互斥锁|
| flag | flag包实现了命令行的解析 |
| encoding/json| encoding/json包实现了编码与解码RFC 4627定义的json对象 |
| html/template | html/template包实现了数据驱动的模版，用于生成文本输出，比如html. 将模板关联到某个数据结构上进行解析。模板内容指向数据结构的元素（通常结构的字段或者map的键）控制解析并且决定某个值会被显示。模板扫描结构以便解析，而“游标”@决定了当前位置在结构中的值 |
| net/http | net/http实现了HTTP请求/响应和URL的解析，并且提供了可扩展的HTTP服务和基本的HTTP客户端 |
| unsafe | unsafe包包含了Golang程序中数据类型上所有不安全的操作。通常无需使用这个 |
| reflect | reflect包实现了运行时反射，运行程序通过抽象类型操作对象。通常用于处理静态类型interface{}的值，并且通过Typeof解析出其动态类型信息，通常会返回一个有接口类型Type的对象 |
| os/exec | os/exec包执行外部命令 |







