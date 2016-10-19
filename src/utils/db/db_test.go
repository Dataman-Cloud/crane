package db

import (
	"os"
	"os/exec"
	"testing"
)

func TestNewDB(t *testing.T) {
	if os.Getenv("BE_NEWDB") == "1" {
		NewDB("foobar", "ff")
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestNewDB")
	cmd.Env = append(os.Environ(), "BE_NEWDB=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("process ran with err %v, want exit status 1", err)
}
