/*
 * IPTV Configuration Factory
 */

package factory

// TLS : information of TLS
type TLS struct {
	PEM string `yaml:"pem"`
	Key string `yaml:"key"`
}
