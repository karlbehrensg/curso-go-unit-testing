package main

import "testing"

func TestAddSucess(t *testing.T) {
	result := Add(20, 2)

	expect := 22

	if result != expect {
		t.Errorf("Expected %d, but got %d", expect, result)
	}
}
