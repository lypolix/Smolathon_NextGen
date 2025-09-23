import { Header } from "../Header/Header";
import "./MainPage.css";

export function MainPage() {
  return (
    <div className="main-page">
      <div className="container">
        <Header />
        <h1>
          Центр организации дорожного движения{" "}
          <span className="highlight">Смоленска</span>
        </h1>
        <div className="info">
          <div className="block1">
            <span className="nameInf">Ситуация на дорогах</span>
            <div className="accidentsInf">
              <span className="accName">Аварии</span>
              <span className="accNum">12</span>
            </div>
            <div className="closedRoads">
              <span className="roadsName">Перекрытые дороги</span>
              <span className="roadsNum">8</span>
            </div>
          </div>
          <div className="traficEstimate">
            <span className="traficName">Оценка пробок</span>
            <span className="traficNum">8</span>
          </div>
        </div>
      </div>
    </div>
  );
}
