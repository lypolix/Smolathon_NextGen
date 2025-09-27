import "./App.css";
import { Routes, Route } from "react-router-dom";
import { MainPage } from "./components/MainPage/MainPage";
import { Team } from "./components/Team/Team";
import { Projects } from "./components/Projects/Projects";
import { News } from "./components/News/News";
import { PopupProvider } from "./components/PopupContext";
import { Statistics } from "./components/Statistics/Statistics";
import { Services } from "./components/Services/Services";
import { Admin } from "./components/Admin/Admin";
import { New } from "./components/New/New";
import { Service } from "./components/Services/Service";
import { Vacancies } from "./components/Vacancies/Vacancies";
function App() {
  return (
    <>
      <PopupProvider>
        <Routes>
          <Route path="/" element={<MainPage />} />
          <Route path="/team" element={<Team />} />
          <Route path="/vacancies" element={<Vacancies />} />
          <Route path="/projects" element={<Projects />} />
          <Route path="/news" element={<News />} />
          <Route path="/statistics" element={<Statistics />} />
          <Route path="/services" element={<Services />} />
          <Route path="/services/:id" element={<Service />} />
          <Route path="news/:id" element={<New />} />
          <Route path="/admin" element={<Admin />} />
        </Routes>
      </PopupProvider>
    </>
  );
}

export default App;
