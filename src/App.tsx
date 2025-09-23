import './App.css'
import {Routes, Route} from "react-router-dom";
import { MainPage } from './components/MainPage/MainPage';
import { Team } from './components/Team/Team';
import { Projects } from './components/Projects/Projects';
import { News } from './components/News/News';
import { Statistics } from './components/Statistics/Statistics';
import { Reestr } from './components/Reestr/Reestr';
import { Entrance } from './components/Entrance/Entrance';
import { Admin } from './components/Admin/Admin';
function App() {
  
  return (
    <>
      <Routes>
        <Route path="/" element={<MainPage />} />
        <Route path="/team" element={<Team />} />
        <Route path="/projects" element={<Projects />} />
        <Route path="/news" element={<News />} />
        <Route path="/statistics" element={<Statistics />} />
        <Route path="/reestr" element={<Reestr />} />
        <Route path="/entrance" element={<Entrance />} />

        <Route path="/admin" element={<Admin />} />
      </Routes>
    </>
  )
}

export default App
