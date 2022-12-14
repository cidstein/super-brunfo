import './App.css';

import {
  BrowserRouter,
  Routes, //replaces "Switch" used till v5
  Route,
  Navigate,
} from "react-router-dom";

import Home from './components/home';
import Round from './components/round';
// import ListCards from './components/list-cards';
import ListMatches from './components/list-matches';
import CardCarousel from './components/card-carousel';
import SignIn from './components/sign-in';
import SignUp from './components/sign-up';

function App() {
  return (
    <div className="App" style={{
      backgroundImage: `url(/background.jpg)`,
      backgroundSize: 'cover',
      backgroundRepeat: 'repeat',
      backgroundPosition: 'center',
      height: '100vh',
      width: '100vw',

    }}>
      <BrowserRouter>
        <Routes>
          <Route path="/home" element={<SignUp />} />
          <Route path="/" element={<Navigate replace to="/home" />} />
          <Route path="/list-cards" element={<CardCarousel />} />
          <Route path="/list-matches" element={<ListMatches />} />
          <Route path="/round/:id" element={<Round />} />
        </Routes>
      </BrowserRouter>
    </div>
  );
}

export default App;
