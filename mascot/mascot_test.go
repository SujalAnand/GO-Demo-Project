package mascot_test

import (
	"testing"

	"example.com/GO/mascot"
)

func TestMascot(t *testing.T) {
	if mascot.BestMascot() != "Go Gooper" {
		t.Fatal("Wrong Mascot")
	}
}
