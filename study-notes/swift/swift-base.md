# swift 笔记：基础部分

1.常量和变量必须在使用前声明，用let来声明常量，用var来声明变量。
``` swift
let maximumNumberOfLoginAttempts = 10
var currentLoginAttempt = 0
```

2.以在一行中声明多个常量或者多个变量，用逗号隔开
``` swift
var x = 0.0, y = 0.0, z = 0.0
```

3.如果你的代码中有不需要改变的值，请使用let关键字将它声明为常量。只将需要改变的值声明为变量。

4.当你声明常量或者变量的时候可以加上类型标注（type annotation），说明常量或者变量中要存储的值的类型。如果要添加类型标注，需要在常量或者变量名后面加上一个冒号和空格，然后加上类型名称
``` swift
var welcomeMessage: String
welcomeMessage = "Hello"
```

5.一般来说你很少需要写类型标注。如果你在声明常量或者变量的时候赋了一个初始值，Swift可以推断出这个常量或者变量的类型，请参考类型安全和类型推断。

6.可以用任何你喜欢的字符作为常量和变量名，包括 Unicode 字符，常量与变量名不能包含数学符号，箭头，保留的（或者非法的）Unicode 码位，连线与制表符。也不能以数字开头，但是可以在常量与变量名的其他地方包含数字。

7.一旦你将常量或者变量声明为确定的类型，你就不能使用相同的名字再次进行声明，或者改变其存储的值的类型。同时，你也不能将常量与变量进行互转。

8.如果你需要使用与Swift保留关键字相同的名称作为常量或者变量名，你可以使用反引号（`）将关键字包围的方式将其作为名字使用。无论如何，你应当避免使用关键字作为常量或变量名，除非你别无选择。

9.使用println()或者print()函数输出常量或者变量的值。Swift 用字符串插值（string interpolation）的方式把常量名或者变量名当做占位符加入到长字符串中，Swift 会用当前常量或变量的值替换这些占位符。将常量或变量名放入圆括号中，并在开括号前使用反斜杠将其转义：
``` swift
println("The current value of friendlyWelcome is \(friendlyWelcome)")
// 输出 "The current value of friendlyWelcome is Bonjour!
```

10.Swift 中的注释与C 语言的注释非常相似。单行注释以双正斜杠（//）作为起始标记，行注释，其起始标记为单个正斜杠后跟随一个星号（/*），终止标记为一个星号后跟随单个正斜杠（*/）。与 C 语言多行注释不同，Swift 的多行注释可以嵌套在其它的多行注释之中。
``` swift
// 这是一个注释

/* 这是一个,
多行注释 */

/* 这是第一个多行注释的开头
/* 这是第二个被嵌套的多行注释 */
这是第一个多行注释的结尾 */

```

11.与其他大部分编程语言不同，Swift 并不强制要求你在每条语句的结尾处使用分号（;），当然，你也可以按照你自己的习惯添加分号。有一种情况下必须要用分号，即你打算在同一行内写多条独立的语句：
``` swift
let cat = "🐱"; println(cat)
```

12.Swift 提供了8，16，32和64位的有符号和无符号整数类型。有符号的Int8,Int16,Int32,Int64还有Int，无符号的UInt8,UInt16,UInt32,UInt64还有UInt。不同整数类型的min和max属性来获取对应类型的最大值和最小值：
``` swift
let minValue = UInt8.min  // minValue 为 0，是 UInt8 类型的最小值
let maxValue = UInt8.max  // maxValue 为 255，是 UInt8 类型的最大值
```

13.Swift 提供了两个特殊的整数类型Int/UInt,长度与当前平台的原生字长相同:
>a.在32位平台上，Int和Int32,UInt和UInt32长度相同

>b.在64位平台上，Int和Int64,UInt和UInt64长度相同

14.尽量不要使用UInt，除非你真的需要存储一个和当前平台原生字长相同的无符号整数。除了这种情况，最好使用Int，即使你要存储的值已知是非负的。统一使用Int可以提高代码的可复用性，避免不同类型数字之间的转换，并且匹配数字的类型推断。

15.Swift 提供了两种有符号浮点数类型,Double精确度很高，至少有15位数字，而Float最少只有6位数字
>a.Double表示64位浮点数。当你需要存储很大或者很高精度的浮点数时请使用此类型

>b.Float表示32位浮点数。精度要求不高的话可以使用此类型

16.Swift 是一个类型安全（type safe）的语言。类型安全的语言可以让你清楚地知道代码要处理的值的类型。由于 Swift 是类型安全的，所以它会在编译你的代码时进行类型检查（type checks），并把不匹配的类型标记为错误。

17.如果你没有显式指定类型，Swift 会使用类型推断（type inference）来选择合适的类型。有了类型推断，编译器可以在编译代码的时候自动推断出表达式的类型。原理很简单，只要检查你赋的值即可。

18.当你声明常量或者变量并赋初值的时候类型推断非常有用。当你在声明常量或者变量的时候赋给它们一个字面量（literal value 或 literal）即可触发类型推断。整数默认推断为Int类型，浮点数默认推断为Double类型
``` swift
let meaningOfLife = 42
// meaningOfLife 会被推测为 Int 类型

