import * as React from 'react';

import { Grid, makeStyles } from "@material-ui/core";

import ActionAreaCard from './card';

interface RoundProps {
  id: string;
  matchId: string;
  playerCardId: string;
  npcCardId: string;
  counter: number;
  victory: boolean;
  finished: boolean;
}

const useStyles = makeStyles({
  root: {
    width: "100%",
    height: "100%",
  },
  form: {
    margin: "16px",
  },
  btnSubmitWrapper: {
    textAlign: "center",
    marginTop: "8px",
  },
  map: {
    width: "100%",
    height: "100%",
  },
});

export default function ActionAreaMatch(props: RoundProps) {
  const classes = useStyles();
  const [playerCard, setPlayerCard] = React.useState<any>([]);
  const [npcCard, setNpcCard] = React.useState<any>([]);
  const { id, matchId, playerCardId, npcCardId, counter, victory, finished } = props;

  React.useEffect(() => {
    fetch('http://localhost:8080/getcard', {
      headers: {
        'Access-Control-Allow-Origin': '*',
        'Content-Type': 'application/json',
        'Accept': 'application/json'
      },
      method: 'GET',
      body: JSON.stringify({
        card_id: playerCardId,
      })
    })
       .then((response) => response.json())
       .then((data) => {
        setPlayerCard(data);
       })
       .catch((err) => {
          console.log(err.message);
       });
  }, []);

  React.useEffect(() => {
    fetch('http://localhost:8080/getcard', {
      headers: {
        'Access-Control-Allow-Origin': '*',
        'Content-Type': 'application/json',
        'Accept': 'application/json'
      },
      method: 'GET',
      body: JSON.stringify({
        card_id: npcCardId,
      })
    })
       .then((response) => response.json())
       .then((data) => {
        setNpcCard(data);
       })
       .catch((err) => {
          console.log(err.message);
       });
  }, []);
  
  return (
    <Grid className={classes.root} container>
      <Grid item xs={12} sm={3}>
        <ActionAreaCard  
          name={playerCard.Name}
          attack={playerCard.Attack}
          defense={playerCard.Defense}
          intelligence={playerCard.Intelligence}
          agility={playerCard.Agility}
          resilience={playerCard.Resilience}
          imageURL={playerCard.ImageURL}
        />
      </Grid>
      <Grid item xs={12} sm={3}>
        <ActionAreaCard  
          name={playerCard.Name}
          attack={playerCard.Attack}
          defense={playerCard.Defense}
          intelligence={playerCard.Intelligence}
          agility={playerCard.Agility}
          resilience={playerCard.Resilience}
          imageURL={playerCard.ImageURL}
        />
      </Grid>
    </Grid>
  );
}
