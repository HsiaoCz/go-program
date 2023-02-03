# zap

zap是一个日志库，由uber开源的
之所以用，相较于logrus，主要是性能优化和内存分配更好

安装

```go
go get go.uber.org/zap
```

zap 提供了两种类型的日志记录器-Sugared Logger和Logger
在性能很好但不是很关键的上下文中,使用SugaredLogger。它比其他结构化日志包快很多
在每一微秒和每一次内存分配都很重要的上下文中,使用Logger，它甚至比SugaredLogger更快，内存分配次数更少

使用Logger
通过调用zap.NewProduction()/zap.NewDevelopment()或者zap.Example()创建一个Logger
每一个函数都会创建一个logger唯一的区别在于它将记录的信息不同。
production logger默认记录调用函数信息，日期和时间等

使用zap.Example()主要在测试代码中使用,Development()在开发环境中使用,product()用于生产环境
zap中使用自己提供的方法记录字段zap.Type()包含多种字段
如果这样觉得麻烦，zap还提供了SugarLogger这种日志记录器，使用printf()格式记录日志字段，但是性能比较低
可以用在非热点函数里面

通过logger调用Info/Error等
默认情况下日志会打印到控制台

这我们首先使用logger这个记录器

```go
var logger *zap.Logger

func main() {
 InitLogger()
  defer logger.Sync()
 simpleHttpGet("www.google.com")
 simpleHttpGet("http://www.google.com")
}

func InitLogger() {
 logger, _ = zap.NewProduction()
}

func simpleHttpGet(url string) {
 resp, err := http.Get(url)
 if err != nil {
  logger.Error(
   "Error fetching url..",
   zap.String("url", url),
   zap.Error(err))
 } else {
  logger.Info("Success..",
   zap.String("statusCode", resp.Status),
   zap.String("url", url))
  resp.Body.Close()
 }
}
```

这里面的zap.String() zap.Error()，就是日志的记录器的方法
每个方法都接受一个消息字符串和任意数量的zapcore.Field场参数。
每个zapcore.Field其实就是一组键值对参数。

使用Sugared Logger 使用这种日志记录器 主要是性能很好但不是很关键的上下文中
使用这种日志记录器会以printf的格式记录日志

```go
var sugarLogger *zap.SugaredLogger

func main() {
 InitLogger()
 defer sugarLogger.Sync()
 simpleHttpGet("www.google.com")
 simpleHttpGet("http://www.google.com")
}

func InitLogger() {
  logger, _ := zap.NewProduction()
 sugarLogger = logger.Sugar()
}

func simpleHttpGet(url string) {
 sugarLogger.Debugf("Trying to hit GET request for %s", url)
 resp, err := http.Get(url)
 if err != nil {
  sugarLogger.Errorf("Error fetching URL %s : Error = %s", url, err)
 } else {
  sugarLogger.Infof("Success! statusCode = %s for URL %s", resp.Status, url)
  resp.Body.Close()
 }
}
```

使用zap记录有嵌套的层级关系,使用zap.Namespace(key string)Field构建一个命名空间

```go
func main() {
  logger := zap.NewExample()
  defer logger.Sync()

  logger.Info("tracked some metrics",
    zap.Namespace("metrics"),
    zap.Int("counter", 1),
  )

  logger2 := logger.With(
    zap.Namespace("metrics"),
    zap.Int("counter", 1),
  )
  logger2.Info("tracked some metrics")
}
```

这里展示了namespace的两种用法，效果是相同的

在gin中如何使用zap

```go
package logger

import (
 "gin_zap_demo/config"
 "net"
 "net/http"
 "net/http/httputil"
 "os"
 "runtime/debug"
 "strings"
 "time"

 "github.com/gin-gonic/gin"
 "github.com/natefinch/lumberjack"
 "go.uber.org/zap"
 "go.uber.org/zap/zapcore"
)

var lg *zap.Logger

// InitLogger 初始化Logger
func InitLogger(cfg *config.LogConfig) (err error) {
 writeSyncer := getLogWriter(cfg.Filename, cfg.MaxSize, cfg.MaxBackups, cfg.MaxAge)
 encoder := getEncoder()
 var l = new(zapcore.Level)
 err = l.UnmarshalText([]byte(cfg.Level))
 if err != nil {
  return
 }
 core := zapcore.NewCore(encoder, writeSyncer, l)

 lg = zap.New(core, zap.AddCaller())
 zap.ReplaceGlobals(lg) // 替换zap包中全局的logger实例，后续在其他包中只需使用zap.L()调用即可
 return
}

func getEncoder() zapcore.Encoder {
 encoderConfig := zap.NewProductionEncoderConfig()
 encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
 encoderConfig.TimeKey = "time"
 encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
 encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
 encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
 return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
 lumberJackLogger := &lumberjack.Logger{
  Filename:   filename,
  MaxSize:    maxSize,
  MaxBackups: maxBackup,
  MaxAge:     maxAge,
 }
 return zapcore.AddSync(lumberJackLogger)
}

// GinLogger 接收gin框架默认的日志
func GinLogger() gin.HandlerFunc {
 return func(c *gin.Context) {
  start := time.Now()
  path := c.Request.URL.Path
  query := c.Request.URL.RawQuery
  c.Next()

  cost := time.Since(start)
  lg.Info(path,
   zap.Int("status", c.Writer.Status()),
   zap.String("method", c.Request.Method),
   zap.String("path", path),
   zap.String("query", query),
   zap.String("ip", c.ClientIP()),
   zap.String("user-agent", c.Request.UserAgent()),
   zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
   zap.Duration("cost", cost),
  )
 }
}

// GinRecovery recover掉项目可能出现的panic，并使用zap记录相关日志
func GinRecovery(stack bool) gin.HandlerFunc {
 return func(c *gin.Context) {
  defer func() {
   if err := recover(); err != nil {
    // Check for a broken connection, as it is not really a
    // condition that warrants a panic stack trace.
    var brokenPipe bool
    if ne, ok := err.(*net.OpError); ok {
     if se, ok := ne.Err.(*os.SyscallError); ok {
      if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
       brokenPipe = true
      }
     }
    }

    httpRequest, _ := httputil.DumpRequest(c.Request, false)
    if brokenPipe {
     lg.Error(c.Request.URL.Path,
      zap.Any("error", err),
      zap.String("request", string(httpRequest)),
     )
     // If the connection is dead, we can't write a status to it.
     c.Error(err.(error)) // nolint: errcheck
     c.Abort()
     return
    }

    if stack {
     lg.Error("[Recovery from panic]",
      zap.Any("error", err),
      zap.String("request", string(httpRequest)),
      zap.String("stack", string(debug.Stack())),
     )
    } else {
     lg.Error("[Recovery from panic]",
      zap.Any("error", err),
      zap.String("request", string(httpRequest)),
     )
    }
    c.AbortWithStatus(http.StatusInternalServerError)
   }
  }()
  c.Next()
 }
}
```

记不得的看
[https://www.liwenzhou.com/posts/Go/zap-in-gin/] 关于在gin框架中使用zap
[https://darjun.github.io/2020/04/23/godailylib/zap/] 关于zap的更多细节

