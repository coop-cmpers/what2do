import React, { useEffect } from "react";
import { Paper } from "@mui/material";
import { helloWorld } from "../clients/hello-world";
import { searchRecommendations } from "../clients/recommendations";

const Home = () => {
  useEffect(() => {
    helloWorld().then((resp) => console.log(resp.message));
    searchRecommendations().then((resp) => console.log(resp));
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