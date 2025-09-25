import { Header } from "../Header/Header";
import "./News.css";
import { useState, useEffect } from "react";
import PublicService from "../../backendConnection/publicInfo/publicInfoService";
import type { News as NewsItem } from "../../types";

export function News() {
  const [news, setNews] = useState<NewsItem[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    let mounted = true;
    const getAllNews = async () => {
      try {
        const result = await PublicService.getNewsInfo(); // ожидается News[]
        if (!mounted) return;
        setNews(result ?? []);
      } catch (e) {
        if (!mounted) return;
        setError("Ошибка загрузки новостей");
      } finally {
        if (!mounted) return;
        setLoading(false);
      }
    };
    getAllNews();
    return () => {
      mounted = false;
    };
  }, []);

  const formatDate = (iso?: string) => {
    if (!iso) return "";
    try {
      const d = new Date(iso);
      // Формат DD.MM.YYYY
      const dd = String(d.getDate()).padStart(2, "0");
      const mm = String(d.getMonth() + 1).padStart(2, "0");
      const yyyy = d.getFullYear();
      return `${dd}.${mm}.${yyyy}`;
    } catch {
      return iso;
    }
  };

  const latest = news.slice(-4).reverse(); // последние 4, от новых к старым

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
            <div className="headingNameNews">Актуальные новости</div>
            <button className="headingButtonNews">Все новости</button>
          </div>

          {loading && <div className="noNewsMessage">Загрузка...</div>}
          {error && !loading && (
            <div className="noNewsMessage">{error}</div>
          )}

          {!loading && !error && (
            <div className="blocksNews">
              {latest.length > 0 ? (
                latest.map((item) => (
                  <div className="blockNews1" key={item.id}>
                    <div className="headingBlockNews">
                      <div className="heading1BlockNews">{item.tag}</div>
                      <div className="headingDateNews">{formatDate(item.date)}</div>
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
          )}
        </div>
      </div>
    </div>
  );
}
