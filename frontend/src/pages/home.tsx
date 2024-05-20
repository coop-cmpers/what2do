import React, { useEffect } from "react";
import { Paper } from "@mui/material";
import { helloBackend } from "../clients/hello-world";
import { searchRecommendations } from "../clients/recommendations";
import { Timestamp } from "@bufbuild/protobuf";

const Home = () => {
  useEffect(() => {
    helloBackend().then((resp) => console.log(resp.message));
    searchRecommendations({searchParam: "Mexican Restaurant", location: "Chinatown, Sydney", eventTime: Timestamp.now()}).then((resp) => console.log(resp));
  }, []);

  return (
    <Paper>
      <div>
        <h1>"THIS IS NODES"</h1>
      </div>
    </Paper>
  );
}

export default Home;