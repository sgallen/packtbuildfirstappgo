package main

import "testing"

func TestCircumference(t *testing.T) {
	t.Log("Testing circumference func with radius of 2.")
	if length := circumference(2); length != 12.56 {
		t.Errorf("Expected a circumference of 12.56, instead got: %.2f", length)
	}
}
