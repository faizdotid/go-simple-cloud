package helpers

type Expiration string

const (
	OneMinute Expiration = "1m"
	OneHour   Expiration = "1h"
	OneDay    Expiration = "1d"
	OneWeek   Expiration = "1w"
	OneMonth  Expiration = "1M"
	Default   Expiration = "default"
)
