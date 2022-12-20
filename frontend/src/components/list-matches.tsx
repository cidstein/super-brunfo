import { useEffect, useState } from "react";


import Grid from "@mui/material/Grid";
import Paper from "@mui/material/Paper";
import { styled } from "@mui/material/styles";

import ActionAreaMatch from "./match";

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
              console.log(data)
              setMatches(data);
           })
           .catch((err) => {
              console.log(err.message);
           });
      }, []);

    return (
        <Grid container spacing={{ xs: 2, md: 2 }} columns={{ xs: 5, sm: 10, md: 15 }}>
        {matches.map((match, index) => (
          <Grid item xs={2} sm={3} md={3} key={index}>
            <ActionAreaMatch  
              id={match.ID }
              victory={match.victory}
              finished={match.finished}
            />
          </Grid>
        ))}
      </Grid>
    );
}
    