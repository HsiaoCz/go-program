package main

import (
	"flag"
	"fmt"
	"os"
	"runtime/pprof"
	"time"
)

// go的性能调优
// profiling是指对应程序的画像，画像就是程序使用CPU和内存的情况
// go 内置的采集性能的工具
// runtine/pprof 采集工具型应用运行数据进行分析
// net/http/pprof 采集服务型应用运行时的数据分析
// pprof 开启后，每隔一段时间就会收集下当前的堆栈信息，获取函数占用CPU的情况，以及内存资源
// pprof 用于性能测试
func main() {
	var isCPUPprof bool
	var isMemPprof bool

	flag.BoolVar(&isCPUPprof, "cpu", false, "turn cpu pprof on")
	flag.BoolVar(&isMemPprof, "mem", false, "turn mem pprof on")
	flag.Parse()
	if isCPUPprof {
		file, err := os.Create("./cpu.pprof")
		if err != nil {
			fmt.Printf("create cpu pprof failed,err:%v\n", err)
			return
		}
		pprof.StartCPUProfile(file)
		defer pprof.StopCPUProfile()
	}
	for i := 0; i < 8; i++ {
		go logicCode()
	}
	time.Sleep(20 * time.Second)
	if isMemPprof {
		file, err := os.Create("./mem.pprof")
		if err != nil {
			fmt.Printf("create mem pprof failed,err%v\n", err)
			return
		}
		pprof.WriteHeapProfile(file)
		file.Close()
	}
}

func logicCode() {
	var c chan int
	for {
		select {
		case v := <-c:
			fmt.Printf("recv from chan,value:%v\n", v)
		default:
			fmt.Println("Hello World")
		}
	}
}
