package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_StringToPointer(t *testing.T) {
	s := "hello"
	ptr := StringToPointer(s)
	assert.NotNil(t, ptr)
	assert.Equal(t, s, *ptr)

	empty := ""
	ptr2 := StringToPointer(empty)
	assert.NotNil(t, ptr2)
	assert.Equal(t, "", *ptr2)
}

func Test_BooleanToPointer(t *testing.T) {
	b := true
	ptr := BooleanToPointer(b)
	assert.NotNil(t, ptr)
	assert.Equal(t, true, *ptr)

	b2 := false
	ptr2 := BooleanToPointer(b2)
	assert.NotNil(t, ptr2)
	assert.Equal(t, false, *ptr2)
}
