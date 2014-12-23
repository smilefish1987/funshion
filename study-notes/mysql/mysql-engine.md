# mysql 引擎

1.查看数据表的创建命令：

``` sql
show create table test;
```

2.MySQL可以将数据以不同的技术存储在文件(内存)中，这种技术就称为存储引擎。每一种存储引擎使用不同的存储机制，索引技巧，锁定水平，最终提供广泛且不同的功能

3.存储引擎：
 >3.1. MyISAM

 >3.2 InnoDB

 >3.3 Memory

 >3.4 CSV

 >3.5 Archive

4.并发控制：当多个连接对记录进行修改时保证数据的一致性和完整性
 
5.锁：
>5.1 共享锁(读锁 ) ：在同一时间段内，多个用户可以读取同一个资源，读取过程中数据不会发生任何变化

>5.2 排他锁（写锁）: 在任何时候只能有一个用户写入资源，当进行写锁时会阻塞其他的读锁或者写锁操作。

6.锁的颗粒（粒度）: 锁策略和锁开销
>6.1 表锁：是一种开销最小的锁策略

>6.2 行锁: 是一种开销最大的锁策略

7.事务是区别数据库和文件系统的区别之一。事务：用于保证数据库的完整性

8.事务的特性：1. 原子性 2.一致性 3.隔离性  4.持久性

9.外键：是保证数据一致性的策略

10.索引：是对数据表中一列或多列的值进行排序的一种结构 (普通索引，唯一索引，全文索引，btree索引，hash索引)

11.各个存储引擎的特点：

| 特性 | MyISAM |  InnoDB | Memory |  Archive |
|:-------|:-----:|:-----:|:-----:|:-----:|
|存储限制|256TB | 64TB  |  有    |    无
|事务安全|- | 支持 | - | -  |
|支持索引| 支持 | 支持 | 支持 | － |
|锁颗粒|表锁 | 行锁 | 表锁 | 行锁 |
|数据压缩| 支持 |－ | － | 支持 |
|支持外键| － | 支持 |－ | － |

12.CSV存储引擎，将数据保持为CSV格式,不支持索引的

13.Blackhole：黑洞引擎，写入的数据都会消失，一般用于做数据复制的中继

14.MyISAM:使用于事务的处理不多的情况

15.InnoDB：适用于事务处理比较多，需要有外键支持的情况

16.设置存储引擎的方法：
>16.1.通过修改MySQL配置文件
``` sql
default-storage-engine = engine;
```

>2.通过创建数据表命令实现
``` sql
CREATE TABLE tbl_name(
...
)ENGINE = MyISAM;
```

>3.通过修改数据表命令实现：
``` sql
ALTER TABLE table_name ENGINE [=] engine_name;
``` 

17.InnoDB和MyISAM是许多人在使用MySQL时最常用的两个表类型，这两个表类型各有优劣，视具体应用而定。基本的差别为：MyISAM类型不支持事务处理等高级处理，而InnoDB类型支持。MyISAM类型的表强调的是性能，其执行数度比InnoDB类型更快，但是不提供事务支持，而InnoDB提供事务支持已经外部键等高级数据库功能。以下是一些细节和具体实现的差别：

>1. InnoDB不支持FULLTEXT类型的索引。

>2. InnoDB 中不保存表的具体行数，也就是说，执行select count(*) from table时，InnoDB要扫描一遍整个表来计算有多少行，但是MyISAM只要简单的读出保存好的行数即可。注意的是，当count(*)语句包含 where条件时，两种表的操作是一样的。

>3. 对于AUTO_INCREMENT类型的字段，InnoDB中必须包含只有该字段的索引，但是在MyISAM表中，可以和其他字段一起建立联合索引。

>4. DELETE FROM table时，InnoDB不会重新建立表，而是一行一行的删除。

>5. LOAD TABLE FROM MASTER操作对InnoDB是不起作用的，解决方法是首先把InnoDB表改成MyISAM表，导入数据后再改成InnoDB表，但是对于使用的额外的InnoDB特性(例如外键)的表不适用。

>6. InnoDB表的行锁也不是绝对的，假如在执行一个SQL语句时MySQL不能确定要扫描的范围，InnoDB表同样会锁全表，例如update table set num=1 where name like “%aaa%”

18.查看mysql服务器支持那些引擎：
``` sql
 $mysql> show engines;
```

下面是我5.5的mysql服务器的输出结果：

|Engine|Support|Comment	|Transactions|	XA |Savepoints|
|:-------|:-----:|:-----:|:-----:|:-----:|:-----:|
|MyISAM|YES|MyISAM|storage|engine|NO|NO|NO|
|InnoDB|DEFAULT|Percona-XtraDB, Supports transactions, row-level locking, and foreign keys|YES|YES|YES|
|MRG_MyISAM|YES|Collection of identical MyISAM tables|NO|NO|NO|
|MEMORY|YES|Hash based, stored in memory, useful for temporary tables|NO|NO|NO|
|CSV|YES|CSV storage engine|NO|NO|NO|
|PERFORMANCE_SCHEMA|YES|Performance Schema|NO|NO|NO|
|Aria|YES|Crash-safe tables with MyISAM heritage|NO|NO|NO|




