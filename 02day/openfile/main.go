package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

// func readForm() {
// 	file, err := os.Open("./23444.txt")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	// 关闭文件
// 	defer file.Close()
// 	fmt.Println("打开文件成功!")
// 	var tmp = make([]byte, 128)
// 	n, err := file.Read(tmp)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	fmt.Println(string(tmp[:n]))
// }

// 打开文件
func main() {
	// readForm()
	// readFromFileBufio()
	readFileByOS()
}

// 利用bufio来读取文件
// func readFromFileBufio() {
// 	file, err := os.Open("./23444.txt")
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	defer file.Close()
// 	//创建一个用来读取文件的缓冲
// 	reader := bufio.NewReader(file)
// 	str, err := reader.ReadString('\n')
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	if err == io.EOF {
// 		return
// 	}
// 	fmt.Println(str)
// }

// 使用ReadFile()来读取文件
func readFileByOS() {
	file, err := os.OpenFile("./23444.txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	str, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln(err)
	}
	if err == io.EOF {
		fmt.Println("读取完毕!")
		return
	}
	fmt.Println(str)
}
