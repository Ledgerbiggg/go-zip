package test

import (
	"fmt"
	"goZipper/util"
	"testing"
	"time"
)

func TestZip(t *testing.T) {
	err := util.Zip("log.zip", "src")
	if err != nil {
		fmt.Println("zip error:", err)
		return
	}
	time.Sleep(3 * time.Second)
}

func TestUnzip(t *testing.T) {
	err := util.Unzip("log.zip", "dest")
	//err := hh("log.zip", "dest")
	if err != nil {
		fmt.Println("unzip error:", err)
		return
	}
	time.Sleep(3 * time.Second)
}
