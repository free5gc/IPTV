/*
 * IPTV Configuration Factory
 */

package factory

type IptvChannel struct {
	ChannelName string `yaml:"ChannelName"`

	VideoPath string `yaml:"VideoPath"`
}
