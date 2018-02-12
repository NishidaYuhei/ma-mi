package lib

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDateStr(t *testing.T) {
	assert.Equal(t, "201802130003", GetDateStr())
}
