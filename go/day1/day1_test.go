package main

import (
	"reflect"
	"testing"
)

func TestHelloName(t *testing.T) {
	line := "pppmfmnfourtworxrqrfhbgx8vvxgrjzhvqmztltwo"
	matches := parseMatches(line)
	want := []Match{{7, "4"}, {11, "2"}, {24, "8"}, {39, "2"}}
	if !reflect.DeepEqual(matches, want) {
		t.Fatalf(`parseMatches("pppmfmnfourtworxrqrfhbgx8vvxgrjzhvqmztltwo") = %+v, want %+v`, matches, want)
	}
}
