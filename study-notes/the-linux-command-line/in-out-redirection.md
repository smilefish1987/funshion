# the linux command line 输入/输出/重定向

一.标准输入，输出，和错误
输出有两种：
>1.程序运行的结果

>2.我们得到状态和错误信息

与 Unix 主题“任何东西都是一个文件”保持一致，程序，比方说 ls，实际上把他们的运行结果 输送到一个叫做标准输出的特殊文件（经常用 stdout 表示），而它们的状态信息则送到另一个 叫做标准错误的文件（stderr）。默认情况下，标准输出和标准错误都连接到屏幕，而不是 保存到磁盘文件。除此之外，许多程序从一个叫做标准输入（stdin）的设备得到输入，默认情况下， 标准输入连接到键盘。
​

二.重定向标准输出
>1.重定向标准输出到另一个文件除了屏幕，我们使用 “>” 重定向符，其后跟着文件名
``` shell
$ls -l /usr/bin > ls-output.txt
```
[注：当我们使用 “>” 重定向符来重定向输出结果时，目标文件总是从开头被重写]
[如果我们需要删除一个文件内容（或者创建一个 新的空文件），可以使用这样的技巧:]
``` shell
$ > filename
```

>2.使用”»“操作符，将导致输出结果添加到文件内容之后。如果文件不存在，文件会 被创建，就如使用了”>”操作符
``` shell
$ls -l /usr/bin >> ls-output.txt
```

三.重定向标准错误

>1.重定向标准错误缺乏专用的重定向操作符。重定向标准错误，我们必须参考它的文件描述符。 一个程序可以在几个编号的文件流中的任一个上产生输出。然而我们必须把这些文件流的前 三个看作标准输入，输出和错误，shell 内部参考它们为文件描述符0，1和2，各自地。shell 提供 了一种表示法来重定向文件，使用文件描述符。因为标准错误和文件描述符2一样，我们用这种 表示法来重定向标准错误：
``` shell
$ ls -l /bin/usr 2 > ls-error.txt
```

[文件描述符”2”，紧挨着放在重定向操作符之前，来执行重定向标准错误到文件 ls-error.txt 任务]


四.重定向标准输出和错误到同一个文件

可能有这种情况，我们希望捕捉一个命令的所有输出到一个文件。为了完成这个，我们 必须同时重定向标准输出和标准错误。有两种方法来完成任务。第一个，传统的方法， 在旧版本 shell 中也有效:
``` shell
$ls -l /bin/usr > ls-output.txt 2>&1
```
[使用这种方法，我们完成两个重定向。首先重定向标准输出到文件 ls-output.txt，然后 重定向文件描述符2（标准错误）到文件描述符1（标准输出）使用表示法2>&1]
[注意重定向的顺序安排非常重要。标准错误的重定向必须总是出现在标准输出 重定向之后，要不然它不起作用]
bash 版本提供了第二种方法，更精简合理的方法来执行这种联合的重定向:

``` shell
$ls -l /bin/usr &> ls-output.txt
```
[我们使用单单一个表示法 &> 来重定向标准输出和错误到文件 ls-output.txt]


五.处理不需要的输出
我们不想要一个命令的输出结果，只想把它们扔掉。这种情况 尤其适用于错误和状态信息。系统为我们提供了解决问题的方法，通过重定向输出结果 到一个特殊的叫做”/dev/null”的文件。这个文件是系统设备，叫做位存储桶，它可以 接受输入，并且对输入不做任何处理.:
``` shell
$ls -l /bin/usr 2> /dev/null
```
[Unix 文化中的/dev/null 位存储桶是个古老的 Unix 概念，由于它的普遍性，它的身影出现在 Unix 文化的 许多部分。当有人说他/她正在发送你的评论到/dev/null，现在你应该知道那是 什么意思了]


六.重定向标准输入
> 1.cat － 连接文件
cat 命令读取一个或多个文件，然后复制它们到标准输出。cat 经常被用来显示简短的文本文件。因为 cat 可以 接受不只一个文件作为参数，所以它也可以用来把文件连接在一起。
如果 cat 没有给出任何参数，它会从标准输入读入数据，因为标准输入，默认情况下，连接到键盘。$cat ,然后输入东西
下一步，输入 Ctrl-d（按住 Ctrl 键同时按下”d”），来告诉 cat，在标准输入中， 它已经到达文件末尾（EOF）.
[我们可以使用这种行为来创建简短的文本文件,$cat > lazy_dog.txt 输入命令，然后输入要放入文件中的文本。记住，最后输入Ctrl+d]


