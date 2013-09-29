回顾会

h1. Sprint 923~929回顾会

参与者：何锵锵、朱文君、郑晓春、程品、阮圆、曾广剑、刘水旺
会议记录：阮圆

h2. 上个迭代问题回顾

* 有些业务的逻辑要到测试的时候才知道是有问题【在story有依赖关系的时候，没有一个全局的考虑】---没遇到
* 数据库的脏数据，影响测试 ---遇到了,并解决
* 原有功能bug导致影响进行新有功能的开发和测试---没遇到
* 开发过程中需求有变动或者需求前后矛盾，开发工作量翻倍，影响心情---没遇到 
* scrum works 的task状态的更新不及时---遇到了，已经喝酸奶
* 迭代story TDR的上传不及时---已解决
* 迭代story上线记录不及时--- 已解决

h2. 本迭代燃尽图

![](images/borndown.png?raw=true)


h2. 本迭代story概括

<table class="table table-condensed table-bordered">
	<tbody>
		<tr>
			<th>编号</th>
			<th>story</th>
			<th>预计点数</th>
			<th>实际点数</th>
			<th>开发人员</th>
			<th>进度</th>
			<th>bug</th>
			<th>备注</th>
		</tr>
		<tr>
			<td>1</td><td>twsee-213 UGC视频增加平台选择及地域策略</td><td>8</td></td>0</td><td>何锵锵、程品</td><td>testing</td><td>0</td>
		</tr>
		<tr>
			<td>2</td><td>twsee-223 UGC视频题材列表页</td><td>2</td><td>0</td><td>何锵锵、程品</td><td>done</td><td>0</td>
		</tr>
		<tr>
			<td>3</td><td>twsee-217 UGC/体育地域策略显示方式</td><td>3</td><td>0</td><td>刘水旺、阮圆</td><td>block</td><td>0</td>
		</tr>
		<tr>
			<td>4</td><td>twsee-219 玩转风行移动版相关修正</td><td>3</td><td>0</td><td>刘水旺、程品</td><td>done</td><td>0</td>
		</tr>
		<tr>
			<td>5</td><td>twsee-218 去别人主页@用户发布的动态信息，需要展现在bms动态管理中</td><td>2</td><td>0</td><td>阮圆、朱文君</td><td>done</td><td>1</td>
		</tr>
		<tr>
			<td>6</td><td>twsee-220 将社区头图更换进入bms管理</td><td>5</td><td>0</td><td>程品、郑晓春</td><td>done</td><td>1</td>
		</tr>
		<tr>
			<td>7</td><td>twsee-221 更换头像/头图时保留用户当前图片</td><td>5</td><td>0</td><td>朱文君、阮圆</td><td>done</td><td>0</td>
		</tr>
		<tr>
			<td>8</td><td>twsee-227 网站订阅应用接入-用户订阅媒体</td><td>3</td><td>0</td><td>何锵锵、程品</td><td>done</td><td>0</td>
		</tr>
		<tr>
			<td>9</td><td>twsee-226 网站订阅应用接入-取消订阅</td><td>3</td><td>0</td><td>何锵锵、程品</td><td>done</td><td>0</td>
		</tr>
		<tr>
			<td>10</td><td>twsee-225 网站订阅应用接入-获取用户订阅列表</td><td>4</td><td>0</td><td>何锵锵、阮圆</td><td>done</td><td>0</td>
		</tr>
		<tr>
			<td>11</td><td>twsee-224 网站订阅应用接入-用户是否订阅、用户订阅数</td><td>3</td><td>0</td><td>何锵锵、阮圆</td><td>done</td><td>0</td>
		</tr>
		<tr>
			<td>12</td><td>twsee-228 点评相关api性能优化调研</td><td>3</td><td>0</td><td>朱文君</td><td>done</td><td>0</td>
		</tr>
		<tr>
			<td>总数</td><td>story：12</td><td>点数：44</td><td></td><td></td><td>未完成story：2</td><td>bug总数：2</td>
		</tr>
	</tbody>
</table>

h2. 本迭代总结


h3. 需保持

* story完成顺利
* bug比较少
* scrumworks 更新比较及时



h3. 待改进

* 迭代内story之间有依赖，影响测试进度 +1
* 复杂的不熟悉的故事，需要调研 +2
* 估点不准 +1
* TDR不细
* 工作量没有体现

h3. 问题解决

* 迭代内story之间有依赖，影响测试进度 +1
** 解除依赖以后，才安排进迭代
** 相互有依赖的故事由一个人take

* 复杂的不熟悉的故事，影响开发进度 +2
** 前期安排调研任务
** 定义复杂的不熟悉的故事
*** 开发估点时间超过5个点
*** story涉及到的存储结构不确定

* 估点不准 +1
** 采取翻牌的制度
** 琐事太多

* TDR不细
** 大家都参与到TDR的细化中去
** TDR应该细化到表，字段 
** TDR写的时候，各种情况case尽量穷尽

* 工作量没有体现
** 添加常规story

h2. bug分析

|开发人员|story总数|TDR内bug|TDR外bug|bug总数|公网bug|
|何锵锵|6|0|0|0|0|
|朱文君|3|0|1|0|0|
|郑晓春|1|0|1|0|0|
|刘水旺|2|0|0|0|0|

h2. 下个迭代会议主持：阮圆