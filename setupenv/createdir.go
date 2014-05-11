package setupenv

import (
	"fmt"
	"os"
)

func CreateDir(pathroot string) {
	programpath := pathroot + "/golanganidb"
	configpath := pathroot + "/.golanganidb"
	cachepath := programpath + "/cache"
	makedir(programpath)
	makedir(configpath)
	makedir(cachepath)
}

// exists returns whether the given file or directory exists or not
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) != true {
		return false, nil
	}
	return false, err
}

func makedir(path string) {
	pathexists, _ := exists(path)
	if pathexists == false {
		patherr := os.Mkdir(path, 0770)
		if patherr == nil {
			fmt.Println(path, " created successfully")
		} else {
			fmt.Println(patherr)
			panic(fmt.Sprintf("%v\n", "Unable to create directory for program.  See above line"))
		}
	} else {
		fmt.Println(path, " already exists")
	}
}