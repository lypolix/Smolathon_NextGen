import { Header } from "../Header/Header";
import "./MainPage.css";
import { useState, useEffect } from "react";
import PublicService from "../../backendConnection/publicInfo/publicInfoService";
import type { StatisticsResponse } from "../../types";

export function MainPage() {
  const [statistics, setStatistics] = useState<StatisticsResponse | undefined>(undefined);

  useEffect(() => {
    const getAllStatistics = async () => {
      try {
        const result = await PublicService.getStatisticsInfo();
        setStatistics(result);
        console.log("stats:", result);
      } catch (error) {
        console.error("Ошибка при получении статистики:", error);
        setStatistics(undefined);
      }
    };
    getAllStatistics();
  }, []);

  return (
    <div className="main-page">
      <div className="container">
        <Header menuBgColor="#18211B66" activeTextColor="white" />
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
          {statistics ? (
            <>
              <div className="block1">
                <span className="nameInf">Ситуация на дорогах</span>
                <div className="accidentsInf">
                  <span className="accName">Аварии</span>
                  <span className="accNum">
                    {statistics.stats.violations_total}
                  </span>
                </div>
                <div className="closedRoads">
                  <span className="roadsName">Перекрытые дороги</span>
                  <span className="roadsNum">
                    {statistics.stats.evacuations_count}
                  </span>
                </div>
              </div>
              <div className="traficEstimate">
                <span className="traficName">Оценка пробок</span>
                <span className="traficNum">{statistics.stats.trips_count}</span>
              </div>
            </>
          ) : (
            <p>Информации о ситуации на дорогах нет</p>
          )}
        </div>
      </div>
    </div>
  );
}
