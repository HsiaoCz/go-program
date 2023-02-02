# zap

zap是一个日志库，由uber开源的
之所以用，相较于logrus，主要是性能优化和内存分配更好

安装
```go
go get go.uber.org/zap
```

zap库的使用与其他的日志库非常相似。先创建一个logger，然后调用各个级别的方法记录日志（Debug/Info/Error/Warn）。zap提供了几个快速创建logger的方法，zap.NewExample()、zap.NewDevelopment()、zap.NewProduction()，还有高度定制化的创建方法zap.New()。创建前 3 个logger时，zap会使用一些预定义的设置，它们的使用场景也有所不同。Example适合用在测试代码中，Development在开发环境中使用，Production用在生成环境。

zap底层 API 可以设置缓存，所以一般使用defer logger.Sync()将缓存同步到文件中。

由于fmt.Printf之类的方法大量使用interface{}和反射，会有不少性能损失，并且增加了内存分配的频次。zap为了提高性能、减少内存分配次数，没有使用反射，而且默认的Logger只支持强类型的、结构化的日志。必须使用zap提供的方法记录字段。zap为 Go 语言中所有的基本类型和其他常见类型都提供了方法。这些方法的名称也比较好记忆，zap.Type（Type为bool/int/uint/float64/complex64/time.Time/time.Duration/error等）就表示该类型的字段，zap.Typep以p结尾表示该类型指针的字段，zap.Types以s结尾表示该类型切片的字段。如：

    zap.Bool(key string, val bool) Field：bool字段
    zap.Boolp(key string, val *bool) Field：bool指针字段；
    zap.Bools(key string, val []bool) Field：bool切片字段。

当然也有一些特殊类型的字段：
    zap.Any(key string, value interface{}) Field：任意类型的字段；
    zap.Binary(key string, val []byte) Field：二进制串的字段。

当然，每个字段都用方法包一层用起来比较繁琐。zap也提供了便捷的方法SugarLogger，可以使用printf格式符的方式。调用logger.Sugar()即可创建SugaredLogger。SugaredLogger的使用比Logger简单，只是性能比Logger低 50% 左右，可以用在非热点函数中。调用SugarLogger以f结尾的方法与fmt.Printf没什么区别，如例子中的Infof。同时SugarLogger还支持以w结尾的方法，这种方式不需要先创建字段对象，直接将字段名和值依次放在参数中即可，如例子中的Infow。

zap默认的输出格式是JSON格式

## 1 、记录层级关系

使用zap.Namespace(key string)Field构建一个命令空间。后续的Field都记录在此命令空间中
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

