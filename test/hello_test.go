package test

import (
	"testing"

	"github.com/diegokrule1/test/app/service1"
)

func TestHello(t *testing.T) {
	want := "Hello, world from go."
	if got := service1.Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
