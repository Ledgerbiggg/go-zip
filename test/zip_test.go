package test

import (
	"fmt"
	"goZipper/util"
	"testing"
)

func TestZip(t *testing.T) {
	err := util.Zip("log.zip")
	if err != nil {
		fmt.Println("zip error:", err)
		return
	}
}

func TestUnzip(t *testing.T) {
	err := util.Unzip("log.zip")
	if err != nil {
		fmt.Println("unzip error:", err)
		return
	}
}
