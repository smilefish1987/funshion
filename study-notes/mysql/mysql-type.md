# MySQL 数据类型

1.整数类型

|类型|说明|字节数|
|:-------|:-----:|:-----:|
|TINYINT | －128～127 ｜ 0 ～255 |  1个字节 |
|SMALLINT | -32768~32767 ｜ 0 ~ 65535 | 2个字节|
|MEDIUMINT|                |3个字节|
|INT |        |4个字节|
|BIGINT  |     |8个字节|


2.浮点类型

|类型|说明|
|:-------|-----:|
|FLOAT  |7位小数位|
|DOUBLE | 不建议使用存储范围大|

3.时间类型

|类型|字节数|
|:-------|-----:|
|YEAR |  1个字节 |
|TIME |  3个字节 |
|DATE |  3个字节 |
|DATETIME | 8个字节|
|TIMESTAMP|  4个字节|

4.字符类型

|类型|说明|
|:-------|:-----|
|CHAR(M) | M个字节 0<=M<=255|
|VARCHAR(M) | L+1个字节 |
|TINYTEXT | L+1个字节 其中L<2的8次方 |
|TEXT  | L+2个字节 ，其中L<2的16次方 |
|MEDIUMTEXT | L+3个字节,其中L<2的24次方 |
|LONGTEXT | L+4个字节, 其中L<2的32次方 |
|ENUM('value1','value2') | 1或2个字节，取决与枚举值的个数（最多65535个）|
|SET('value1','value2') | 1,2,3,4或者8个字节取决于set的大小（最多64个）|