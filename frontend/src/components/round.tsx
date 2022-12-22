import { Container, Row, Col } from 'react-bootstrap';

import ActionAreaCard from './card';
import { useEffect, useState } from 'react';

interface RoundProps {
  id: string;
  matchId: string;
  playerCardId: string;
  npcCardId: string;
  counter: number;
  victory: boolean;
  finished: boolean;
}

export default function Round(props: RoundProps) {
  const [cards, setCards] = useState<any[]>([]);
  const { id, playerCardId, npcCardId, counter, victory, finished } = props;

  useEffect(() => {
    fetch(`http://localhost:8080/getroundcards?id=${id}`, {
      headers: {
        'Access-Control-Allow-Origin': '*',
        'Content-Type': 'application/json',
        'Accept': 'application/json'
      },
      method: 'GET',
    })
       .then((response) => response.json())
       .then((data) => {
        setCards(data);
       })
       .catch((err) => {
          console.log(err.message);
       });
  }, [cards]);
  
  return (
    <Container className="container">
      <Row>
        {cards.map((card, index) => (
          <Col xxl={3} key={index} className="mx-auto my-2">
            <ActionAreaCard  
              name={card.Name}
              attack={card.Attack}
              defense={card.Defense}
              intelligence={card.Intelligence}
              agility={card.Agility}
              resilience={card.Resilience}
              imageURL={card.ImageURL}
            />
          </Col>
        ))}
      </Row>
    </Container>
  );
}
