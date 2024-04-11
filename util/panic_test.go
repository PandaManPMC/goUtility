package util

import "testing"

func TestPanic(t *testing.T) {
	defer func() {
		t.Log("defer")
	}()

	t.Log("a")
	panic("b")

}
