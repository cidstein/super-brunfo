import * as React from 'react';

import { experimentalStyled as styled } from '@mui/material/styles';
import Paper from '@mui/material/Paper';
import Grid from '@mui/material/Grid';

interface MatchProps {
  id: string;
  victory: boolean;
  finished: boolean;
}

const Item = styled(Paper)(({ theme }) => ({
  backgroundColor: theme.palette.mode === 'dark' ? '#1A2027' : '#fff',
  ...theme.typography.body2,
  padding: theme.spacing(1),
  textAlign: 'center',
  color: theme.palette.text.secondary,
}));

export default function ActionAreaMatch(props: MatchProps) {
  const { id, victory, finished } = props;
  return (
    <Grid container rowSpacing={1} columnSpacing={{ xs: 1, sm: 1, md: 1 }}>
      <Grid item xs={6} md={8}>
        <Item>ID</Item>
      </Grid>
      <Grid item xs={6} md={4}>
        <Item>{id}</Item>
      </Grid>
      <Grid item xs={6} md={8}>
        <Item>Vitória</Item>
      </Grid>
      <Grid item xs={6} md={4}>
        <Item>{victory}</Item>
      </Grid>
      <Grid item xs={6} md={8}>
        <Item>Terminada</Item>
      </Grid>
      <Grid item xs={6} md={4}>
        <Item>{finished}</Item>
      </Grid>
    </Grid>
  );
}
