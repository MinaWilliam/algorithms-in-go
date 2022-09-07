package sort

import (
	"reflect"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	tests := [...]struct {
		given []int
		want  []int
	}{
		{given: []int{4, 2, 3, 1}, want: []int{1, 2, 3, 4}},
	}

	for _, tc := range tests {
		got := SelectionSort(tc.given)
		if !reflect.DeepEqual(got, tc.want) {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func TestFindSmallest(t *testing.T) {

	got := findSmallest([]int{4, 2, 3, 1})
	want := 3
	if got != want {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}
