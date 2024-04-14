//go:build testing
package assert

import "testing"

func AssertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
}


func AssertNil[T comparable](t *testing.T, got T) {
    t.Helper()

    var n T
    if got != n {
        t.Errorf("got %v, wanted nil", got)
    }
}


func AssertNotNil[T comparable](t *testing.T, got T) {
    t.Helper()

    var n T
    if got == n {
        t.Errorf("got %v, wanted not nil", got)
    }
}
