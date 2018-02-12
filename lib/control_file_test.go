package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadFile(t *testing.T) {
	assert.Equal(t, "aaaa", ReadFile(GetHomePath()+"/.ma-mi/config"))
}
