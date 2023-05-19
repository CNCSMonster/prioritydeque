package main

import (
	"fmt"
	"testing"

	"github.com/cncsmonster/prioritydeque"
)

func TestReplace(t *testing.T) {
	pdq := prioritydeque.FromSlice(func(a, b any) bool {
		return a.(int) < b.(int)
	}, 1, 23, 55, 21)
	// now max and min should be 1 and 55
	want := "1,55"
	get := fmt.Sprintf("%d,%d", pdq.Min(), pdq.Max())
	if get != want {
		t.Error("want", want, ",get", get)
	}
	pdq.ReplaceMax(22)
	want = "1,23"
	get = fmt.Sprintf("%d,%d", pdq.Min(), pdq.Max())
	if get != want {
		t.Error("want", want, ",get", get)
	}
	pdq.ReplaceMax(92)
	want = "1,92"
	get = fmt.Sprintf("%d,%d", pdq.Min(), pdq.Max())
	if get != want {
		t.Error("want", want, ",get", get)
	}
	pdq.ReplaceMin(11)
	want = "11,92"
	get = fmt.Sprintf("%d,%d", pdq.Min(), pdq.Max())
	if get != want {
		t.Error("want", want, ",get", get)
	}
	pdq.ReplaceMin(-23)
	want = "-23,92"
	get = fmt.Sprintf("%d,%d", pdq.Min(), pdq.Max())
	if get != want {
		t.Error("want", want, ",get", get)
	}
}
