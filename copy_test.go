// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package spirv

import (
	"reflect"
	"testing"
)

func TestCopy(t *testing.T) {
	testCopy(t, nil)
	testCopy(t, []uint32{1, 2, 3})
}

func testCopy(t *testing.T, want []uint32) {
	have := Copy(want)

	if !reflect.DeepEqual(have, want) {
		t.Fatalf("copy mismatch:\nHave: %v\nWant: %v", have, want)
	}
}
