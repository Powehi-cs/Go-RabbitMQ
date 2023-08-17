package utils

import (
	"log"
	"os"
	"strings"
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%v: %v", msg, err)
	}
}

func BodyFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[1:], " ")
	}
	return s
}
