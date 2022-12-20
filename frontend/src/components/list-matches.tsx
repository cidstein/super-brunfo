import { useEffect, useState } from "react";


import Paper from "@mui/material/Paper";
import { styled } from "@mui/material/styles";
import List from "@mui/material/List";
import ListItemText from "@mui/material/ListItemText/ListItemText";
import ListItem from "@mui/material/ListItem";
import MilitaryTechIcon from '@mui/icons-material/MilitaryTech';
import SentimentVeryDissatisfiedIcon from '@mui/icons-material/SentimentVeryDissatisfied';
import PlayCircleOutlineIcon from '@mui/icons-material/PlayCircleOutline';
import ListItemAvatar from "@mui/material/ListItemAvatar";
import Avatar from "@mui/material/Avatar";
import { IconButton } from "@material-ui/core";

export default function ListMatches() {
    const [matches, setMatches] = useState<any[]>([]);
    const Item = styled(Paper)(({ theme }) => ({
        backgroundColor: theme.palette.mode === 'dark' ? '#1A2027' : '#fff',
        ...theme.typography.body2,
        padding: theme.spacing(2),
        textAlign: 'center',
        color: theme.palette.text.secondary,
      }));

    useEffect(() => {
        fetch('http://localhost:8080/listmatches', {
          headers: {
            'Access-Control-Allow-Origin': '*',
            'Content-Type': 'application/json',
            'Accept': 'application/json'
          },
          method: 'GET'
        })
           .then((response) => response.json())
           .then((data) => {
              setMatches(data);
           })
           .catch((err) => {
              console.log(err.message);
           });
      }, []);

    function victory(match: any) {
      if (match.Victory) {
        return <MilitaryTechIcon />
      } else if (match.Finished) {
        return <SentimentVeryDissatisfiedIcon />
      } else {
        return <PlayCircleOutlineIcon />
      }
    }

    return (
      <List sx={{ width: '100%', maxWidth: 360 }}>
        {matches.map((match, index) => (
          <ListItem key={index}>
            <IconButton edge="start" color="inherit" aria-label="menu">
              <Avatar>
                {victory(match)}
              </Avatar>
            </IconButton>
            <ListItemText primary={`Partida #${match.Counter}`} />
          </ListItem>
        ))}
      </List>
    );
}
    