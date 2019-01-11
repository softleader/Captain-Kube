package app

import (
	"testing"
)

func TestActiveContext(t *testing.T) {
	ctxs := []string{"CTX_local=-e localhost", "CTX_93=-e 192.168.1.93 --color"}
	if err := initContext(ctxs); err != nil {
		t.Error(err)
		t.FailNow()
	}
	currentCtx, err := newActiveContext("93")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if currentCtx.Endpoint.Host != "192.168.1.93" {
		t.Error("endpoint host must be 192.168.1.93")
	}
}
