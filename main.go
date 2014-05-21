package main

import (
	"fmt"
	"github.com/chet1211221/golanganidb/setupenv"
)

func main() {
	programdatapath, programconfigpath := setupenv.CreateDir(setupenv.GetHomeDir())
	fmt.Println(programdatapath)
	fmt.Println(programconfigpath)
	initialconfig := setupenv.Config{"golanganidb", 1, 1, "http://api.anidb.net", 9001}
	setupenv.CreateConfig(programconfigpath+"/golanganidb.conf", &initialconfig)
	fmt.Printf("%v\n", initialconfig)
	//runningconfig := new(setupenv.Config)	
	setupenv.ReadConfig(programconfigpath + "/golanganidb.conf")

}
