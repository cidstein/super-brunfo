import { useEffect, useState } from "react";

import { Button, ButtonGroup } from 'react-bootstrap';
import { useNavigate } from "react-router-dom";

import Round from "./round";

interface Match {
    ID: string;
    Counter: number;
    Finished: boolean;
    Victory: boolean;
}

export default function ListMatches() {
    const [matches, setMatches] = useState<Match[]>([]);
    let navigate = useNavigate();

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

    function loadGame(match_id: string) {
        navigate(`/round/${match_id}`);
    }

    return (
      <ButtonGroup vertical size='lg'>
        {matches.map((match, index) => (
          <Button
            key={index}
            variant="outline-dark"
            size="lg"
            onClick={() => loadGame(match.ID)}
          >
            {`Partida #${match.Counter}`}
          </Button>
        ))}
      </ButtonGroup>
    );
}
    