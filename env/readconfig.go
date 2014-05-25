package env

import (
	//"os"
	"fmt"
	"io/ioutil"
	"strings"
)

//Read the preexisting config file into a Config struct to be used later on in the program. Need to work on adding the read config into the struck
func ReadConfig(configfile string) {
	configfilebyte, err := ioutil.ReadFile(configfile)
	if err != nil {
		fmt.Println(err)
	}
	configfilestring := strings.SplitAfter(string(configfilebyte[:]), "\r\n")
	fmt.Println(configfilestring)
	configparams := []string{"client", "clientver", "protover", "url", "port"}
	for a := range configparams {
		b := stringsearch(configfilestring, configparams[a])
		fmt.Println(b)

	}

}

//func (h *Config) configtostruct(configstring []string) *Config {

//}

func stringsearch(searchstrings []string, substring string) string {
	for g := range searchstrings {
		result := strings.Contains(searchstrings[g], substring)
		if result == true {
			configlinesplit := strings.SplitAfter(searchstrings[g], "=")
			//configlinesplit = strings.TrimSpace(configlinesplit)
			//fmt.Println(strings.TrimSpace(configlinesplit[1]))
			return strings.TrimSpace(configlinesplit[1])
		}
	}
	return ""
}
