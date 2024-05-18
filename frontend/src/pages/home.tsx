import React, { useEffect } from "react";
import DateButton from "../components/date_button";
import { Grid, Stack, ThemeProvider, Typography, useTheme } from "@mui/material";
import { theme } from "../helpers/theme";
import TimeButton from "../components/time_button";
import CuisinePicker from "../components/cuisine";
import LocationPicker from "../components/location";
import { helloWorld } from "../clients/hello-world";
import { searchRecommendations } from "../clients/recommendations";



const Home = () => {
    useEffect(() => {
    helloWorld().then((resp) => console.log(resp.message));
    searchRecommendations().then((resp) => console.log(resp));
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