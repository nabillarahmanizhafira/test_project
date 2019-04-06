package log

import (
	"bytes"
	"runtime"
	"strconv"
	"strings"
)

type (
	// Logger is interface to any implementation of log
	Logger interface {
		Println(...interface{})
		Printf(string, ...interface{})
		Error(error, ...interface{})
		Errorf(error, string, ...interface{})
		Panic(...interface{})
		Fatal(...interface{})
		Debugln(...interface{})
		Debugf(string, ...interface{})
		SetVariable(Fields) Logger
	}

	//Fields type for standardize custom field input
	Fields map[string]interface{}
)

// singleton instantiation
var (
	logger Logger
)

//Init logger with logrus
func Init() {
	logger = NewLogrus(true)
}

// Println output Info
func Println(msg ...interface{}) {
	logger.Println(msg...)
}

// Printf info with format
func Printf(fmt string, msg ...interface{}) {
	logger.Printf(fmt, msg...)
}

// Error output error
func Error(err error, msg ...interface{}) {
	logger.Error(err, msg...)
}

// Errorf output error with format
func Errorf(err error, fmt string, msg ...interface{}) {
	logger.Errorf(err, fmt, msg...)
}

// Panic log
func Panic(msg ...interface{}) {
	logger.Panic(msg...)
}

// Fatal log
func Fatal(msg ...interface{}) {
	logger.Fatal(msg...)
}

// Debugln log debug
func Debugln(msg ...interface{}) {
	logger.Debugln(msg...)
}

//Debugf debug with format
func Debugf(fmt string, msg ...interface{}) {
	logger.Debugf(fmt, msg...)
}

//SetVariable returns logger object with variables
func SetVariable(v map[string]interface{}) Logger {
	return logger.SetVariable(v)
}

// getParentCaller is function to determine current file which produce the log
func getParentCaller(callDepth int) string {
	var buffer bytes.Buffer

	pc, file, line, ok := runtime.Caller(callDepth)
	fnc := runtime.FuncForPC(pc)
	if ok {
		separator := strings.Split(file, "test_project/")
		if len(separator) > 1 {
			buffer.WriteString(separator[1])
		} else {
			buffer.WriteString(file)
		}
		buffer.WriteString(":")
		buffer.WriteString(strconv.Itoa(line))
		buffer.WriteString(" @")

		funcName := fnc.Name()
		funcName = funcName[strings.LastIndex(funcName, ".")+1:] // remove the function detailed path
		buffer.WriteString(funcName)
	}

	return buffer.String()
}
