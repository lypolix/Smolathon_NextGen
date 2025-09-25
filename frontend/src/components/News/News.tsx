import { Header } from "../Header/Header";
import "./News.css";
import { useState, useEffect } from "react";
import PublicService from "../../backendConnection/publicInfo/publicInfoService";
import type { News } from "../../types";
export function News() {
  const [news, setNews] = useState<News[] | undefined>(undefined);
  useEffect(() => {
    const getAllNews = async () => {
      const result = await PublicService.getNewsInfo();
      setNews(result);
      console.log(result);
    };
    getAllNews();
  }, []);

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
              <button className="headingButtonNews">Все новости</button>
            </div>
            <div className="blocksNews">
              {news && news.length > 0 ? (
                news.slice(-4).map((item) => (
                  <div className="blockNews1" key={item.id}>
                    <div className="headingBlockNews">
                      <div className="heading1BlockNews">{item.tag}</div>
                      <div className="headingDateNews">{item.date}</div>
                    </div>
                    <div className="contentBlockNews">
                      <div className="contentBlockNews1">
                        <h2 className="headingContentBlockNews">
                          {item.title}
                        </h2>
                        <p className="textContentBlockNews">{item.content}</p>
                      </div>
                      <button className="contentBlockNewsButton">
                        Подробнее
                      </button>
                    </div>
                  </div>
                ))
              ) : (
                <div className="noNewsMessage">Новостей нет</div>
              )}
            </div>
          </div>
        </div>
      </div>
    </>
  );
}
