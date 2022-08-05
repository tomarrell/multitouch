package multitouch

import (
	"testing"
)

func TestMultitouch(t *testing.T) {
	mt, err := NewMultitouch("./testdata/evdump")
	if err != nil {
		t.Fatal(err)
	}

	go mt.Begin()

	e := mt.Next()
	if e.Action != ActionBegin {
		t.Fatal("action not begin")
	}

	e = mt.Next()
	if e.Action != ActionEnd {
		t.Fatal("action not end")
	}
}
