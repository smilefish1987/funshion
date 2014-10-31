# swift学习笔记  初见

1.Swift 是一种新的编程语言，用于编写 iOS 和 OS X 应用。swift的特性：函数多返回值，闭包，范型

2.Swift的源码文件以.swift为后缀。

3.Swift可以当脚本使用，不需要main函数，局作用域中的代码会被自动当做程序的入口点。语句结尾的(;)分号也不是必须，并且不建议书写。像下面这样:

``` swift
//文件hello.swift
println("Hello,world")
```

4.使用 **let** 来声明常量，使用 **var** 来声明变量。一个常量的值，在编译的时候，并不需要有明确的值，但是你只能为它赋值一次。常量就是这样一个值：你只需要决定一次，但是需要使用很多次。

5.变量或者常量声明时类型是可选的，声明的同时赋值的话，编译器会自动推断类型。如果初始值没有提供足够的信息（或者没有初始值），那你需要在变量后面声明类型，用冒号分割。

``` swift
var age = 23
age = 24
let name = "smile"
let pi:Double = 3.14
```

6.值永远不会被隐式转换为其他类型。如果你需要把一个值转换成其他类型，请显式转换
``` swift
let label = "the age is "
let age = 23
//通过类型初始化转换
let msg = label + String(age)
//值转换成字符串的方法：把值写到括号中，并且在括号之前写一个反斜杠
let msg2 = "the age is \(age) !"
```

7.使用方括号[]来创建数组和字典的字面量，并使用下标或者键（key）来访问元素
```swift
//数组
var shoppingList = ["catfish", "water", "tulips", "blue paint"]
shoppingList[1] = "bottle of water"

//字典
var occupations = [
    "Malcolm": "Captain",
    "Kaylee": "Mechanic",
]
occupations["Jayne"] = "Public Relations"
```

8.要创建一个空数组或者字典，使用初始化语法：
``` swift
//字符串数组
let emptyArray = String[]()

//键类型为string,值类型为Float的字典
let emptyDictionary = Dictionary<String, Float>()
```

9.如果类型信息可以被推断出来，你可以用[]和[:]来创建空数组和空字典，其实也就是如果上下文已经声明过变量，可以使用简单的语法将其设置为空值再使用。
``` swift
//数组
var shoppingList = ["catfish", "water", "tulips", "blue paint"]
shoppingList[1] = "bottle of water"

//字典
var occupations = [
    "Malcolm": "Captain",
    "Kaylee": "Mechanic",
]
occupations["Jayne"] = "Public Relations"

//将两者置为空
shoppingList = []
occupations = [:]

```

10.使用if和switch来进行条件操作，使用for-in、for、while和do-while来进行循环。包裹条件和循环变量括号可以省略，但是语句体的大括号是必须的。
``` swift
let individualScores = [75, 43, 103, 87, 12]
var teamScore = 0
for score in individualScores {
    if score > 50 {
        teamScore += 3
    } else {
        teamScore += 1
    }
}
println(teamScore)
```

11.在if语句中，条件必须是一个布尔表达式。

12.有些变量的值是可选的。一个可选的值可能是一个具体的值或者是nil，表示值缺失。在类型后面加一个问号来标记这个变量的值是可选的。可以一起使用if和let来处理值缺失的情况。
``` swift
var optionalName: String? = "Hello"
//如果变量的可选值是nil，条件会判断为false;如果不是nil，会将值赋给let后面的常量，这样代码块中就可以使用这个值了
if let name = optionalName {  //当可选值不为nil时执行if
    println("Hello, \(name)")
}
```

