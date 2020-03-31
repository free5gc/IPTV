/*
 * IPTV Configuration Factory
 */

package factory

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

var IptvConfig Config

func checkErr(err error) {
	if err != nil {
		err = fmt.Errorf("[Configuration] %s", err.Error())
	}
}

// TODO: Support configuration update from REST api
func InitConfigFactory(f string) {
	content, err := ioutil.ReadFile(f)
	checkErr(err)

	IptvConfig = Config{}

	err = yaml.Unmarshal([]byte(content), &IptvConfig)
	checkErr(err)

	fmt.Println("Successfully initialize configuration %s", f)
}


