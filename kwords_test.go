package kwords

import (
	"reflect"
	"strings"
	"testing"
)

func TestKwords(t *testing.T) {
	s := "hello how are you today hello hello hello how how how how are are you"
	words := strings.Split(s, " ")

	r := KWords(3, words)
	e := []OccurringWord{
		{"how", 5},
		{"hello", 4},
		{"are", 3},
	}
	if !reflect.DeepEqual(r, e) {
		t.Errorf("Kwords(3, %s)=%v expect %v", s, r, e)
	}

	r = KWords(10, words)
	e = []OccurringWord{
		{"how", 5},
		{"hello", 4},
		{"are", 3},
		{"you", 2},
		{"today", 1},
	}
	if !reflect.DeepEqual(r, e) {
		t.Errorf("Kwords(10, %s)=%v expect %v", s, r, e)
	}
}
