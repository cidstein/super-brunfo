import { useEffect, useState } from "react";

import Container from 'react-bootstrap/Container';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';

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
        console.log(data);
      })
      .catch((err) => {
        console.log(err.message);
      });
  }, []);

  return (
    <Container>
      <Row >
        {cards.map((card, index) => (
          <Col xxl={6} key={index} className="mx-auto my-2">
            <ActionAreaCard
              name={card.Name}
              attack={card.Attack}
              defense={card.Defense}
              intelligence={card.Intelligence}
              agility={card.Agility}
              resilience={card.Resilience}
              flavourText={card.FlavourText}
              imageURL={card.ImageURL}
            />
          </Col>
        ))}
      </Row>
    </Container>
  );
}
