package utils

import (
	"fmt"
	"time"
)

func TimestampToTime(timestamp int64) *time.Time {
	t := time.Unix(timestamp, 0)
	return &t
}

func ToString(v interface{}) string {
	return fmt.Sprint(v)
}
