package Utils

import (
	"time"
)

var locale, _ = time.LoadLocation("Asia/Jakarta")

func Now() time.Time {
	var now time.Time
	now = time.Now().In(locale)
	return now
}
