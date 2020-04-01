# free5GC IPTV
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Ffree5gc%2FIPTV.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Ffree5gc%2FIPTV?ref=badge_shield)


This is an simple IPTV Server that you can self host on Linux / Windows base on [FFmpeg](https://www.ffmpeg.org/) and [Go](https://golang.org/).

This project is maintained by [free5GC](https://free5gc.org) for validation of IPTV applications in 5G core network.

## Installation

### Install Dependencies

- FFmpeg

    ```shell
    sudo apt install -y ffmpeg
    ```

- node.js
  
    ```shell
    sudo apt-get install nodejs
    ```

- yarn

    ```shell
    sudo apt install -y curl
    curl -sS https://dl.yarnpkg.com/debian/pubkey.gpg | sudo apt-key add -
    echo "deb https://dl.yarnpkg.com/debian/ stable main" | sudo tee /etc/apt/sources.list.d/yarn.list
    sudo apt update && sudo apt install yarn
    ```

- npm

    ```shell
    sudo apt-get install python-software-properties
    sudo add-apt-repository ppa:gias-kay-lee/npm
    sudo apt-get update
    sudo apt-get install npm
    ```

- golang

    ```shell
    wget https://dl.google.com/go/go1.12.9.linux-amd64.tar.gz
    sudo tar -C /usr/local -zxvf go1.12.9.linux-amd64.tar.gz
    mkdir -p ~/go/{bin,pkg,src}
    echo 'export GOPATH=$HOME/go' >> ~/.bashrc
    echo 'export GOROOT=/usr/local/go' >> ~/.bashrc
    echo 'export PATH=$PATH:$GOPATH/bin:$GOROOT/bin' >> ~/.bashrc
    echo 'export GO111MODULE=off' >> ~/.bashrc
    source ~/.bashrc
    ```

### Build Project

- Install go package

    ```shell
    go get -u github.com/gin-contrib/static
    go get -u github.com/gin-gonic/gin
    go get -u github.com/urfave/cli
    go get -u gopkg.in/yaml.v2
    ```

- Build Web Client

    ```shell
    cd web-client
    yarn install
    yarn build
    cd ..
    vim iptvcfg.conf # Configure iptv details
    mkdir hls        # Create chache folder, default is ./hls
    go run iptv.go
    ```

## How to watch IPTV

The IPTV channel playlist's location depends on what you configure in `iptvcfg.conf`'s IPTVServer -> ServerAddr

And the playlist's location will be: `IPv4:Port/iptv/index.m3u`

1. Generally
    Using Web to check IPTV channel is an general way to do it. You can use our web to view IPTV channel.

    Open your web browser, goto the ip:port you have configured in `iptvcfg.conf`'s ServerAddr.

2. On PC
    You can use:
    - [VLC](https://www.videolan.org/vlc/index.zh-TW.html)
    - [Kodi](https://kodi.tv/)
    - [IINA](https://iina.io/) (only on OSX)

3. On Android
    - [NET IP TV](https://play.google.com/store/apps/details?id=com.dnamedya.netiptv)
    - [Kodi](https://play.google.com/store/apps/details?id=org.xbmc.kodi)

4. On iOS
     - [GSE SMART IPTV](https://apps.apple.com/us/app/gse-smart-iptv/id1028734023)
     - [Movie Stream: Cast & Streaming](https://apps.apple.com/us/app/movie-stream-ip-tv-films/id1450912244)
  
## Contact Information

You can contact [free5gc.org@gmail.com](mailto:free5gc.org@gmail.com).

## License

We are using [Apache 2.0](./LICENSE.txt) for the project.


[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Ffree5gc%2FIPTV.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Ffree5gc%2FIPTV?ref=badge_large)

## Release Note

### v2020-03-31-01

+ Implement unicast IPTV server
+ Verify IPTV application running on 5G core network