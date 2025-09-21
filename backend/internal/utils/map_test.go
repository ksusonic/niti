package utils

import (
	"strconv"
	"testing"
)

func TestMap(t *testing.T) {
	// Test with int to string conversion
	nums := []int{1, 2, 3}
	result := Map(nums, strconv.Itoa)
	expected := []string{"1", "2", "3"}

	if len(result) != len(expected) {
		t.Fatalf("expected length %d, got %d", len(expected), len(result))
	}

	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("at index %d: expected %s, got %s", i, expected[i], result[i])
		}
	}
}

func TestMapEmpty(t *testing.T) {
	var nums []int
	result := Map(nums, strconv.Itoa)

	if len(result) != 0 {
		t.Errorf("expected empty slice, got length %d", len(result))
	}
}
