import { useEffect, useState } from 'react';

import { useParams } from 'react-router-dom';
import { Card, Container, Row, Col, ListGroup, Button, ButtonGroup } from 'react-bootstrap';

export default function Round() {
  let params = useParams();
  let id = params.id;
  
  const [ round, setRound ] = useState({
    ID: "",
    Match: {
      ID: "",
      Counter: 0,
      Finished: false,
      Victory: false,
    },
    Cards: [
      {
        ID: "",
        Name: "",
        ImageURL: "",
        Attack: 0,
        Defense: 0,
        Intelligence: 0,
        Agility: 0,
        Resilience: 0,
      },
    ],
    Counter: 0,
    Victory: false,
    Attribute: "",
  })

  useEffect(() => {
    fetch(`http://localhost:8080/loadround?match_id=${id}`, {
      headers: {
        'Access-Control-Allow-Origin': '*',
        'Content-Type': 'application/json',
        'Accept': 'application/json'
      },
      method: 'GET',
    })
       .then((response) => response.json())
       .then((data) => {
        setRound(data);
       })
       .catch((err) => {
          console.log(err.message);
       });
  }, [id]);

  function handleAttributeClick(attribute: string) {
    const requestOptions = {
      method: 'PUT',
      headers: {
        'Access-Control-Allow-Origin': '*',
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ round_id: round.ID, attribute: attribute })

    };
    fetch('http://localhost:8080/playround', requestOptions)
        .then(response => response.json())
        .then(data => setRound(data))
        .catch((err) => {
          console.log(err.message);
        });
  }
  
  return (
    <Container className="container">
      <Row>
        <Col md={2} className="mx-auto my-2">
          <ListGroup className="list-group-flush">
            <ListGroup.Item>{`Match #${round.Match.Counter}`}</ListGroup.Item>
            <ListGroup.Item>{`Round #${round.Counter}`}</ListGroup.Item>
            <ListGroup.Item>{`Victory: ${round.Victory}`}</ListGroup.Item>
            <ListGroup.Item>{`Attribute: ${round.Attribute}`}</ListGroup.Item>
          </ListGroup>
        </Col>
        {round.Cards.map((card, index) => (
          <Col md={5} key={index} className="mx-auto my-2">
              <Card style={{ width: '15rem' }}>
              <Card.Header as="h6">{card.Name}</Card.Header>
              <Card.Img src={card.ImageURL} bsPrefix="customCardImg" />
              <ButtonGroup vertical className="list-group-flush" >
                <Button variant="light" onClick={() => handleAttributeClick("attack")}> {`Attack ${card.Attack}`}</Button>
                <Button variant="light">{`Defense ${card.Defense}`}</Button>
                <Button variant="light">{`Intelligence ${card.Intelligence}`}</Button>
                <Button variant="light">{`Agility ${card.Agility}`}</Button>
                <Button variant="light">{`Resilience ${card.Resilience}`}</Button>
              </ButtonGroup>
            </Card>
          </Col>
        ))}
      </Row>
    </Container>
  );
}