let pi = 3.14159
// pi 会被推测为 Double 类型
```

19.整数字面量可以被写作：
> a.一个十进制数，没有前缀

> b.一个二进制数，前缀是0b

> c.一个八进制数，前缀是0o

> d.一个十六进制数，前缀是0x

``` swift
let decimalInteger = 17
let binaryInteger = 0b10001       // 二进制的17
let octalInteger = 0o21           // 八进制的17
let hexadecimalInteger = 0x11     // 十六进制的17
```

20.浮点字面量可以是十进制（没有前缀）或者是十六进制（前缀是0x）。小数点两边必须有至少一个十进制数字（或者是十六进制的数字）。浮点字面量还有一个可选的指数（exponent），在十进制浮点数中通过大写或者小写的e来指定，在十六进制浮点数中通过大写或者小写的p来指定。
>a. 如果一个十进制数的指数为exp，那这个数相当于基数和10^exp的乘积
>``` swift
1.25e2 表示 1.25 × 10^2，等于 125.0。
1.25e-2 表示 1.25 × 10^-2，等于 0.0125。
>```
>b.如果一个十六进制数的指数为exp，那这个数相当于基数和2^exp的乘积
>``` swift
0xFp2 表示 15 × 2^2，等于 60.0。
0xFp-2 表示 15 × 2^-2，等于 3.75
>```

21.数值类字面量可以包括额外的格式来增强可读性。整数和浮点数都可以添加额外的零并且包含下划线，并不会影响字面量:
``` swift
let paddedDouble = 000123.456
let oneMillion = 1_000_000
let justOverOneMillion = 1_000_000.000_000_1
```

22.通常来讲，即使代码中的整数常量和变量已知非负，也请使用Int类型。总是使用默认的整数类型可以保证你的整数常量和变量可以直接被复用并且可以匹配整数类字面量的类型推断。 只有在必要的时候才使用其他整数类型，比如要处理外部的长度明确的数据或者为了优化性能、内存占用等等。使用显式指定长度的类型可以及时发现值溢出并且可以暗示正在处理特殊数据。

23.要将一种数字类型转换成另一种，你要用当前值来初始化一个期望类型的新数字，这个数字的类型就是你的目标类型:
``` swift
let twoThousand: UInt16 = 2_000
let one: UInt8 = 1
let twoThousandAndOne = twoThousand + UInt16(one)
```

24.SomeType(ofInitialValue)是调用 Swift 构造器并传入一个初始值的默认方法。你并不能传入任意类型的值，只能传入UInt16内部有对应构造器的值。不过你可以扩展现有的类型来让它可以接收其他类型的值（包括自定义类型）。

25.结合数字类常量和变量不同于结合数字类字面量。字面量3可以直接和字面量0.14159相加，因为数字字面量本身没有明确的类型。它们的类型只在编译器需要求值的时候被推测。

26.类型别名（type aliases）就是给现有类型定义另一个名字。你可以使用typealias关键字来定义类型别名。当你想要给现有类型起一个更有意义的名字时，类型别名非常有用
``` swift
typealias AudioSample = UInt16
var maxAmplitudeFound = AudioSample.min
// maxAmplitudeFound 现在是 0
```

27.Swift 有一个基本的布尔（Boolean）类型，叫做Bool。布尔值指逻辑上的（logical），因为它们只能是真或者假。Swift 有两个布尔常量，true和false
``` swift
let orangesAreOrange = true
let turnipsAreDelicious = false
```

28.初始化常量或者变量的时候如果所赋的值类型已知，就可以触发类型推断，这让 Swift 代码更加简洁并且可读性更高。

29.元组（tuples）把多个值组合成一个复合值。元组内的值可以是任意类型，并不要求是相同类型
``` swift
let http404Error = (404, "Not Found")
// http404Error 的类型是 (Int, String)，值是 (404, "Not Found")
```
30.以将一个元组的内容分解（decompose）成单独的常量和变量，然后你就可以正常使用它们了。如果你只需要一部分元组值，分解的时候可以把要忽略的部分用下划线（_）标记。还可以通过下标来访问元组中的单个元素，下标从零开始。
``` swift
let (statusCode, statusMessage) = http404Error
println("The status code is \(statusCode)")
// 输出 "The status code is 404"
println("The status message is \(statusMessage)")
// 输出 "The status message is Not Found"

