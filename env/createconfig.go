package env

import (
	"fmt"
	"os"
	"strconv"
	"log"
)

func CreateConfig(configfile string, initialconfig *Config) {

	configfileexists, err := exists(configfile)

	if err != nil {
		log.Fatal(err)
		
	} else {
		configfilecreated, err := os.Create(configfile)
		if err != nil {
			fmt.Println(err)
			panic(fmt.Sprintf("%v\n", "Unable to create configuration file for program.  See above line"))
		}
		defer configfilecreated.Close()
		configfilecreated.WriteString("client=" + initialconfig.Client + "\r\n")
		configfilecreated.WriteString("clientver=" + strconv.Itoa(initialconfig.Clientver) + "\r\n")
		configfilecreated.WriteString("protover=" + strconv.Itoa(initialconfig.Protover) + "\r\n")
		configfilecreated.WriteString("url=" + initialconfig.Url + "\r\n")
		configfilecreated.WriteString("port=" + strconv.Itoa(initialconfig.Port) + "\r\n")
	}

}
