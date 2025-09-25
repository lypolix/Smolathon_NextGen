import { Header } from "../Header/Header";
import "./MainPage.css";
import { useState, useEffect } from "react";
import PublicService from "../../backendConnection/publicInfo/publicInfoService";
import type { Traffic } from "../../types";

export function MainPage() {
  const [traffic, setTraffic] = useState<Traffic[] | undefined>(undefined);
    useEffect(() => {
      const getAllTraffic = async () => {
        const result = await PublicService.getTrafficInfo();
        setTraffic(result);
        console.log(result);
      };
      getAllTraffic();
    }, []);
  return (
    <>
      <div className="main-page">
        <div className="container">
          <Header menuBgColor="#18211B66" activeTextColor="white"/>
          <div className="cont">
            <div className="cont1">
            <h1>
              Центр организации дорожного движения{" "}
              <span className="highlight">Смоленска</span>
            </h1>
            <p className="mainPageText">
              Центр организации дорожного движения Смоленской области (ЦОДД) —
              это современное региональное учреждение, созданное для обеспечения
              безопасности, эффективности и устойчивого развития транспортной
              системы региона.
            </p>
          </div>
          <img className="d3MainPage" src="/3d1.png" />
          </div>
          
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
    </>
  );
}