let (justTheStatusCode, _) = http404Error
println("The status code is \(justTheStatusCode)")
// 输出 "The status code is 404"

println("The status code is \(http404Error.0)")
// 输出 "The status code is 404"
println("The status message is \(http404Error.1)")
// 输出 "The status message is Not Found"
```

31.可以在定义元组的时候给单个元素命名，给元组中的元素命名后，你可以通过名字来获取这些元素的值
``` swift
let http200Status = (statusCode: 200, description: "OK")
println("The status code is \(http200Status.statusCode)")
// 输出 "The status code is 200"
println("The status message is \(http200Status.description)")
// 输出 "The status message is OK"
```

32.元组在临时组织值的时候很有用，但是并不适合创建复杂的数据结构。如果你的数据结构并不是临时使用，请使用类或者结构体而不是元组。

33.使用可选类型（optionals）来处理值可能缺失的情况。可选类型表示：
>a.有值，等于 x

>b.没有值

>Swift 的可选类型可以让你暗示任意类型的值缺失，并不需要一个特殊值。

34.toInt方法可能会失败，所以它返回一个可选类型（optional）Int，而不是一个Int。一个可选的Int被写作Int?而不是Int。问号暗示包含的值是可选类型，也就是说可能包含Int值也可能不包含值。
``` swift
let possibleNumber = "123"
let convertedNumber = possibleNumber.toInt()
// convertedNumber 被推测为类型 "Int?"， 或者类型 "optional Int"
```

35.以使用if语句来判断一个可选是否包含值。如果可选类型有值，结果是true；如果没有值，结果是false。当你确定可选类型确实包含值之后，你可以在可选的名字后面加一个感叹号（!）来获取值，这被称为可选值的强制解析（forced unwrapping）。
``` swift
if convertedNumber {
    println("\(possibleNumber) has an integer value of \(convertedNumber!)")
} else {
    println("\(possibleNumber) could not be converted to an integer")
}
// 输出 "123 has an integer value of 123"
```

36.使用!来获取一个不存在的可选值会导致运行时错误。使用!来强制解析值之前，一定要确定可选包含一个非nil的值。

37.使用可选绑定（optional binding）来判断可选类型是否包含值，如果包含就把值赋给一个临时常量或者变量。可选绑定可以用在if和while语句中来对可选类型的值进行判断并把值赋给一个常量或者变量。
``` swift
if let actualNumber = possibleNumber.toInt() {
    println("\(possibleNumber) has an integer value of \(actualNumber)")
} else {
    println("\(possibleNumber) could not be converted to an integer")
}
// 输出 "123 has an integer value of 123"
```

38.可以给可选变量赋值为nil来表示它没有值。nil不能用于非可选的常量和变量。如果你的代码中有常量或者变量需要处理值缺失的情况，请把它们声明成对应的可选类型。如果你声明一个可选常量或者变量但是没有赋值，它们会自动被设置为nil。
``` swift
var serverResponseCode: Int? = 404
// serverResponseCode 包含一个可选的 Int 值 404
serverResponseCode = nil
// serverResponseCode 现在不包含值

