# the linux command line 从shell眼中看世界

1.echo － 显示一行文本

>1.1(字符)展开

>每一次你输入一个命令，然后按下 enter 键，在 bash 执行你的命令之前，bash 会对输入 的字符完成几个步骤处理。我们已经知道两三个案例，怎样一个简单的字符序列，例如”*”, 对 shell 来说，有很多的涵义。使这个发生的过程叫做（字符）展开。通过展开， 你输入的字符，在 shell 对它起作用之前，会展开成为别的字符。
[echo 是一个 shell 内部命令，来完成非常简单的任务。 它在标准输出中打印出它的文本参数]

>1.2 echo *

>为什么 echo 不打印”*”呢？随着你回想起我们所学过的 关于通配符的内容，这个”*”字符意味着匹配文件名中的任意字符，但是在原先的讨论 中我们却不知道 shell 是怎样实现这个功能的。最简单的答案就是shell把”*”展开成了 另外的东西（在这种情况下，就是在当前工作目录下的文件名字），在echo命令被执行 前。当回车键被按下时，shell在命令被执行前在命令行上自动展开任何符合条件的字符， 所以echo命令从不会发现”*”,只把它展开成结果.


>1.3路径名展开

>这种通配符工作机制叫做路径名展开。
``` shell
$ echo D*
$ echo *.s
$ echo [[:upper:]]*
$ echo /usr/*/share
```
[隐藏文件路径名展开
正如我们知道的，以圆点字符开头的文件名是隐藏文件。路径名展开也尊重这种 行为。像这样的展开：不会显示隐藏文件]

>1.4 波浪线展开
当它用在一个 单词的开头时，它会展开成指定用户的主目录名，如果没有指定用户名，则是当前用户的主目录.
``` shell
$ echo ~user_name
```

>1.5算术表达式展开
shell 允许算术表达式通过展开来执行。这允许我们把 shell 提示当作计算器来使用:
``` shell
$ echo $((2+2))
```
算术表达式展开使用这种格式：$((expression))
（以上括号中的）表达式是指算术表达式，它由数值和算术操作符组成。
算术表达式只支持整数（全部是数字，不带小数点），但是能执行很多不同的操作。这里是 一些

>它支持的操作符：

>|操作符 | 说明|
|:-----|:-------|
|+ |  加|
|- |   减|
|* | 乘 |
|/ | 除（但是记住，因为展开只是支持整数除法，所以结果是整数。）|
|% | 取余，只是简单的意味着，“余数”|
|**|  取幂|

>[在算术表达式中空格并不重要，并且表达式可以嵌套]
``` shell
$ echo $(($((5**2)) * 3))
```
[一对括号可以用来把多个子表达式括起来。通过这个技术，我们可以重写上面的例子， 同时用一个展开代替两个，来得到一样的结果]
``` shell
$ echo $(((5**2) * 3))
```

>1.6花括号展开
可能最奇怪的展开是花括号展开。通过它，你可以从一个包含花括号的模式中 创建多个文本字符串。
``` shell
$ echo Front-{A,B,C}-Back
```
[花括号展开模式可能包含一个开头部分叫做报头，一个结尾部分叫做附言。花括号表达式本身可 能包含一个由逗号分开的字符串列表，或者一系列整数，或者单个的字符串。这种模式可能 不包括嵌入的空白]
``` shell
$ echo {Z..A}
[花括号的展开可以嵌套]
$ echo a{A{1,2},B{3,4}}b
$ mkdir {2007..2009}-0{1..9} {2007..2009}-{10..12}
```

>1.7参数展开
``` shell
$ echo $USER
$ printenv | less
```
[通过参数展开，如果你拼写错了一个变量名， 展开仍然会进行，只是展成一个空字符串]


>1.8命令替换

>命令替换允许我们把一个命令的输出作为一个展开模式来使用：
``` shell
$echo $(ls)
$ls -l $(which cp)
[这里我们把 which cp 的执行结果作为一个参数传递给 ls 命令，因此要想得到 cp 程序的 输出列表，不必知道它完整的路径名。我们不只限制于简单命令。也可以使用整个管道线]
$ file $(ls /usr/bin/* | grep zip)
[在旧版 shell 程序中，有另一种语法也支持命令替换，可与刚提到的语法轮换使用。 bash 也支持这种语法。它使用倒引号来代替美元符号和括号：
$ls -l `which cp`
]
```

>1.9引用
``` shell
$echo this is a      test
$echo The total is $100.00
```
shell 提供了一种 叫做引用的机制，来有选择地禁止不需要的展开:
>>a.双引号

>>如果你把文本放在双引号中， shell 使用的特殊字符，除了 $，\ (反斜杠），和 `（倒引号）之外， 则失去它们的特殊含义，被当作普通字符来看待。这意味着单词分割，路径名展开， 波浪线展开，和花括号展开都被禁止，然而参数展开，算术展开，和命令替换 仍然执行。使用双引号，我们可以处理包含空格的文件名。

>>使用双引号，我们可以阻止单词分割，得到期望的结果；进一步，我们甚至可以修复 破损的文件名。【记住，在双引号中，参数展开，算术表达式展开，和命令替换仍然有效】
``` shell
$echo "$USER $((2+2)) $(cal)"
```
[在默认情况下，单词分割机制会在单词中寻找空格，制表符，和换行符，并把它们看作 单词之间的界定符。它们只作为分隔符使用。因为它们把单词分为不同的参数]
单词分割被禁止，内嵌的空格也不会被当作界定符，它们成为参数的一部分。 一旦加上双引号，我们的命令行就包含一个带有一个参数的命令。


>>事实上，单词分割机制把换行符看作界定符，对命令替换产生了一个，虽然微妙，但有趣的影响。 考虑下面的例子:
``` shell
$echo $(cal)
$echo "$(cal)"
```

>>b.单引号
如果需要禁止所有的展开，我们使用单引号。


>>c.转义字符

>>有时候我们只想引用单个字符。我们可以在字符之前加上一个反斜杠，在这个上下文中叫做转义字符。 经常在双引号中使用转义字符，来有选择地阻止展开。
``` shell
echo "The balance for user is: \$5.00"
[使用转义字符来消除文件名中一个字符的特殊含义，是很普遍的]
[为了允许反斜杠字符出现，输入”\”来转义。注意在单引号中，反斜杠失去它的特殊含义，它 被看作普通字符]
[echo 命令带上"-e"选项，能够解释转义序列。你可以把转义序列放在$' '里面, 以下例子，使用 sleep 命令，一个简单的程序，它会等待指定的秒数，然后退出。 我们可以创建一个简单的倒数计数器：
 sleep 10; echo -e "Time's up\a" 也可以这样
 sleep 10; echo "Time's up" $'\a'
]
```