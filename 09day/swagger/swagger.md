# swagger

swagger 用来制作接口文档

swagger 会根据 openAPI 规范去生成各式各类的接口相关联的内容
常见的流程是编写注解==>调用生成库==>生成标准描述文件==>生成/导入到对应的 swagger 工具

安装 swagger

```go
go get -u github.com/swaggo/swag/cmd/swag
go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files
go get -u github.com/alecthomas/template
```

这里需要使用 go install 将 swag 安装到 bin 目录
单纯 go get 起不到效果

注解的描述
@Summary 摘要内容
@Produce API 可以产生的 MIME 类型列表,MIME 类型可以简单理解为响应类型，例如：
@Param 参数格式，从左到右分别为：参数名，入参类型，数据类型/是否必填/注释
@Success 响应成功，从左到右依次是：状态码，参数类型，数据类型，注释
@Failure 响应失败，从左到右分别为：路由地址,HTTP 方法

看几个例子:

```go
// @Summary 获取多个标签(这里对函数功能进行摘要)
// @Produce json (说明产生json格式的响应类型)
// @Param name query string false "标签名称" maxlength(100)  (指定响应参数，包含参数名称，入参类型，参数的类型，是否是必填项 注释)
// @Param state query int false "状态" Enums(0,1) default(1)
// @Success 200 {object} model.Tag "成功" (响应成功 状态码 响应成功后传递的参数的类型，参数的数据类型，注释)
// @failure 400 {object} errcode.Error "请求错误"
// @failure 500 {object} errcode.Error "内部错误"
// @Router /api/v1/tags  [get]  (路由路径，请求方法)
func (t Tag)List(c *gin.Context){}

// 这里create是post方法，有一些不同点
// @Param name body string true "标签名称" minlength(3) maxlength(100)
// post 方法的入参类型为body
func (t Tag)Create(c *gin.Context){}

// 对于put方法
// 入参类型也为body 不同于get方法的query
// @Param id path int true "标签id"  这里URL上的参数 注意是path的入参类型
// @Router /api/v1/tags/{id} [put]
func (t Tag)Update(c *gin.Context){}

// 含有路由参数的情况 和put一样
// @Router /api/v1/tags/{id} [delete]
func (t Tag)Delete(c *gin.Context){}
```

main 函数的接口文档

```go
// @title 博客系统
// @version 1.0
// description go project let's do go project
// termsOfServer https://github.com/HsiaoCz/go-program
// main函数上主要是标题 版本 描述信息 以及服务地质
func main(){}
```

接口文档的生成:

```bash
swag init
```

使用 swag 生成接口文档之后，就是要在路由中使用了
使用时需要在路由导入相应的包文件，以及注册相应的路由

```go
import (
    _ "github.com/HsiaoCz/go-program/docs"
    // 这里导入生成的docs文件夹
    ginSwagger "github.com/swaggo/gin-swagger"
    "github.com/swaggo/gin-swagger/swaggerFiles"
)

func NewRouter()*gin.Engine{
    // 将swagger路由注册进来
    r:=gin.Defalut()
    r.GET("/swagger/*any",ginSwagger.WrapHandler(swaggerFiles.Handler))
    return r
}

```

访问 http://localhsot:9090/swagger/insex.html 可以看到 swagger 页面
