package util

import (
	"errors"
	"os"
	"os/exec"
	"testing"
)

func TestExitIfErrWithoutError(t *testing.T) {
	ExitIfErr(nil)
	// if there is an error it stop the test.
}

func TestExitIfErrWithError(t *testing.T) {
	if os.Getenv("BE_CRASHER") == "1" {
		ExitIfErr(errors.New("err"))
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestExitIfErr")
	cmd.Env = append(os.Environ(), "BE_CRASHER=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("process ran with err %v, want exit status 1", err)
}