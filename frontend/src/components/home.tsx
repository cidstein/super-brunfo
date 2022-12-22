import { useState } from 'react';

import { Button, ButtonGroup, Container, Navbar } from 'react-bootstrap';
import ToggleButton from 'react-bootstrap/ToggleButton';
import Stack from 'react-bootstrap/Stack';

import ListCards from './list-cards';
  export default function Home() {
    const [loadMatches, setLoadMatches] = useState(false);
    const [loadListCard, setLoadListCard] = useState(false);
    const [loadStatistics, setLoadStatistics] = useState(false);
    const [radioValue, setRadioValue] = useState('1');
  
    const radios = [
      { name: 'Matches', value: '1' },
      { name: 'Card list', value: '2' },
      { name: 'Statistics', value: '3' },
    ];

    return (
        <Stack gap={0}>
            <Navbar bg="dark" variant="dark">
                <Container>
                    <Navbar.Brand href="#home">Super Brunfo</Navbar.Brand>
                </Container>
            </Navbar>
            <Stack direction="horizontal" gap={2}>
                <div className="d-flex flex-column border-right sticky-top">
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
                                        setLoadStatistics(false);
                                    } else if (radio.value === '2') {
                                        setLoadMatches(false);
                                        setLoadListCard(true);
                                        setLoadStatistics(false);
                                    } else if (radio.value === '3') {
                                        setLoadMatches(false);
                                        setLoadListCard(false);
                                        setLoadStatistics(true);
                                    }
                                }}
                                >
                                    {radio.name}
                                </ToggleButton>
                            ))}
                        </ButtonGroup>
                    </Stack>
                </div>
                <div className="bg-light border">
                    { loadMatches && <div>Matches</div>}
                    { loadListCard && <ListCards />}
                    { loadStatistics && <div>Statistics</div>}
                </div>
            </Stack>
        </Stack>
    )
}