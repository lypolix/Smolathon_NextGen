import { Header } from "../Header/Header";
import "./Statistics.css";
import { useState, useEffect } from "react";
import PublicService from "../../backendConnection/publicInfo/publicInfoService";
import type { Statistics } from "../../types";
export function Statistics() {
  const [statistics, setStatistics] = useState<Statistics[] | undefined>(
    undefined
  );
  useEffect(() => {
    const getAllStatistics = async () => {
      const result = await PublicService.getStatisticsInfo();
      setStatistics(result);
      console.log(result);
    };
    getAllStatistics();
  }, []);
  const [chosenOption, setChosenOption] = useState<string>("dtp");
  return (
    <>
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
              <h2 className="statContentHeadingText">
                Показатели за текущий год
              </h2>
              <div className="statContentHeadingSections">
                <div
                  onClick={() => setChosenOption("dtp")}
                  style={
                    chosenOption == "dtp"
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
                    chosenOption == "evacuation"
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
                    chosenOption == "fines"
                      ? { backgroundColor: "#E65252" }
                      : { backgroundColor: "#E652524D" }
                  }
                  className="SCHSection3"
                >
                  Штрафы
                </div>
              </div>
            </div>
            {chosenOption === "dtp" && (
              <div className="statContentInfo">
                <div className="statContentInfoAmountDTP">
                  <div className="statContentInfoAmountTextDTP">Кол-во ДТП</div>
                  <div className="statContentInfoAmountNumberDTP">7 812</div>
                </div>
                <div className="statContentInfoAmountInjured">
                  <div className="statContentInfoAmountTextInjured">
                    Кол-во пострадавших
                  </div>
                  <div className="statContentInfoAmountNumberInjured">812</div>
                </div>
                <div className="statContentInfoAmountDead">
                  <div className="statContentInfoAmountTextDead">
                    Погибшие в автокатастрофах
                  </div>
                  <div className="statContentInfoAmountNumberDead">132</div>
                </div>
              </div>
            )}
            {chosenOption === "evacuation" && (
              <div className="statContentInfo">
                <div className="statContentInfoAmountDTP">
                  <div className="statContentInfoAmountTextDTP">Кол-во Эвакуаций</div>
                  <div className="statContentInfoAmountNumberDTP">7 784</div>
                </div>
                <div className="statContentInfoAmountInjured">
                  <div className="statContentInfoAmountTextInjured">
                    Кол-во эвакуаторов
                  </div>
                  <div className="statContentInfoAmountNumberInjured">153</div>
                </div>
                <div className="statContentInfoAmountDead">
                  <div className="statContentInfoAmountTextDead">
                    Доход с плтных штраф-стоянок
                  </div>
                  <div className="statContentInfoAmountNumberDead">12 345 ₽</div>
                </div>
              </div>
            )}
            {chosenOption === "fines" && (
              <div className="statContentInfo">
                <div className="statContentInfoAmountFines1">
                  <div className="statContentInfoAmountTextDTP">Общая сумма штрафов</div>
                  <div className="statContentInfoAmountNumberDTP">123 784 ₽</div>
                </div>
                <div className="statContentInfoAmountFines2">
                  <div className="statContentInfoAmountTextInjured">
                    Сумма собранных штрафов
                  </div>
                  <div className="statContentInfoAmountNumberInjured">12 345 ₽</div>
                </div>
              </div>
            )}
          </div>
        </div>
      </div>
    </>
  );
}
