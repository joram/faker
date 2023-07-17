package faker

import (
	"math/rand"
	"time"
)

func GetAge(seed int64) int {
	return getAge().GetRandomValue(seed)
}

func GetBirthday(seed int64) time.Time {
	rand.Seed(seed)
	year := time.Now().Year() - GetAge(seed)
	minDate := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC).Unix()
	maxDate := time.Date(year, 12, 31, 0, 0, 0, 0, time.UTC).Unix()
	delta := maxDate - minDate
	sec := rand.Int63n(delta) + minDate
	return time.Unix(sec, 0)
}
