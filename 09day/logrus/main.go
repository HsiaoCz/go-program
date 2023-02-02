package main

import "github.com/sirupsen/logrus"

func main() {
	logrus.SetLevel(logrus.TraceLevel)
	logrus.Trace("trace message")
	logrus.Debug("debug msg")
	logrus.Info("info msg")
	logrus.Warn("warn msg")
	logrus.Error("error msg")
	logrus.Fatal("fatal msg")
	logrus.Panic("panic msg")

	logrus.SetReportCaller(true)
	logrus.Info("info msg")

	// 添加字段
	logrus.WithFields(logrus.Fields{
		"name": "zhangsan",
		"age":  20,
	}).Info("info msg")

	// 为所有日志都加上user_id和ip字段
	requestLogger := logrus.WithFields(logrus.Fields{
		"user_id": 1900,
		"ip":      "127.0.0.1:9090",
	})
	requestLogger.Info("info massage")
	requestLogger.Error("error msg")

}
