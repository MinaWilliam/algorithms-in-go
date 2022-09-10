package divideconquer

import "testing"

func TestCount(t *testing.T) {
	got := Count([]int{1, 2, 3})
	want := 3

	if got != want {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}
