import React, { useEffect } from "react";
import DateButton from "../components/date_button";
import { Grid, ThemeProvider, Typography, useTheme } from "@mui/material";
import { theme } from "../helpers/theme";
import TimeButton from "../components/time_button";
import CuisinePicker from "../components/cuisine";
import LocationPicker from "../components/location";
import { helloBackend } from "../clients/hello-world";
import { searchRecommendations } from "../clients/recommendations";
import { Timestamp } from "@bufbuild/protobuf";


const Home = () => {
  useEffect(() => {
    helloBackend().then((resp: { message: any; }) => console.log(resp.message));
    searchRecommendations({searchParam: "Mexican Restaurant", location: "Chinatown, Sydney", eventTime: Timestamp.now()}).then((resp: any) => console.log(resp));
  }, []);

  return (
    <ThemeProvider theme={(theme)}>
      
      <div style={{ backgroundColor: theme.palette.primary.dark }}>
      <Grid container alignItems="center" flexDirection="column" spacing={6}>
      <Grid item>
        <Typography align="center" variant="h1" style={{ color: theme.palette.primary.main }}><b>
          Where To Meet
          </b> </Typography>
      </Grid>
        
          <Grid item> 
            <DateButton />
          </Grid>

          <Grid item> 
            <TimeButton />
          </Grid>

          <Grid item>
            <CuisinePicker /> 
          </Grid>

          <Grid item> 
            <LocationPicker/>
          </Grid>

        </Grid>
      </div>
      </ThemeProvider>
     
  );
}

export default Home;