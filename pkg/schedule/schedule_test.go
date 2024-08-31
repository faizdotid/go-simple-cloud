package schedule_test

import (
	"fmt"
	"go-simple-cloud/pkg/schedule"
	"testing"
)

type testStruct struct {
	s string
	l int
}

func (t *testStruct) Do() error {
	if t.l == 10 {
		return fmt.Errorf("Error: %v", t.l)
	}
	fmt.Println(t.s)
	t.l++
	return nil
}

func TestSchedule(t *testing.T) {
	errorChan := make(chan error)
	defer close(errorChan)

	s := &testStruct{
		s: "Hello, World!",
		l: 0,
	}

	if err := schedule.Schedule(1, s, errorChan); err != nil {
		t.Fatalf("Failed to schedule task: %v", err)
	}
	go func() {
		t.Logf("Task scheduled")
	}()

}
