import React, { Component } from "react";

import "./VideoIndexComponent.css";

type IndexListSrc = {
  src: string;
  channelHandler: (newChannel: string) => void;
};
type IndexItemContent = {
  channelName: string;
  channelSrc: string;
};

class VideoIndex extends Component<IndexListSrc> {
  playlist: IndexItemContent[];
  listItems: JSX.Element[];
  state = {
    activeChannel: -1,
  };

  constructor(props: IndexListSrc) {
    super(props);
    this.playlist = [];
    this.listItems = [];
  }

  componentDidMount() {
    fetch(this.props.src, { method: "get" })
      .then((response) => {
        this._parseM3U8(response.text());
      })
      .catch((error) => console.log(error));
  }

  shouldComponentUpdate(nextProps: IndexListSrc, nextState: any) {
    if (nextProps.src !== this.props.src) {
      fetch(nextProps.src, { method: "get" })
        .then((response) => {
          this._parseM3U8(response.text());
        })
        .catch((error) => console.log(error));
      return true;
    }
    if (nextState.activeChannel !== this.state.activeChannel) {
      return true;
    }
    return false;
  }

  private _changeChannel = (index: number) => {
    if (index >= this.playlist.length) {
      return;
    }
    this.setState({ activeChannel: index });
    this.props.channelHandler(this.playlist[index].channelSrc);
  };

  private _parseM3U8(rsp: Promise<string>) {
    // force re-render on line 54 if load new index
    this.setState({ activeChannel: -1 });
    rsp
      .then((value) => {
        var index: string = value;
        //console.log(value);
        if (value.match(/#EXTM3U/g)) {
          //console.log("M3U8 file");
          if (value.match(/#EXTINF:[ ]*(0|-1)/g)) {
            //console.log("is index");
            var localIdxList: string[] = index
              .split("\n#")
              .filter(function (line) {
                return line.match(/EXTINF+/g);
              });
            var channelList: IndexItemContent[] = localIdxList.map(
              (item, idx) => {
                var name: string = item.substring(
                  item.indexOf(",") + 1,
                  item.indexOf("\n")
                );
                var url: string = item.split("\n")[1];
                return { channelName: name, channelSrc: url };
              }
            );
            this.playlist = channelList;
            //console.log("playlist collect complete");
            this._changeChannel(0);
          } else {
            console.log("is not index");
            this.playlist = [
              { channelName: "IPTV", channelSrc: this.props.src },
            ];
            this._changeChannel(0);
          }
        } else {
          console.log("not m3u8 file");
          this.playlist = [
            { channelName: "No IPTV found", channelSrc: "Not an m3u8 file" },
          ];
          this._changeChannel(0);
        }
      })
      .catch((error) => console.log(error));
  }

  render() {
    this.listItems = this.playlist.map((d: IndexItemContent, idx: number) => (
      <button
        className={`list-group-item list-group-item-action + ${
          this.state.activeChannel === idx ? "active" : ""
        }`}
        onClick={() => this._changeChannel(idx)}
        key={idx.toString()}
      >
        {" "}
        {d.channelName}{" "}
      </button>
    ));

    return <div className="player-index list-group">{this.listItems}</div>;
  }
}

export default VideoIndex;
