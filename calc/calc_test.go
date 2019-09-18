package calc

import "testing"

func TestSum(t *testing.T) {
	get := sum(1, 2)
	want := 3
	if get != want {
		t.Fatalf("sum(1, 2)should be 3, but doesn't match")
	}

	get = sum(7, 2)
	want = 9
	if get != want {
		t.Fatalf("sum(7, 2) should be 9, but doesn't match")
	}
}
