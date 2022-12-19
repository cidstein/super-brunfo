import { AppBar, IconButton, Toolbar, Typography } from "@material-ui/core";
import { FunctionComponent } from "react";

import PlayCircleFilledWhiteOutlinedIcon from '@material-ui/icons/PlayCircleFilledWhiteOutlined';

export const Navbar: FunctionComponent = () => {
  return (
    <AppBar position="static">
      <Toolbar>
        <IconButton edge="start" color="inherit" aria-label="menu">
          <PlayCircleFilledWhiteOutlinedIcon />
        </IconButton>
        <Typography variant="h6">Super Brunfo</Typography>
      </Toolbar>
    </AppBar>
  );
};
