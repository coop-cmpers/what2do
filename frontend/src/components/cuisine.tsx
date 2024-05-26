import React from "react";
import Box from '@mui/material/Box';
import TextField from '@mui/material/TextField';
import { Grid, Typography } from "@mui/material";

const CuisinePicker = () => {
    return (
    <Grid container alignItems="center" flexDirection="column">
      <Grid item>
        <Typography variant="h3" component="h3">
        Pick a cuisine
    </Typography>
      </Grid>
    <Grid item>
    <Box
        component="form"
        sx={{
            '& > :not(style)': { m: 1, width: '25ch' },
        }}
        noValidate
        autoComplete="off"
        >
        <TextField id="outlined-basic" label="Cuisine" variant="outlined" />
      </Box>
      </Grid>
    </Grid>

    );

}

export default CuisinePicker;