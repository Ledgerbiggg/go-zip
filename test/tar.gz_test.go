package test

import (
	"goZipper/util"
	"testing"
)

func TestTarGz(t *testing.T) {
	util.TarGz()
}

func TestUnTarGz(t *testing.T) {
	err := util.UnTarGz()
	if err != nil {
		t.Error(err)
	}
}