> 2.$cat < lazy_dog.txt
使用“<”重定向操作符，我们把标准输入源从键盘改到文件 lazy_dog.tx。我们看到结果 和传递单个文件名作为参数的执行结果一样。把这和传递一个文件名参数作比较，尤其没有意义， 但它是用来说明把一个文件作为标准输入源


> 3.管道线
命令可以从标准输入读取数据，然后再把数据输送到标准输出，命令的这种能力被 一个 shell 特性所利用，这个特性叫做管道线。使用管道操作符”|”（竖杠），一个命令的 标准输出可以管道到另一个命令的标准输入：
``` shell
$ command1 | command2
$ ls -l /usr/bin | less
```

>4.过滤器
管道线经常用来对数据完成复杂的操作。有可能会把几个命令放在一起组成一个管道线。 通常，以这种方式使用的命令被称为过滤器。过滤器接受输入，以某种方式改变它，然后 输出它。第一个我们想试验的过滤器是 sort[排序].
>>4.1 sort 排序
``` shell
$ ls /bin  /usr/bin | sort | less
```

>>4.2 uniq - 报道或忽略重复行
uniq 命令经常和 sort 命令结合在一起使用。uniq 从标准输入或单个文件名参数接受数据有序 列表（详情查看 uniq 手册页），默认情况下，从数据列表中删除任何重复行。所以，为了确信 我们的列表中不包含重复句子（这是说，出现在目录/bin 和/usr/bin 中重名的程序），我们添加 uniq 到我们的管道线中.
``` shell
$ls /bin /usr/bin | sort | uniq | less
```

[如果我们想只看到重复的数据列表，让 uniq 命令带上”-d”选项]


>> 4.3 wc －打印行，字和字节数
wc（字计数）命令是用来显示文件所包含的行，字和字节数：
``` shell
$wc ls-output.txt
```
wc 打印出来三个数字：包含在文件 ls-output.txt 中的行数，单词数和字节数， 正如我们先前的命令，如果 wc 不带命令行参数，它接受标准输入。”-l”选项限制命令输出只能 报道行数。添加 wc 到管道线来统计数据，是个很便利的方法。查看我们的有序列表中程序个数.

``` shell
$ls /bin /usr/bin |sort | uniq | wc -l
```

>>4.4 grep －打印匹配行
grep 是个很强大的程序，用来找到文件中的匹配文本。这样使用 grep 命令：
``` shell
grep pattern [file...]
``` 

>>当 grep 遇到一个文件中的匹配”模式”，它会打印出包含这个类型的行。grep 能够匹配的模式可以 很复杂，但是现在我们把注意力集中在简单文本匹配上面.
``` shell
$ ls /bin /usr/bin | sort |uniq | grep zip
```
[grep 有－对方便的选项：”-i”导致 grep 忽略大小写当执行搜索时（通常，搜索是大小写 敏感的），”-v”选项会告诉 grep 只打印不匹配的行。]


>>4.5 head / tail －打印文件开头部分/结尾部分
head 命令打印文件的前十行，而 tail 命令打印文件的后十行。默认情况下，两个命令 都打印十行文本，但是可以通过”-n”选项来调整命令打印的行数.
``` shell
$ head -n 5 ls-output.txt
$ tail -n 5 ls-output.txt
$ ls /usr/bin | tail -n 5
```
[tail 有一个选项允许你实时的浏览文件,使用”-f”选项，tail 命令继续监测这个文件，当新的内容添加到文件后，它们会立即 出现在屏幕上。这会一直继续下去直到你输入 Ctrl-c]
``` shell
$ tail -f /var/log/message
```

>>4.6 tee － 从 Stdin 读取数据，并同时输出到 Stdout 和文件
为了和我们的管道隐喻保持一致，Linux 提供了一个叫做 tee 的命令，这个命令制造了 一个”tee”，安装到我们的管道上。tee 程序从标准输入读入数据，并且同时复制数据 到标准输出（允许数据继续随着管道线流动）和一个或多个文件。当在某个中间处理 阶段来捕捉一个管道线的内容时，这很有帮助。这里，我们重复执行一个先前的例子， 这次包含 tee 命令，在 grep 过滤管道线的内容之前，来捕捉整个目录列表到文件 ls.txt：
``` shell
$ ls /usr/bin | tee ls.txt | grep zip
```