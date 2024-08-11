package helpers

import "time"

func CreateExpirationTime(exp Expiration) int64 {
	now := time.Now().Unix()
	switch exp {
	case OneMinute:
		return now + 60
	case OneHour:
		return now + 3600
	case OneDay:
		return now + 86400
	case OneWeek:
		return now + 604800
	case OneMonth:
		return now + 2592000
	default:
		return 0
	}
}
