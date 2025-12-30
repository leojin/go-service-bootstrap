package utils

import (
	"log"
	"os"
)

var Out *log.Logger
var Err *log.Logger

func init() {
	Out = log.New(os.Stdout, "", log.LstdFlags)
	Err = log.New(os.Stderr, "", log.LstdFlags)
}
