import './App.css';

import { MuiThemeProvider } from '@material-ui/core';

import { Navbar } from './components/navbar';
import theme from "./theme";
import MenuBar from './components/menu';
import ListCards from './components/list-cards';
import ListMatches from './components/list-matches';
import Round from './components/round';


function App() {

  
  return (
    <MuiThemeProvider theme={theme}>
      <Navbar />
      <MenuBar />
      {/* <ListCards /> */}
      <Round 
        id={'9a6ae5d4-a8e7-4967-8c5a-dab776089e83'}
        matchId={'76af9dd2-6cfa-4d10-bb5f-c25a08d7517d'}
        playerCardId={'9f240160-ed87-44db-b9d9-07a4ea180c28'}
        npcCardId={'d0798de7-968b-4a7d-975a-a43fa39f07c6'}
        counter={1}
        victory={false}
        finished={false}
      />
    </MuiThemeProvider>
  );
}

export default App;
