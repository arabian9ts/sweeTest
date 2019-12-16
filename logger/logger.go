package logger

import (
	"github.com/sirupsen/logrus"

	"github.com/arabian9ts/sweeTest/app/domain/value"
)

type Logger func(msg string, entry interface{})

func LogrusWriter() Logger {
	return func(msg string, entry interface{}) {
		var fields logrus.Fields
		switch data := entry.(type) {
		case value.QueryLogEntry:
			fields = logrus.Fields{
				"timestamp": data.Timestamp,
				"query":     data.Query,
			}
		case value.AccessLogEntry:
			fields = logrus.Fields{
				"timestamp": data.Timestamp,
				"userId":    data.UserId,
				"userType":  data.UserType,
			}
		default:
			logrus.Error("Log Type is Invalid")
		}
		logrus.WithFields(fields).Info(msg)
	}
}
