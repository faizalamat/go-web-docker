package main

import "testing"

func TestSum(t *testing.T) {
	if add(2, 5) != 7 {
		t.Fail()
	}

	if add(2, 100) != 102 {
		t.Fail()
	}

	if add(222, 100) != 322 {
		t.Fail()
	}

}

func TestProduct(t *testing.T) {
	if multiplay(2, 5) != 10 {
		t.Fail()
	}
	if multiplay(2, 100) != 200 {
		t.Fail()
	}
	if multiplay(222, 3) != 666 {
		t.Fail()
	}
}
