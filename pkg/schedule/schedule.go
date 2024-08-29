package schedule

import (
	"errors"
	"time"
)

func Schedule(d time.Duration, i ScheduleInterface) error {
	if d <= 0 {
		return errors.New("invalid duration, must be greater than 0")
	}

	ticker := time.NewTicker(d * time.Second)
	defer ticker.Stop()

	errChan := make(chan error, 1)
	doneChan := make(chan bool, 1)
	go func() {
		for {
			select {
			case <-doneChan:
				return
			case <-ticker.C:
				err := i.Do()
				if err != nil {
					errChan <- err
					doneChan <- true
					return
				}
			}
		}
	}()

	return <-errChan
}
