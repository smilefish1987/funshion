# MySQL 表

1.查看当前那个数据库被打开：
``` sql
SELECT DATABASE();
```

2.创建表
``` sql
CREATE TABLE [IF NOT EXISTS] table_name(
	column_name data_type,
	....
)

CREATE TABLE tb1(
	username VARCHAR(20),
	age    TINYINT UNSIGNED,
        salary  FLOAT(8,2) UNSIGNED
);
```

3.查看数据表：
``` sql
SHOW TABLES [FROM db_name]
[LIKE `pattern` | WHERE expr]
```

4.查看数据表的结构：
``` sql
SHOW COLUMNS FROM tbl_name;
```

5.插入记录
``` sql
INSERT [INTO] tbl_name [(col_name,...)] VALUES (val1,...)
```

6.查找记录：
``` sql
SELECT expr,.. FROM tbl_name
```

7.空值与非空值
>NULL , 字段值可以为空

>NOT NULL,字段值禁止为空


8.AUTO_INCREMENT
>1.自动编号，且必须与主键组合使用

>2.默认情况下，起始值为1，每次的增量为1


9.PRIMARY KEY
>1.主键约束

>2.每张数据表只能存在一个主键

>3.主键保证记录的唯一性

>4.主键自动为NOT NULL

10.UNIQUE KEY
>1.唯一约束

>2.唯一约束可以保证记录的唯一性

>3.唯一约束的字段可以为空值（NULL）

>4.每张数据表可以存在多个唯一约束

11.DEFAULT
> 1.默认值

> 2.当插入记录时，如果没有明确为字段赋值，则自动赋予默认值

12.约束
>1.约束保证数据的完整性和一致性
>2.约束分为表级约束和列级约束
>3.约束类型包括：
>>a. NOT NULL (非空约束)

>>b.PRIMARY KEY (主键约束)

>>c.UNIQUE KEY (唯一约束)

>>d.DEFAULT （默认约束）

>>e.FOREIGN KEY (外键约束)

13.FOREIGN KEY
>1.保持数据一致性，完整性

>2.实现一对一或一对多关系

14.FOREIGN KEY 的要求(拥有外键的表为子表，子表参照的表为父表)
>1.父表和子表必须使用相同的存储引擎，而且禁止使用临时表

>2.数据表的存储引擎只能为InnoDB

>3.外键列和参照列必须具有相似的数据类型，其中数字的长度或者是否有符号位必须相同；而字符的长度则可以不同

>4.外键列和参照列必须创建索引。如果外键列不存在索引的话，MySQL将自动创建索引。

15.编译数据表的默认存储引擎,MySQL配置文件
``` sql
default-storage-engine = INNODB
```

16.创建外键约束：
``` sql
FOREIGN KEY (pid) REFERENCES tbl_name(pid)
```

17.查看表的索引
``` sql
show indexes from tbl_name;
```

18.外键约束的参照，用于父表的操作时子表的相应操作
>1.CASCADE: 从父表删除或更新且自动删除或更新子表中匹配的行

>2.SET NULL :从父表删除或更新列，并设置子表中的外键列位NULL。如果使用该选项，必须保证子表列没有指定NOT NULL

>3.RESTRICT:拒绝对父表的删除或更新操作

>4.NO ACTION: 标准SQL的关键字，在MySQL中与RESTRICT相同

19.表级约束与列级约束
>1.对一个数据列建立的约束，称为列级约束

>2.对多个数据列建立的约束，称为表级约束

>3.列级约束既可以在列定义时声明，也可以在列定义后声明，表级约束只能在列定义后声明。

>4.NOT NULL 和DEFUALT只存在列级约束

20.修改数据表

20.1添加单列
``` sql
ALTER TABLE tbl_name ADD [COLUMN] col_name column_definition 
[FIRST | AFTER col_name]
```

20.2添加多列
``` sql
ALTER TABLE tbl_name ADD [COLUMN]
(col_name column_definition,...)
```

20.3删除列
``` sql
ALTER TABLE tbl_name DROP [COLUMN] col_name;
ALTER TABLE tbl_name DROP [COLUMN] col_name1,DROP [COLUMN] col_name2;
```

20.4添加主键约束
``` sql
ALTER TABLE tbl_name ADD [CONSTRAINT [symbol]] PRIMARY KEY [index_type](index_col_name,...)
```

20.5添加唯一约束
``` sql
ALTER TABLE tbl_name ADD [CONSTRAINT [symbol]] UNIQUE [INDEX|KEY] [index_name] [index_type]
(index_col_name,...)
```

20.6添加外键约束
``` sql
ALTER TABLE tbl_name ADD [CONSTRAINT [symbol]] FOREIGN KEY [index_name] (index_col_name,...)
reference_definition
```

20.7添加／删除默认约束
``` sql
ALTER TABLE tbl_name ALTER [COLUMN] col_name {SET DEFAULT litera_value | DROP DEFAULT}
```

20.8删除主键约束
``` sql
ALTER TABLE tbl_name DROP PRIMERY KEY
```

20.9删除唯一约束
``` sql
ALTER TABLE tbl_name DROP {KEY|INDEX} index_name
```

20.10.删除外键约束
``` sql
ALTER TABLE tbl_name DROP FOREIGN KEY fk_symbol
```

20.11修改列定义
``` sql
ALTER TABLE tbl_name MODIFY [COLUMN] col_name column_definition [FIRST | AFTER col_name]
```

20.12修改列的名称
``` sql
ALTER TABLE tbl_name CHANGE [COLUMN] old_col_name new_col_name column_definition [FIRST | AFTER col_name] 
```

20.13修改数据表的名字
``` sql
ALTER TABLE tbl_name RENAME [TO|AS] new_tbl_name

RENAME TABLE tbl_name TO new_tbl_name [,tbl_name2 TO new_tbl_name2] ...
```









