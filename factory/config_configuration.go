/*
 * IPTV Configuration Factory
 */

package factory

// Configuration : information of IPTV Application and AF
type Configuration struct {
	IPTVServer IptvServer `yaml:"IPTVServer"`

	AFConfig AFConfig `yaml:"AFConfig"`
}
