package utl

import (
	"log"
	"os"
)

var (
	Log *log.Logger
)

func init() {
	Log = log.New(os.Stdout, "", log.LstdFlags)
}
