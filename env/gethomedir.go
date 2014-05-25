package env

import (
	"fmt"
	"os/user"
)

func GetHomeDir() string {
	usr, err := user.Current()
	if err != nil {
		fmt.Println(err)
	}
	homedir := usr.HomeDir
	return homedir
}
