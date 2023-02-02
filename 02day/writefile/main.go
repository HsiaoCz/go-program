package main

import (
	"fmt"
	"log"
	"os"
)

// 写入文件
// 将一个文件里的内容复制到另一个文件
func writeFile() {
	file, err := os.OpenFile("./2344.txt", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	var s string = "Hello World!"
	_, err = file.WriteString(s)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("写入完毕")
}

func main() {
	writeFile()
}
