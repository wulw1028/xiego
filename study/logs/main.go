package main

import "github.com/wulw1028/xiego/study/logs"

var log logs.Logger

func main() {
	//log = logs.NewFileLog("debug", "./", "wlw.log", 10*1024*1024)
	log = logs.NewConsoleLog("debug")
	//defer fileLog.Close()
	log.Debug("fileLog msg")
	log.Error("fileLog msg")
}
