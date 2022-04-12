package main

import (
	"testing"
)

func TestImportProblems(t *testing.T) {
	want := []Problem{Problem{question: "6+2", answer: "8"}, Problem{question: "7+1", answer: "8"}}
	got := ImportProblems("test_problems.csv")
	if want[0] != got[0] {
		t.Fatalf(`ImportProblems("test_problems.csv")`)
	}
}
