package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	// 注册默认的路由
	// 默认的路由 使用两个中间件
	// 有一个引擎
	// 修改ip为内网ip
	// r.Run("0.0.0.0:9090")
	r := gin.Default()
	// 加载模板文件
	r.LoadHTMLFiles("templates/login.html", "templates/index.html")
	r.Use(m1)
	r.GET("/hello/:name", HelloHandler)
	r.GET("/product", QueryHandler)
	r.GET("/user/login", m1, GetUserLoginIndex)
	r.POST("/user/login", UserLogin)
	r.POST("/loadfile", LoadFile)
	r.POST("/loadfies", LoadFiles)
	r.GET("/redirect", HTTPRedirect)
	// r.POST("/user/info", BindJson)
	// 绑定query
	r.GET("/user/info", BindJson)
	r.Run(":9090")
}

func HelloHandler(c *gin.Context) {
	name := c.Param("name")
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"name": name,
	})
}

// 获取路由查询的参数
// 在gin中有这样三种方式
// 1、c.Query("key")
// 2、c.DefaultQuery("key","默认值"),这种获取方式在没有取到值得时候，会返回默认值
// 3、c.GetQuery("key") 没有值会返回false
// 这里需要注意得是，查询得内容不要跟在路由注册的后面
// 比如r.GET("product/?name&age")
// 这种是不对的
func QueryHandler(c *gin.Context) {
	name, ok := c.GetQuery("name")
	if !ok {
		name = "张三"
	}
	age, ok := c.GetQuery("age")
	if !ok {
		age = "2001"
	}
	c.JSON(http.StatusOK, gin.H{
		"name": name,
		"age":  age,
	})
}

// 获取form表单信息
// 获取表单有三种方式
// 1. username:=c.PostForm("username")
// 第二种，没拿到返回默认值
// username:=c.DefaultPostForm("username","zhangsan")
// 第三种方式
// 使用GetPostForm("key") 这种方式返回一个string，一个布尔值
func UserLogin(c *gin.Context) {
	u := User{}
	u.Username = c.PostForm("username")
	u.Password = c.PostForm("password")
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Username": u.Username,
		"Password": u.Password,
	})
}

func GetUserLoginIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

type User struct {
	Username string `json:"username"`
	Password string `json:"pssword"`
}

// 绑定参数
// 绑定json需要加json的tag
// 绑定query查询参数 需要加form的tag
// 还可以绑定URI参数 ，加uri的tag
// 请求post 格式为
// r.POST("/user/:name/:age/:gender")
// shouldBindUri

// 绑定formdata 使用shouldbind
// 绑定传递过来的表单信息
type UserInfo struct {
	Name   string `json:"name" form:"name"`
	Age    int    `json:"age" form:"age"`
	Gender string `json:"gender" form:"gender"`
}

func BindJson(c *gin.Context) {
	var userInfo UserInfo
	// 绑定json
	// err := c.ShouldBindJSON(&userInfo)
	// 绑定query查询参数
	err := c.ShouldBindQuery(&userInfo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "wrong",
		})
		return
	}
	c.JSON(http.StatusOK, userInfo)
}

// gin的参数校验
// 在结构体里加tag binding
// 参数校验使用validata库

// http请求重定向

func HTTPRedirect(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "http://www.bilibili.com")
}

// 路由重定向
// 网页需要跳转到登录
// 使用r.HandleContext(c) 处理上下文
func RouterRedirect(c *gin.Context) {
	// 1.修改请求的URL
	c.Request.URL.Path = "/b"
	//继续处理
	// r.HandleContext(c)
}

// 文件上传
// 单文件上传
// 处理multipart forms提交文件时默认的内存限制是32 MiB
// 可以通过下面的方式修改
// router.MaxMultipartMemory = 8 << 20  // 8 MiB
func LoadFile(c *gin.Context) {
	file, err := c.FormFile("f1")
	if err != nil {
		log.Fatal(err)
	}
	// 将文件上传到指定目录
	dst := fmt.Sprintf("C:/tmp/%s", file.Filename)
	c.SaveUploadedFile(file, dst)
}

// 上传多个文件
// 上传多个文件使用Multipart form
func LoadFiles(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["file"]

	for index, file := range files {
		log.Println(file.Filename)
		dst := fmt.Sprintf("C:/tmp/%s_%d", file.Filename, index)
		c.SaveUploadedFile(file, dst)
	}
}

// any方法可以匹配所有的路由

// 中间件的使用 gin中的中间件必须是一个gin.HandlerFunc类型
// 定义一个中间件
func m1(c *gin.Context) {
	//这个中间件计算处理事件
	start := time.Now()
	//c.Next()进行下一步处理
	c.Next()
	cost := time.Since(start)
	fmt.Println("处理花费了:", cost)
}

// 使用中间件 使用r.use 全局使用
// 为单个路由添加中间件
// r.GET("/hello",m1,HelloHandle)
// 跨中间件取值和设置值
// c.Set("name","zhangsan")
// c.Get("name")
// 为路由组设置中间件
// sgroup:=r.Group("/shop",m1)

// 为模板添加函数
// r.SetFuncMap()
