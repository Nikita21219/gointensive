package main

import "testing"

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestMinCoins(t *testing.T) {
	t1 := minCoins(100, []int{1, 50, 90})
	a1 := []int{50, 50}
	if !Equal(t1, a1) {
		t.Errorf("got %v, wanted %v", t1, a1)
	}

	t2 := minCoins(18750, []int{50, 100, 500, 1000, 5000})
	a2 := []int{5000, 5000, 5000, 1000, 1000, 1000, 500, 100, 100, 50}
	if !Equal(t2, a2) {
		t.Errorf("got %v, wanted %v", t2, a2)
	}

	t3 := minCoins(100, []int{})
	var a3 []int
	if !Equal(t3, a3) {
		t.Errorf("got %v, wanted %v", t3, a3)
	}

	t4 := minCoins(100, []int{1, 90})
	a4 := []int{90, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	if !Equal(t4, a4) {
		t.Errorf("got %v, wanted %v", t4, a4)
	}

	t5 := minCoins(5, []int{10, 90})
	a5 := []int{}
	if !Equal(t5, a5) {
		t.Errorf("got %v, wanted %v", t5, a5)
	}

	t6 := minCoins(6000, []int{500, 3000, 5000})
	a6 := []int{3000, 3000}
	if !Equal(t6, a6) {
		t.Errorf("got %v, wanted %v", t6, a6)
	}

	t7 := minCoins(100, []int{90, 10, 50})
	a7 := []int{90, 10}
	if !Equal(t7, a7) {
		t.Errorf("got %v, wanted %v", t7, a7)
	}

	t8 := minCoins(13, []int{1, 5, 10})
	a8 := []int{10, 1, 1, 1}
	if !Equal(t8, a8) {
		t.Errorf("got %v, wanted %v", t8, a8)
	}
}

func TestMinCoins2(t *testing.T) {
	t1 := minCoins2(100, []int{1, 50, 90})
	a1 := []int{50, 50}
	if !Equal(t1, a1) {
		t.Errorf("got %v, wanted %v", t1, a1)
	}

	t2 := minCoins2(18750, []int{50, 100, 500, 1000, 5000})
	a2 := []int{50, 100, 100, 500, 1000, 1000, 1000, 5000, 5000, 5000}
	if !Equal(t2, a2) {
		t.Errorf("got %v, wanted %v", t2, a2)
	}

	t3 := minCoins2(100, []int{})
	var a3 []int
	if !Equal(t3, a3) {
		t.Errorf("got %v, wanted %v", t3, a3)
	}

	t4 := minCoins2(100, []int{1, 90})
	a4 := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 90}
	if !Equal(t4, a4) {
		t.Errorf("got %v, wanted %v", t4, a4)
	}

	t5 := minCoins2(5, []int{10, 90})
	a5 := []int{}
	if !Equal(t5, a5) {
		t.Errorf("got %v, wanted %v", t5, a5)
	}

	t6 := minCoins2(6000, []int{500, 3000, 5000})
	a6 := []int{3000, 3000}
	if !Equal(t6, a6) {
		t.Errorf("got %v, wanted %v", t6, a6)
	}

	t7 := minCoins2(100, []int{90, 10, 50})
	a7 := []int{90, 10}
	if !Equal(t7, a7) {
		t.Errorf("got %v, wanted %v", t7, a7)
	}

	t8 := minCoins2(13, []int{1, 5, 10})
	a8 := []int{1, 1, 1, 10}
	if !Equal(t8, a8) {
		t.Errorf("got %v, wanted %v", t8, a8)
	}
}
