import { Card, Container, Row, Col, ListGroup, Button, ButtonGroup } from 'react-bootstrap';

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
  const { id, playerCardId, npcCardId, counter, victory, finished } = props;
  const [cards, setCards] = useState<any[]>([]);
  const [ attributeSelected, setAttributeSelected ] = useState<boolean>(false);
  const [ attributeWon, setAttributeWon ] = useState<string>("light");

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

  function handleAttributeClick(attribute: string) {
    setAttributeSelected(true);
    setAttributeWon(attribute);
  }
  
  return (
    <Container className="container">
      <Row>
        <Col md={2} className="mx-auto my-2">
          <ListGroup className="list-group-flush">
            <ListGroup.Item>{`Match #1`}</ListGroup.Item>
            <ListGroup.Item>{`Round #4`}</ListGroup.Item>
            <ListGroup.Item>{`Rounds won: 2`}</ListGroup.Item>
          </ListGroup>
        </Col>
        {cards.map((card, index) => (
          <Col md={5} key={index} className="mx-auto my-2">
            { (index === 0 || attributeSelected) &&
              <Card style={{ width: '15rem' }}>
              <Card.Header as="h6">{card.Name}</Card.Header>
              <Card.Img src={card.ImageURL} bsPrefix="customCardImg" />
              <ButtonGroup vertical className="list-group-flush" >
                <Button
                  variant={
                    (index === 0) ? attributeWon : "danger"
                  }
                  onClick={() => 
                    handleAttributeClick("success")
                  }
                >
                  {`Attack ${card.Attack}`}
                </Button>
                <Button variant="light">{`Defense ${card.Defense}`}</Button>
                <Button variant="light">{`Intelligence ${card.Intelligence}`}</Button>
                <Button variant="light">{`Agility ${card.Agility}`}</Button>
                <Button variant="light">{`Resilience ${card.Resilience}`}</Button>
              </ButtonGroup>
            </Card>
            }
          </Col>
        ))}
      </Row>
    </Container>
  );
}
