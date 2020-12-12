package testdata

import (
	"encoding/json"
	"testing"
)

var SerializeStruct = func(t *testing.T, val interface{}) []byte {
	vbytes, err := json.Marshal(val)
	if err != nil {
		t.Fatal(err)
	}
	return vbytes
}
