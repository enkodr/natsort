package natsort

import (
	"testing"
)

// TestExtractNumber tests the extractNumber function
func TestExtractNumber(t *testing.T) {
	testCases := []struct {
		s           string
		start       int
		expectedNum int
		expectedLen int
	}{
		// Basic numbers
		{"123abc", 0, 123, 3},
		{"abc123", 3, 123, 3},
		{"file10.txt", 4, 10, 2},
		{"file01.txt", 4, 1, 2},

		// Leading zeros
		{"file001.txt", 4, 1, 3},
		{"file0001.txt", 4, 1, 4},

		// No numbers
		{"file.txt", 4, 0, 0},
		{"abc", 0, 0, 0},

		// Numbers at the end
		{"report2023", 6, 2023, 4},

		// Empty string
		{"", 0, 0, 0},

		// Start index out of bounds
		{"file1", 10, 0, 0},

		// Non-digit characters after numbers
		{"file12a34", 4, 12, 2},
		{"file12a34", 7, 34, 2},
	}

	for _, tc := range testCases {
		num, length := extractNumber(tc.s, tc.start)
		if num != tc.expectedNum || length != tc.expectedLen {
			t.Errorf("extractNumber(%q, %d) = (%d, %d); want (%d, %d)",
				tc.s, tc.start, num, length, tc.expectedNum, tc.expectedLen)
		}
	}
}

// TestCompare tests the Compare function
func TestCompare(t *testing.T) {
	testCases := []struct {
		a        string
		b        string
		expected bool
	}{
		// Basic comparisons
		{"file1", "file2", true},
		{"file2", "file10", true},
		{"file10", "file1", false},

		// Leading zeros
		{"file01", "file10", true},
		{"file001", "file01", false},
		{"file0001", "file001", false},

		// Identical strings
		{"file1", "file1", false},
		{"file10", "file10", false},

		// Different prefixes
		{"alpha1", "beta1", true},
		{"beta1", "alpha1", false},

		// Mixed numeric and non-numeric
		{"file1a", "file1b", true},
		{"file1b", "file1a", false},
		{"file1a", "file1a", false},

		// Numbers vs. letters
		{"file10a", "file10", false},
		{"file10", "file10a", true},

		// Multiple numeric sections
		{"file1version2", "file1version10", true},
		{"file1version10", "file1version2", false},
		{"file2version2", "file10version1", true},

		// Empty strings
		{"", "file1", true},
		{"file1", "", false},
		{"", "", false},

		// Non-digit characters after numbers
		{"file12a34", "file12a35", true},
		{"file12a34", "file12a34", false},
		{"file12a35", "file12a34", false},
	}

	for _, tc := range testCases {
		result := Compare(tc.a, tc.b)
		if result != tc.expected {
			t.Errorf("Compare(%q, %q) = %v; want %v",
				tc.a, tc.b, result, tc.expected)
		}
	}
}

func BenchmarkCompare(b *testing.B) {
	a := "file123"
	c := "file124"
	for i := 0; i < b.N; i++ {
		Compare(a, c)
	}
}
