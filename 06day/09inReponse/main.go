package main

// 内置的响应
// NotFound函数，包装一个404状态码和一个额外的信息
// ServeFile函数，从文件系统提供文件，返回给请求者
// ServeContent函数，它可以把实现了io.ReadSeeker接口的任何东西里面的内容返回给请求者
// 这个函数还可以处理Range请求(范围请求)，如果只请求了资源的一部分内容，那么ServeContent就可以如此响应。而ServeFile或io.Copy则不行
// Redirect 函数，告诉客户端重定向到另外一个URL
func main(){

}