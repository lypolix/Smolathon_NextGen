import "./New.css";
import { Header } from "../Header/Header";
import { useNavigate, useParams } from "react-router-dom";
import PublicService from "../../backendConnection/publicInfo/publicInfoService";
import type { New1Response } from "../../types";
import { useState, useEffect } from "react";

export function New() {
  const [new1, setNew1] = useState<New1Response | null>(null);
  const [loading, setLoading] = useState(true);
  const { id } = useParams<{ id: string }>();
  const navigate = useNavigate();

  useEffect(() => {
    if (id) {
      const getNew = async () => {
        try {
          const result = await PublicService.getNewInfo(Number(id));
          setNew1(result ?? null);
        } catch (error) {
          console.error("Ошибка при загрузке новости:", error);
          setNew1(null);
        } finally {
          setLoading(false);
        }
      };
      getNew();
    }
  }, [id]);

  if (loading) {
    return (
      <div className="News">
        <div className="container">
          <Header
            menuBgColor="white"
            textColor="#203716"
            activeTextColor="white"
          />
          <p>Загрузка...</p>
        </div>
      </div>
    );
  }

  if (!new1) {
    return (
      <div className="News">
        <div className="container">
          <Header
            menuBgColor="white"
            textColor="#203716"
            activeTextColor="white"
          />
          <div className="contentNews">
            <div className="headingNews">
              <div className="headingNameNews">
                Новость не найдена или была удалена
              </div>
              <button
                onClick={() => navigate("/news")}
                className="headingButtonNews"
              >
                Все новости
              </button>
            </div>
          </div>
        </div>
      </div>
    );
  }

  return (
    <div className="News">
      <div className="container">
        <Header
          menuBgColor="white"
          textColor="#203716"
          activeTextColor="white"
        />
        <div className="contentNews">
          <div className="headingNews">
            <div className="headingNameNews">
              {new1.news?.title ?? "Заголовок отсутствует"}
            </div>
            <button
              onClick={() => navigate("/news")}
              className="headingButtonNews"
            >
              Все новости
            </button>
          </div>
          <div className="undHeadingNew">
            <div className="tagUndHeadingNew">
              {new1.news?.tag ?? "Тег отсутствует"}
            </div>
            <div className="dateUndHeadingNew">
              {new1.news?.date
                ? new Date(new1.news.date).toLocaleDateString("ru-RU")
                : "Дата отсутствует"}
            </div>
          </div>
          <div className="contentNew">
            {new1.news?.content ?? "Нет описания"}
          </div>
          <a className="linkNew">Отчет об установлении новых камер.pdf</a>
        </div>
      </div>
    </div>
  );
}
