package tests

import (
	"testing"

	"github.com/moatasemgamal/gofreerdp"
)

func TestInit(t *testing.T) {
	if freeRDP, err := gofreerdp.Init(); freeRDP == nil || err != nil {
		t.Error(err)
	}
}