var surveyAnswer: String?
// surveyAnswer 被自动设置为 nil
```

39.Swift 的nil和 Objective-C 中的nil并不一样。在 Objective-C 中，nil是一个指向不存在对象的指针。在 Swift 中，nil不是指针——它是一个确定的值，用来表示值缺失。任何类型的可选状态都可以被设置为nil，不只是对象类型。

40.隐式解析可选类型（implicitly unwrapped optionals）。把想要用作可选的类型的后面的问号（String?）改成感叹号（String!）来声明一个隐式解析可选类型。当可选类型被第一次赋值之后就可以确定之后一直有值的时候，隐式解析可选类型非常有用。隐式解析可选类型主要被用在 Swift 中类的构造过程中。
``` swift
let possibleString: String? = "An optional string."
println(possibleString!) // 需要惊叹号来获取值
// 输出 "An optional string."

let assumedString: String! = "An implicitly unwrapped optional string."
println(assumedString)  // 不需要感叹号
// 输出 "An implicitly unwrapped optional string."
//一个隐式解析可选类型其实就是一个普通的可选类型，但是可以被当做非可选类型来使用，并不需要每次都使用解析来获取可选值。如果你在隐式解析可选类型没有值的时候尝试取值，会触发运行时错误。和你在没有值的普通可选类型后面加一个惊叹号一样
```

41.仍然可以把隐式解析可选类型当做普通可选类型来判断它是否包含值，也可以在可选绑定中使用隐式解析可选类型来检查并解析它的值
``` swift
if assumedString {
    println(assumedString)
}
// 输出 "An implicitly unwrapped optional string."

if let definiteString = assumedString {
    println(definiteString)
}
// 输出 "An implicitly unwrapped optional string."
```

42.如果一个变量之后可能变成nil的话请不要使用隐式解析可选类型。如果你需要在变量的生命周期中判断是否是nil的话，请使用普通可选类型。

43.在某些情况下，如果值缺失或者值并不满足特定的条件，你的代码可能并不需要继续执行。这时，你可以在你的代码中触发一个断言（assertion）来结束代码运行并通过调试来找到值缺失的原因。

44.断言会在运行时判断一个逻辑条件是否为true。从字面意思来说，断言“断言”一个条件是否为真。你可以使用断言来保证在运行其他代码之前，某些重要的条件已经被满足。如果条件判断为true，代码运行会继续进行；如果条件判断为false，代码运行停止，你的应用被终止。可以使用全局assert函数来写一个断言。向assert函数传入一个结果为true或者false的表达式以及一条信息，当表达式为false的时候这条信息会被显示
``` swift
let age = -3
assert(age >= 0, "A person's age cannot be less than zero")
// 因为 age < 0，所以断言会触发

//断言信息不能使用字符串插值，但断言信息可以省略
assert(age >= 0)
```

45.当条件可能为假时使用断言，但是最终一定要保证条件为真，这样你的代码才能继续运行。断言的适用情景：
>a.整数类型的下标索引被传入一个自定义下标脚本实现，但是下标索引值可能太小或者太大。

>b.需要给函数传入一个值，但是非法的值可能导致函数不能正常执行。

>c.一个可选值现在是nil，但是后面的代码运行需要一个非nil值。

46.断言可能导致你的应用终止运行，所以你应当仔细设计你的代码来让非法条件不会出现。然而，在你的应用发布之前，有时候非法条件可能出现，这时使用断言可以快速发现问题。






