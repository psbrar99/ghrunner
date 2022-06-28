package main

import "testing"

func TestSum(t *testing.T) {
	if Sum(2, 3, 4, 6, 9) != 54 {
		t.Error("expected", 24, "got", Sum(2, 3, 4, 6, 9))
	}

}
