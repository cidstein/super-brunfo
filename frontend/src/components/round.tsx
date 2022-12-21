import * as React from 'react';

import { Grid, makeStyles } from "@material-ui/core";

import ActionAreaCard from './card';
import { useEffect } from 'react';

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

export default function Round(props: RoundProps) {
  const classes = useStyles();
  const [matchId, setMatchId] = React.useState<string>("");
  const [roundId, setRoundId] = React.useState<string>("");
  const [npcCardId, setNpcCardId] = React.useState<string>("");
  const [playerCard, setPlayerCard] = React.useState<any>([]);
  const [npcCard, setNpcCard] = React.useState<any>([]);
  const { playerCardId, counter, victory, finished } = props;

  useEffect(() => {
    fetch(`http://localhost:8080/getcard?id=${playerCardId}`, {
      headers: {
        'Access-Control-Allow-Origin': '*',
        'Content-Type': 'application/json',
        'Accept': 'application/json'
      },
      method: 'GET',
    })
       .then((response) => response.json())
       .then((data) => {
        setPlayerCard(data);
       })
       .catch((err) => {
          console.log(err.message);
       });
  }, []);

  // useEffect(() => {
  //   fetch('http://localhost:8080/getcard', {
  //     headers: {
  //       'Access-Control-Allow-Origin': '*',
  //       'Content-Type': 'application/json',
  //       'Accept': 'application/json'
  //     },
  //     method: 'GET',
  //     body: JSON.stringify({
  //       card_id: npcCardId,
  //     })
  //   })
  //      .then((response) => response.json())
  //      .then((data) => {
  //       setNpcCard(data);
  //      })
  //      .catch((err) => {
  //         console.log(err.message);
  //      });
  // }, []);
  
  return (
    <Grid className={classes.root} container>
      <Grid item xs={6}>
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
      {/* <Grid item xs={6}>
        <ActionAreaCard  
          name={npcCard.Name}
          attack={npcCard.Attack}
          defense={npcCard.Defense}
          intelligence={npcCard.Intelligence}
          agility={npcCard.Agility}
          resilience={npcCard.Resilience}
          imageURL={npcCard.ImageURL}
        />
      </Grid> */}
    </Grid>
  );
}
