# GinBlog

# 前后端交互

定义好API文档，错误的参数处理(比如用户名冲突)。<br>

通过Json格式传输数据。 <br>

```Go
ctx.JSON(http.StatusOK, gin.H{
	"status": code,
	"data":   data,
	"msg":    errmsg.Get_Error_Msg(code),
})
```

<br>

API文档: <br>
记录每个controller下的API，请求方式，URL，路径参数，响应实例。<br>

<br>

# 配置 
使用ini读取配置信息。节（section）、键（key）和值（value）组成，编写方便，表达性强，并能实现基本的配置分组功能<br>

[config_name] <br>
k = v <br>

读取参数: <br>
```Go
*ini.File.Section("config_name").Key("Key").String()
```

<br>

# 初始化路由组
Restful API 风格, REST描述的是在网络中client和server的一种交互形式。URL定位资源，用HTTP动词（GET,POST,DELETE,DETC）描述操作。<br>

```Go
func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()
	v1 := r.Group("api/v1")
	{
		//用户模块的路由接口
		//分类模块的路由接口
		//文章模块的路由接口
	}
}
```
分成两个路由组，一个需要鉴权，一个不需要 <br>

<br>

# 数据库

GORM，ORM（Object Relation Mapping）,对象关系映射，实际上就是对数据库的操作进行封装，对上层开发人员屏蔽数据操作的细节，开发人员看到的就是一个个对象，大大简化了开发工作，提高了生产效率。它可以节省相当多的繁琐编码。sql.DB 对象是许多数据库连接的池，其中包含 ' 使用中 ' 和' 空闲 ' 两种连接状态。当您使用连接来执行数据库任务 (例如执行 SQL 语句或查询行) 时，该连接会被标记为正在使用中。任务完成后，连接将被标记为空闲 <br>

1. 初始化工作，连接数据库，最大闲置连接，最大连接，复用时间，数据库迁移
2. 设计模型结构体，tag用于反射，动态获取该结构体的属性(gorm, json)
3. 钩子函数，一个模型下的方法
   
```Go
db.Select("id").Where("username = ?", username)
Select Id from users where username = ?
```

<br>

# 错误处理

设置一系列状态码常量。建立字典，状态码 ---> 对应信息 <br>
```Go
const (
	//status codes
	SUCCESS = 200
	ERROR   = 500

	//code = 1000... 用户模块错误
	ERROR_UserName_Used  = 1001
	ERROR_Password_WRONG = 1002
	//code = 2000... 文章模块错误

	//code = 3000... 分类模块错误
)
var codemsg = map[int]string{}

func Get_Error_Msg(code int) string {
	return codemsg[code]
}
```

<br>

# 实现API接口

ApiPost进行调试 <br>

model层实现数据库层面的增删改查 <br>

controller层实现具体的业务逻辑，接收前端发送过来的请求，调用model层方法与数据库交互，将结果返回给前端。<br>

