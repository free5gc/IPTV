/*
 * IPTV Configuration Factory
 */

package factory

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// IptvConfig : Save all information of IPTV
var IptvConfig Config

func checkErr(err error) {
	if err != nil {
		fmt.Printf("[Configuration] %s", err.Error())
	}
}

// InitConfigFactory : read configuration file and parse into structure
// TODO: Support configuration update from REST api
func InitConfigFactory(f string) {
	content, err := ioutil.ReadFile(f)
	checkErr(err)

	IptvConfig = Config{}

	err = yaml.Unmarshal([]byte(content), &IptvConfig)
	checkErr(err)

	fmt.Printf("Successfully initialize configuration %s\n", f)
}
