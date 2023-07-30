package str2

import (
	//"mmgreiner/utils/str2"

	"testing"

	"gotest.tools/assert"
)

func TestIsInt(t *testing.T) {
	assert.Assert(t, IsInt("1"))
	assert.Assert(t, !IsInt("3.1415"))
	assert.Assert(t, !IsInt("hallo"))
}

func TestIsFloat(t *testing.T) {
	assert.Assert(t, IsFloat("1.0"))
	assert.Assert(t, IsFloat("1"))
	assert.Assert(t, !IsFloat("hallo"))
}

func TestToFloat(t *testing.T) {
	Warn = true
	assert.Equal(t, ToFloat("1.0"), 1.0)
	assert.Equal(t, ToFloat("1"), 1.0)
	assert.Equal(t, ToFloat("hallo"), 0.0)
}
