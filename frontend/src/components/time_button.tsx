import React, { useState } from "react";
import { Grid, Typography } from '@mui/material';
import '@fontsource/roboto/300.css';
import '@fontsource/roboto/400.css';
import '@fontsource/roboto/500.css';
import '@fontsource/roboto/700.css';
import dayjs, { Dayjs } from 'dayjs';
import { DemoContainer } from '@mui/x-date-pickers/internals/demo';
import { LocalizationProvider } from '@mui/x-date-pickers/LocalizationProvider';
import { AdapterDayjs } from '@mui/x-date-pickers/AdapterDayjs';
import { TimePicker } from '@mui/x-date-pickers/TimePicker';


const TimeButton = () => {
    const [value, setValue] = useState<Dayjs | null>(dayjs('2022-04-17T15:30'));

    return (
    <Grid container alignItems="center" flexDirection="column">
        <Grid item>
            <Typography variant="h3" component="h3">
                What time?
            </Typography>
            </Grid>
        <Grid item> 
            <LocalizationProvider dateAdapter={AdapterDayjs}>
            <DemoContainer components={['TimePicker', 'TimePicker']}>
            
            <TimePicker
                label="Pick a time"
                value={value}
                onChange={(newValue) => setValue(newValue)}
            />
            </DemoContainer>
            </LocalizationProvider>
        </Grid>
  </Grid>

    );
}

export default TimeButton;