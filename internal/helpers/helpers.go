package helpers

import (
	"errors"
	"time"
)

func CreateExpirationTime(exp Expiration) (int64, error) {
	now := time.Now().Unix()
	switch exp {
	case OneMinute:
		return now + 60, nil
	case OneHour:
		return now + 3600, nil
	case OneDay:
		return now + 86400, nil
	case OneWeek:
		return now + 604800, nil
	case OneMonth:
		return now + 2592000, nil
	default:
		return 0, errors.New("invalid expiration time")
	}
}
