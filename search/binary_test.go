package search

import (
	"testing"
)

func TestBinaryExist(t *testing.T) {
	tests := [...]struct {
		input  []int
		search int
		want   bool
	}{
		{input: []int{1, 2, 3, 4, 5}, search: 3, want: true},
		{input: []int{1, 2, 3, 4, 5}, search: 10, want: false},
	}

	for _, tc := range tests {
		got := BinaryExist(tc.input, tc.search)
		if got != tc.want {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestBinarySearch(t *testing.T) {
	tests := [...]struct {
		input  []int
		search int
		want   int
	}{
		{input: []int{1, 2, 3, 4, 5}, search: 3, want: 2},
		{input: []int{1, 2, 3, 4, 5}, search: 10, want: -1},
	}

	for _, tc := range tests {
		got := BinarySearch(tc.input, tc.search)
		if got != tc.want {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}
