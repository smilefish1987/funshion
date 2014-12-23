# MySQL 初见

1.修改mysql的提示符：
>a.登录的时候 mysql -uroot -proot --prompt 提示符

>b.登录之后通过prompt命令：
``` sql
	mysql>prompt 提示符
```

2.提示符的参数:
\D  完整的日期
\d  当前数据库
\h  服务器名称
\u  当前用户

3.mysql常用命令:
``` sql
SELECT VERSION();  显示当前服务器的版本
SELECT NOW();  显示当前日期时间
SELECT USER();  显示当前用户
```
4.MYSQL语句的规范
>a. 关键字与函数名称全部大写

>b.数据库名称，表名称，字段名称全部小写

>c.SQL语句必须以分号结尾

>d.mysql不区分大小写，但是应该按规范来操作

5.创建数据库
``` sql
CREATE {DATABASE | SCHEMA} [IF NOT EXISTS] db_name 
[DEFAUT] CHARACTER SET [=] charset_name
```

6.修改数据库
``` sql
ALTER {DATABASE | SCHEMA} [db_name] 
[DEFAULT] CHARACTER SET [=] charset_name
```

7.查看数据库的信息
``` sql
SHOW CREATE DATABASE db_name;
```

8.查看系统的数据库
``` sql
SHOW DATABASES like 'pattern'
```

9.删除数据库
``` sql
DROP {DATABASE | SCHEMA} [IF EXISTS] db_name
```

10.查看数据库的警告信息
``` sql
SHOW WARNINGS;
``` 