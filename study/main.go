package main

import "github.com/wulw1028/xiego/study/logs"

func main() {
	fileLog := logs.NewFileLog("debug", "./", "wlw.log", 10*1024*1024)
	//defer fileLog.Close()
	fileLog.Debug("fileLog msg")
	fileLog.Error("fileLog msg")
}
