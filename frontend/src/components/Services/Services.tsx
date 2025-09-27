import { Header } from "../Header/Header";
import "./Services.css";
import { useState, useEffect } from "react";
import PublicService from "../../backendConnection/publicInfo/publicInfoService";
import type { ServicesResponse } from "../../types";
import { useNavigate } from "react-router-dom";

export function Services() {
  const navigate = useNavigate();
  const [services, setServices] = useState<ServicesResponse | undefined>(undefined);

  useEffect(() => {
    const getAllServices = async () => {
      const result = await PublicService.getServicesInfo();
      setServices(result);
      console.log(result);
    };
    getAllServices();
  }, []);

  // последние 4 услуги
  const lastServices = services?.services.slice(-4) ?? [];

  return (
    <div className="serve-page">
      <div className="container">
        <Header
          menuBgColor="white"
          textColor="#203716"
          activeTextColor="white"
        />
        <h1 className="serve-pageHeading">Услуги</h1>

        {/* Блок 1 */}
        <div className="serve1">
          <div className="serve1Block1">
            <div className="serve1Block1Heading">
              <h2 className="serve1Block1HeadingText">
                {lastServices[0]?.title ?? "Название услуги"}
              </h2>
              <div className="serve1Block1HeadingCost">
                {lastServices[0]?.price
                  ? `${lastServices[0].price} ₽`
                  : "Цена не указана"}
              </div>
            </div>
            <p className="serve1Block1Description">
              {lastServices[0]?.description ?? "Описание услуги"}
            </p>
          </div>
          <button
            onClick={() =>
              lastServices[0] && navigate(`/services/${lastServices[0].id}`)
            }
            className="serve1Button"
          >
            Купить
          </button>
        </div>

        {/* Блоки 2–4 */}
        <div className="servesInLine">
          <div className="serve2">
            <div className="serve2Container">
              <div className="serve2ContainerHeading">
                <div className="serve2ContainerHeadingName">
                  {lastServices[1]?.title ?? "Название услуги"}
                </div>
                <div className="serve2ContainerHeadingCost">
                  {lastServices[1]?.price
                    ? `${lastServices[1].price} ₽`
                    : "Цена не указана"}
                </div>
              </div>
              <img className="serve2car" src="/car3d.png" />
              <div className="serve2Content">
                {lastServices[1]?.description ?? "Описание услуги"}
              </div>
            </div>
            <button
              onClick={() =>
                lastServices[1] && navigate(`/services/${lastServices[1].id}`)
              }
              className="serve2Button"
            >
              Купить
            </button>
          </div>

          <div className="serve3">
            <div className="serve2Container">
              <div className="serve2ContainerHeading">
                <div className="serve2ContainerHeadingName">
                  {lastServices[2]?.title ?? "Название услуги"}
                </div>
                <div className="serve2ContainerHeadingCost">
                  {lastServices[2]?.price
                    ? `${lastServices[2].price} ₽`
                    : "Цена не указана"}
                </div>
              </div>
              <img className="serve2car" src="/car3d2.png" />
              <div className="serve2Content">
                {lastServices[2]?.description ?? "Описание услуги"}
              </div>
            </div>
            <button
              onClick={() =>
                lastServices[2] && navigate(`/services/${lastServices[2].id}`)
              }
              className="serve2Button"
            >
              Купить
            </button>
          </div>

          <div className="serve4">
            <div className="serve2Container">
              <div className="serve2ContainerHeading">
                <div className="serve2ContainerHeadingName">
                  {lastServices[3]?.title ?? "Название услуги"}
                </div>
                <div className="serve2ContainerHeadingCost">
                  {lastServices[3]?.price
                    ? `${lastServices[3].price} ₽`
                    : "Цена не указана"}
                </div>
              </div>
              <img className="serve2car" src="thing.png" />
              <div className="serve2Content">
                {lastServices[3]?.description ?? "Описание услуги"}
              </div>
            </div>
            <button
              onClick={() =>
                lastServices[3] && navigate(`/services/${lastServices[3].id}`)
              }
              className="serve2Button"
            >
              Купить
            </button>
          </div>
        </div>
      </div>
    </div>
  );
}
