/*
 * IPTV Configuration Factory
 */

package factory

// Config : All detail configuration of IPTV
type Config struct {
	Info Info `yaml:"info"`

	Configuration Configuration `yaml:"configuration"`
}
