package handler

import (
	"testing"
)

func TestParseYaml(t *testing.T) {
	data := `
    - path: /google
      url: https://google.com
  `

	expect := []PathUrl{PathUrl{Path: "/google", URL: "https://google.com"}}

	got, err := ParseYaml([]byte(data))

	if err != nil {
		t.Fatal(err)
	}

	if got[0] != expect[0] {
		t.Fatalf("Expected %v but got %v", expect[0], got[0])
	}
}

func TestBuildMap(t *testing.T) {
	data := []PathUrl{PathUrl{Path: "/google", URL: "https://google.com"}}

	expect := map[string]string{
		"/google": "https://google.com",
	}

	got := BuildMap(data)

	if got["/google"] != expect["/google"] {
		t.Fatalf("Expected %v but got %v", expect["/google"], got["/google"])
	}
}

func TestParseJson(t *testing.T) {
	data := `[{"path":"/google","url":"https://google.com"}]`

	expect := []PathUrl{PathUrl{Path: "/google", URL: "https://google.com"}}

	got, err := ParseJson([]byte(data))

	if err != nil {
		t.Fatal(err)
	}

	if got[0] != expect[0] {
		t.Fatalf("Expected %v but got %v", expect[0], got[0])
	}
}
