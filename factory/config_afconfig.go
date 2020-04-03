/*
 * IPTV Configuration Factory
 */

package factory

// AFConfig : detail information of AF
type AFConfig struct {
	AFName string `yaml:"AFName"`

	Sbi Sbi `yaml:"Sbi"`

	NrfURI string `yaml:"NrfUri"`

	LADN string `yaml:"LADN"`

	Subscriber []string `yaml:"Subscriber"`
}
