import React from "react";
import Divider from '@mui/material/Divider';
import '@fontsource/roboto/300.css';
import '@fontsource/roboto/400.css';
import '@fontsource/roboto/500.css';
import '@fontsource/roboto/700.css';
import Typography from '@mui/material/Typography';
import { DemoContainer } from '@mui/x-date-pickers/internals/demo';
import { AdapterDayjs } from '@mui/x-date-pickers/AdapterDayjs';
import { LocalizationProvider } from '@mui/x-date-pickers/LocalizationProvider';
import { DatePicker } from '@mui/x-date-pickers/DatePicker';
import { Grid } from "@mui/material";

const DateButton = () => {
  return (
    <Grid container alignItems="center" flexDirection="column">
      <Grid item>
        <Typography variant="h3" component="h3">
          Which date are you looking for?
        </Typography>
      </Grid>
      <Grid item>
        <LocalizationProvider dateAdapter={AdapterDayjs}>
          <DemoContainer components={['DatePicker']}>
            <DatePicker label="Pick a Date" />
          </DemoContainer>
        </LocalizationProvider>
      </Grid>
    </Grid>
  );
}

export default DateButton;