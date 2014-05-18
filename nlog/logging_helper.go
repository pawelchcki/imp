package nlog

import (
	"log"
	"os"
)

var emergLogger, alertLogger, critLogger, errLogger, warningLogger, noticeLogger, infoLogger, debugLogger *log.Logger

func init() {
	useStandardLogging()
}

func useStandardLogging() {
	flag := 0
	emergLogger = log.New(os.Stdout, "EMERG ", flag)
	alertLogger = log.New(os.Stdout, "ALERT ", flag)
	critLogger = log.New(os.Stdout, "CRIT ", flag)
	errLogger = log.New(os.Stdout, "ERR ", flag)
	warningLogger = log.New(os.Stdout, "WARNING ", flag)
	noticeLogger = log.New(os.Stdout, "NOTICE ", flag)
	infoLogger = log.New(os.Stdout, "INFO ", flag)
	debugLogger = log.New(os.Stdout, "DEBUG ", flag)
}

func Emerg(v ...interface{}) {
	emergLogger.Print(v...)
}
func Alert(v ...interface{}) {
	alertLogger.Print(v...)
}
func Crit(v ...interface{}) {
	critLogger.Print(v...)
}
func Err(v ...interface{}) {
	errLogger.Print(v...)
}
func Warning(v ...interface{}) {
	warningLogger.Print(v...)
}
func Notice(v ...interface{}) {
	noticeLogger.Print(v...)
}
func Info(v ...interface{}) {
	infoLogger.Print(v...)
}
func Debug(v ...interface{}) {
	debugLogger.Print(v...)
}

func Emergf(f string, v ...interface{}) {
	emergLogger.Printf(f, v...)
}
func Alertf(f string, v ...interface{}) {
	alertLogger.Printf(f, v...)
}
func Critf(f string, v ...interface{}) {
	critLogger.Printf(f, v...)
}
func Errf(f string, v ...interface{}) {
	errLogger.Printf(f, v...)
}
func Warningf(f string, v ...interface{}) {
	warningLogger.Printf(f, v...)
}
func Noticef(f string, v ...interface{}) {
	noticeLogger.Printf(f, v...)
}
func Infof(f string, v ...interface{}) {
	infoLogger.Printf(f, v...)
}
func Debugf(f string, v ...interface{}) {
	debugLogger.Printf(f, v...)
}
