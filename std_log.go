package loginterface

import (
	"fmt"
	"log"
	"os"
)

const logDepth = 2

// The NewDefaultLog return a new instance of "log" to stdout
func NewDefaultLog() Logger {
	return NewStdLog(log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile))
}

type StdLogger interface {
	Printf(format string, v ...interface{})
	Println(v ...interface{})
	Print(v ...interface{})
	Output(calldepth int, s string) error
}

func NewStdLog(logger StdLogger) Logger {
	return &stdLog{
		logger: logger,
	}
}

type stdLog struct {
	logger StdLogger
}

func (s *stdLog) Debug(args ...interface{}) {
	s.logger.Output(logDepth, "DEBUG: "+fmt.Sprint(args...))
}

func (s *stdLog) Info(args ...interface{}) {
	s.logger.Output(logDepth, "INFO: "+fmt.Sprint(args...))
}

func (s *stdLog) Warn(args ...interface{}) {
	s.logger.Output(logDepth, "WARN: "+fmt.Sprint(args...))
}

func (s *stdLog) Error(args ...interface{}) {
	s.logger.Output(logDepth, "ERROR: "+fmt.Sprint(args...))
}

func (s *stdLog) Debugln(args ...interface{}) {
	s.logger.Output(logDepth, "DEBUG: "+fmt.Sprintln(args...))
}

func (s *stdLog) Infoln(args ...interface{}) {
	s.logger.Output(logDepth, "INFO: "+fmt.Sprintln(args...))
}

func (s *stdLog) Warnln(args ...interface{}) {
	s.logger.Output(logDepth, "WARN: "+fmt.Sprintln(args...))
}

func (s *stdLog) Errorln(args ...interface{}) {
	s.logger.Output(logDepth, "ERROR: "+fmt.Sprintln(args...))
}

func (s *stdLog) Debugf(format string, args ...interface{}) {
	s.logger.Output(logDepth, "DEBUG: "+fmt.Sprintf(format, args...))
}

func (s *stdLog) Infof(format string, args ...interface{}) {
	s.logger.Output(logDepth, "INFO: "+fmt.Sprintf(format, args...))
}

func (s *stdLog) Warnf(format string, args ...interface{}) {
	s.logger.Output(logDepth, "WARN: "+fmt.Sprintf(format, args...))
}

func (s *stdLog) Errorf(format string, args ...interface{}) {
	s.logger.Output(logDepth, "ERROR: "+fmt.Sprintf(format, args...))
}