13.switch支持任意类型的数据以及各种比较操作——不仅仅是整数以及测试相等
``` swift
let vegetable = "red pepper"
switch vegetable {
case "celery": //传统的c-switch
    let vegetableComment = "Add some raisins and make ants on a log."
case "cucumber", "watercress": //可以使用逗号(,)匹配多个值中的一个
    let vegetableComment = "That would make a good tea sandwich."
case let x where x.hasSuffix("pepper"): //使用where语句支持更复杂的逻辑
    let vegetableComment = "Is it a spicy \(x)?"
default:  //defalut是必须的
    let vegetableComment = "Everything tastes good in soup."
}
```

14.运行switch中匹配到的子句之后，程序会退出switch语句，并不会继续向下运行，所以不需要在每个子句结尾写break。

15.可以使用for-in来遍历字典，需要两个变量来表示每个键值对
``` swift
let interestingNumbers = [
    "Prime": [2, 3, 5, 7, 11, 13],
    "Fibonacci": [1, 1, 2, 3, 5, 8],
    "Square": [1, 4, 9, 16, 25],
]
var largest = 0
for (kind, numbers) in interestingNumbers {
    for number in numbers {
        if number > largest {
            largest = number
        }
    }
}
println(largest)
```

16.使用while来重复运行一段代码直到不满足条件。循环条件可以在开头也可以在结尾
``` golang
var index = 2
while index < 100 {
    index = index * 2
}
println(index)

var doindex = 2
do {
    doindex = doindex * 2
} while doindex < 100
println(doindex)
```

17.可以在循环中使用..来表示范围,不包括上界；可以使用...表示范围，包括上下界
``` swift
var sum = 0
for i in 0..3 {
    sum += i
}
println(sum) //3

var sum = 0
for i in 0...3 {
    sum += i
}
println(sum) //6
```

18.使用func来声明一个函数，使用名字和参数来调用函数。使用->来指定函数返回值。函数可以返回多个值，也可以带可变个数的参数，当然也可以嵌套。
``` swift
func greet(name: String, day: String) -> String {
    return "Hello \(name), today is \(day)."
}

//使用一个元组来返回多个值
func getGasPrices() -> (Double, Double, Double) {
    return (3.59, 3.69, 3.79)
}

//可变参数函数
func sumOf(numbers: Int...) -> Int {
    var sum = 0
    //参数在函数内表现为数组的形式
    for number in numbers {
        sum += number
    }
    return sum
}

//嵌套函数
func returnFifteen() -> Int {
    var y = 10
    func add() {
        //被嵌套的函数可以访问外侧函数的变量，你可以使用嵌套函数来重构一个太长或者太复杂的函数
        y += 5
    }
    add()
    return y
}

//函数的调用：函数名和实参
greet("fish", "Tuesday")
getGasPrices()
sumOf()
sumOf(1, 2, 3)
returnFifteen()
```

19.函数是第一等类型,既可以作为参数，也可以作为返回值
``` swift
//函数作为另一个函数的参数
func hasAnyMatches(list: Int[], condition: Int -> Bool) -> Bool {
    for item in list {
        if condition(item) {
            return true
        }
    }
    return false
}
func lessThanTen(number: Int) -> Bool {
    return number < 10
}
var numbers = [20, 19, 7, 12]
hasAnyMatches(numbers, lessThanTen)

//函数作为返回值
func makeIncrementer() -> (Int -> Int) {
    func addOne(number: Int) -> Int {
        return 1 + number
    }
    return addOne
}
var increment = makeIncrementer()
increment(7)
```
20.使用{}来创建一个匿名闭包。使用in将参数和返回值类型声明与闭包函数体进行分离。函数实际上是一种特殊的闭包。
``` swift
numbers.map({
    (number: Int) -> Int in
    let result = 3 * number
    return result
})

//如果一个闭包的类型已知，比如作为一个回调函数，你可以忽略参数的类型和返回值。单个语句闭包会把它语句的值当做结果返回。
numbers.map({ number in 3 * number })
//通过参数位置而不是参数名字来引用参数——这个方法在非常短的闭包中非常有用。当一个闭包作为最后一个参数传给一个函数的时候，它可以直接跟在括号后面
sort([1, 5, 3, 12, 2]) { $0 > $1 }

```

