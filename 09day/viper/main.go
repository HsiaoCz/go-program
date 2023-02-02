package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func main() {
	// 设置文件名
	// 这里需要注意的是 设置文件名 不要带有后缀
	viper.SetConfigName("config")
	// 设置文件类型
	viper.SetConfigType("toml")
	// 设置文件路径
	// 文件路径可以存在多个
	// viper会根据顺序自动查找
	viper.AddConfigPath(".")
	// 设置默认值
	viper.SetDefault("redis.port", 6379)
	// 读取文件
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	// viper.Get("key")  获得值
	// // viper获取值时以 section.key的形式，即传入嵌套的键的名称获取
	// fmt.Println(viper.Get("app_name"))
	// fmt.Println(viper.Get("log_level"))
	// fmt.Println("mysql ip: ", viper.Get("mysql.ip"))
	// fmt.Println("mysql port: ", viper.Get("mysql.port"))
	// fmt.Println("mysql user: ", viper.Get("mysql.user"))
	// fmt.Println("mysql password: ", viper.Get("mysql.password"))
	// fmt.Println("mysql database: ", viper.Get("mysql.database"))

	// fmt.Println("redis ip: ", viper.Get("redis.ip"))
	// fmt.Println("redis port: ", viper.Get("redis.port"))

	fmt.Println("protocols: ", viper.GetStringSlice("server.protocols"))
	fmt.Println("ports: ", viper.GetIntSlice("server.ports"))
	fmt.Println("timeout: ", viper.GetDuration("server.timeout"))

	fmt.Println("mysql ip: ", viper.GetString("mysql.ip"))
	fmt.Println("mysql port: ", viper.GetInt("mysql.port"))

	if viper.IsSet("redis.port") {
		fmt.Println("redis.port is set")
	} else {
		fmt.Println("redis.port is not set")
	}

	fmt.Println("mysql settings: ", viper.GetStringMap("mysql"))
	fmt.Println("redis settings: ", viper.GetStringMap("redis"))
	fmt.Println("all settings: ", viper.AllSettings())
}
