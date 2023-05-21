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

func TestReplaceMatched(t *testing.T) {
	pq := prioritydeque.FromSlice(
		func(a, b any) bool {
			return a.(int) < b.(int)
		}, 12, 3, 3, 55, 55,
	)
	if fmt.Sprintf("%d,%d", pq.Min(), pq.Max()) != "3,55" {
		t.Error(fmt.Sprintf("%d,%d", pq.Min(), pq.Max()), "3,55")
	}
	match := func(v int) func(any) bool {
		return func(x any) bool {
			return x.(int) == v
		}
	}
	pq.Replace(match(3), 4)
	if fmt.Sprintf("%d,%d", pq.Min(), pq.Max()) != "3,55" {
		t.Error(fmt.Sprintf("%d,%d", pq.Min(), pq.Max()), "3,55")
	}
	pq.Replace(match(55), 32)
	if fmt.Sprintf("%d,%d", pq.Min(), pq.Max()) != "3,55" {
		t.Error(fmt.Sprintf("%d,%d", pq.Min(), pq.Max()), "3,55")
	}
}

func TestReplaceAllMatched(t *testing.T) {
	pq := prioritydeque.FromSlice(
		func(a, b any) bool {
			return a.(int) < b.(int)
		}, 12, 3, 3, 55, 55,
	)
	if fmt.Sprintf("%d,%d", pq.Min(), pq.Max()) != "3,55" {
		t.Error(fmt.Sprintf("%d,%d", pq.Min(), pq.Max()), "3,55")
	}
	match := func(v int) func(any) bool {
		return func(x any) bool {
			return x.(int) == v
		}
	}
	pq.ReplaceAll(match(3), 4)
	if fmt.Sprintf("%d,%d", pq.Min(), pq.Max()) != "4,55" {
		t.Error(fmt.Sprintf("%d,%d", pq.Min(), pq.Max()), "4,55")
	}
	pq.ReplaceAll(match(55), 32)
	if fmt.Sprintf("%d,%d", pq.Min(), pq.Max()) != "4,32" {
		t.Error(fmt.Sprintf("%d,%d", pq.Min(), pq.Max()), "4,32")
	}
}
