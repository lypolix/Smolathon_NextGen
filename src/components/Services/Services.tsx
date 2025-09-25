import { Header } from "../Header/Header";
import "./Services.css";
import { useState, useEffect } from "react";
import PublicService from "../../backendConnection/publicInfo/publicInfoService";
import type { Services } from "../../types";
export function Services() {
  const [services, setServices] = useState<Services[] | undefined>(undefined);
  useEffect(() => {
    const getAllServices = async () => {
      const result = await PublicService.getServicesInfo();
      setServices(result);
      console.log(result);
    };
    getAllServices();
  }, []);
  return (
    <>
      <div className="serve-page">
        <div className="container">
          <Header
            menuBgColor="white"
            textColor="#203716"
            activeTextColor="white"
          />
          <h1 className="serve-pageHeading">Услуги</h1>
          <div className="serve1">
            <div className="serve1Block1">
              <div className="serve1Block1Heading">
                <h2 className="serve1Block1HeadingText">
                  Платная справка «чистоты автомобиля»
                </h2>
                <div className="serve1Block1HeadingCost">10 000 ₽</div>
              </div>
              <p className="serve1Block1Description">
                Проверка авто перед покупкой на наличие штрафов, ограничений,
                долгов, арестов. Составление отчёта с рекомендациями по
                дальнейшим действиям. Проверка по базе утилизации и розыска
              </p>
            </div>
            <button className="serve1Button">Купить</button>
          </div>
          <div className="serve2"></div>
        </div>
      </div>
    </>
  );
}
