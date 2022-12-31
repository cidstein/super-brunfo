import { useEffect, useState } from 'react';

import { Carousel, Container, Row, Col, ButtonGroup, Navbar, ToggleButton, Stack } from 'react-bootstrap';

import ListCards from './list-cards';
import ListMatches from './list-matches';
import ActionAreaCard from "./card";

export default function CardCarousel() {
  const [cards, setCards] = useState<any[]>([]);
  const [idx, setIdx] = useState(0);
  const [loadMatches, setLoadMatches] = useState(false);
  const [loadListCard, setLoadListCard] = useState(false);
  const [radioValue, setRadioValue] = useState('1');

  const radios = [
    { name: 'Matches', value: '1' },
    { name: 'Card list', value: '2' },
    { name: 'Game', value: '3' },
  ];

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
      })
      .catch((err) => {
        console.log(err.message);
      });
  }, []);

  return (
    <Container className="container">
      <Row>
        <Navbar bg="light" variant="light">
          <Container>
            <Navbar.Brand >Super Brunfo</Navbar.Brand>
          </Container>
        </Navbar>
      </Row>
      <Row className="show-grid">
        <Col md={2} className="border-right sticky-top">
          <Stack gap={2}>
            <ButtonGroup vertical size='sm'>
              {radios.map((radio, idx) => (
                <ToggleButton
                  key={idx}
                  id={`radio-${idx}`}
                  type="radio"
                  variant="light"
                  name="radio"
                  value={radio.value}
                  checked={radioValue === radio.value}
                  onChange={(e) => setRadioValue(e.currentTarget.value)}
                  onClick={() => {
                    if (radio.value === '1') {
                      setLoadMatches(true);
                      setLoadListCard(false);
                    } else if (radio.value === '2') {
                      setLoadMatches(false);
                      setLoadListCard(true);
                    } else if (radio.value === '3') {
                      setLoadMatches(false);
                      setLoadListCard(false);
                    }
                  }}
                >
                  {radio.name}
                </ToggleButton>
              ))}
            </ButtonGroup>
          </Stack>
        </Col>
        <Col md={10}>
          <Carousel variant="dark">
            {cards.map((card, index) => (
              <Carousel.Item>
                <Carousel.Caption bsPrefix='customCarouselCaption' >
                  <h2>{card.Name}</h2>
                </Carousel.Caption>

                <img
                  className="customImgCarousel"
                  src={card.ImageURL}
                  alt={card.Name}
                />
                <div style={{ position: 'absolute', top: '33%', right: '22%', width: '80px' }}>
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
                </div>

              </Carousel.Item>

            ))}
          </Carousel>
        </Col>
      </Row>
    </Container>
  );
}
