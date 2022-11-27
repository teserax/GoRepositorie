package math

import "testing"

func TestAdd(t *testing.T) {

	got := Add(4, 6)
	want := 10

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
func TestSubtract(t *testing.T) {
	got := Subtract(5, 5)
	want := 0
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
func TestMulti(t *testing.T) {
	got := Multi(4, 4)
	want := 16
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
