package logrus_log

import (
	"runtime"

	"github.com/sirupsen/logrus"
	"technodom/internal/util/logger"
)

type LogrusLogger struct {
	logrus *logrus.Logger
}

func New() logger.Logger {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})

	return &LogrusLogger{
		log,
	}
}

func (logger *LogrusLogger) Info(message string, args map[string]interface{}) {
	fields := logrus.Fields{
		"method": checkFuncName(2),
	}
	for key, value := range args {
		fields[key] = value
	}

	logger.logrus.WithFields(fields).Info(message)
}

func (logger *LogrusLogger) Error(message string, args map[string]interface{}) {
	fields := logrus.Fields{
		"method": checkFuncName(2),
	}
	for key, value := range args {
		fields[key] = value
	}

	logger.logrus.WithFields(fields).Error(message)
}

func checkFuncName(skip int) string {
	pc, _, _, ok := runtime.Caller(skip)
	if !ok {
		return "unknown"
	}
	me := runtime.FuncForPC(pc)
	if me == nil {
		return "unnamed"
	}
	return me.Name()
}
