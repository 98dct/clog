package main

import (
	"clog"
	"log"
	"os"
)

func main() {
	clog.Info("std log")
	clog.SetOptions(clog.WithLevel(clog.DebugLevel))
	clog.Debug("change std log to debug level")
	clog.SetOptions(clog.WithFormatter(&clog.JsonFormatter{IgnoreBasicFields: false}))
	clog.Debug("log in json format")
	clog.Info("another log in json format")
	// 输出到文件
	fd, err := os.OpenFile("test.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln("create file test.log failed")
	}
	defer fd.Close()
	l := clog.New(clog.WithLevel(clog.InfoLevel),
		clog.WithOutput(fd), clog.WithFormatter(&clog.JsonFormatter{IgnoreBasicFields: false}))
	l.Info("custom log with json formatter")
}
