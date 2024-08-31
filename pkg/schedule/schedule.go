package schedule

import (
	"errors"
	"log"
	"reflect"
	"strings"
	"time"
)

func getStructName(i interface{}) string {
	return strings.TrimPrefix(reflect.TypeOf(i).String(), "*")
}

// Schedule schedules tasks to run every specified duration (in minutes).
func Schedule(d time.Duration, i ScheduleInterface, errorChan chan<- error) error {
	if d <= 0 {
		return errors.New("invalid duration, must be greater than 0")
	}
	ticker := time.NewTicker(d * time.Minute) // Changed from Second to Minute
	log.Printf("[Schedule] Scheduled task [%s] to run every %d minutes", getStructName(i), int(d.Minutes()))
	go func() {
		for range ticker.C {
			log.Printf("[Schedule] Running task [%s]", getStructName(i))
			if err := i.Do(); err != nil {
				errorChan <- err
			}
		}
	}()
	return nil
}
