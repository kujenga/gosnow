package gosnow

import (
	"testing"
)

const simpleSrc = "# My API\n## GET /message\n + Response 200 (text/plain)\n\n        Hello World\n"

func TestBasicParse(t *testing.T) {
	res, err := Parse(simpleSrc)
	if err != nil {
		t.Fatalf("Parse failed with error: %v", err)
	}
	if res == "" {
		t.Fatal("Parse returned empty result")
	}
}

func TestParseEquality(t *testing.T) {
	res1, err := Parse(simpleSrc)
	if err != nil {
		t.Fatalf("Parse failed with err: %v", err)
	}
	res2, err := OptionParse(simpleSrc, 0)
	if err != nil {
		t.Fatalf("OptionParse failed with err: %v", err)
	}

	if res1 != res2 {
		t.Error("Results should be equal")
	}
}
