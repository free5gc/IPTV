import React, { Component } from "react";
//import { InputGroup, Button, FormControl } from "react-bootstrap";
import logo from "./Asserts/free5gc.png";
import VideoIndex from "Components/VideoIndexComponent";
import VideoScreen from "Components/VideoScreenComponent";

import "./App.css";
import "bootstrap/dist/css/bootstrap.min.css";

//const defaultPlaylistUrl: string = "http://10.10.2.29/index.m3u";
const defaultPlaylistUrl: string = window.location.href + "iptv/index.m3u";

class App extends Component {
  state = {
    channel: "",
    input: "",
    indexSrc: "",
    version: "",
  };

  constructor(props: any) {
    super(props);
    this.state = {
      channel: "Not found",
      input: defaultPlaylistUrl,
      indexSrc: defaultPlaylistUrl,
      version: "0000-00-00-00",
    };
    fetch(window.location.href + "version", { method: "get" })
      .then((response: Response) => response.text())
      .then((value) => {
        if (value.match(/^\d{4}-\d{2}-\d{2}-\d{2}$/)) {
          this.setState({ version: value });
        } else {
          this.setState({ version: "version error" });
        }
      })
      .catch((error: Error) => {
        console.log(error);
        this.setState({ version: "server not found" });
      });
  }

  ChangeChannel = (newChannel: string) => {
    //console.log(newChannel);
    this.setState({ channel: newChannel });

    this._loadIndexSrc = this._loadIndexSrc.bind(this);
    this._handleChange = this._handleChange.bind(this);
  };

  private _loadIndexSrc() {
    this.setState({ indexSrc: this.state.input });
  }

  private _handleChange(event: React.FormEvent<HTMLSelectElement>) {
    this.setState({ input: event.currentTarget.value });
  }

  render() {
    return (
      <div className="App">
        <div className="App-header">
          <img className="logo" src={logo} alt="Logo" />
          {/*<InputGroup className="mb-4">*/}
          {/*  <FormControl*/}
          {/*    type="text"*/}
          {/*    placeholder="IPTV Index"*/}
          {/*    aria-label="IPTV Index"*/}
          {/*    value={this.state.input}*/}
          {/*    onChange={(e: React.FormEvent<HTMLSelectElement>) =>*/}
          {/*      this._handleChange(e)*/}
          {/*    }*/}
          {/*  />*/}
          {/*  <InputGroup.Append>*/}
          {/*    <Button*/}
          {/*      variant="outline-secondary"*/}
          {/*      onClick={() => this._loadIndexSrc()}*/}
          {/*    >*/}
          {/*      Load*/}
          {/*    </Button>*/}
          {/*  </InputGroup.Append>*/}
          {/*</InputGroup>*/}
        </div>
        <div className="App-body container">
          <div className="row">
            <div className="col-lg-8 col-md-12">
              <VideoScreen src={this.state.channel} />
            </div>
            <div className="col-lg-4 col-md-12">
              <VideoIndex
                src={this.state.indexSrc}
                channelHandler={this.ChangeChannel}
              />
            </div>
          </div>
        </div>
        <div className="App-footer">
          CopyRight @free5GC <br />
          IPTV version: {this.state.version}
        </div>
      </div>
    );
  }
}

export default App;
