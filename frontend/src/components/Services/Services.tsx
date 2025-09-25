import { Header } from "../Header/Header";
import "./Services.css";
import { useState, useEffect } from "react";
import PublicService from "../../backendConnection/publicInfo/publicInfoService";
import type { Services as Service } from "../../types";

// Тип ответа от API: { services: Service[] }
type ServicesResponse = { services: Service[] };

export function Services() {
  const [services, setServices] = useState<Service[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const getAllServices = async () => {
      try {
        setLoading(true);
        const result = await PublicService.getServicesInfo();

        // Нормализация: поддерживаем и { services: [...] }, и просто массив на всякий случай
        const data: Service[] = Array.isArray(result)
          ? (result as Service[])
          : Array.isArray((result as ServicesResponse)?.services)
          ? (result as ServicesResponse).services
          : [];

        setServices(data);
      } catch (e: any) {
        setError(e?.message || "Ошибка загрузки услуг");
        setServices([]);
      } finally {
        setLoading(false);
      }
    };
    getAllServices();
  }, []);

  return (
    <div className="serve-page">
      <div className="container">
        <Header menuBgColor="white" textColor="#203716" activeTextColor="white" />

        <h1 className="serve-pageHeading">Услуги</h1>

        {loading && <div>Загрузка...</div>}
        {error && <div className="error">{error}</div>}

        {/* Важно: services-list станет grid/flex с gap в CSS */}
        <div className="services-list">
          {services.map((s) => (
            // Важно: serve1 получит фикс ширину/паддинги/бордер-радиус и margin/gap в CSS
            <div className="serve1" key={s.id}>
              <div className="serve1Block1">
                <div className="serve1Block1Heading">
                  <h2 className="serve1Block1HeadingText">{s.title}</h2>
                  <div className="serve1Block1HeadingCost">
                    {Intl.NumberFormat("ru-RU").format(s.price)} ₽
                  </div>
                </div>

                <p className="serve1Block1Description">{s.description}</p>

                {s.icon_url && s.icon_url.trim().length > 0 && (
                  <div
                    className="serve1Icon"
                    style={{ backgroundImage: `url(${s.icon_url})` }}
                    aria-label={s.title}
                  />
                )}
              </div>

              <button className="serve1Button">Купить</button>
            </div>
          ))}

          {!loading && !error && services.length === 0 && (
            <div>Список услуг пуст.</div>
          )}
        </div>

        {/* Синий блок снизу создаёт артефакт — в Services.css нужно убрать фон/высоту у .serve2 */}
        <div className="serve2" aria-hidden="true"></div>
      </div>
    </div>
  );
}
