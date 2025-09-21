package utils

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMap(t *testing.T) {
	t.Run("int to string conversion", func(t *testing.T) {
		nums := []int{1, 2, 3}
		result := Map(nums, strconv.Itoa)
		expected := []string{"1", "2", "3"}

		assert.Equal(t, expected, result, "should convert all integers to strings")
	})
}

func TestMapEmpty(t *testing.T) {
	var nums []int
	result := Map(nums, strconv.Itoa)

	assert.Empty(t, result, "should return empty slice for empty input")
}
