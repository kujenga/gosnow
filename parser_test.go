package gosnow

import (
	"testing"
)

func TestBasicParse(t *testing.T) {
	src := "# My API\n## GET /message\n + Response 200 (text/plain)\n\n        Hello World\n"
	res, err := Parse(src)
	if err != nil {
		t.Fatalf("Parse failed with error: %v", err)
	}
	if res == "" {
		t.Fatal("Parse returned empty result")
	}
}
