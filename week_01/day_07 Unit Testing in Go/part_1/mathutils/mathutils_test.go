// Part 1: Unit Testing in Go
// mathutils_test.go

// ✅ รัน test
// go test or go test <moudule>_test.go
package mathutils

import "testing"

// ✅ เขียน test:
func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5

	if result != expected {
		t.Errorf("Add(2, 3) = %d; want %d", result, expected)
	}
}

// ✅ Table-Driven Test (เขียน test หลาย ๆ กรณี)
func TestAddCases(t *testing.T) {
	cases := []struct {
		a, b, expected int
	}{
		{2, 3, 5},
		{0, 0, 0},
		{-1, -1, -2},
	}

	for _, c := range cases {
		got := Add(c.a, c.b)
		if got != c.expected {
			t.Errorf("Add(%d, %d) = %d; want %d", c.a, c.b, got, c.expected)
		}
	}
}
