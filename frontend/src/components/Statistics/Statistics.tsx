import { Header } from "../Header/Header";
import "./Statistics.css";
import { useState, useEffect } from "react";
import PublicService from "../../backendConnection/publicInfo/publicInfoService";
import type { StatisticsResponse } from "../../types";

export function Statistics() {
  const [statistics, setStatistics] = useState<StatisticsResponse | undefined>(
    undefined
  );
  const [chosenOption, setChosenOption] = useState<string>("dtp");

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

  const fmtNum = (v?: number | null) =>
    v === undefined || v === null ? "–" : v.toLocaleString("ru-RU");

  const fmtMoney = (v?: number | null) =>
    v === undefined || v === null ? "–" : `${v.toLocaleString("ru-RU")} ₽`;

  // короткие алиасы для удобства
  const s = statistics?.stats;

  return (
    <div className="stat-page">
      <div className="container">
        <Header
          menuBgColor="#62A7444D"
          buttonColor="#62A744"
          textColor="#F3FFED"
        />
        <h1 className="statHeading">Статистика</h1>
        <p className="statText">
          Анализ и сравнение данных о ДТП, произошедших в Смоленске
        </p>

        <div className="statContent">
          <div className="statContentHeading">
            <h2 className="statContentHeadingText">Показатели за текущий год</h2>
            <div className="statContentHeadingSections">
              <div
                onClick={() => setChosenOption("dtp")}
                style={
                  chosenOption === "dtp"
                    ? { backgroundColor: "#62A744" }
                    : { backgroundColor: "#62A74499" }
                }
                className="SCHSection1"
              >
                ДТП
              </div>
              <div
                onClick={() => setChosenOption("evacuation")}
                style={
                  chosenOption === "evacuation"
                    ? { backgroundColor: "#E6A152" }
                    : { backgroundColor: "#E6A1524D" }
                }
                className="SCHSection2"
              >
                Эвакуации
              </div>
              <div
                onClick={() => setChosenOption("fines")}
                style={
                  chosenOption === "fines"
                    ? { backgroundColor: "#E65252" }
                    : { backgroundColor: "#E652524D" }
                }
                className="SCHSection3"
              >
                Штрафы
              </div>
            </div>
          </div>

          {/* Блок ДТП */}
          {chosenOption === "dtp" && (
            <div className="statContentInfo">
              <div className="statContentInfoAmountDTP">
                <div className="statContentInfoAmountTextDTP">Кол-во ДТП</div>
                <div className="statContentInfoAmountNumberDTP">
                  {fmtNum(s?.violations_total)}
                </div>
              </div>

              <div className="statContentInfoAmountInjured">
                <div className="statContentInfoAmountTextInjured">
                  Кол-во пострадавших
                </div>
                <div className="statContentInfoAmountNumberInjured">
                  {fmtNum(s?.orders_total)}
                </div>
              </div>

              <div className="statContentInfoAmountDead">
                <div className="statContentInfoAmountTextDead">
                  Погибшие в автокатастрофах
                </div>
                <div className="statContentInfoAmountNumberDead">
                  {/* в модели нет явного поля для погибших */}
                  {"–"}
                </div>
              </div>
            </div>
          )}

          {/* Блок Эвакуации */}
          {chosenOption === "evacuation" && (
            <div className="statContentInfo">
              <div className="statContentInfoAmountDTP">
                <div className="statContentInfoAmountTextDTP">
                  Кол-во эвакуаций
                </div>
                <div className="statContentInfoAmountNumberDTP">
                  {fmtNum(s?.evacuations_count)}
                </div>
              </div>

              <div className="statContentInfoAmountInjured">
                <div className="statContentInfoAmountTextInjured">
                  Кол-во эвакуаторов
                </div>
                <div className="statContentInfoAmountNumberInjured">
                  {fmtNum(s?.evacuators_count)}
                </div>
              </div>

              <div className="statContentInfoAmountDead">
                <div className="statContentInfoAmountTextDead">
                  Доход с платных штраф-стоянок
                </div>
                <div className="statContentInfoAmountNumberDead">
                  {fmtMoney(s?.fine_lot_income)}
                </div>
              </div>
            </div>
          )}

          {/* Блок Штрафы */}
          {chosenOption === "fines" && (
            <div className="statContentInfo">
              <div className="statContentInfoAmountFines1">
                <div className="statContentInfoAmountTextDTP">
                  Общая сумма штрафов
                </div>
                <div className="statContentInfoAmountNumberDTP">
                  {fmtMoney(s?.fines_amount_total)}
                </div>
              </div>

              <div className="statContentInfoAmountFines2">
                <div className="statContentInfoAmountTextInjured">
                  Сумма собранных штрафов
                </div>
                <div className="statContentInfoAmountNumberInjured">
                  {fmtMoney(s?.collected_amount_total)}
                </div>
              </div>
            </div>
          )}
        </div>
      </div>
    </div>
  );
}
