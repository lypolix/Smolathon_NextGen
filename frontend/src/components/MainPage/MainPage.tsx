import { Header } from "../Header/Header";
import "./MainPage.css";
import { useState, useEffect } from "react";
import PublicService from "../../backendConnection/publicInfo/publicInfoService";
import type { Statistics, Traffic } from "../../types";
import { usePopup } from "../PopupContext";

// Ответы бэкенда имеют вложенные ключи { stats: ... } и { traffic: ... }.
// PublicService должен уже возвращать распакованные объекты: stats и traffic.

export function MainPage() {
  const [stats, setStats] = useState<Statistics | null>(null);
  const [traffic, setTraffic] = useState<Traffic | null>(null);
  const { showPopup, closePopup } = usePopup();

  useEffect(() => {
    const load = async () => {
      try {
        // Параллельные запросы к /api/stats и /api/traffic
        const [s, t] = await Promise.all([
          PublicService.getStatisticsInfo(),  // возврат response.data.stats
          PublicService.getTrafficInfo()      // возврат response.data.traffic
        ]);
        setStats(s);
        setTraffic(t);
        console.log("stats:", s, "traffic:", t);
      } catch (e) {
        console.error("Ошибка загрузки данных главной страницы", e);
      }
    };
    load();
  }, []);

  // Подстановка значений без изменения вёрстки:
  // - Аварии: используем доступный агрегат нарушений как показательный индикатор
  const accidentsValue = stats?.violations_total ?? "—";
  // - Перекрытые дороги: в API нет прямого поля; показываем число активных светофоров как инфраструктурный индикатор
  const closedRoadsValue = stats?.traffic_lights_active ?? "—";
  // - Оценка пробок: в API нет единого индекса; берём сумму основных типов светофоров как индикатор загрузки
  const trafficEstimateValue =
    traffic
      ? (traffic.light_types["Т.1"] ?? 0) +
        (traffic.light_types["Т.2"] ?? 0) +
        (traffic.light_types["Т.3"] ?? 0)
      : "—";

  return (
    <>
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
                Центр организации дорожного движения Смоленской области (ЦОДД)
                — это современное региональное учреждение, созданное
                для обеспечения безопасности, эффективности и устойчивого
                развития транспортной системы региона.
              </p>
            </div>
            <img className="d3MainPage" src="/3d1.png" />
          </div>

          <div className="info">
            <div className="block1">
              <span className="nameInf">Ситуация на дорогах</span>
              <div className="accidentsInf">
                <span className="accName">Аварии</span>
                <span className="accNum">{accidentsValue}</span>
              </div>
              <div className="closedRoads">
                <span className="roadsName">Перекрытые дороги</span>
                <span className="roadsNum">{closedRoadsValue}</span>
              </div>
            </div>
            <div className="traficEstimate">
              <span className="traficName">Оценка пробок</span>
              <span className="traficNum">{trafficEstimateValue}</span>
            </div>
          </div>
        </div>
      </div>

      {showPopup === "editor" && (
        <>
          <div className="overlay" onClick={closePopup}></div>
          <div className="mainPagePopup">
            <img onClick={closePopup} className="popupClose" src="/close.png" />
            <div className="popupHeading">
              <h2 className="popupHeadingName">Вход как</h2>
              <div className="popupHeadingName1">Редактор</div>
            </div>
            <input placeholder="Логин" className="popupInputLogin" />
            <input placeholder="Пароль" className="popupInputPassword" type="password" />
            <button className="popupButton">Войти</button>
          </div>
        </>
      )}

      {showPopup === "admin" && (
        <>
          <div className="overlay" onClick={closePopup}></div>
          <div className="mainPagePopup">
            <img onClick={closePopup} className="popupClose" src="/close.png" />
            <div className="popupHeading1">
              <h2 className="popupHeadingName">Вход как</h2>
              <div className="popupHeadingName11">Администратор</div>
            </div>
            <input placeholder="Логин" className="popupInputLogin" />
            <input placeholder="Пароль" className="popupInputPassword" type="password" />
            <button className="popupButton">Войти</button>
          </div>
        </>
      )}
    </>
  );
}
