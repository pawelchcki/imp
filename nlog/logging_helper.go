package nlog

import (
	"log"
	"log/syslog"
	"os"
)

func init() {
	initializeStdOutLoggers()
}

var prioStringMap = map[syslog.Priority]string{
	syslog.LOG_EMERG:   "EMERG",
	syslog.LOG_ALERT:   "ALERT",
	syslog.LOG_CRIT:    "CRIT",
	syslog.LOG_ERR:     "ERR",
	syslog.LOG_WARNING: "WARNING",
	syslog.LOG_NOTICE:  "NOTICE",
	syslog.LOG_INFO:    "INFO",
	syslog.LOG_DEBUG:   "DEBUG",
}

var prioLoggers map[syslog.Priority]*log.Logger = make(map[syslog.Priority]*log.Logger)

func initializeStdOutLoggers() {
	flag := 0
	for prio, lvl := range prioStringMap {
		prioLoggers[prio] = log.New(os.Stdout, lvl+" ", flag)
	}
}

func levelPrintf(prio syslog.Priority, f string, v ...interface{}) {
	prioLoggers[prio].Printf(f, v...)
}

func levelPrint(prio syslog.Priority, v ...interface{}) {
	prioLoggers[prio].Print(v...)
}

func Emerg(v ...interface{}) {
	levelPrint(syslog.LOG_EMERG, v...)
}
func Alert(v ...interface{}) {
	levelPrint(syslog.LOG_ALERT, v...)
}
func Crit(v ...interface{}) {
	levelPrint(syslog.LOG_CRIT, v...)
}
func Err(v ...interface{}) {
	levelPrint(syslog.LOG_ERR, v...)
}
func Warning(v ...interface{}) {
	levelPrint(syslog.LOG_WARNING, v...)
}
func Notice(v ...interface{}) {
	levelPrint(syslog.LOG_NOTICE, v...)
}
func Info(v ...interface{}) {
	levelPrint(syslog.LOG_INFO, v...)
}
func Debug(v ...interface{}) {
	levelPrint(syslog.LOG_DEBUG, v...)
}

func Emergf(f string, v ...interface{}) {
	levelPrintf(syslog.LOG_EMERG, f, v...)
}
func Alertf(f string, v ...interface{}) {
	levelPrintf(syslog.LOG_ALERT, f, v...)
}
func Critf(f string, v ...interface{}) {
	levelPrintf(syslog.LOG_CRIT, f, v...)
}
func Errf(f string, v ...interface{}) {
	levelPrintf(syslog.LOG_ERR, f, v...)
}
func Warningf(f string, v ...interface{}) {
	levelPrintf(syslog.LOG_WARNING, f, v...)
}
func Noticef(f string, v ...interface{}) {
	levelPrintf(syslog.LOG_NOTICE, f, v...)
}
func Infof(f string, v ...interface{}) {
	levelPrintf(syslog.LOG_INFO, f, v...)
}
func Debugf(f string, v ...interface{}) {
	levelPrintf(syslog.LOG_DEBUG, f, v...)
}
