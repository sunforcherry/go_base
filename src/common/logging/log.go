//high level log wrapper, so it can output different log based on level
package logging

import (
	"fmt"
	"io"
	"log"
	"os"
)

const (
	Ldate         = log.Ldate
	Llongfile     = log.Llongfile
	Lmicroseconds = log.Lmicroseconds
	Lshortfile    = log.Lshortfile
	LstdFlags     = log.LstdFlags
	Ltime         = log.Ltime
)

type (
	LogLevel int
	LogType  int
)

const (
	LOG_FATAL   = LogType(0x1)
	LOG_ERROR   = LogType(0x2)
	LOG_WARNING = LogType(0x4)
	LOG_DEBUG   = LogType(0x8)
	LOG_INFO    = LogType(0x10)
)

const (
	LOG_LEVEL_NONE  = LogLevel(0x0)
	LOG_LEVEL_FATAL = LOG_LEVEL_NONE | LogLevel(LOG_FATAL)
	LOG_LEVEL_ERROR = LOG_LEVEL_FATAL | LogLevel(LOG_ERROR)
	LOG_LEVEL_WARN  = LOG_LEVEL_ERROR | LogLevel(LOG_WARNING)
	LOG_LEVEL_DEBUG = LOG_LEVEL_WARN | LogLevel(LOG_DEBUG)
	LOG_LEVEL_INFO  = LOG_LEVEL_DEBUG | LogLevel(LOG_INFO)
	LOG_LEVEL_ALL   = LOG_LEVEL_INFO
)

var _log *logger = New()

func init() {
	SetFlags(Ldate | Ltime | Lshortfile)
}

func Logger() *log.Logger {
	return _log._log
}

func SetLevel(level LogLevel) {
	_log.SetLevel(level)
}
func GetLogLevel() LogLevel {
	return _log.level
}

func SetOutput(out io.Writer) {
	_log._log = log.New(out, _log._log.Prefix(), _log._log.Flags())
}

func SetFlags(flags int) {
	_log._log.SetFlags(flags)
}

func Info(v ...interface{}) {
	_log.Info(v...)
}

func Infof(format string, v ...interface{}) {
	_log.Infof(format, v...)
}

func Debug(v ...interface{}) {
	_log.Debug(v...)
}

func Debugf(format string, v ...interface{}) {
	_log.Debugf(format, v...)
}

func Warning(v ...interface{}) {
	_log.Warning(v...)
}

func Warningf(format string, v ...interface{}) {
	_log.Warningf(format, v...)
}

func Error(v ...interface{}) {
	_log.Error(v...)
}

func Errorf(format string, v ...interface{}) {
	_log.Errorf(format, v...)
}

func Fatal(v ...interface{}) {
	_log.Fatal(v...)
}

func Fatalf(format string, v ...interface{}) {
	_log.Fatalf(format, v...)
}

func SetLevelByString(level string) {
	_log.SetLevelByString(level)
}

func SetHighlighting(highlighting bool) {
	_log.SetHighlighting(highlighting)
}

type logger struct {
	_log         *log.Logger
	level        LogLevel
	highlighting bool
}

func (l *logger) SetHighlighting(highlighting bool) {
	l.highlighting = highlighting
}

func (l *logger) SetLevel(level LogLevel) {
	l.level = level
}

func (l *logger) SetLevelByString(level string) {
	l.level = StringToLogLevel(level)
}

func (l *logger) log(t LogType, v ...interface{}) {
	if l.level|LogLevel(t) != l.level {
		return
	}

	v1 := make([]interface{}, len(v)+2)
	logStr, logColor := LogTypeToString(t)
	if l.highlighting {
		v1[0] = "\033" + logColor + "m[" + logStr + "]"
		copy(v1[1:], v)
		v1[len(v)+1] = "\033[0m"
	} else {
		v1[0] = "[" + logStr + "]"
		copy(v1[1:], v)
		v1[len(v)+1] = ""
	}

	s := fmt.Sprintln(v1...)
	l._log.Output(4, s)
}

func (l *logger) logf(t LogType, format string, v ...interface{}) {
	if l.level|LogLevel(t) != l.level {
		return
	}

	logStr, logColor := LogTypeToString(t)
	var s string
	if l.highlighting {
		s = "\033" + logColor + "m[" + logStr + "] " + fmt.Sprintf(format, v...) + "\033[0m"
	} else {
		s = "[" + logStr + "] " + fmt.Sprintf(format, v...)
	}
	l._log.Output(4, s)
}

func (l *logger) Fatal(v ...interface{}) {
	l.log(LOG_FATAL, v...)
	os.Exit(-1)
}

func (l *logger) Fatalf(format string, v ...interface{}) {
	l.logf(LOG_FATAL, format, v...)
	os.Exit(-1)
}

func (l *logger) Error(v ...interface{}) {
	l.log(LOG_ERROR, v...)
}

func (l *logger) Errorf(format string, v ...interface{}) {
	l.logf(LOG_ERROR, format, v...)
}

func (l *logger) Warning(v ...interface{}) {
	l.log(LOG_WARNING, v...)
}

func (l *logger) Warningf(format string, v ...interface{}) {
	l.logf(LOG_WARNING, format, v...)
}

func (l *logger) Debug(v ...interface{}) {
	l.log(LOG_DEBUG, v...)
}

func (l *logger) Debugf(format string, v ...interface{}) {
	l.logf(LOG_DEBUG, format, v...)
}

func (l *logger) Info(v ...interface{}) {
	l.log(LOG_INFO, v...)
}

func (l *logger) Infof(format string, v ...interface{}) {
	l.logf(LOG_INFO, format, v...)
}

func StringToLogLevel(level string) LogLevel {
	switch level {
	case "fatal":
		return LOG_LEVEL_FATAL
	case "error":
		return LOG_LEVEL_ERROR
	case "warn":
		return LOG_LEVEL_WARN
	case "warning":
		return LOG_LEVEL_WARN
	case "debug":
		return LOG_LEVEL_DEBUG
	case "info":
		return LOG_LEVEL_INFO
	}
	return LOG_LEVEL_ALL
}

func LogTypeToString(t LogType) (string, string) {
	switch t {
	case LOG_FATAL:
		return "fatal", "[0;31"
	case LOG_ERROR:
		return "error", "[0;31"
	case LOG_WARNING:
		return "warning", "[0;33"
	case LOG_DEBUG:
		return "debug", "[0;36"
	case LOG_INFO:
		return "info", "[0;37"
	}
	return "unknown", "[0;37"
}

func New() *logger {
	return Newlogger(os.Stdout, "")
}

func Newlogger(w io.Writer, prefix string) *logger {
	return &logger{log.New(w, prefix, LstdFlags), LOG_LEVEL_ALL, true}
}
