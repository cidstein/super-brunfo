import { useEffect, useState } from "react";

// import { Button } from 'primereact/button';
import { useNavigate } from "react-router-dom";

import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faGamepad } from '@fortawesome/free-solid-svg-icons'
{/* <FontAwesomeIcon icon="fa-regular fa-gamepad" /> */ }
import Round from "./round";
import { Button, ButtonGroup } from "react-bootstrap";



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

  function startNewGame() {
    fetch('http://localhost:8080/start', {
      headers: {
        'Access-Control-Allow-Origin': '*',
        'Content-Type': 'application/json',
        'Accept': 'application/json'
      },
      method: 'POST'
    })
      .then((response) => response.json())
      .then((data) => {
        loadGame(data.ID);
      })
      .catch((err) => {
        console.log(err.message);
      });
  }

  return (
    <div>
      {matches.length === 0 &&
        <div className="surface-0 text-700 text-center">
          <button
            className="font-bold px-5 py-3 p-button-raised p-button-rounded white-space-nowrap"
            onClick={startNewGame}
          >
            <FontAwesomeIcon icon={faGamepad} />
            Start a new game
          </button>
        </div>
      }
      {
        matches.length > 0 &&
        <ButtonGroup vertical size='lg'>
          {matches.map((match, index) => (
            <Button
              key={index}
              variant="light"
              size="lg"
              onClick={() => loadGame(match.ID)}
            >
              {`Partida #${match.Counter}`}
            </Button>
          ))}
        </ButtonGroup>
      }
    </div >
  );
}
