package value

import (
	"time"

	"github.com/arabian9ts/sweeTest/app/domain/model"
)

type (
	QueryLogEntry struct {
		Timestamp time.Time
		Query     string
	}

	AccessLogEntry struct {
		Timestamp time.Time
		UserId    int64
		UserType  model.UserType
	}
)
