package main

import (
	"reflect"
	"testing"
)


func TestSumAllTails(t *testing.T) {

	checkSum := func(t testing.TB, got, want []int) {
		
		t.Helper()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}


	t.Run("make the sum of some slices", func(t *testing.T) {

		got := SumAllTails([]int{1,2}, []int{0,9})
		want := []int{2,9}

		checkSum(t, got, want)
	})

	t.Run("Safely sum empty slices", func(t *testing.T) {

		got := SumAllTails([]int{}, []int{3,4,5})
		want := []int{0, 9}

		checkSum(t, got, want)
	})
}

// func TestSumAllTails(t *testing.T) {

// 	got := SumAllTails([]int{1, 2}, []int{0, 9})
// 	want := []int{2, 9}

// 	if !reflect.DeepEqual(got, want) {
// 		t.Errorf("got %v want %v", got, want)
// 	}
// }