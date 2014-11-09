package modlog

import (
	"github.com/bmizerany/assert"
	"os"
	"testing"
)

func TestCanCreateLogger(t *testing.T) {
	_ = New(os.Stdout, "", 0)
}

func TestCanCreateLoggerThatWritesToStderr(t *testing.T) {
	a := New(os.Stdout, "", 0)
	b := New(os.Stderr, "", 0)

	assert.NotEqual(t, a, b)
}
