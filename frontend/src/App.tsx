import React, { useEffect, useState } from 'react';
import './App.css';

import { MuiThemeProvider } from '@material-ui/core';
import { experimentalStyled as styled } from '@mui/material/styles';
import Box from '@mui/material/Box';
import Paper from '@mui/material/Paper';
import Grid from '@mui/material/Grid';

import { Navbar } from './components/navbar';
import theme from "./theme";
import ActionAreaCard from './components/card';



function App() {
  const [cards, setCards] = useState<any[]>([]);
  
  const Item = styled(Paper)(({ theme }) => ({
    backgroundColor: theme.palette.mode === 'dark' ? '#1A2027' : '#fff',
    ...theme.typography.body2,
    padding: theme.spacing(2),
    textAlign: 'center',
    color: theme.palette.text.secondary,
  }));


  useEffect(() => {
    fetch('http://localhost:8080/listcards', {
      headers: {
        'Access-Control-Allow-Origin': '*',
        'Content-Type': 'application/json',
        'Accept': 'application/json'
      },
      method: 'GET'
    })
       .then((response) => response.json())
       .then((data) => {
          setCards(data);
       })
       .catch((err) => {
          console.log(err.message);
       });
  }, []);

  return (
    <MuiThemeProvider theme={theme}>
      <Navbar />
      <Grid container spacing={{ xs: 2, md: 2 }} columns={{ xs: 5, sm: 10, md: 15 }}>
        {cards.map((card, index) => (
          <Grid item xs={2} sm={3} md={3} key={index}>
            <ActionAreaCard  
              name={card.Name}
              attack={card.Attack}
              defense={card.Defense}
              intelligence={card.Intelligence}
              agility={card.Agility}
              resilience={card.Resilience}
              imageURL={card.ImageURL}
            />
          </Grid>
        ))}
      </Grid>
    </MuiThemeProvider>
  );
}

export default App;
