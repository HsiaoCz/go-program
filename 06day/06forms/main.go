package main

import (
	"flag"
	"fmt"
	"net/http"
)

// 1.如何通过表单来发送POST请求
// 这里的action 就是处理的路径  method是请求方法
// <form action="" method="post">
//	<input type="text" name="first_name" />
//	<input type="text" name="last_name" />
//	<input type="submit"/>
// </form>
// type 的text 会以键值对的形式提交
// 表单的数据会以name-value对的形式，以post请求发送出去
// 它的数据内容会放在POST请求的Body里面
// 这个name-value对在body里面的格式
// 通过POST发送的name-value数据对的格式可以通过表单的COntent Type来指定,也就是enctype属性
// 这个属性的默认值为"application/x-www-form-urlencoded"
// 浏览器被要求支持这种格式和multipart/form-data
// 如果是html5 还要求支持text/plain这种格式
// 如果enctype是application/x-www-form-urlencoded,那么浏览器会将表单数据编码到查询字符串里面，例如：
// first_name=sau%20shengg&last_name=cheng

// 如果enctype是multipart/form-data，那么
// 每一个name-value对都会被转换为一个MME消息部分
// 每一个部分都有自己的Content Type 和Content Disposition

// 如何选择这两种格式
// 简单文本：表单URL编码
// 大量数据，例如上传文件:Multipart-MIME这种格式
// 通过这种格式 甚至可以把二进制数据通过选择Base64编码，来当作文本进行发送

// 通过表单的method属性，可以设置POST还是GET
// GET请求没有body 所有的数据都通过URL的name-value对来发送

// Form字段
// Request上的函数允许我们从URL中提取数据，通过这些字段：
// Form
// PostForm
// MultipartForm
// FOrm里面的数据是Key-value对
// 使用它 通常需要这样做：
// 先调用ParseForm或ParseMultipartForm来解析Request
// 然后相应的访问Form、PostForm或MultipartForm字段
func main() {
	listenAddr := flag.String("listenAddr", ":9090", "set listenAddr")
	flag.Parse()
	http.HandleFunc("/hello", Hello)
	http.HandleFunc("/", StaticForm)
	server := http.Server{
		Addr:    *listenAddr,
		Handler: nil,
	}
	server.ListenAndServe()
}

func Hello(w http.ResponseWriter, r *http.Request) {
	// 解析r.ParseForm来解析request的form
	// r.ParseForm()
	// 解析mutlipart 使用ParseMutlipartForm
	r.ParseMultipartForm(1024) //需要一个长度
	fmt.Println(r.Form.Get("first_name"))
	fmt.Println(r.Form.Get("last_name"))
	// fmt.Fprintln(w, r.Form)
	fmt.Fprintln(w, r.MultipartForm)
}

func StaticForm(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "postform.html")
}

// PostForm字段
// 在之前的那个例子中，我们如果只想得到first_name这个Key的value
// 可以使用r.Form["first_name"]
// 但是form存在一个问题，当URL里面有相同的字段时，它们会放在一个slice里面，表单的值靠前，URL里面的值靠后
// 如果只想要表单的key-value对，不要URL的，可以使用PostForm
// 不过PostForm 只支持表单格式为application/x-www-form-urlencodeda 的form表单
// 想要得到multipart key-value，必须使用multipartForm字段

// 想要使用Multipartform字段，首先需要调用ParseMultiForm这个方法
// 该方法会在必要时调用ParseForm方法
// 参数需要读取数据的长度
// multipartForm 只包含表单的key-value对
// 返回类型是一个struct而不是map 这个struct里有两个map
// 第一个map key is string value : []srring
// 第二个map 空的 (key tring value:文件)

// FormValue 和 PostFormValue方法
// FormValue 方法会返回form字段中指定key对应的第一个value
// 无需调用ParseFoem或ParseMultipartForm
// PostFormValue 方法也一样，但只能读取PostForm
// FormValue和PostFormValue 都会调用ParseMultipartForm方法
// 但如果表单的enctype设为multipart/form-data，那么即使你调用PareMultipartForm方法，也无法通过FormValue获得想要的值
