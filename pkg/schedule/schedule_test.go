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
	s := testStruct{s: "Hello, World!"}
	err := schedule.Schedule(1, &s)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
}