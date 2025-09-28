package loggerutils

import "go.uber.org/zap"

var Logger *zap.Logger

func init() {
	// 初始化日志系统
	var err error
	Logger, err = zap.NewProduction()
	if err != nil {
		panic("无法初始化日志系统: " + err.Error())
	}
	defer Logger.Sync() // 确保日志被刷新到输出
}
