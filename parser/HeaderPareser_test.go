package parser

import (
	"testing"
)

func testHeaderParser(t *testing.T, hp HeaderParser) {
	if sh, err := hp.Parse(); err != nil {
		t.Log(err)
		t.Fail()
	} else {
		t.Log("encoded = " + sh.String())
	}
}