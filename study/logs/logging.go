package logs

import (
	"errors"
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
	"time"
)

type LogLevel uint16

const (
	UNKNOWN LogLevel = iota
	DEBUG
	TRACE
	INFO
	WARNING
	ERROR
	FATAL
)

func parserLogLevel(s string) (LogLevel, error) {
	s = strings.ToUpper(s)
	switch s {
	case "DEBUG":
		return DEBUG, nil
	case "TRACE":
		return TRACE, nil
	case "INFO":
		return INFO, nil
	case "WARNING":
		return WARNING, nil
	case "ERROR":
		return ERROR, nil
	case "FATAL":
		return FATAL, nil
	default:
		err := errors.New("无效的常数")
		return UNKNOWN, err
	}
}

func getLogString(lv LogLevel) string {
	switch lv {
	case DEBUG:
		return "DEBUG"
	case TRACE:
		return "TRACE"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARNING"
	case ERROR:
		return "ERROR"
	case FATAL:
		return "FATAL"
	}
	return ""
}

type ConsoleLog struct {
	Level LogLevel
}

type FileLog struct {
	Level         LogLevel
	filePath      string
	fileName      string
	maxFileSize   int64
	fileObject    *os.File
	errFileObject *os.File
}

func NewConsoleLog(levelStr string) ConsoleLog {
	level, err := parserLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	return ConsoleLog{
		Level: level,
	}
}

func NewFileLog(levelStr, filePath, fileName string, maxFileSize int64) *FileLog {
	level, err := parserLogLevel(levelStr)
	if err != nil {
		panic(err)
	}
	f := &FileLog{
		Level:       level,
		filePath:    filePath,
		fileName:    fileName,
		maxFileSize: maxFileSize,
	}
	err = f.initFile()
	if err != nil {
		panic(err)
	}
	return f
}

func getFuncInfo(skip int) (funcName, file string, line int) {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		fmt.Println("runtime.Caller(skip) error")
	}
	funcName = strings.Split(runtime.FuncForPC(pc).Name(), ".")[0]
	return
}

func (c ConsoleLog) log(lv LogLevel, msg string, errList []interface{}) {
	if c.Level <= lv {
		for _, value := range errList {
			msg = fmt.Sprint(msg, value)
		}

		t := time.Now()
		timeString := t.Format("2006-01-02 15:04:05")
		funcName, fileName, lineNo := getFuncInfo(3)
		fmt.Printf("[%s] [%s] [Func(%s) Path(%s) Line(%d)] : %s\n", timeString, getLogString(lv), funcName, fileName, lineNo, msg)
	}
}

func (f FileLog) file(lv LogLevel, msg string, errList []interface{}) {
	if f.Level <= lv {
		for _, value := range errList {
			msg = fmt.Sprintln(msg, value)
		}

		t := time.Now()
		timeString := t.Format("2006-01-02 15:04:05")
		funcName, fileName, lineNo := getFuncInfo(3)
		fmt.Fprintf(f.fileObject, "[%s] [%s] [Func(%s) Path(%s) Line(%d)] : %s\n", timeString, getLogString(lv), funcName, fileName, lineNo, msg)
		if lv >= ERROR {
			fmt.Fprintf(f.errFileObject, "[%s] [%s] [Func(%s) Path(%s) Line(%d)] : %s\n", timeString, getLogString(lv), funcName, fileName, lineNo, msg)
		}
	}
}

func (f *FileLog) initFile() error {
	fullFileName := path.Join(f.filePath, f.fileName)
	errFullFileName := fullFileName + ".err"
	fileObject, err := os.OpenFile(fullFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return err
	}

	errFileObject, err := os.OpenFile(errFullFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return err
	}
	f.fileObject = fileObject
	f.errFileObject = errFileObject
	return nil
}

func (f FileLog) CheckFileSize(file *os.File) bool {
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println(err)
	}
	return f.maxFileSize >= fileInfo.Size()
}

func (f FileLog) Close() {
	f.fileObject.Close()
	f.errFileObject.Close()
}

func (c ConsoleLog) Debug(msg string, errList ...interface{}) {
	c.log(DEBUG, msg, errList)
}

func (c ConsoleLog) Error(msg string, errList ...interface{}) {
	c.log(ERROR, msg, errList)
}

func (f FileLog) Debug(msg string, errList ...interface{}) {
	f.file(DEBUG, msg, errList)
}

func (f FileLog) Error(msg string, errList ...interface{}) {
	f.file(ERROR, msg, errList)
}
