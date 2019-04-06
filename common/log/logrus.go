package log

import (
	"os"

	"github.com/sirupsen/logrus"
)

//Log custom logger. This is an object which implements Logger interface
type Log struct {
	log    *logrus.Logger
	fields Fields
}

//NewLogrus return log object
func NewLogrus(enableDebug bool) Logger {
	logObj := &Log{}
	if enableDebug {
		logrus.SetLevel(logrus.DebugLevel)
	}

	// set formatter
	accLogger := logrus.New()
	formatter := &logrus.JSONFormatter{}

	accLogger.Out = os.Stdout
	accLogger.Formatter = formatter
	logObj.log = accLogger
	return logObj
}

// Println output Info
func (lg *Log) Println(msg ...interface{}) {
	lg.log.WithFields(lg.getFields()).Infoln(msg...)
}

// Printf info with format
func (lg *Log) Printf(fmt string, msg ...interface{}) {
	lg.log.WithFields(lg.getFields()).Infof(fmt, msg...)
}

// Error output error
func (lg *Log) Error(err error, msg ...interface{}) {
	lg.log.WithFields(lg.getFields(err)).Errorln(msg...)
}

// Errorf output error with format
func (lg *Log) Errorf(err error, fmt string, msg ...interface{}) {
	lg.log.WithFields(lg.getFields(err)).Errorf(fmt, msg...)
}

// Panic log
func (lg *Log) Panic(msg ...interface{}) {
	lg.log.WithFields(lg.getFields()).Panicln(msg...)
}

// Fatal log
func (lg *Log) Fatal(msg ...interface{}) {
	lg.log.WithFields(lg.getFields()).Fatalln(msg...)
}

// Debugln log debug
func (lg *Log) Debugln(msg ...interface{}) {
	lg.log.WithFields(lg.getFields()).Debugln(msg...)
}

//Debugf debug with format
func (lg *Log) Debugf(fmt string, msg ...interface{}) {
	lg.log.WithFields(lg.getFields()).Debugf(fmt, msg...)
}

func (lg *Log) getFields(err ...error) (fields logrus.Fields) {
	fields = logrus.Fields{}

	if len(err) > 0 {
		fields["err"] = err[0]
	}

	callDepth := 4
	if lg.fields != nil {
		for k, v := range lg.fields {
			fields[k] = v
		}
		callDepth = 3
	}

	fields["file"] = getParentCaller(callDepth)

	return fields
}

//SetVariable returns logger object with variables
func (lg *Log) SetVariable(v Fields) Logger {
	varLog := &Log{
		fields: v,
		log:    lg.log,
	}

	return varLog
}
