package main

//proto编码解码和json很像
import (
	"fmt"
	study "go-program/10day/protobuf/demo"
	"log"

	"google.golang.org/protobuf/proto"
)

func main() {
	// 初始化proto中的消息
	studyInfo := &study.StudyInfo{}

	//常规赋值
	studyInfo.Id = 1
	studyInfo.Name = "study prtyon"
	studyInfo.Duration = 180
	studyInfo.Score = map[string]int32{
		"shizhan": 100,
		"hongl":   200,
	}

	//用字符串的方式打印
	fmt.Printf("字符串输出结果:%v\n", studyInfo.String())

	//转化成二进制文件
	marshal, err := proto.Marshal(studyInfo)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Marshal 转成二进制文件的结果:%v\n", marshal)

	//将二进制文件转成结构体
	newStudyInfo := study.StudyInfo{}
	err = proto.Unmarshal(marshal, &newStudyInfo)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("二进制转成结构体的结果:%v\n", &newStudyInfo)
}
