package log

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/sjmillington/golang-microservices/github-api/src/api/config"
)

var (
	Log *logrus.Logger
)

func init() {

	level, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		level = logrus.DebugLevel
	}

	Log = &logrus.Logger{
		Level: level,
		Out:   os.Stdout,
	}

	if config.IsProduction() {
		Log.Formatter = &logrus.JSONFormatter{}
	} else {
		Log.Formatter = &logrus.TextFormatter{}
	}

}

//can use tags to record customer IDs and REST calls to charge customers for services used.
//i.e log clientId:123421423 status:success. this will get stored in ES and can be searched for charges
func Info(msg string, tags ...string) {
	if Log.Level < logrus.InfoLevel {
		return
	}

	Log.WithFields(parseFields(tags...)).Info(msg)

}

func Debug(msg string, tags ...string) {
	if Log.Level < logrus.DebugLevel {
		return
	}

	Log.WithFields(parseFields(tags...)).Debug(msg)

}

func Error(msg string, tags ...string) {
	if Log.Level < logrus.ErrorLevel {
		return
	}

	Log.WithFields(parseFields(tags...)).Error(msg)

}

func parseFields(tags ...string) logrus.Fields {
	result := make(logrus.Fields, len(tags))

	for _, tag := range tags {
		if strings.Contains(tag, ":") {
			els := strings.Split(tag, ":")
			result[strings.TrimSpace(els[0])] = strings.TrimSpace(els[1])
		}
	}

	return result
}
