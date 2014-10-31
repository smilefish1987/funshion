# golang笔记 array/slice/map

**array**

1.array的声明 **[n]type**  n表示数组的长度，而type表示数组中存储数据的类型

``` golang
    var iarr [3]int
```

2.对数组元素的赋值和访问，通过方括号[]和索引完成，索引从0开始

``` golang
    var iarr [3]int
	iarr[0] = 1
	iarr[1] = 2
	println(iarr[0], iarr[1], iarr[2])
```
	
3.也可以使用类型推断的带初始值的声明和初始化
``` golang
    iarr2 := [2]int{1, 2}
	println(iarr2[0], iarr2[1])
```

4.当带初始化值时，数组的长度可以用...代替
``` golang
    iarr3 := [...]int{1, 2, 3}
	println(iarr3[0], iarr3[2])
```
	
5.当声明和初始化多维数组时，所有项目必须都指定
``` golang
    iarr4 := [3][2]int{[2]int{1, 2}, [2]int{3, 4}, [2]int{5, 6}}
	println(iarr4[0][1])
```
	
6.声明一个array时，你必须在方括号内输入些内容，数字或者三个点(...)
7.array的大小是类型的一部分，大小是固定的，不能改变
8.数组是值类型，当赋值或者作为参数时会发生拷贝
9.复合类型array,slice和map,元素复合声明的类型与外部一致，则可以省略

``` golang
    iarr5 := [3][2]int{{7, 8}, {9, 10}, {11, 12}}
	println(iarr5[1][1])
```
	
	
**slice**

1.slice与array相似，但是在新的元素加入时可以增加长度。slice只是指向底层的一个array，所以可以理解slice是一个指向array的指针。不同于array,slice是引用类型,也就是说当把slice赋值给另一个变量时，两个变量的引用指向同一个底层的array；当把slice作为参数传递给函数时，在函数内对slice的修改会体现在函数的外面

2.通过内建的make函数声明一个slice,slice的使用和array是一样的
``` golang
    sl := make([]int, 10)
	sl[0] = 1
	sl[1] = 2
	println(sl[0], sl[1], sl[2])  // 1 2 0
	println(len(sl), cap(sl))     // 10  10
```
3.通过已有的array和slice创建一个新的slice:
``` golang
    a := [...]int{1, 2, 3, 4, 5, 6, 7}
	s1 := a[2:4]
	println(s1[0], s1[1], len(s1), cap(s1))  //3 4 2 5
	s1 = append(s1, 4)
	println(s1[0], s1[1], len(s1), cap(s1))  //3 4 3 5
	s2 := a[2:]
	println(len(s2), cap(s2))  // 5 5
	s3 := a[:5]
	println(len(s3), cap(s3))  // 5 7
	s4 := a[:]
	println(len(s4), cap(s4))  // 7  7
	s5 := s4[:]
	println(len(s5), cap(s5))  // 7 7 
```

4.append函数用于往slice里面追加元素或其他slice:

``` golang
   func append(slice []Type, elems ...Type) []Type
//The append built-in function appends elements to the end of a slice. If it has sufficient capacity, the destination is resliced to accommodate the new elements. If it does not, a new underlying array will be allocated. Append returns the updated slice. It is therefore necessary to store the result of append, often in the variable holding the slice itself.

    slice = append(slice, elem1, elem2)
    slice = append(slice, anotherSlice...) //最后这三个点不能少
    
    slice0 := []int{1,2}
    slice1 := append(slice0,3)
    slice2 := append(slice1,4,5,6)
    slice3 := append(slice2,slice0...) //最后这三个点不能少
```

5.copy函数用于复制slice:
``` golang
    func copy(dst, src []Type) int
//The copy built-in function copies elements from a source slice into a destination slice. (As a special case, it also will copy bytes from a string to a slice of bytes.) The source and destination may overlap. Copy returns the number of elements copied, which will be the minimum of len(src) and len(dst).
    var a = [...]int{0, 1, 2, 3, 4, 5, 6, 7}
    var s = make([]int, 6)
    n1 := copy(s, a[0:])
    println(n1) // 6
    n2 := copy(s, s[2:])
    println(n2) //4
```
 
 
 **map**   

1.map就是一个key:value的对，定义的类型如下：
``` golang
    map[key-type]value-type


    tmap := make(map[string]int)
    tmap["mm"] = 1
    fmt.Println(tmap["mm"])
    
    monthdays := map[string]int{
    	"Jan": 31, "Feb": 28, "Mar": 31,  //最后这个逗号不能少
    }
    
    //通过for和range遍历array,slice,map,string
    for key, val := range monthdays {
    	fmt.Printf("%s\t%d \n", key, val)
    	fmt.Println(key + "\t" + strconv.Itoa(val))
    }
    
    //判断元素是否存在
    _, ok := monthdays["Feb"]

	if ok {
		fmt.Println("monthdays has the key: Feb")
	}
	
    //删除map中指定的key
func delete(m map[Type]Type1, key Type)
//The delete built-in function deletes the element with the specified key (m[key]) from the map. If m is nil or there is no such element, delete is a no-op.
    
	delete(monthdays, "Jan")
```	




