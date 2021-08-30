package main

import (
	"testing"
)

func TestisPalindrome(t *testing.T) {
	out := isPalindrome(121)
	want := true
	if out != want {
		t.Errorf("got %t, want %t", out, want)
	}
}