21.使用class和类名来创建一个类。类中属性的声明和常量、变量声明一样，唯一的区别就是它们的上下文是类。同样，方法和函数声明也一样。
``` swift
//类的定义
class NamedShape {
    //属性
    var numberOfSides: Int = 0
    var name: String
    let age:Int = 23

    //构造函数，完成属性的初始化
    init(name: String) {
        //self用于区分实例变量和函数的参数
        self.name = name
    }
    
    //析构函数，不需要参数列表的小括号
    deinit{
        println("deinit")
    }

    //类的方法，用于类属性的操作
    func simpleDescription() -> String {
        return "A shape with \(numberOfSides) sides."
    }
}

//要创建一个类的实例，在类名后面加上括号
var shape = NamedShape(name:"fish")
//用点语法来访问实例的属性和方法
shape.numberOfSides = 7
var shapeDescription = shape.simpleDescription()

var opshape:NamedShape? = NamedShape(name:"smile")
//可选值需要在变量后面添加!来引用值
opshape!.numberOfSides = 4
println(opshape!.simpleDescription())
opshape = nil //调用析构函数
```
22.如果你需要在删除对象之前进行一些清理工作，使用deinit创建一个析构函数

23.子类的定义方法是在它们的类名后面加上父类的名字，用冒号分割。创建类的时候并不需要一个标准的根类，所以你可以忽略父类。子类如果要重写父类的方法的话，需要用override标记——如果没有添加override就重写父类方法的话编译器会报错。编译器同样会检测override标记的方法是否确实在父类中。
``` swift
//定义子类Square继承 NamedShape
class Square: NamedShape {
    //定义子类自己的属性
    var sideLength: Double
    
    //子类的构造方法
    init(sideLength: Double, name: String) {
        //先初始化子类自己的属性
        self.sideLength = sideLength
        //调用父类的构造方法
        super.init(name: name)
        //初始化从父类继承类的属性
        numberOfSides = 4
    }
    
    //定义子类自己的方法
    func area() ->  Double {
        return sideLength * sideLength
    }
    //通过override关键字重写父类的方法
    override func simpleDescription() -> String {
        return "A square with sides of length \(sideLength)."
    }
}
let test = Square(sideLength: 5.2, name: "my test square")
println(test.area())  //27.04
println(test.simpleDescription()) // A square with sides of length 5.2. 
```
24.属性可以有 getter 和 setter
``` swift
class EquilateralTriangle: NamedShape {
    var sideLength: Double = 0.0

    init(sideLength: Double, name: String) {
        //设置子类声明的属性值
        self.sideLength = sideLength
        //调用父类的构造器
        super.init(name: name)
        //改变父类定义的属性值。其他的工作比如调用方法、getters和setters也可以在这个阶段完成
        numberOfSides = 3
    }
    
    //计算属性perimeter定义了getter和setter
    var perimeter: Double {
    get {
        return 3.0 * sideLength
    }
    set {//setter默认情况下外部传入的参数是newValue
        sideLength = newValue / 3.0
    }
    }

    override func simpleDescription() -> String {
        return "An equilateral triagle with sides of length \(sideLength)."
    }
}
var triangle = EquilateralTriangle(sideLength: 3.1, name: "a triangle")
triangle.perimeter  // 调用perimeter属性的getter
triangle.perimeter = 9.9  // 调用perimeter属性的setter
triangle.sideLength
```

25.属性观察器，如果你不需要计算属性，但是仍然需要在设置一个新值之前或者之后运行代码，使用willSet和didSet

