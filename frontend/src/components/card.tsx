import * as React from 'react';

import Card from '@mui/material/Card';
import CardContent from '@mui/material/CardContent';
import CardMedia from '@mui/material/CardMedia';
import Typography from '@mui/material/Typography';
import { CardActionArea } from '@mui/material';
import { experimentalStyled as styled } from '@mui/material/styles';
import Box from '@mui/material/Box';
import Paper from '@mui/material/Paper';
import Grid from '@mui/material/Grid';

interface CardProps {
  name: string;
  attack: string;
  defense: string;
  intelligence: string;
  agility: string;
  resilience: string;
  imageURL: string;
}

const Item = styled(Paper)(({ theme }) => ({
  backgroundColor: theme.palette.mode === 'dark' ? '#1A2027' : '#fff',
  ...theme.typography.body2,
  padding: theme.spacing(2),
  textAlign: 'center',
  color: theme.palette.text.secondary,
}));

export default function ActionAreaCard(props: CardProps) {
  const { name, attack, defense, intelligence, agility, resilience, imageURL } = props;
  return (
    <Card sx={{ maxWidth: 200 }}>
      <CardActionArea>
        <CardMedia
          component="img"
          height="140"
          image={imageURL}
          alt={name}
        />
        <CardContent>
          <Typography gutterBottom variant="h5" component="div">
            {name}
          </Typography>
          <Grid container rowSpacing={1} columnSpacing={{ xs: 1, sm: 1, md: 1 }}>
            <Grid item xs={6} md={8}>
              <Item>Ataque</Item>
            </Grid>
            <Grid item xs={6} md={4}>
              <Item>{attack}</Item>
            </Grid>
            <Grid item xs={6} md={8}>
              <Item>Defesa</Item>
            </Grid>
            <Grid item xs={6} md={4}>
              <Item>{defense}</Item>
            </Grid>
            <Grid item xs={6} md={8}>
              <Item>Inteligência</Item>
            </Grid>
            <Grid item xs={6} md={4}>
              <Item>{intelligence}</Item>
            </Grid>
            <Grid item xs={6} md={8}>
              <Item>Agilidade</Item>
            </Grid>
            <Grid item xs={6} md={4}>
              <Item>{agility}</Item>
            </Grid>
            <Grid item xs={6} md={8}>
              <Item>Resiliência</Item>
            </Grid>
            <Grid item xs={6} md={4}>
              <Item>{resilience}</Item>
            </Grid>
          </Grid>
        </CardContent>
      </CardActionArea>
    </Card>
  );
}
