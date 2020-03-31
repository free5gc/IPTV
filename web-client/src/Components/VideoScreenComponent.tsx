import React, { Component } from "react";
import ReactPlayer from "react-player";

import "./VideoScreenComponent.css";

type Channel = {
  src: string;
};

class VideoScreen extends Component<Channel> {
  render() {
    return (
      <div className="player-wrapper">
        <ReactPlayer
          key={this.props.src}
          url={this.props.src}
          className="react-player"
          width={"100%"}
          height={"auto"}
          controls
        />
        <p>SRC: {this.props.src}</p>
      </div>
    );
  }
}

export default VideoScreen;
