import * as React from 'react';

import Card from 'react-bootstrap/Card';
import ListGroup from 'react-bootstrap/ListGroup';

interface CardProps {
  name: string;
  attack: string;
  defense: string;
  intelligence: string;
  agility: string;
  resilience: string;
  imageURL: string;
}

export default function ActionAreaCard(props: CardProps) {
  const { name, attack, defense, intelligence, agility, resilience, imageURL } = props;
  return (
    <Card style={{ width: '18rem' }}>
      <Card.Header as="h5">{name}</Card.Header>
      <Card.Img variant="top" src={imageURL} />
      <ListGroup className="list-group-flush">
        <ListGroup.Item>{`Attack ${attack}`}</ListGroup.Item>
        <ListGroup.Item>{`Defense ${defense}`}</ListGroup.Item>
        <ListGroup.Item>{`Intelligence ${intelligence}`}</ListGroup.Item>
        <ListGroup.Item>{`Agility ${agility}`}</ListGroup.Item>
        <ListGroup.Item>{`Resilience ${resilience}`}</ListGroup.Item>
      </ListGroup>
    </Card>
  );
}
