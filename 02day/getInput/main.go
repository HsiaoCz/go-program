package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// 获取用户输入如果有空格 或出现问题

// func useScan() {
// 	var s string
// 	fmt.Println("请输入一句话:")
// 	_, err := fmt.Scanln(&s)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	fmt.Println(s)
// }

func main() {
	// useScan()
	useBufioForGetInput()
}

// 使用bufio来获取用户输入
func useBufioForGetInput() {
	var s string
	reader := bufio.NewReader(os.Stdout)
	fmt.Println("请输入一句话:")
	s, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(s)

}
