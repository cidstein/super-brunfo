import './App.css';

import { MuiThemeProvider } from '@material-ui/core';

import theme from "./theme";
import Home from './components/home';

function App() {
  return (
    <MuiThemeProvider theme={theme}>
      {/* <ListCards /> */}
      <Home />
      {/* <Round 
        id={'9a6ae5d4-a8e7-4967-8c5a-dab776089e83'}
        matchId={'76af9dd2-6cfa-4d10-bb5f-c25a08d7517d'}
        playerCardId={'9f240160-ed87-44db-b9d9-07a4ea180c28'}
        npcCardId={'d0798de7-968b-4a7d-975a-a43fa39f07c6'}
        counter={1}
        victory={false}
        finished={false}
      /> */}
    </MuiThemeProvider>
  );
}

export default App;
