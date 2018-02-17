package lib

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadFile(t *testing.T) {
	assert.Equal(t, "/Users/bellocha/Documents/test", ReadFile(GetHomePath()+"/.ma-mi/config"))
}

func TestExists(t *testing.T) {
	assert.Equal(t, true, Exists(GetHomePath()+"/Desktop"))
}

func TestGetHomePath(t *testing.T) {
	assert.Equal(t, "/Users/bellocha", GetHomePath())
}

func TestReplacePath(t *testing.T) {
	assert.Equal(t, "/Users/bellocha/Deskto", strings.Replace("~/Desktop", "~", GetHomePath(), 1))
}
