# GinBlog

# 前后端交互

定义好API文档，错误的参数处理(比如用户名冲突)。

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

<br>

# 数据库

GORM，ORM（Object Relation Mapping）,对象关系映射，实际上就是对数据库的操作进行封装，对上层开发人员屏蔽数据操作的细节，开发人员看到的就是一个个对象，大大简化了开发工作，提高了生产效率。它可以节省相当多的繁琐编码。<br>

sql.DB 对象是许多数据库连接的池，其中包含 ' 使用中 ' 和' 空闲 ' 两种连接状态。当您使用连接来执行数据库任务 (例如执行 SQL 语句或查询行) 时，该连接会被标记为正在使用中。任务完成后，连接将被标记为空闲 <br>

1. 初始化工作，连接数据库，最大闲置连接，最大连接，复用时间，数据库迁移
2. 设计模型结构体，tag用于反射，动态获取该结构体的属性(gorm, json)

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

controller层，实现具体的业务逻辑，与数据库交互，返回给前端。<br>
