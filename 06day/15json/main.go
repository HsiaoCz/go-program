package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// json 与go 的类型映射
// Go bool:JSON boolean
// Go float64:JSON 数值
// Go string:Json strings
// Go nil:Json null

// 对于未知结构的Json
// 可以使用map[string]interface{}可以存储任意JSON对象
// []interface{} 可以存储任意JSON数组

// 读取JSON 需要一个解码器:dec:=json.NewDecoder(r.Body)
// 这个Decodeer参数需要实现Reader接口
// 在解码器上进行解码:dec.Decode(&query)

// 写入json 需要一个编码器:enc:=json.NewEncoder(w)
// 参数需实现Writer接口
// 使用编码器:enc.Encode(results)
func main() {
	JsonEnDe()
	http.HandleFunc("/json", JsonHandler)
	http.ListenAndServe(":9090", nil)
}

func JsonHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		dec := json.NewDecoder(r.Body)
		company := company{}
		err := dec.Decode(&company)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		enc := json.NewEncoder(w)
		err = enc.Encode(company)
		if err != nil {
			log.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

type company struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
}

// go语言里编解码json还可以使用Marshal和UnMarshal
// 把struct转化为json格式
// unmarshal 将json格式的数据转化成go struct

func JsonEnDe() {
	jsonStr := `
	{
		"id":123,
		"name":"google",
		"country":"USA"
	}`

	c := company{}
	_ = json.Unmarshal([]byte(jsonStr), &c)
	fmt.Println(c)

	// 使用json.marshal是没有缩进的
	bytes, _ := json.Marshal(c)
	fmt.Println(string(bytes))

	// 使用有缩进的Marshal
	bytes1, _ := json.MarshalIndent(c, "", "  ")
	fmt.Println(string(bytes1))
}

// 两种区别
// 针对 string 或bytes 使用marshal 或unmarshal
// 针对stream 使用Encode /Decode 将数据写入到io.Writer或者从io.Reader读数据
