/*
 * IPTV Configuration Factory
 */

package factory

// IptvChannel : detail of each channel
type IptvChannel struct {
	ChannelName string `yaml:"ChannelName"`

	VideoPath string `yaml:"VideoPath"`
}
