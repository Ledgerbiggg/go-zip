package log

import (
	"log"
)

// InitLogStyle 初始化日志格式
func InitLogStyle() {
	log.SetPrefix("[go-zipper-log] ")
	log.SetFlags(log.LstdFlags)
}

func Println(a ...any) {
	log.Println(a)
}

func Print(a ...any) {
	log.Print(a)
}
