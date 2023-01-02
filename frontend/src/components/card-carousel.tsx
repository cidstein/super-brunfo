import { useEffect, useState } from 'react';

import { Carousel } from 'react-bootstrap';

import ActionAreaCard from "./card";

export default function CardCarousel() {
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
      })
      .catch((err) => {
        console.log(err.message);
      });
  }, []);

  return (
    <Carousel variant="dark">
      {cards.map((card, index) => (
        <Carousel.Item>
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
        </Carousel.Item>
      ))}
    </Carousel>
  );
}
