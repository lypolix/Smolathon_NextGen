import "./Services.css";
import { Header } from "../Header/Header";
import { useParams } from "react-router-dom";
import PublicService from "../../backendConnection/publicInfo/publicInfoService";
import type { Service1Response } from "../../types";
import { useEffect, useState } from "react";

export function Service() {
  const [service1, setService1] = useState<Service1Response | null>(null);
  const [loading, setLoading] = useState(true);
  const { id } = useParams<{ id: string }>();

  useEffect(() => {
    if (id) {
      const getService = async () => {
        try {
          const result = await PublicService.getServiceInfo(Number(id));
          setService1(result ?? null);
        } catch (error) {
          console.error("Ошибка при загрузке услуги:", error);
          setService1(null);
        } finally {
          setLoading(false);
        }
      };
      getService();
    }
  }, [id]);

  if (loading) {
    return <div className="container">Загрузка...</div>;
  }

  if (!service1) {
    return <div className="container">Услуга не найдена</div>;
  }

  const { service } = service1; 

  return (
    <div className="serve-page">
      <div className="container">
        <Header
          menuBgColor="white"
          textColor="#203716"
          activeTextColor="white"
        />
        <h1 className="oneServeHeading">
          <span className="highlight1">Услуга</span>
          <br />
          {service.title}
        </h1>
        <div className="oneServePrice">
          <div className="oneServePriceName">Цена</div>
          <div className="oneServePricePrice">{service.price} ₽</div>
        </div>
        <div className="oneServeContent">
          <span className="highlight3">{service.title}</span>
          <br />
          {service.description}
        </div>
        <div className="oneServeHeading2">
          Оставить <span className="highlight4">заявку</span>
        </div>
        <div className="oneServeForm">
          <div className="oneServeFormBlock">
            <div className="oneServeFormBlockHeading">ФИО</div>
            <input
              placeholder="Иванов Иван Алексеевич"
              className="oneServeFormBlockInput"
            />
          </div>
          <div className="oneServeFormBlock">
            <div className="oneServeFormBlockHeading">Номер телефона</div>
            <input
              placeholder="+7 (950) 737 18-22"
              className="oneServeFormBlockInput"
            />
          </div>
          <div className="oneServeFormBlock">
            <div className="oneServeFormBlockHeading">Email</div>
            <input
              placeholder="hacksmolathon@gmail.com"
              className="oneServeFormBlockInput"
            />
          </div>
        </div>
        <button className="oneSeveFormButton">Подтвердить</button>
      </div>
    </div>
  );
}
