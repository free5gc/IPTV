/*
 * IPTV Configuration Factory
 */

package factory

type IptvServer struct {
	ServerAddr Ipv4Port `yaml:"ServerAddr"`

	Channel []IptvChannel `yaml:"Channels"`

	CacheFolder string `yaml:"CacheFolder"`

	WebClientFolder string `yaml:"WebClientFolder"`
}
