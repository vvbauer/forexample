package hello

import (
	"testing"
)

func TestHello(t *testing.T) {
	want := "Hello Leroy!"
	if got := Hello(); got != want {
		t.Errorf("hello() = %q, want = %q", got, want)
	}
}
