package main

import (
	//"fmt"
	"github.com/chet1211221/golanganidb/setupenv"
)

func main() {
	setupenv.CreateDir(setupenv.GetHomeDir())
	setupenv.CreateDir("/")
}