``` swift
class TriangleAndSquare {
    var triangle: EquilateralTriangle {
    willSet {
        square.sideLength = newValue.sideLength
    }
    }
    var square: Square {
    willSet { //willSet表示将要设置属性时，调用的观察器，新值的默认名称为newValue;didSet表示已经设置完的那一刻调用的观察器，属性的旧值的默认名称为oldValue
        triangle.sideLength = newValue.sideLength
    }
    }
    init(size: Double, name: String) {
        square = Square(sideLength: size, name: name)
        triangle = EquilateralTriangle(sideLength: size, name: name)
    }
}
var triangleAndSquare = TriangleAndSquare(size: 10, name: "another test shape")
triangleAndSquare.square.sideLength
triangleAndSquare.triangle.sideLength
triangleAndSquare.square = Square(sideLength: 50, name: "larger square")
triangleAndSquare.triangle.sideLength
```

26.类中的方法和一般的函数有一个重要的区别，函数的参数名只在函数内部使用，但是方法的参数名需要在调用的时候显式说明（除了第一个参数）。默认情况下，方法的参数名和它在方法内部的名字一样，不过你也可以定义第二个名字，这个名字被用在方法内部。
``` swift
class Counter {
    var count: Int = 0
    //方法的第二个参数有两个名字，内部名字times,外部名字numberOfTimes
    func incrementBy(amount: Int, numberOfTimes times: Int) {
        count += amount * times
    }
}
var counter = Counter()
//对于指定了外部参数名字的方法的调用，必须带指定的外部参数名
counter.incrementBy(2, numberOfTimes: 7)
```

27.处理变量的可选值时，你可以在操作（比如方法、属性和子脚本）之前加?。如果?之前的值是nil，?后面的东西都会被忽略，并且整个表达式返回nil。否则，?之后的东西都会被运行。在这两种情况下，整个表达式的值也是一个可选值。
``` swift
let optionalSquare: Square? = Square(sideLength: 2.5, name: "optional square")
let sideLength = optionalSquare?.sideLength
```

28.使用enum来创建一个枚举。就像类和其他所有命名类型一样，枚举可以包含方法
``` swift
//枚举Rank的原始值类型是Int,需要设置第一个原始值。剩下的原始值会按照顺序赋值。你也可以使用字符串或者浮点数作为枚举的原始值.
enum Rank: Int {
    case Ace = 1
    case Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten
    case Jack, Queen, King
    func simpleDescription() -> String {
        switch self {
        case .Ace:
            return "ace"
        case .Jack:
            return "jack"
        case .Queen:
            return "queen"
        case .King:
            return "king"
        default:
            return String(self.toRaw())
        }
    }
}
let ace = Rank.Ace
let aceRawValue = ace.toRaw()

//使用toRaw和fromRaw函数来在原始值和枚举值之间进行转换
if let convertedRank = Rank.fromRaw(3) {
    let threeDescription = convertedRank.simpleDescription()
}
```

29.枚举的成员值是实际值，并不是原始值的另一种表达方法。实际上，如果原始值没有意义，你不需要设置。有两种方式可以引用枚举成员：给枚举常量赋值时，枚举成员Suit.Hearts需要用全名来引用，因为常量没有显式指定类型。在switch里，枚举成员使用缩写.Hearts来引用，因为self的值已经知道是一个suit。已知变量类型的情况下你可以使用缩写。
``` swift
//没有指定原始值
enum Suit {
    case Spades, Hearts, Diamonds, Clubs
    func simpleDescription() -> String {
        switch self {
        //在已知类型的情况下，可以使用缩写的形式
        case .Spades:
            return "spades"
        case .Hearts:
            return "hearts"
        case .Diamonds:
            return "diamonds"
        case .Clubs:
            return "clubs"
        }
    }

}
//通过全名给枚举常量的赋值
let hearts = Suit.Hearts
let heartsDescription = hearts.simpleDescription()
```

30.使用struct来创建一个结构体。结构体和类有很多相同的地方，比如方法和构造器。它们之间最大的一个区别就是 结构体是传值，类是传引用。
``` swift
struct Card {
    var rank: Rank
    var suit: Suit
    func simpleDescription() -> String {
        return "The \(rank.simpleDescription()) of \
        (suit.simpleDescription())"
    }
}
let threeOfSpades = Card(rank: .Three, suit: .Spades)
let threeOfSpadesDescription = threeOfSpades.simpleDescription()
```

