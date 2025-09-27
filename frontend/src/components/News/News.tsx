import { Header } from "../Header/Header";
import "./News.css";
import { useState, useEffect } from "react";
import PublicService from "../../backendConnection/publicInfo/publicInfoService";
import type { NewsResponse } from "../../types";
import { useNavigate } from "react-router-dom";

export function News() {
  const navigate = useNavigate();
  const [newsResp, setNewsResp] = useState<NewsResponse | undefined>(undefined);

  useEffect(() => {
    const getAllNews = async () => {
      const result = await PublicService.getNewsInfo(); 
      setNewsResp(result);
      console.log(result);
    };
    getAllNews();
  }, []);

  const items = newsResp?.news ?? []; 

  return (
    <>
      <div className="News">
        <div className="container">
          <Header
            menuBgColor="white"
            textColor="#203716"
            activeTextColor="white"
          />
          <div className="contentNews">
            <div className="headingNews">
              <div className="headingNameNews">Актуальные новости</div>
              <button
                onClick={() => navigate("/news")}
                className="headingButtonNews"
              >
                Все новости
              </button>
            </div>

            <div className="blocksNews">
              {items.length ? (
                items
                  .slice(-4)      
                  .reverse()      
                  .map((item) => (
                    <div className="blockNews1" key={item.id}>
                      <div className="headingBlockNews">
                        <div className="heading1BlockNews">{item.tag}</div>
                        <div className="headingDateNews">
                          {item.date ? new Date(item.date).toLocaleDateString("ru-RU") : ""}
                        </div>
                      </div>
                      <div className="contentBlockNews">
                        <div className="contentBlockNews1">
                          <h2 className="headingContentBlockNews">{item.title}</h2>
                          <p className="textContentBlockNews">
                            {item.content.length > 200 ? item.content.slice(0, 200) + "…" : item.content}
                          </p>
                        </div>
                        <button
                          onClick={() => navigate(`/news/${item.id}`)}
                          className="contentBlockNewsButton"
                        >
                          Подробнее
                        </button>
                      </div>
                    </div>
                  ))
              ) : (
                <div className="headingBlockNews">Новостей нет</div>
              )}
            </div>
          </div>
        </div>
      </div>
    </>
  );
}
