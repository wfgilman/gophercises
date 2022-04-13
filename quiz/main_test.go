package main

import (
	"testing"
)

func TestImportProblems(t *testing.T) {
	want := [][]string{
		{"6+2", "8"},
		{"7+1", "8"},
	}
	got := ImportProblems("test_problems.csv")
	if want[0][0] != got[0][0] {
		t.Fatalf("Wanted %s and got %s", want[0][0], got[0][0])
	}
}

func TestParseRecords(t *testing.T) {
	want := []Problem{Problem{question: "6+2", answer: "8"}, Problem{question: "7+1", answer: "8"}}
	testRecords := [][]string{
		{"6+2", "8"},
		{"7+1", "   8"}, // Test elimination of whitespace
	}
	got := ParseRecords(testRecords)
	if want[0] != got[0] {
		t.Fatalf("Wanted %s and got %s", want[0], got[0])
	}
	if want[1] != got[1] {
		t.Fatalf("Wanted %s and got %s", want[1], got[1])
	}
}
