package mascot_test

import (
	"testing"

	"example.com/GO/mascot"
)

func TestMascot(t *testing.T) {
	if mascot.BestMascot() != "TUX" {
		t.Fatal("Wrong Mascot")
	}

}
