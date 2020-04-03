/*
 * IPTV Configuration Factory
 */

package factory

// Ipv4Port : Ipv4 and Port
type Ipv4Port struct {
	IPv4Addr string `yaml:"IPv4"`

	Port int `yaml:"Port"`
}
