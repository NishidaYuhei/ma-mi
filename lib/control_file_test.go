package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadFile(t *testing.T) {
	assert.Equal(t, "/Users/bellocha/Documents/test", ReadFile(GetHomePath()+"/.ma-mi/config"))
}

func TestExists(t *testing.T) {
	assert.Equal(t, true, Exists(GetHomePath()+"/Desktop"))
}
