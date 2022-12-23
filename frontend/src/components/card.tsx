import * as React from 'react';
import { Button, ButtonGroup } from 'react-bootstrap';

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
    <Card style={{ width: '15rem' }}>
      <Card.Header as="h6">{name}</Card.Header>
      <Card.Img src={imageURL} bsPrefix="customCardImg" />
      <ButtonGroup vertical className="list-group-flush" >
        <Button variant="light">{`Attack ${attack}`}</Button>
        <Button variant="light">{`Defense ${defense}`}</Button>
        <Button variant="light">{`Intelligence ${intelligence}`}</Button>
        <Button variant="light">{`Agility ${agility}`}</Button>
        <Button variant="light">{`Resilience ${resilience}`}</Button>
      </ButtonGroup>
    </Card>
  );
}
