package sort

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	got := QuickSort([]int{4, 2, 3, 1, 5, 7, 6})
	want := []int{1, 2, 3, 4, 5, 6, 7}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}
