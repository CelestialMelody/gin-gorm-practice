package temp

// package log

import (
	"fmt"
	"log"
	"os"
	"runtime"
)

type Level int

var (
	F                  *os.File
	DefaultCallerDepth = 2  // 调用层级
	DefaultPrefix      = "" // 日志前缀
	loggo              *log.Logger
	logPrefix          = "" // 日志前缀
	levelFlags         = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

//func SetUp() {
//	var err error
//	filePath := log.getLogFilePath()
//	fileName := log.getLogFileName()
//	F, err = log.openLogFile(fileName, filePath)
//	if err != nil {
//		log.Fatalf("log.Setup err: %v", err)
//	}
//	loggo = log.New(F, DefaultPrefix, log.LstdFlags)
//}

// 设置日志前缀; 有点难理解
func setPrefix(level Level) {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth) // 获取调用层级
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], file, line) // 格式化前缀
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}
	loggo.SetPrefix(logPrefix)
}

func Debug(v ...interface{}) {
	setPrefix(DEBUG)
	loggo.Println(v)
}

func Info(v ...interface{}) {
	setPrefix(INFO)
	loggo.Println(v)
}

func Warn(v ...interface{}) {
	setPrefix(WARNING)
	loggo.Println(v)
}

func Error(v ...interface{}) {
	setPrefix(ERROR)
	loggo.Println(v)
}

func Fatal(v ...interface{}) {
	setPrefix(FATAL)
	loggo.Fatalln(v)
}
