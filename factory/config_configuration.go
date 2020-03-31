/*
 * IPTV Configuration Factory
 */

package factory

type Configuration struct {
	IPTVServer IptvServer `yaml:"IPTVServer"`

	AFConfig AFConfig `yaml:"AFConfig"`
}
