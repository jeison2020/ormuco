package test

import (
	"fmt"
	"ormuco.go/internal/handler"
	"testing"
)

func TestCompareVersions(t *testing.T) {
	testCases := []struct {
		v1       string
		v2       string
		expected string
	}{
		{"1.2.3", "1.2.4", "The version 1.2.3 is lower than version. 1.2.4"},
		{"2.0", "1.5.1", "The version 2.0 is greater than version. 1.5.1"},
		{"3", "3.0.0", "they are equal"},
		{"1.2.a", "1.2.3", "The version 1 is not a number"},
		{"1.2", "1.2.b", "The version 2 is not a number"},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s vs %s", tc.v1, tc.v2), func(t *testing.T) {
			result := handler.CompareVersions(tc.v1, tc.v2)
			if result != tc.expected {
				t.Errorf("Expected '%s', but got '%s'", tc.expected, result)
			}
		})
	}
}
