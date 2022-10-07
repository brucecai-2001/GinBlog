# GinBlog

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

```Go
func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()
	v1 := r.Group("api/v1")
    v1.GET(.....)
}
```

<br>

# 数据库

GORM，ORM（Object Relation Mapping）,对象关系映射，实际上就是对数据库的操作进行封装，对上层开发人员屏蔽数据操作的细节，开发人员看到的就是一个个对象，大大简化了开发工作，提高了生产效率。它可以节省相当多的繁琐编码。<br>

sql.DB 对象是许多数据库连接的池，其中包含 ' 使用中 ' 和' 空闲 ' 两种连接状态。当您使用连接来执行数据库任务 (例如执行 SQL 语句或查询行) 时，该连接会被标记为正在使用中。任务完成后，连接将被标记为空闲 <br>

1. 初始化工作，连接数据库，最大闲置连接，最大连接，复用时间，数据库迁移
2. 设计模型结构体，tag用于反射，动态获取该结构体的属性(gorm, json)

<br>