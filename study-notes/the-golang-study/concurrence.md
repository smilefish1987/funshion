# golang笔记  并发

1.Golang从语言层面就支持了并行。goroutine是Golang并行设计的核心要素。

2.叫做goroutine是因为已有的短语——线程、协程、进程等等都传递不了准确的含义。goroutine有简单的模型：它是与其他goroutine并行执行的，有着相同地址空间的函数。它是轻量的，仅比分配栈空间多一点点消耗，而初始时栈是很小的，所以它们也是廉价的，并且随着需要在堆空间上分配（和释放）

3.goroutine说到底其实就是线程，但是它比线程更小，十几个goroutine可能体现在底层就是五六个线程，Go语言内部帮你实现了这些goroutine之间的内存共享。执行goroutine只需极少的栈内存(大概是4~5KB)，当然会根据相应的数据伸缩。也正因为如此，可同时运行成千上万个并发任务。goroutine比thread更易用、更高效、更轻便。

4.goroutine是通过Go的runtime管理的一个线程管理器。goroutine通过go关键字实现了，其实就是一个普通的函数。
``` golang
package main

import (
	"fmt"
	"time"
)

func ready(w string, sec int) {
	time.Sleep(time.Duration(sec) * time.Second)
	fmt.Println(w, " is ready!")
}

func main() {

	go ready("tea", 2)
	go ready("coffee", 1)
	fmt.Println("I'm waiting")
	//main函数等待的时间足够长，每个goroutine会打印各自的文本到屏幕，如果不等待goroutine的执行，程序立刻终止，而任何正在执行的goroutine都会停止。为了修复这个，需要一些能够同goroutine通信的机制，这一机制通过channels的形式使用
	time.Sleep(3 * time.Second)
}
```

``` golang
package main

import (
	"runtime"
)

func main() {
	go say("world")
	say("hello")
}

func say(s string) {
	for i := 0; i < 5; i++ {
		//runtime.Gosched()表示让CPU把时间片让给别人,下次某个时候继续恢复执行该goroutine
		runtime.Gosched()
		println(s)
	}
}
//上面的多个goroutine运行在同一个进程里面，共享内存数据，不过设计上我们要遵循：不要通过共享来通信，而要通过通信来共享.
```

5.channel可以与Unix Shell中的双向管道做类比：可以通过它发送或者接收值。这些值只能是特定的类型：channel类型。定义一个channel时，也需要定义发送到channel中的值的类型。注意：必须使用make创建channel:
``` golang
ci := make(chan int) //用于发送和接收int
cs := make(chan string) //用于发送和接收string
cf := make(chan interface{}) //用于发送和接收任意类型
```
6.向channel发送或者接收数据，是通过操作符 <- 完成的，具体是发送还是接收依赖于操作符的位置：

``` golang
ci <- 1  //发送1到ci
<-ci  //从ci 接收数据并扔掉
i:= <-ci //从ci接收，并保存到i中

```

7.channel用于通信的实例：
``` golang
package main

import (
	"fmt"
	"time"
)

var c chan int    //定义c作为int型的channel,并且是全局的，goroutine可以访问它

func ready(w string, sec int) {
	time.Sleep(time.Duration(sec) * time.Second)
	fmt.Println(w, " is ready!")
	c <- 1   //发送数据到c ,位置在打印之后
}

func main() {
	c = make(chan int)   //初始化 c
	go ready("tea", 2)   //用go 开始一个goroutine
	go ready("coffee", 1)

	fmt.Println("I'm waiting,but not too long")
	<-c   //等待，直到从channel上接收一直，然后把值扔掉
	<-c   //接收另一个值 (阻塞式的)
}

//代码运行的结果：
// I'm waiting,but not too long
//	coffee  is ready!
//	tea  is ready!
```
8.Golang内建的关键字:select,通过select(和其东西)可以监听channel上输入的数据。
``` golang
package main

import (
	"fmt"
	"time"
)

var c chan int

func ready(w string, sec int) {
	time.Sleep(time.Duration(sec) * time.Second)
	fmt.Println(w, " is ready!")
	c <- 1
}

func main() {
	c = make(chan int)
	go ready("tea", 2)
	go ready("coffee", 1)

	fmt.Println("I'm waiting,but not too long")
	i := 0

L:
	for {
		select { //通过select监听channel上的输入，现在将会一直等待，直到从channel c上收到多个响应时才会退出循环L
		case <-c:
			i++
			if i > 1 {
				break L
			}
		}
	}
}
``` 

