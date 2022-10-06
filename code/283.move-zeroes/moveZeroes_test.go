package moveZeroes

import (
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	got := []int{-1, 0, 0, -4, 5, 6, 0}
	want := []int{-1, -4, 5, 6, 0, 0, 0}
	moveZeroes(got)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("[MoveZeroes] Failed got %v, want %v", got, want)
	}
}
