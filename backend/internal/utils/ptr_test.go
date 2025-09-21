package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPtr(t *testing.T) {
	val := 42
	ptr := Ptr(val)

	require.NotNil(t, ptr, "expected non-nil pointer")
	assert.Equal(t, val, *ptr, "pointer should contain the original value")
}

func TestNilIfEmpty(t *testing.T) {
	t.Run("empty string", func(t *testing.T) {
		emptyStr := ""
		result := NilIfEmpty(emptyStr)
		assert.Nil(t, result, "expected nil for empty string")
	})

	t.Run("non-empty string", func(t *testing.T) {
		nonEmpty := "test"
		result := NilIfEmpty(nonEmpty)
		require.NotNil(t, result, "expected non-nil for non-empty string")
		assert.Equal(t, nonEmpty, *result, "pointer should contain the original string")
	})
}

func TestDeref(t *testing.T) {
	t.Run("non-nil pointer", func(t *testing.T) {
		val := 42
		ptr := &val
		result := Deref(ptr)
		assert.Equal(t, val, result, "should return the dereferenced value")
	})

	t.Run("nil pointer", func(t *testing.T) {
		var nilPtr *int
		result := Deref(nilPtr)
		assert.Equal(t, 0, result, "should return zero value for nil pointer")
	})
}
