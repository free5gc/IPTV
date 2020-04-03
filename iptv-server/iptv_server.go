package iptvserver

import (
	"fmt"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"

	"../factory"
	"./hls-channel"
)

// Server : server information read from config file and pass into this process
type Server struct {
	IptvServerIpv4Port factory.Ipv4Port
	Channels           []factory.IptvChannel
	CacheFolder        string
	WebClient          string
	Version            string
}

// Start : Start of IPTV Application Server
func (s Server) Start() {
	// Start M3U8 encoder
	m3uChannelList := hlschannel.ChannelList{}
	// IPTV Channel List Complie -> m3u
	for _, channel := range s.Channels {
		m3uChannelList = append(m3uChannelList, hlschannel.Channel{Name: channel.ChannelName, VideoPath: channel.VideoPath})
	}
	err := m3uChannelList.Compile(s.CacheFolder)
	if err != nil {
		fmt.Println(err)
		return
	}

	// IPTV Video Source compile -> m3u8
	for _, channel := range m3uChannelList {
		go func(channel hlschannel.Channel) {
			err := channel.Compile(s.CacheFolder)
			if err != nil {
				fmt.Printf("IPTV Channel create fail: %s %+v\n", channel.Name, err)
			} else {
				fmt.Printf("IPTV Channel create successful: %s\n", channel.Name)
			}
		}(channel)
	}

	// Start gin server
	router := gin.Default()
	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
	})
	// Run WebClient Server
	router.Use(static.Serve("/", static.LocalFile(s.WebClient, false)))
	// Run M3U8 Server
	router.Use(static.Serve("/iptv", static.LocalFile(s.CacheFolder, false)))
	router.GET("/version", func(c *gin.Context) {
		c.String(http.StatusOK, s.Version)
	})
	if err := router.Run(s.IptvServerIpv4Port.IPv4Addr + ":" + strconv.Itoa(s.IptvServerIpv4Port.Port)); err != nil {
		fmt.Println(err)
	}
}
