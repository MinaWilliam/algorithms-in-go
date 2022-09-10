package divideconquer

import "testing"

func TestSum(t *testing.T) {
	got := Sum([]int{1, 2, 3})
	want := 6

	if got != want {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}
