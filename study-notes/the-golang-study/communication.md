# golang笔记 通信

1.Golang文件操作不带缓冲的I/O：
``` golang
package main

import (
	"os"  //导入os包
)

func main() {
	buf := make([]byte, 1024)
	f, _ := os.Open("interface.go")  //打开要读取的文件
	defer f.Close()   //确保文件会关闭
	for {
		n, _ := f.Read(buf)  // 一次性读取1024个字节
		if n == 0 {   //读到了文件末尾
			break   
		}
		os.Stdout.Write(buf[:n])  //将读取的内容写入os.Stdout
	}
}
```

2.Golang使用带缓冲的I/O:
``` golang
package main

import (
	"bufio"   //导入带缓存的I/O包
	"os"
)

func main() {
	buf := make([]byte, 1024)
	f, _ := os.Open("interface.go")
	defer f.Close()
	r := bufio.NewReader(f)   //转换f为有缓冲的Reader
	w := bufio.NewWriter(os.Stdout)

	defer w.Flush()
	for {
		n, _ := r.Read(buf) //从Reader读取
		if n == 0 {
			break
		}
		w.Write(buf[0:n]) //向屏幕输出文件内容
	}
}
```

3.io.Reader接口对于Golang来说非常重要。许多函数需要通过io.Reader读取一些数据作为输入。为了满足这个接口，只需要实现一个方法: Read(p []byte) (n int, err error)。写入则是实现了Write方法的io.Writer接口。如果你让自己的程序或者包中的类型实现了io.Reader或者io.Writer接口，那么整个Go标准库都可以使用这个类型。
``` golang
//io包中的Reader和Writer接口
type Reader interface {
        Read(p []byte) (n int, err error)
}
type Writer interface {
        Write(p []byte) (n int, err error)
}
```

4.bufio包中有ReadString()和ReadLine()方法，可以完成更加通用的一次读取一行的功能

5.检查一个目录或者文件是否存在
``` golang
package main

import (
	"os"
)

func main() {
    //通过os.Stat可以检测文件或者目录是否存在
	if _, e := os.Stat("name"); e != nil {
		os.Mkdir("name", 0755)  //不存在，通过Mkdir()创建目录
	} else {
		println("exist")
	}
}
```

6.命令行参数可以通过OS包，或者flag包进行解析：
``` golang
package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	for k, v := range os.Args {
		println(k, v)
	}

	dnsses := flag.Bool("dnssec", false, "Request DNSSEC records")
	port := flag.String("port", "53", "Set the query port")
	flag.Usage = func() { //重新定义Usage函数
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS] [name ...]\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()    //解析标识，并填充变量
}
```

7.可以通过os/exec包中的函数执行外部命令，这也是Golang中主要的执行命令的方法。通过定义一个有着数个方法的*exec.Cmd结构来使用

``` golang
package main

import (
	"os/exec"
)

func main() {
	cmd := exec.Command("/bin/ls", "-l")
	err := cmd.Run()
	if err != nil {
		println("error")
	} else {
		println(err)
	}


	buf, err := cmd.Output()
	//buf是一个[]byte,命令执行返回的数据
	if err != nil {
		println("error")
	} else {
		println(buf)
	}
}
``` 

8.所有网络相关的类型和函数可以在net包中找到。这其中最重要的函数是Dial()。当Dial到远程系统，这个函数返回Conn接口类型，可以用于发送和接收信息。函数Dial()简洁的抽象了网络层和传输层。所以IPv4或者IPv6，TCP或者UDP可以共用一个接口。
``` golang
conn,e := Dial("tcp","192.168.1.23:80")
conn,e := Dial("udp","192.168.1.23:80")
conn,e := Dial("tcp","[2620:0:2d0:200::10]:80") //ipv6方括号是必须的
``` 

9.conn是一个io.ReadWriter,可以通过它发送和接收消息

10.net是比较底层的包，还有更高层次的包，比如：http
``` golang
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	r, err := http.Get("http://www.baidu.com/robots.txt") //使用http的Get获取html
	if err != nil {  // 错误处理
		fmt.Printf("%s\n", err)
		return
	}

	b, err := ioutil.ReadAll(r.Body)  //将内容读入b
	r.Body.Close()
	if err == nil {
		fmt.Printf("%s", string(b))
	}
}
```




