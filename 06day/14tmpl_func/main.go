package main

import (
	"fmt"
	"net/http"
	"regexp"
)

// 参数就是模板里面的值
// 可以是bool、整数、string..
// 也可以是struct、struct的字段、数组的key等等
// 参数可以是变量、方法(返回单个值或返回一个值和一个错误)或函数
// 参数也可以是一个点，也就是传入模板引擎的那个值

// 在action中设置变量，变量以$开头
// {{range $key,$value:=.}}
// key {{$key}} value {{$value}}
// {{end}}

// 管道是按顺序连接到一起的参数、函数和方法
// {{p1|p2|p3}}

// 函数

// 静态路由：一个路径对应一个页面
// /home /about 这种路由，一个路由对应一个页面

// 带参数的路由：根据路由参数，创建出一族不同的页面
// 比如/companies/123
// companies/Microsoft
// 页面固定 但是数据不同
func main() {
	http.HandleFunc("/don/", DonHandle)
	http.ListenAndServe(":9090", nil)
}

func DonHandle(w http.ResponseWriter, r *http.Request) {
	pattern, _ := regexp.Compile(`/don/(\w+)`)
	matches := pattern.FindStringSubmatch(r.URL.Path)
	if len(matches) > 0 {
		name := matches[1]
		fmt.Fprintln(w, "Hello:"+name)
	}
	w.WriteHeader(http.StatusNotFound)
}
