package logger

import (
	"fmt"
	"runtime"
	"strings"
	"time"

	"go.uber.org/zap"
)

/*
DEBUG --> may print either just string ('preparing public data') or json object, make sure json objects printed
this case do not exceed 50 json objects, you can do this by making sure you use this only in places where you persume
json printed is smalled in size

For DEBUG levels, make sure they printed inside code, not start or end of a function.
*/
func LogDebug(msg string, session string, val interface{}) {
	if configuration.IsEnabled {
		_, file, line, _ := runtime.Caller(1)
		// ZapLogger.Debug(formatLogMessage(DEBUG, file, fmt.Sprintf("%v", line), session, msg), zap.Any("data", val))
		logMsg := formatLogMsg(DEBUG, file, session, msg, line)
		ZapLogger.Debug(logMsg, zap.Any("data", val))
		//fmt.Println(logMsg)
	}
	//formatLogMsg(DEBUG, file, session, msg, line), zap.Any("data", val)
}

/*
DEBUG_2 --> may print both string and or json object of any size
*/
func LogDebug2(msg string, session string, val interface{}) {
	if configuration.IsEnabled {
		_, file, line, _ := runtime.Caller(1)
		//ZapLogger.Debug(formatLogMessage(DEBUG2, file, fmt.Sprintf("%v", line), session, msg), zap.Any("data", val))
		logMsg := formatLogMsg(DEBUG2, file, session, msg, line)
		ZapLogger.Debug(logMsg, zap.Any("data", val))
		//fmt.Println(logMsg)
	}
}

// TODO: we are not printing any data using val so it is just there to keep backward compatibility becasue alot of logs
// have empty val parameter

/*
INFO --> must always print just string (like, "file entered" or "function xyz called" or similiar)
*/
func LogInfo(msg string, session string, val ...any) {
	if configuration.IsEnabled {
		_, file, line, _ := runtime.Caller(1)
		//ZapLogger.Info(formatLogMessage(INFO, file, fmt.Sprintf("%v", line), session, msg))
		logMsg := formatLogMsg(INFO, file, session, msg, line)
		ZapLogger.Info(logMsg)
		//fmt.Println(logMsg)
	}
}

/*
WARNING --> prints exact warning in string format, this is used when non breaking exception may happen because of user input
*/
func LogWarning(session string, val interface{}) {
	if configuration.IsEnabled {
		_, file, line, _ := runtime.Caller(1)
		//ZapLogger.Warn(formatLogMessage(WARNING, file, fmt.Sprintf("%v", line), session, ""), zap.Any("data", val))
		logMsg := formatLogMsg(WARNING, file, session, "", line)
		ZapLogger.Warn(logMsg, zap.Any("data", val))
		//fmt.Println(logMsg)
	}
}

/*
ERROR --> prints exact warning in string format, this is used when breaking exception may happen because of user input
*/
func LogError(session string, val interface{}) {
	if configuration.IsEnabled {
		_, file, line, _ := runtime.Caller(1)
		//ZapLogger.Error(formatLogMessage(ERROR, file, fmt.Sprintf("%v", line), session, ""), zap.Any("data", val))
		logMsg := formatLogMsg(ERROR, file, session, "", line)
		ZapLogger.Error(logMsg, zap.Any("data", val))
		//fmt.Println(logMsg)
	}
}

func LogFatal(session string, val interface{}) {
	if configuration.IsEnabled {
		_, file, line, _ := runtime.Caller(1)
		//ZapLogger.Fatal(formatLogMessage(FATAL, file, fmt.Sprintf("%v", line), session, ""), zap.Any("data", val))
		logMsg := formatLogMsg(FATAL, file, session, "", line)
		ZapLogger.Fatal(logMsg, zap.Any("data", val))
		//fmt.Println(logMsg)
	}
}

func LogPanic(session string, val interface{}) {
	if configuration.IsEnabled {
		_, file, line, _ := runtime.Caller(1)
		//ZapLogger.Panic(formatLogMessage(PANIC, file, fmt.Sprintf("%v", line), session, ""), zap.Any("data", val))
		logMsg := formatLogMsg(PANIC, file, session, "", line)
		ZapLogger.Panic(logMsg, zap.Any("data", val))
		//fmt.Println(logMsg)
	}
}

func formatLogMessage(level, path, line, session, message string) string {
	formatedPath := formatPath(path)
	return fmt.Sprintf("\n%v =>    ::%v    ::%v[%v]    ::%v    ::%v    ::%v\n", level, internalVersion, formatedPath, line, time.Now().Format(time.RFC1123), session, message)
}

// creates short path
func formatPath(path string) string {
	formattedPath := path
	splitedString := strings.Split(path, "/")
	pathLen := len(splitedString)

	if pathLen > 2 {
		formattedPath = fmt.Sprintf("%v", strings.Join(splitedString[pathLen-2:], "/"))
	}

	return formattedPath
}

func formatLogMsg(level, path, sessionId, msg string, line int) string {
	filePath := formatPath(path)
	curDateStr := time.Now().Format("2006-01-02 15:04:05.000")
	fileDetail := fmt.Sprintf("%s:%d", filePath, line)
	//return fmt.Sprintf("%-8s  =>  %-35s  :::  %s   %s   %s\n", strings.ToUpper(level), fileDetail, curDateStr, sessionId, msg)
	return fmt.Sprintf("%-8s    =>  %-35s  :::  %s   %s   %s", strings.ToUpper(level), fileDetail, curDateStr, sessionId, msg)
}