1. 加密密码，scrpt加盐哈希算法，在密码中混入一段“随机”的字符串再进行哈希加密，这个被字符串被称作盐值。如同上面例子所展示的，这使得同一个密码每次都被加密为完全不同的字符串。为了校验密码是否正确，我们需要储存盐值。通常和密码哈希值一起存放在账户数据库中，或者直接存为哈希字符串的一部分。<br>
2. JWT验证，使应用程序知道向应用程序发送请求的人是谁。JWT 的原理是，服务器认证以后，生成一个 JSON 对象，发回给用户， 以后，用户与服务端通信的时候，都要发回这个 JSON 对象。服务器完全只靠这个对象认定用户身份。为了防止用户篡改数据，服务器在生成这个对象的时候，会加上签名。JWT 的三个部分依次如下。Header（头部，Payload（负载，Signature（签名）。当用户登录后，会将用户的信息进行加密，然后返回客户端一个加密后的字符串，可以存储在客户端的Cookie里，此后每一次请求都会带上它(放在请求头），如果此字符串和服务端使用私钥验证后一致，则认证成功，否则失败。

<br>

# 日志

logrus<br>
注册一个gin中间件，记录执行时间，客户端信息，IP<br>

# 优化

## 数据库并发优化

1. 分表，防止单表数据量太大。按照某种规则（RANGE,HASH取模等），切分到多张表里面去。但是这些表还是在同一个库中，所以库级别的数据库操作还是有IO瓶颈。该项目不适合分库，表之间依赖比较强 <br>
	创建一个新的表，CREATE TABLE TABLE_01 LIKE TABLE_00 <br>
	Insert into TABLE_01 <br>
	SELECT * FROM TBALE_00 WHERE ID % 2 = 0; <br>
	根据请求的ID来搜索表名 <br>

2. 读写分离，主从数据库，binlog + docker <br>

	binlog记录了MySQL所有写的操作，存储在 ~/datadir 下的 mysql-bin 中。<br>
	打开binlog log_bin on <br>
	查看binlog mysqlbinlog --base64 DATABASE mysql-bin <br>
	
	使用docker搭建主从服务器 <br>
	    首先拉取docker 镜像： docker pull mysql:latest <br>
	    主服务器: docker run -p 3339:3306 --name mysql-master -e MYSQL_ROOT_PASSWORD=123456 -d mysql:5.7 <br>
	    从服务器: docker run -p 3340:3306 --name mysql-slave -e MYSQL_ROOT_PASSWORD=123456 -d mysql:5.7
        进入到 Master 容器内部： docker exec -it mysql-master /bin/bash <br>
	    创建用户，配置server-id，开启binlog <br>
	    Docker允许通过外部访问容器或者容器之间互联的方式来提供网络服务。 <br>
	    从配置主从(主停机): change master to master_host = ip, master_user='slave', 'master_password', 'master_port', 'log_file', 'log_pos'<br>
	    show slave status 查看同步状态 <br>
	    
	数据迁移 <br>
	   源服务器 mysqldump -u root -p --opt src_database > src_database.sql <br>
	   主服务器 CREATE DATABASE new_database<br>
	           mysql -u root -p kalacloud_new_database < /tmp/kalacloud-data-export.sql <br>

	修改后端代码 <br>
	    改造配置 <br>
	    改造model层，不影响业务逻辑代码，新增一个db对象，一个连接主服务器，一个连接从服务器。<br>

	修改结果: 读业务走读从服务器，写业务走主服务器 <br>
	
3. 搜索文档
   以前模糊匹配，使用 LIKE。在文本比较少时是合适的，但是对于大量的文本数据检索，是不可想象的。 <br>
   全文检索 <br>
       create fulltext index content_tag_fulltext on fulltext_test(content,tag); <br>
	   MATCH(content,tag) AGAINST('key word')  <br>
	   全文索引，有两个变量，最小搜索长度和最大搜索长度，对于长度小于最小搜索长度和大于最大搜索长度的词语，都不会被索引。通俗点就是说，想对一个词语使用全文索引搜索，那么这个词语的长度必须在以上两个变量的区间内。<br>

4. 缓存
   把热数据存放到Redis缓存中，内存读写快，减小MySQL服务器压力，和业务部署在一个服务器上 ，Dial方法和Redis建立TCP连接，使用DO函数发送一个command <br>
       目标是实现文章 id ：article的缓存，使用Redis的string数据类型，最大可容纳512MB <br>
	   Model层初始化动态缓存，读取配置，初始化Redis连接池，避免建立和关闭连接的开销，管理空闲连接。封装面向连接池的DO方法(读取，缓存)，pool.Get从连接池中获取一个空闲连接，收取参数发送命令<br>
	   
	   Controller层，将参数id转为string类型，先向缓存发起请求，缓存有数据就返回，转为Article类型返回。没有则从MySQL拿数据返回给用户再异步发送到Redis里 <br>
  
5. Devops
   之前部署都是本地push github，然后服务器pull，最后运行，但是这样效率不高。<br>
   使用devop方案<br>
      写一个devop脚本，先杀进程，git pull，再启动 ./webserver &<br>
	  后端新增一个函数来重启服务器，运行这个脚本 <br>
	  注册到Webhook，一旦有push就自动pull

# 坑

调用查询接口的时候报错，出现了如下error，runtime error: invalid memory address or nil pointer dereference， 当在nil上调用一个属性或者方法的时候 , 会报空指针<br>
查看调用栈找到出错的那一行，db变量调用了查询函数。所以有可能发生错误的地方就是db == nil。然后进行判断，发现db 果然== nil。找到db定义的那一行，发现使用了 := ，本来db应该是全局的，但这样一来就变成了局部变量，最后使用 = 解决问题。<br>

<br>

# 未来的优化

1.  ElasticSearch <br>
	   Elasticsearch是一个分布式、高扩展、高实时的搜索与数据分析引擎。它能很方便的使大量数据具有搜索、分析和探索的能力。<br>
	   基础概念：索引，拥有相似特征的文档的集合。类型，在一个索引中，你可以定义一种或多种类型。一个类型是你的索引的一个逻辑上的分类。文档，一个文档是一个可被索引的基础信息单元，以JSON格式来表示。/index/type1 <br>
	   ES使Restful API来进行交互。新增文档，PUT ip:port/index/type {body: json} 。搜索，GET ip:port/index/type/_search。<br>
	   ElasticSearch和MySQL该如何同步 <br>
	
2. 缓存超过512MB? <br>

3. 缓存数据库一致性，两个业务同时修改数据库和缓存可能会造成冲突。写业务不修改缓存而是直接删掉缓存，读业务先找缓存，发现不存在于是从数据库更新缓存。**写业务先更新数据库，再删除缓存，因为缓存的写入通常要远远快于数据库的写入，保证更新操作读完数据库后会删除旧的缓存值**

4. 负载均衡
   用户的请求解析到负载均衡服务器上，负载均衡器转发到应用服务器上。反向代理是指以代理服务器来接受网络上的连接请求，然后将请求转发给内部网络上的服务器，并将从服务器上得到的结果返回给请求连接的客户端，此时代理服务器对外就表现为一个反向代理服务器。