9.select默认是阻塞的，只有当监听的channel中有发送或接收可以进行时才会运行，当多个channel都准备好的时候，select是随机的选择一个执行的。
``` golang
package main

import "fmt"

func fibonacci(c, quit chan int) {
    x, y := 1, 1
    for {
        select {
        case c <- x:
            x, y = y, x + y
        case <-quit:
            fmt.Println("quit")
            return
        }
    }
}

func main() {
    c := make(chan int)
    quit := make(chan int)
    go func() {
        for i := 0; i < 10; i++ {
            fmt.Println(<-c)
        }
        quit <- 0
    }()
    fibonacci(c, quit)
}


//在select里面还有default语法，select其实就是类似switch的功能，default就是当监听的channel都没有准备好的时候，默认执行的（select不再阻塞等待channel）

select {
case i := <-c:
    // use i
default:
    // 当c阻塞的时候执行这里
}
```

10.虽然goroutine是并发执行的，但是他们并不是并行执行的。如果不告诉Golang额外的东西，同一时刻只会有一个goroutine执行。利用runtime.GOMAXPROCS(n)可以设置goroutine并行执行的数量。【GOMAXPROCS设置了同时运行的CPU的最大数量，并返回之前的设置。如果n<1,不会改变当前设置。当调度得到改进后，这将被移除】。如果不希望修改源代码，同样可以通过设置环境变量GOMAXPROCS为目标值。

11.在Golang中通过 ch :=make(chan int) 创建的channel是无缓冲的channel。无缓冲channel是在多个goroutine之间同步很棒的工具。无缓冲的channel主要体现在两方面：

>a.当从channel读取数据 （value := <-ch）,它将被阻塞，直到有数据被接收。

>b. 往channel中发送数据(ch <- 5)，将会被阻塞，直到数据被读出

12.Golang也提供了带缓冲的channel,缓冲的大小指定了channel中可以存储多少个元素。 ch :=make(chan int,4) ,创建了一个可以存储4个元素的int型channel,在这个channel中前4个元素可以无阻塞的写入，当写入第5个元素时，代码会被阻塞，直到其他goroutine从channel中读取了一些元素，腾出了空间。
``` golang
ch := make(chan type,value)
// value ==0  无缓冲
// value >0   缓冲value个元素
```

13.通过内置函数close来关闭channel当channel被关闭后，读取端需要知道这个事情。
``` golang
package main

import (
    "fmt"
)

func fibonacci(n int, c chan int) {
    x, y := 1, 1
    for i := 0; i < n; i++ {
        c <- x
        x, y = y, x + y
    }
    close(c)  //通过close关闭channel
}

func main() {
    c := make(chan int, 10)
    go fibonacci(cap(c), c)
    for i := range c { //range c 能不断的读取channel里面的数据，直到该channel被显式的关闭
        fmt.Println(i)
    }
}

//通过下面的代码测试channel是否被关闭
x , ok = <- ch
//当ok被赋值为true意味着channel尚未关闭，同时可以读取数据。当ok被赋值为false，意味着channel被关闭
```

14.channel的两点说明：
>a. 记住应该在生产者的地方关闭channel，而不是消费的地方去关闭它，这样容易引起panic

>b.记住一点的就是channel不像文件之类的，不需要经常去关闭，只有当你确实没有任何发送数据了，或者你想显式的结束range循环之类的

15.有时候会出现goroutine阻塞的情况，那么我们如何避免整个程序进入阻塞的情况呢？我们可以利用select来设置超时，通过如下的方式实现：
``` golang
func main() {
    c := make(chan int)
    o := make(chan bool)
    go func() {
        for {
            select {
                case v := <- c:
                    println(v)
                case <- time.After(5 * time.Second):
                    println("timeout")
                    o <- true
                    break
            }
        }
    }()
    <- o
}

```

16.runtime包中有几个处理goroutine的函数:

| 函数 |   用途 |
|:------| :------|
|func Goexit()|退出当前执行的goroutine，但是defer函数还会继续调用|
|func Gosched()|让出当前goroutine的执行权限，调度器安排其他等待的任务运行，并在下次某个时候从该位置恢复执行|
|func NumCPU() int |返回 CPU 核数量|
|func NumGoroutine() int|返回正在执行和排队的任务总数|
|func GOMAXPROCS(n int) int|用来设置可以并行计算的CPU核数的最大值，并返回之前的值|








