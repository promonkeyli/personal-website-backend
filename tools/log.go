package tools

import (
	"fmt"
	"time"

	"github.com/fatih/color"
)

var colorMap = map[string]*color.Color{
	"debug": color.New(color.FgWhite),
	"info":  color.New(color.FgGreen),
	"warn":  color.New(color.FgYellow),
	"error": color.New(color.FgRed),
}

// generateLog: 日志内容生成
func generateLog(info string, level string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	logMsg := fmt.Sprintf("[log-%s] %s %s", level, timestamp, info)
	c := colorMap[level]
	c.Println(logMsg)
}

// 控制信息常规日志输出:白色
func LogDebug(info string) {
	generateLog(info, "debug")
}

// 控制信息着重/成功日志输出:绿色
func LogInfo(info string) {
	generateLog(info, "info")
}

// 控制信息警告日志输出: 黄色
func LogWarn(info string) {
	generateLog(info, "warn")
}

// 控制信息错误日志输出: 红色
func LogError(info string) {
	generateLog(info, "error")
}
