import './App.css';

import { MuiThemeProvider } from '@material-ui/core';

import { Navbar } from './components/navbar';
import theme from "./theme";
import MenuBar from './components/menu';
import ListCards from './components/list-cards';
import ListMatches from './components/list-matches';


function App() {

  
  return (
    <MuiThemeProvider theme={theme}>
      <Navbar />
      <MenuBar />
      <ListMatches />
    </MuiThemeProvider>
  );
}

export default App;
