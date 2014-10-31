# golang笔记 函数

1.函数是golang程序的基础构件

2.golang中没有c中函数原型声明的概念。在源文件中函数定义的顺序可以随意安排，因为编译在执行时会扫描每个文件。golang不允许函数嵌套，但是可以通过匿名函数实现嵌套。

3.函数通过关键字func来定义：

``` golang
    type myType int //新类型的定义
    //func 函数的接收者 函数名(参数列表) 返回值{}
    func (p myType) addStep(step int) int {
    	return int(p) + step
    }
```
    
4.有接收者的函数称为方法，当然函数也可以没有接收者，就是我们通常说的函数：

``` golang
    func timeTwo(step int) int {  //单返回值
	    return step * step
    }
    
    func printMsg(msg string) {  //没有返回值
    	println(msg)
    }
    
    func returnTwo(in int) (int, int) {  //都返回值
    	return in * 2, in * 3
    }
    
    func checkError(errorno int) (errorlevel int, errormsg string) {    //命令返回值的函数
    	errorlevel = errorno
    	errormsg = "test msg"
    	//return errorlevel, errormsg  //可以这样子返回
    	return     //推荐的方式
    }
```
5.golang的函数和方法可以返回多个返回值
6.golang函数的返回值或者结果参数可以指定一个名字，并且像原始的变量那样使用，就像输入参数那样。如果对其命名，在函数开始时，它们会用其类型的零值初始化。如果函数在不加参数的情况下使用了return语句，结果参数会返回。

7.函数也可以递归

``` golang

    func recFunc(in int) {
    	if in == 0 {
    		return
    	}
    	recFunc(in - 1)
    	println(in)
    }
```
8.在Golang中，定义在函数外的变量是全局的，俗称全局变量。那些定义在函数内部的变量，对于函数来说是局部的，也就是局部变量。如果一个局部变量的名字和一个全局变量的名字相同，在函数执行的时候，局部变量将覆盖全局变量（命名覆盖）

9.局部变量仅仅在执行定义的函数时是有效的。

10.通过defer的延迟列表，可以在函数返回前执行一些资源回收的代码，defer列表以后进先出的方式执行。

``` golang
    
    for i:=0;i<5;i++ {
        defer println(i)
    }
```

11.defer后面可以接匿名函数

``` golang

    defer func(){
        //do-something
    }()

    defer func(x int){
        //do-something
    }(5)   //带参数的匿名函数调用
    
    func f()(retval int){
        defer func(){
            retval++   //匿名函数可以访问任何命名返回值
        }()
        return 0
    }
```

12.函数可以接受不定数量的参数，这样子的函数就是变参函数。golang中变参函数的定义如下：

``` golang
func myfunc(arg ...int) {
   // arg ...int 告诉编译器函数接受不定数量的参数，这些参数全部都是int型，在函数体中，变量arg是一个int类型的slice.
	for key, val := range arg {
		println(key, val)
	}
}

func main() {
	myfunc(1, 2)
	myfunc(1, 2, 3, 4)
}
```

13.如果不指定变参的类型，默认是空的接口 interface{}

14.函数可以跟其他类型的值一样，复制给变量或者作为参数传递,也就是说在Golang中函数是一等公民

``` golang

    f := func() {
		println("func")
	}

	f()

	fmt.Printf("%T\n", f)  // 变量f的类型为 func()

	var funcMap = map[int]func() int{
		1: func() int { println(1); return 1 },
		2: func() int { println(2); return 2 },
	}

	println(funcMap[1]())
	fmt.Printf("%T\n", funcMap[2]) //变量funcMap[2]的类型为  func() int

```

15.把函数作为参数传递给其他的函数，就可以实现传统意义上的回调

```golang
    func funcprint(x int) {
    	println(x)
    }
    
    func callback(ret int, f func(int)) { //将函数作为参数
    	f(ret)
    }
    
    callback(10,funcprint)
    
```

16.golang中的异常处理，没有传统意义上的try/catch的结构，使用的是panic-and-recover机制，建议只是把panic-and-recover作为最后的手段被使用，不到万不得以不建议使用。panic和recover都是内建函数，说明如下：

**func panic(v interface{})**
*The panic built-in function stops normal execution of the current goroutine. When a function F calls panic, normal execution of F stops immediately. Any functions whose execution was deferred by F are run in the usual way, and then F returns to its caller. To the caller G, the invocation of F then behaves like a call to panic, terminating G's execution and running any deferred functions. This continues until all functions in the executing goroutine have stopped, in reverse order. At that point, the program is terminated and the error condition is reported, including the value of the argument to panic. This termination sequence is called panicking and can be controlled by the built-in function recover.*

panic是一个内建函数，可以中断原有的控制流程，进入一个令人恐慌的流程中。当函数F调用panic,函数F的执行被中断，接着F中的延迟函数会正常执行，最后F返回到调用它的地方。在调用的地方，F的行为就像调用了panic。这一过程继续向上，直到程序崩溃时的所有goroutine返回。恐慌可以直接调用panic产生，也可以由运行时错误产生，比如数组访问越界

**func recover() interface{}**
*The recover built-in function allows a program to manage behavior of a panicking goroutine. Executing a call to recover inside a deferred function (but not any function called by it) stops the panicking sequence by restoring normal execution and retrieves the error value passed to the call of panic. If recover is called outside the deferred function it will not stop a panicking sequence. In this case, or when the goroutine is not panicking, or if the argument supplied to panic was nil, recover returns nil. Thus the return value from recover reports whether the goroutine is panicking.*

recover是一个内建的函数，可以让进入令人恐慌的流程中的goroutine恢复过来。
** recover 仅在延迟函数中有效** 。在正常的执行过程中，调用recover会返回nil并且没有任何其他效果。如果当前的goroutine陷入恐慌，调用recover可以捕获到panic的输入值，并且恢复政策的执行。

``` golang

    //在defer函数中通过检测recover函数的返回值判断是否发生panic
    func isthrowPanic(f func())(b bool){
        defer func(){
            if x:=recover();x!=nil{
                b = true
            }
        }()
        f()
        return
    }
```