31.一个枚举成员的实例可以有实例值。相同枚举成员的实例可以有不同的值。创建实例的时候传入值即可。实例值和原始值是不同的：枚举成员的原始值对于所有实例都是相同的，而且你是在定义枚举的时候设置原始值。
``` swift
enum ServerResponse {
    case Result(String, String)
    case Error(String)
}

let success = ServerResponse.Result("6:00 am", "8:09 pm")
let failure = ServerResponse.Error("Out of cheese.")

switch success {
//通过命名实例值的名字，提取枚举的实例值
case let .Result(sunrise, sunset):
    let serverResponse = "Sunrise is at \(sunrise) and sunset is at \(sunset)."
case let .Error(error):
    let serverResponse = "Failure...  \(error)"
}
```

32.使用protocol来声明一个协议(协议是方法的集合)，类、枚举和结构体都可以实现协议。你可以像使用其他命名类型一样使用协议名,当你处理类型是协议的值时，协议外定义的方法不可用。
``` swift
protocol ExampleProtocol {
    var simpleDescription: String { get }
    mutating func adjust() //mutating关键字用于值类型的方法修改实例的属性
}

class SimpleClass: ExampleProtocol {
    var simpleDescription: String = "A very simple class."
    var anotherProperty: Int = 69105
    //类不需要mutating关键字
    func adjust() {
        simpleDescription += "  Now 100% adjusted."
    }
}
var a = SimpleClass()
a.adjust()
let aDescription = a.simpleDescription
//像其他类型一样使用协议类型
let protocolValue: ExampleProtocol = a
protocolValue.simpleDescription
//协议外定义的方法不可用和属性不能使用
// protocolValue.anotherProperty  // Uncomment to see the error

struct SimpleStructure: ExampleProtocol {
    var simpleDescription: String = "A simple structure"
    //mutating关键字用来标记一个会修改结构体的方法
    mutating func adjust() {
        simpleDescription += " (adjusted)"
    }
}
var b = SimpleStructure()
b.adjust()
let bDescription = b.simpleDescription
```

33.使用extension来为现有的类型添加功能，比如新的方法和参数。你可以使用扩展来改造定义在别处，甚至是从外部库或者框架引入的一个类型，使得这个类型遵循某个协议。
``` swift
extension Int: ExampleProtocol {
    var simpleDescription: String {
    return "The number \(self)"
    }
    mutating func adjust() {
        self += 42
    }
}
7.simpleDescription
```

34.在尖括号里写一个名字来创建一个泛型函数或者类型
``` swift
//范型函数
func repeat<ItemType>(item: ItemType, times: Int) -> ItemType[] {
    var result = ItemType[]()
    for i in 0..times {
        result += item
    }
    return result
}
repeat("knock", 4)

//范型类型
// Reimplement the Swift standard library's optional type
enum OptionalValue<T> {
    case None
    case Some(T)
}
var possibleInteger: OptionalValue<Int> = .None
possibleInteger = .Some(100)

```

35.在类型名后面使用where来指定对类型的需求，比如，限定类型实现某一个协议，限定两个类型是相同的，或者限定某个类必须有一个特定的父类。
``` swift
func anyCommonElements <T, U where T: Sequence, U: Sequence, T.GeneratorType.Element: Equatable, T.GeneratorType.Element == U.GeneratorType.Element> (lhs: T, rhs: U) -> Bool {
    for lhsItem in lhs {
        for rhsItem in rhs {
            if lhsItem == rhsItem {
                return true
            }
        }
    }
    return false
}
anyCommonElements([1, 2, 3], [3])

//<T: Equatable>和<T where T: Equatable>是等价的
```









