package test

import (
	"goZipper/util"
	"testing"
)

func TestTarGz(t *testing.T) {
	err := util.TarGz("src", "log")
	if err != nil {
		t.Error(err)
	}
}

func TestUnTarGz(t *testing.T) {
	err := util.UnTarGz("log.tar.gz", "dest")
	if err != nil {
		t.Error(err)
	}
}
