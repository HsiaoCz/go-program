# logrus日志库

主要用来作为gin的日志管理的
不过有一个更好用的zap

安装:

```go
go get -u github.com/sirupsen/logrus
```

logrus支持更多的日志级别
- Panic :记录日志，然后panic
- Fatal :致命错误，出现错误时程序无法正常运转。输出日志后，程序退出
- Error :错误日志，需要查看原因
- Warn  :警告信息，提醒程序员注意
- Info  :关键操作，核心流程的日志
- Debug :一般程序中输出的调试信息
- Trace :很细粒度的信息，一般用不到

日志的级别从上到下依次增大，默认的日志级别为InfoLevel
为了能够看到Trace和Debug，我们需要在函数的第一行设置日志级别
这里有一点要注意，如果设置了日志级别，高于这个级别的日志不会显示

日志输出，一般会输出time/level和msg

我们还可以对日志进行定制

```go
// 使用logrus.setReportCaller(true) 设置日志输出中显示文件名称和方法信息
	logrus.SetReportCaller(true)
	logrus.Info("info msg")

```

可以往日志中添加字段

```go
// 通过调用logrus.WithField和logrus.WithFields实现添加字段
// logrus.WithFileds接收一个logrus.Fileds类型的参数，底层为map[string]interface{}
	logrus.WithFields(logrus.Fields{
		"name": "zhangsan",
		"age":  20,
	}).Info("info msg")
// 如果在一个函数中的所有日志都需要添加某些字段，可以使用WithFields的返回值
// 在web请求中我们可以为所有请求的处理器都加上user_id和ip字段
	requestLogger := logrus.WithFields(logrus.Fields{
		"user_id": 1900,
		"ip":      "127.0.0.1:9090",
	})
	requestLogger.Info("info massage")
	requestLogger.Error("error msg")
// withFields返回一个logrus.Entry类型的值，它将logrus.Logger和设置的logrus.Fileds保存下来,调用Entry相关方法输出日志时，保存下来的logrus.Fields也会随之输出

```

重定向输出

```go
// 默认情况下，日志输出到io.Stder。输出到控制台
// 调用logrus.SetOutput()传入一个io.Writer参数，后续调用相关方法写入到io.Writer中
// 可以将日志写入到文件，标准输出，buffer中
 writer1 := &bytes.Buffer{}
  writer2 := os.Stdout
  writer3, err := os.OpenFile("log.txt", os.O_WRONLY|os.O_CREATE, 0755)
  if err != nil {
    log.Fatalf("create file log.txt failed: %v", err)
  }
  // io.MultiWriter(可以输出到多个io.Writer
  logrus.SetOutput(io.MultiWriter(writer1, writer2, writer3))
  logrus.Info("info msg")

```

自定义logrus对象
自定义对象之后，我们可以设置一些函数设置日志输出级别
```go
log := logrus.New()

  log.SetLevel(logrus.InfoLevel)
  log.SetFormatter(&logrus.JSONFormatter{})

  log.Info("info msg")
```

日志格式

logrus支持两种日志格式
文本格式和json格式
可以通过logrus.SetFormatter设置日志的格式

```go
logrus.SetLevel(logrus.TraceLevel)
  logrus.SetFormatter(&logrus.JSONFormatter{})

  logrus.Trace("trace msg")
  logrus.Debug("debug msg")
  logrus.Info("info msg")
  logrus.Warn("warn msg")
  logrus.Error("error msg")
  logrus.Fatal("fatal msg")
  logrus.Panic("panic msg")
```

除了这两种格式，还支持一些第三方格式
```go
 go get github.com/antonfisher/nested-logrus-formatter
```

```go
  logrus.SetFormatter(&nested.Formatter{
    HideKeys:    true,
    FieldsOrder: []string{"component", "category"},
  })

  logrus.Info("info msg")
```
nested格式提供了多个字段用来定制行为

```go
// github.com/antonfisher/nested-logrus-formatter/formatter.go
type Formatter struct {
  FieldsOrder     []string
  TimestampFormat string  
  HideKeys        bool    
  NoColors        bool    
  NoFieldsColors  bool    
  ShowFullLevel   bool    
  TrimMessages    bool    
}
```

默认的logrus 输出日志中字段是key-value格式，使用nested格式，我们可以通过设置hidekeys为true隐藏键，只输出值
默认的logurs是按键的字母顺序输出字段，可以设置FieldOrder定义输出字段顺序
通过设置TimestampFormat设置日期格式

```go
logrus.SetFormatter(&nested.Formatter{
    // HideKeys:        true,
    TimestampFormat: time.RFC3339,
    FieldsOrder:     []string{"name", "age"},
  })

  logrus.WithFields(logrus.Fields{
    "name": "dj",
    "age":  18,
  }).Info("info msg")
```

设置钩子
每条日志输出前都会执行钩子的特定方法。
可以添加输出字段，根据日志级别输出到不同的目的地,logrus内置了一个钩子syslog
将日志输出到syslog中
我们可以自定义钩子，在输出的日志中增加一个app=awesome-web字段

钩子需要实现logrus.Hook接口
```go
// github.com/sirupsen/logrus/hooks.go
type Hook interface {
  Levels() []Level
  Fire(*Entry) error
}
```

levels()方法返回感兴趣的日志级别，输出其他日志时不会触发钩子。fire是日志输出前调用的钩子方法

```go
package main

import (
  "github.com/sirupsen/logrus"
)

type AppHook struct {
  AppName string
}

func (h *AppHook) Levels() []logrus.Level {
  return logrus.AllLevels
}

func (h *AppHook) Fire(entry *logrus.Entry) error {
  entry.Data["app"] = h.AppName
  return nil
}

func main() {
  h := &AppHook{AppName: "awesome-web"}
  logrus.AddHook(h)

  logrus.Info("info msg")
}
```

还有一些将日志放送到不同数据库的钩子库

[https://github.com/weekface/mgorus] :发送到mongoDB
[https://github.com/rogierlommers/logrus-redis-hook] ：发送到redis
[https://github.com/vladoatanasov/logrus_amqp] :将日志发送到ActiveMQ
