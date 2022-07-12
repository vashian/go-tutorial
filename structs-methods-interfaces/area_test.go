package main

import "testing"


func TestArea(t *testing.T) {

	rectangle := Rectangle{12.0, 6.0}

	got := Area(rectangle)
	want := 72.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}