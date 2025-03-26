package log

import "log"

var (
	info *log.Logger
)

func init() {
	log.Printf("init call directly")
}

func Info(message string) {
}

func Debug() {}
