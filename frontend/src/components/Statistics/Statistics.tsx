import { Header } from "../Header/Header";
import "./Statistics.css";
import { useState, useEffect } from "react";
import PublicService from "../../backendConnection/publicInfo/publicInfoService";

// Локальный тип под структуру stats из бэкенда
type StatsPayload = {
  collected_amount_total: number;
  evacuations_count: number;
  evacuators_count: number;
  fine_lot_income: number;
  fines_amount_total: number;
  orders_total: number;
  traffic_lights_active: number;
  trips_count: number;
  violations_total: number;
};

function fmtNumber(n?: number) {
  if (typeof n !== "number" || !isFinite(n)) return "—";
  return n.toLocaleString("ru-RU");
}

function fmtCurrency(n?: number) {
  if (typeof n !== "number" || !isFinite(n)) return "—";
  return `${n.toLocaleString("ru-RU")} ₽`;
}

export function Statistics() {
  const [stats, setStats] = useState<StatsPayload | null>(null);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);
  const [chosenOption, setChosenOption] = useState<string>("dtp");

  useEffect(() => {
    let mounted = true;

    const getAllStatistics = async () => {
      try {
        const result: any = await PublicService.getStatisticsInfo();
        if (!mounted) return;
        const payload: StatsPayload | null =
          (result?.stats as StatsPayload) ?? (result as StatsPayload) ?? null;
        if (!payload) {
          setError("Данные статистики отсутствуют");
        } else {
          setStats(payload);
        }
      } catch {
        if (!mounted) return;
        setError("Ошибка загрузки статистики");
      } finally {
        if (!mounted) return;
        setLoading(false);
      }
    };

    getAllStatistics();

    return () => {
      mounted = false;
    };
  }, []);

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
            <h2 className="statContentHeadingText">
              Показатели за текущий год
            </h2>
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

          {loading && <div className="noNewsMessage">Загрузка...</div>}
          {error && !loading && <div className="noNewsMessage">{error}</div>}

          {!loading && !error && stats && (
            <>
              {chosenOption === "dtp" && (
                <div className="statContentInfo">
                  <div className="statContentInfoAmountDTP">
                    <div className="statContentInfoAmountTextDTP">Кол-во ДТП</div>
                    <div className="statContentInfoAmountNumberDTP">
                      {fmtNumber(stats.violations_total)}
                    </div>
                  </div>
                  <div className="statContentInfoAmountInjured">
                    <div className="statContentInfoAmountTextInjured">
                      Активных светофоров
                    </div>
                    <div className="statContentInfoAmountNumberInjured">
                      {fmtNumber(stats.traffic_lights_active)}
                    </div>
                  </div>
                  <div className="statContentInfoAmountDead">
                    <div className="statContentInfoAmountTextDead">Всего поездок</div>
                    <div className="statContentInfoAmountNumberDead">
                      {fmtNumber(stats.trips_count)}
                    </div>
                  </div>
                </div>
              )}

              {chosenOption === "evacuation" && (
                <div className="statContentInfo">
                  <div className="statContentInfoAmountDTP">
                    <div className="statContentInfoAmountTextDTP">Кол-во эвакуаций</div>
                    <div className="statContentInfoAmountNumberDTP">
                      {fmtNumber(stats.evacuations_count)}
                    </div>
                  </div>
                  <div className="statContentInfoAmountInjured">
                    <div className="statContentInfoAmountTextInjured">Кол-во эвакуаторов</div>
                    <div className="statContentInfoAmountNumberInjured">
                      {fmtNumber(stats.evacuators_count)}
                    </div>
                  </div>
                  <div className="statContentInfoAmountDead">
                    <div className="statContentInfoAmountTextDead">Доход со штраф-стоянок</div>
                    <div className="statContentInfoAmountNumberDead">
                      {fmtCurrency(stats.fine_lot_income)}
                    </div>
                  </div>
                </div>
              )}

              {chosenOption === "fines" && (
                <div className="statContentInfo">
                  <div className="statContentInfoAmountFines1">
                    <div className="statContentInfoAmountTextDTP">Общая сумма штрафов</div>
                    <div className="statContentInfoAmountNumberDTP">
                      {fmtCurrency(stats.fines_amount_total)}
                    </div>
                  </div>
                  <div className="statContentInfoAmountFines2">
                    <div className="statContentInfoAmountTextInjured">Сумма собранных штрафов</div>
                    <div className="statContentInfoAmountNumberInjured">
                      {fmtCurrency(stats.collected_amount_total)}
                    </div>
                  </div>
                </div>
              )}
            </>
          )}
        </div>
      </div>
    </div>
  );
}
