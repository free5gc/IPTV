/*
 * IPTV Configuration Factory
 */

package factory

type AFConfig struct {
	AFName string `yaml:"AFName"`

	Sbi Sbi `yaml:"Sbi"`

	NrfUri string `yaml:"NrfUri"`

	LADN string `yaml:"LADN"`

	Subscriber []string `yaml:"Subscriber"`
}
