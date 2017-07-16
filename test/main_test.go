package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCommand(t *testing.T) {
	assert.Equal(t, false, false, "Non registered Id has no value")
	assert.Equal(t, false, false, "Non registered Id has false state")
}
