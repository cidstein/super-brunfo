import { useEffect, useState } from "react";

import Grid from "@mui/material/Grid";

import ActionAreaCard from "./card";

export default function ListCards() {
    const [cards, setCards] = useState<any[]>([]);

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
    );
}
    