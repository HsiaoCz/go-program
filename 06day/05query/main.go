package main

import (
	"log"
	"net/http"
)

// 查询参数
// URL Query
// http://www.example.com/post?id=1238thread_id=456
// 这是一个URL的例子，其中id=1238thread_id=456就是查询参数
// r.URL.RawQuery会提供实际查询的原始字符串
// 上面例子的RawQuery的值就是id=1238thread_id=456
// r.URL.Query()，会提供查询字符串对应的map[string][]string
// 查询参数的key 允许重复
func main() {
	http.HandleFunc("/home", QueryThing)
	http.ListenAndServe(":9090", nil)
}

func QueryThing(w http.ResponseWriter, r *http.Request) {
	url := r.URL
	query := url.Query()
	//使用key来获取 如果key相同，会返回一个字符串的切片
	id := query["id"]
	log.Println(id)
	// 使用get方式来获取 key相同 会返回第一个key的值
	name := query.Get("name")
	log.Println(name)
}
