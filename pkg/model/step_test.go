package model_test

import (
	"github.com/thomaspoignant/api-scenario/pkg/model"
	"github.com/thomaspoignant/api-scenario/test"
	"testing"
	"time"
)

func Test_PauseStep (t *testing.T) {
	step := model.Step{
		StepType: model.Pause,
		Duration: 1,
	}
	start := time.Now()
	step.Run()
	end := time.Now()

	if start.Add(1 * time.Second).After(end) {
		t.Errorf("We should wait for at least 1 seconds, start:%v - end:%v", start, end)
	}
}

func Test_OutputPause(t *testing.T) {
	test.SetupLog()
	step := model.Step{
		StepType: model.Pause,
		Duration: 1,
	}
	want := "------------------------\nWaiting for 1s\n"
	got := test.CaptureOutput(func(){
		step.Run()
	})
	test.Equals(t, "Output messages are different", want, got)
}