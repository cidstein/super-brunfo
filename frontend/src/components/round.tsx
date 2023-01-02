import { useEffect, useState } from 'react';

import { useParams } from 'react-router-dom';
import { Card, Container, Row, Col, ListGroup, Button, ButtonGroup, Navbar, Modal } from 'react-bootstrap';
import ActionAreaCard from './card';
import { Sidebar } from 'primereact/sidebar';
import { RadioButton } from 'primereact/radiobutton';

export default function Round() {
  let params = useParams();
  let id = params.id;

  const [round, setRound] = useState({
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
        Attack: 0,
        Defense: 0,
        Intelligence: 0,
        Agility: 0,
        Resilience: 0,
        FlavourText: "",
        ImageURL: "",
      },
    ],
    Counter: 0,
    Victory: false,
    Attribute: "",
  })
  const [visible, setVisible] = useState(false);
  const categories = [{ name: "Attack", value: "Attack" }, { name: "Defense", value: "Defense" }, { name: "Intelligence", value: "Intelligence" }, { name: "Agility", value: "Agility" }, { name: "Resilience", value: "Resilience" }]
  const [attribute, setAttribute] = useState(null);

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

  function handleAttributeClick() {
    console.log("handleAttributeClick");
    // StaticExample();
    // const requestOptions = {
    //   method: 'PUT',
    //   headers: {
    //     'Access-Control-Allow-Origin': '*',
    //     'Content-Type': 'application/json',
    //   },
    //   body: JSON.stringify({ round_id: round.ID, attribute: attribute })

    // };
    // fetch('http://localhost:8080/playround', requestOptions)
    //   .then(response => response.json())
    //   .then(data => setRound(data))
    //   .catch((err) => {
    //     console.log(err.message);
    //   });
  }

  return (
    <Container className="container">
      <Sidebar visible={visible} position="top" className="p-sidebar-lg" onHide={() => setVisible(false)}>
        <h5>Basic</h5>
        <div className="field-radiobutton">
          <RadioButton inputId="attribute1" name="attribute" value="Attack" onChange={(e) => setAttribute(e.value)} checked={attribute === 'Attack'} />
          <label htmlFor="attribute1">Attack</label>
        </div>
        <div className="field-radiobutton">
          <RadioButton inputId="attribute2" name="attribute" value="Defense" onChange={(e) => setAttribute(e.value)} checked={attribute === 'Defense'} />
          <label htmlFor="attribute2">Defense</label>
        </div>
        <div className="field-radiobutton">
          <RadioButton inputId="attribute3" name="attribute" value="Intelligence" onChange={(e) => setAttribute(e.value)} checked={attribute === 'Intelligence'} />
          <label htmlFor="attribute3">Intelligence</label>
        </div>
        <div className="field-radiobutton">
          <RadioButton inputId="attribute4" name="attribute" value="Agility" onChange={(e) => setAttribute(e.value)} checked={attribute === 'Agility'} />
          <label htmlFor="attribute4">Agility</label>
        </div>
        <div className="field-radiobutton">
          <RadioButton inputId="attribute4" name="attribute" value="Resilience" onChange={(e) => setAttribute(e.value)} checked={attribute === 'Resilience'} />
          <label htmlFor="attribute4">Resilience</label>
        </div>
      </Sidebar>
      <Row>
        <Navbar bg="dark" variant="dark">
          <Container>
            {/* <Navbar.Brand >Super Brunfo</Navbar.Brand> */}
            <Navbar.Text className="mx-auto">
              {`Match #${round.Match.Counter} - Round #${round.Counter}`}
            </Navbar.Text>
          </Container>
        </Navbar>
      </Row>
      <Row>
        {/* <Col md={5}>
          <ActionAreaCard
            name={round.Cards[0].Name}
            attack={round.Cards[0].Attack}
            defense={round.Cards[0].Defense}
            intelligence={round.Cards[0].Intelligence}
            agility={round.Cards[0].Agility}
            resilience={round.Cards[0].Resilience}
            flavourText={round.Cards[0].FlavourText}
            imageURL={round.Cards[0].ImageURL}
          />
        </Col> */}
        <Col md={2}>
          <Button
            variant="outline-light"
            type='submit'
            onClick={() => setVisible(true)}
            bsPrefix="button-78"
          >
            Combat
          </Button>
        </Col>
        {/* <Col md={5}>
          <ActionAreaCard
            name={round.Cards[1].Name}
            attack={round.Cards[1].Attack}
            defense={round.Cards[1].Defense}
            intelligence={round.Cards[1].Intelligence}
            agility={round.Cards[1].Agility}
            resilience={round.Cards[1].Resilience}
            flavourText={round.Cards[1].FlavourText}
            imageURL={round.Cards[1].ImageURL}
          />
        </Col> */}
      </Row>
    </Container>
  );
}
