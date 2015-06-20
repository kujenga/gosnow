package gosnow

import (
	"io/ioutil"
	"testing"
)

const simpleSrc string = `
# My API
## GET /message
 + Response 200 (text/plain)

         Hello World
`

const namelessSrc string = `
FORMAT: 1A

# Group Messages

# Message [/messages/{id}]

## Retrieve Message [GET]
+ Response 200 (text/plain)

        Hello World!

`

var (
	apibFile      = "test/fixtures/sample-api.apib"
	astFile       = "test/fixtures/sample-api-ast.json"
	sourcemapFile = "test/fixtures/sample-api-sourcemap.json"
)

// replace the variables with the contents of the file they point to
func init() {
	if c, err := ioutil.ReadFile(apibFile); err != nil {
		panic("apibFile not found")
	} else {
		apibFile = string(c)
	}

	if c, err := ioutil.ReadFile(astFile); err != nil {
		panic("astFile not found")
	} else {
		astFile = string(c)
	}

	if c, err := ioutil.ReadFile(sourcemapFile); err != nil {
		panic("sourcemapFile not found")
	} else {
		sourcemapFile = string(c)
	}
}

func TestParse(t *testing.T) {
	res, err := Parse(simpleSrc)
	if err != nil {
		t.Fatalf("Parse failed with error: %v", err)
	}
	if res == "" {
		t.Fatal("Parse returned empty result")
	}
}

// ensure that the option parse with a 0 does the same thing as the simple parse
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

func TestParseError(t *testing.T) {
	junk := "*#(*(R$#&)$#)R*(Y@#_RH"
	res, err := OptionParse(junk, -1)
	if err == nil {
		t.Errorf("OptionParse did not fail for junk input")
	}
	if res != "" {
		t.Errorf("OptionParse returned non=empty result for junk input")
	}
}

func TestFilesOptionParse(t *testing.T) {
	res, err := OptionParse(apibFile, ScRenderDescriptionsOptionKey)
	if err != nil {
		t.Errorf("OptionParse failed for key ScRenderDescriptionsOptionKey with error: %v", err)
	} else if res == "" {
		t.Errorf("OptionParse for key ScRenderDescriptionsOptionKey returned empty result")
	}

	_, err = OptionParse(namelessSrc, RequireBlueprintNameOptionKey)
	if err == nil {
		t.Errorf("strict OptionParse did not fail for key RequireBlueprintNameOptionKey")
	}

	res, err = OptionParse(apibFile, ExportSourcemapOptionKey)
	if err != nil {
		t.Errorf("OptionParse failed for ExportSourcemapOptionKey with error: %v", err)
	} else if res == "" {
		t.Errorf("OptionParse for key ExportSourcemapOptionKey returned empty result")
	}
}
