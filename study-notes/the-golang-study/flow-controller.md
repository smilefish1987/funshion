# golang笔记 流程控制


在Golang里只有少数几个控制结构。语法跟c语言相比，圆括号可以省略，语句体的大括号必须要有。

1. if-else
    
``` golang
    monthdays := map[string]int{
    		"Jan": 31, "Feb": 28, "Mar": 31,
    }
    _, ok := monthdays["Feb"]
    if  ok { //语句体的左大括号必须和if在一行
    		fmt.Println("monthdays has the key: Feb")
    }

    
    
    if _, ok := monthdays["Feb"];ok { //if语句体可以带初始化语句，用于局部变量的初始化
    		fmt.Println("monthdays has the key: Feb")
    }
    
    
    const c2 = 46

	if c2 > 0 && !false { // 可以使用逻辑运算符
		println("c2 is larger then 0")
	} else {
		println("c2 is less then 0")
	}
```	

2. goto

``` golang
    package main
    func main() {
	    myfun()
    }

    func myfun() {
	    i := 0
    Here:
	    println(i)
	    i++
	    if i <= 5 {
		    goto Here  // 用goto调转到当前函数内定义的标签，标签是区分大小写的
	    }
    }
```
    
3. for

``` golang

    package main
    
    func main() {

    	for i := 0; i < 5; i++ {  //传统c-for
    		println(i)
    	}


    	j := 0
    	for j < 5 {       // c-while ，golang中没有while/do-while
    		println(j)
    		j++
    	}


    	k := 0
    	for {      // c-while(true)  golang死循环 
    		println(k)
    		k++
    		if k >= 5 {
    			break  //跳出当前循环
    		}
    	}
    }
```

4. switch 
   
   switch的表达式不必是常量或整数，自上而下执行，直到找到匹配项，如果switch没有表达式，它会匹配true
   
``` golang

    func main() {

	const c = 2

	switch {    //switch没有表达式，匹配true,(switch当if-else使用)
	case c < 0:
		println("c<0")
	case c > 1 && c < 3:
		println("c==2")
	}

	switch c {   //c-switch c风格的switch
	case 1:
		println("c==1")
	case 2:
		println("c==2")
		fallthrough   //默认情况下匹配执行后会直接退出，使用fallthrough可以强制往下执行
	case 3:
		println("c==3")
	}

	switch b := 1; b {  //带初始化表达式的switch
	case -1:
		println("-1")
	case 0:
		println("0")
	default:    //default指定其他所有的分支都不匹配时的行为
		println("default")
	}

	switch a := 'c'; a {
	case 'a', 'b', 'c':   //case分支可以使用逗号分隔的列表，这些值是or的关系
		println("not operator")
	case '+', '-':
		println("operators")
	default:
		println("default")
	}

}

```

5. break和continue

``` golang

    func main() {

    	for i := 0; i < 10; i++ {
    		println(i)
    		if i == 5 {
    			break      //终止循环
    		}
    	}

    Out:
    	for i := 0; i < 5; i++ {
    		for j := 0; j < 10; j++ {
    			println(i, j)
    			if j == 5 {
    				break Out   //break后面接标签，跳出多层循环
    			}
    		}
    	}

    	for i := 0; i < 10; i++ {
    		if i%2 == 0 {
    			continue   //终断本轮循环，让循环进入下一轮
    		}
    		println(i)
    	}
    }
```

	
6. range
range关键字是一个迭代器，可用于循环，可以用于slice,array,map,string,chan类型，当被调用时，从它循环的内容中返回一个键值对

``` golang

    func main() {

    	list := []string{"ab", "abc", "abcd"}
    
    	for k, v := range list {  //用于slice
    		println(k, v)
    	}
    
    	arr := [3]int{1, 2, 3}
    	for k, v := range arr {  //用于array
    		println(k, v)
    	}
    
    	for pos, char := range "Hello,world" {  //用于string
    		println(pos, char)
    	}
    
    	smap := map[string]int{      
    		"a": 12, "b": 14, "c": 15,
    	}
    
    	for key, val := range smap {  //用于map
    		println(key, val)
    	}
    }
```




