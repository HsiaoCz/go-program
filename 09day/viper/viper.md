# viper

viper 主要用来管理配置
不过不是那种分布式的配置管理

viper 可以监控配置文件的修改，可以从环境变量/命令行和 io.Reader 中读取配置
可以从 etcd/Consul 读取配置

安装

```go
go get github.com/spf13/viper

```

viper 有多种读取配置的方式，get 方法返回的值是 interface{}类型
可以使用 GetType()方法
对应的 type 有多种：bool/float64/int/string/time/duration/intslice/stringslice
指定的键不存在或者类型不正确会返回对应的零值

要判断某个键是否存在，使用 IsSet()方法 GetStringMap 和 GetStringMapString 直接以 map 返回某个键下面所有的键值对，前者返回 map[string]interface{}，后者返回 map[string]string。AllSettings 以 map[string]interface{}返回所有设置。

```go
func main() {
  viper.SetConfigName("config")
  viper.SetConfigType("toml")
  viper.AddConfigPath(".")
  err := viper.ReadInConfig()
  if err != nil {
    log.Fatal("read config failed: %v", err)
  }

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
```

viper 使用 set 还可以设置值

viper.Set("redis.port",5344)
viper 设置值的优先级很高

命令行选项
如果一个键没有通过 viper.Set()显示设置值，那么获取值的时候会尝试从命令行获取
如果有，优先使用

```go
func init() {
  pflag.Int("redis.port", 8381, "Redis port to connect")

  // 绑定命令行
  viper.BindPFlags(pflag.CommandLine)
}
```

如果没有获取到，viper 将会从环境变量获取值

```go
func init() {
  // 绑定环境变量
  viper.AutomaticEnv()
  // 单独绑定
  viper.BindEnv("redis.port")
}
```

调用 BindEnv 方法，如果只传入一个参数，则这个参数既表示键名，又表示环境变量名。如果传入两个参数，则第一个参数表示键名，第二个参数表示环境变量名
还可以通过 viper.SetEnvPrefix 方法设置环境变量前缀，这样一来，通过 AutomaticEnv 和一个参数的 BindEnv 绑定的环境变量，在使用 Get 的时候，viper 会自动加上这个前缀再从环境变量中查找
如果对应的环境变量不存在，viper 会自动将键名全部转为大写再查找一次。所以，使用键名 gopath 也能读取环境变量 GOPATH 的值

读取配置 1.从 io.Reader 读取配置

```go
package main

import (
  "bytes"
  "fmt"
  "log"

  "github.com/spf13/viper"
)

func main() {
  viper.SetConfigType("toml")
  tomlConfig := []byte(`
app_name = "awesome web"

# possible values: DEBUG, INFO, WARNING, ERROR, FATAL
log_level = "DEBUG"

[mysql]
ip = "127.0.0.1"
port = 3306
user = "dj"
password = 123456
database = "awesome"

[redis]
ip = "127.0.0.1"
port = 7381
`)
  err := viper.ReadConfig(bytes.NewBuffer(tomlConfig))
  if err != nil {
    log.Fatal("read config failed: %v", err)
  }

  fmt.Println("redis port: ", viper.GetInt("redis.port"))
}
```

viper 的 unmashal
将配置给 unmarshal 到 struct 中

```go
package main

import (
  "fmt"
  "log"

  "github.com/spf13/viper"
)

type Config struct {
  AppName  string
  LogLevel string

  MySQL    MySQLConfig
  Redis    RedisConfig
}

type MySQLConfig struct {
  IP       string
  Port     int
  User     string
  Password string
  Database string
}

type RedisConfig struct {
  IP   string
  Port int
}

func main() {
  viper.SetConfigName("config")
  viper.SetConfigType("toml")
  viper.AddConfigPath(".")
  err := viper.ReadInConfig()
  if err != nil {
    log.Fatal("read config failed: %v", err)
  }

  var c Config
  viper.Unmarshal(&c)

  fmt.Println(c.MySQL)
}
```

保存配置

有时候，我们想要将程序中生成的配置，或者所做的修改保存下来。viper 提供了接口！
WriteConfig：将当前的 viper 配置写到预定义路径，如果没有预定义路径，返回错误。将会覆盖当前配置；
SafeWriteConfig：与上面功能一样，但是如果配置文件存在，则不覆盖；
WriteConfigAs：保存配置到指定路径，如果文件存在，则覆盖；
SafeWriteConfig：与上面功能一样，但是入股配置文件存在，则不覆盖。

```go
package main

import (
  "log"

  "github.com/spf13/viper"
)

func main() {
  viper.SetConfigName("config")
  viper.SetConfigType("toml")
  viper.AddConfigPath(".")

  viper.Set("app_name", "awesome web")
  viper.Set("log_level", "DEBUG")
  viper.Set("mysql.ip", "127.0.0.1")
  viper.Set("mysql.port", 3306)
  viper.Set("mysql.user", "root")
  viper.Set("mysql.password", "123456")
  viper.Set("mysql.database", "awesome")

  viper.Set("redis.ip", "127.0.0.1")
  viper.Set("redis.port", 6381)

  err := viper.SafeWriteConfig()
  if err != nil {
    log.Fatal("write config failed: ", err)
  }
}
```

监控配置文件更改
实现热加载配置 不需要重启服务

```go
package main

import (
  "fmt"
  "log"
  "time"

  "github.com/spf13/viper"
)

func main() {
  viper.SetConfigName("config")
  viper.SetConfigType("toml")
  viper.AddConfigPath(".")
  err := viper.ReadInConfig()
  if err != nil {
    log.Fatal("read config failed: %v", err)
  }

  // 观察配置 派个小弟看着
  // 如果配置文件发生了修改 viper会自动加载
  viper.WatchConfig()

  fmt.Println("redis port before sleep: ", viper.Get("redis.port"))
  time.Sleep(time.Second * 10)
  fmt.Println("redis port after sleep: ", viper.Get("redis.port"))
}
```

还可以为配置修改增加一个回调函数

```go
viper.OnConfigChange(func(e fsnotify.Event) {
  fmt.Printf("Config file:%s Op:%s\n", e.Name, e.Op)
})
```
