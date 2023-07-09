package utl

import (
	"flag"
	"log"
	"os"
)

var (
	Log *log.Logger
)

func init() {
	var logpath = "./info.log"

	flag.Parse()
	var file, err1 = os.Create(logpath)

	if err1 != nil {
		panic(err1)
	}
	Log = log.New(file, "", log.LstdFlags|log.Lshortfile)
}
