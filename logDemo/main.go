package main

import (
	"github.com/cihub/seelog"
)

func main() {
	logger, _ := seelog.LoggerFromConfigAsFile("seelog.xml")
	logger.Debug("调试信息的日志") //
	logger.Info("一般信息的日志")
	logger.Warn("警告信息的日志")
	logger.Error("错误信息的日志")
	logger.Critical("致命错误的日志")

}
