# golang笔记 运算符和内建函数

**运算符**

1.golang支持的运算符

    *(乘)    /(除)    %(取余)     <<(左移)      >>(右移)      &(按位与)      &^(按位清除)
    +(加)    -（减）    | (按位或)     ^（按位异或）
    ==(等于)   !=（不等于）   <(小于)     <=（小于等于）      >（大于）       >=（大于等于）
    <-（通道）
    && (逻辑与)
    || （逻辑或）

2.golang不支持运算符重载(方法的重载),但是一些内置的运算符却可以重载，比如+号

**内建函数**

[内建函数，无需引入任何包就可以使用它们]

1.close 函数  用于channel通信，使用它来关闭channel
``` golang
func close(c chan<- Type)
```
*The close built-in function closes a channel, which must be either bidirectional or send-only. It should be executed only by the sender, never the receiver, and has the effect of shutting down the channel after the last sent value is received. After the last value has been received from a closed channel c, any receive from c will succeed without blocking, returning the zero value for the channel element.*

2.delete 函数 用于删除map中指定键的值
``` golang
func delete(m map[Type]Type1, key Type)
```
*The delete built-in function deletes the element with the specified key (m[key]) from the map. If m is nil or there is no such element, delete is a no-op.*

3.len 函数 返回指定变量的长度
``` golang
func len(v Type) int
```
*The len built-in function returns the length of v, according to its type:*

    Array: the number of elements in v.
    Pointer to array: the number of elements in *v (even if v is nil).
    Slice, or map: the number of elements in v; if v is nil, len(v) is zero.
    String: the number of bytes in v.
    Channel: the number of elements queued (unread) in the channel buffer;
    if v is nil, len(v) is zero.
    
4.cap 函数 返回指定变量的容量
``` golang
func cap(v Type) int
```
*The cap built-in function returns the capacity of v, according to its type:*

    Array: the number of elements in v (same as len(v)).
    Pointer to array: the number of elements in *v (same as len(v)).
    Slice: the maximum length the slice can reach when resliced;
    if v is nil, cap(v) is zero.
    Channel: the channel buffer capacity, in units of elements;
    if v is nil, cap(v) is zero.
    
5.new 函数  用于各种类型的内存分配
``` golang
func new(Type) *Type
```
*The new built-in function allocates memory. The first argument is a type, not a value, and the value returned is a pointer to a newly allocated zero value of that type.*

6.make 函数 用于内建类型(slice,map,chan)的内存分配
``` golang
func make(Type, size IntegerType) Type
```
*The make built-in function allocates and initializes an object of type slice, map, or chan (only). Like new, the first argument is a type, not a value. Unlike new, make's return type is the same as the type of its argument, not a pointer to it. The specification of the result depends on the type:*

    Slice: The size specifies the length. The capacity of the slice is equal to its length. A second integer argument may be provided to specify a different capacity; it must be no smaller than the length, so make([]int, 0, 10) allocates a slice of length 0 and capacity 10.
    Map: An initial allocation is made according to the size but the resulting map has length 0. The size may be omitted, in which case a small starting size is allocated.
    Channel: The channel's buffer is initialized with the specified buffer capacity. If zero, or the size is omitted, the channel is unbuffered.
    
7.copy 函数 用于slice的复制
``` golang
func copy(dst, src []Type) int
```
*The copy built-in function copies elements from a source slice into a destination slice. (As a special case, it also will copy bytes from a string to a slice of bytes.) The source and destination may overlap. Copy returns the number of elements copied, which will be the minimum of len(src) and len(dst).*

8.append 函数 用于往slice的尾部追加元素
``` golang
func append(slice []Type, elems ...Type) []Type
```
*The append built-in function appends elements to the end of a slice. If it has sufficient capacity, the destination is resliced to accommodate the new elements. If it does not, a new underlying array will be allocated. Append returns the updated slice. It is therefore necessary to store the result of append, often in the variable holding the slice itself:*

``` golang
    slice = append(slice, elem1, elem2)
    slice = append(slice, anotherSlice...)
```

*As a special case, it is legal to append a string to a byte slice, like this:*

``` golang
    slice = append([]byte("hello "), "world"...)
```
    
9.panic 函数 用于终止当前goroutine的执行
``` golang
func panic(v interface{})
```
*The panic built-in function stops normal execution of the current goroutine. When a function F calls panic, normal execution of F stops immediately. Any functions whose execution was deferred by F are run in the usual way, and then F returns to its caller. To the caller G, the invocation of F then behaves like a call to panic, terminating G's execution and running any deferred functions. This continues until all functions in the executing goroutine have stopped, in reverse order. At that point, the program is terminated and the error condition is reported, including the value of the argument to panic. This termination sequence is called panicking and can be controlled by the built-in function recover.*


10.recover 函数 用于恢复被panic函数终止的goroutine的行为
``` golang
func recover() interface{}
```
*The recover built-in function allows a program to manage behavior of a panicking goroutine. Executing a call to recover inside a deferred function (but not any function called by it) stops the panicking sequence by restoring normal execution and retrieves the error value passed to the call of panic. If recover is called outside the deferred function it will not stop a panicking sequence. In this case, or when the goroutine is not panicking, or if the argument supplied to panic was nil, recover returns nil. Thus the return value from recover reports whether the goroutine is panicking.*

11.print 函数 用于打印信息
``` golang
func print(args ...Type)
```
*The print built-in function formats its arguments in an implementation- specific way and writes the result to standard error. Print is useful for bootstrapping and debugging; it is not guaranteed to stay in the language.*


12.println 函数 用于打印信息并换行
``` golang
func println(args ...Type)
```
*The println built-in function formats its arguments in an implementation- specific way and writes the result to standard error. Spaces are always added between arguments and a newline is appended. Println is useful for bootstrapping and debugging; it is not guaranteed to stay in the language.*

13.complex 函数 用于构造一个复数
``` golang
func complex(r, i FloatType) ComplexType
```
*The complex built-in function constructs a complex value from two floating-point values. The real and imaginary parts must be of the same size, either float32 or float64 (or assignable to them), and the return value will be the corresponding complex type (complex64 for float32, complex128 for float64).*

14.real 函数 用于返回一个复数的实部
``` golang
func real(c ComplexType) FloatType
```
*The real built-in function returns the real part of the complex number c. The return value will be floating point type corresponding to the type of c.*

15.imag 函数 用于返回一个复数的虚部
``` golang
func imag(c ComplexType) FloatType
```
*The imag built-in function returns the imaginary part of the complex number c. The return value will be floating point type corresponding to the type of c.*





