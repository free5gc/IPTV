package hlschannel

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// ChannelList : group of Channel structure
type ChannelList []Channel

// Channel : single channel structure
type Channel struct {
	Name      string
	VideoPath string
}

// Compile : compile group of channel into m3u playlist index
func (channelList *ChannelList) Compile(cacheFolder string) (err error) {
	file, err := os.OpenFile(filepath.Join(cacheFolder, "index.m3u"), os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return
	}
	defer file.Close()

	for _, channel := range *channelList {
		fmt.Fprintf(file, "#EXTM3U\n#EXTINF:-1,%s\n/iptv/%s/%s.m3u8\n", channel.Name, channel.Name, channel.Name)
	}

	return
}

// Compile : for each single channel, encode to m3u8 playlist and ts frame and run streaming
func (channel *Channel) Compile(cacheFolder string) (err error) {
	// TODO: Use relative path now, maybe use absolute path in the future
	videoStoreRootPath := fmt.Sprintf(cacheFolder+"/%s/", channel.Name)
	_ = os.Mkdir(videoStoreRootPath, 0770)

	command := exec.Command("ffmpeg", "-re",
		"-stream_loop", "-1",
		"-i", channel.VideoPath,
		"-codec:", "copy",
		"-bsf:v", "h264_mp4toannexb",
		"-hls_time", "6",
		"-hls_segment_filename", fmt.Sprintf("%s_%%03d.ts", videoStoreRootPath+channel.Name),
		"-hls_segment_type", "mpegts",
		"-hls_flags", "delete_segments+append_list+omit_endlist+temp_file",
		"-f", "hls",
		fmt.Sprintf("%s%s.m3u8", videoStoreRootPath, channel.Name),
	)

	err = command.Start()
	if err != nil {
		return
	}
	err = command.Wait()
	return
}
