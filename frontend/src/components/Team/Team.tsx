import { Header } from "../Header/Header";
import "./Team.css";
import { useState, useEffect } from "react";
import PublicService from "../../backendConnection/publicInfo/publicInfoService";
import type { Team } from "../../types";

// Файл лежит в public: доступен по /ava.webp
const DEFAULT_AVATAR = "/ava.webp";

export function Team() {
  const [team, setTeam] = useState<Team[]>([]);
  const [error, setError] = useState<string | null>(null);
  const [loading, setLoading] = useState<boolean>(true);

  useEffect(() => {
    const getAllTeam = async () => {
      try {
        setLoading(true);
        const result = await PublicService.getTeamInfo(); // ожидаем Team[]
        // Если PublicService иногда вернет { team: [...] }, подстрахуемся:
        const normalized: Team[] = Array.isArray(result)
          ? result
          : Array.isArray((result as any)?.team)
          ? (result as any).team
          : [];
        setTeam(normalized);
      } catch (e: any) {
        setError(e?.message || "Ошибка загрузки команды");
        setTeam([]);
      } finally {
        setLoading(false);
      }
    };
    getAllTeam();
  }, []);

  return (
    <div className="team-page">
      <div className="container">
        <Header />
        <h1 className="teamHeading">Команда</h1>
        <p className="teamDescribe">
          Мы стараемся, чтобы на дорогах Смоленска Вам было приятно ездить
        </p>

        {loading && <div>Загрузка...</div>}
        {error && <div className="error">{error}</div>}

        <div className="teamPeople">
          {team.map((person) => {
            const bgUrl =
              person.photo_url && person.photo_url.trim() !== ""
                ? person.photo_url
                : DEFAULT_AVATAR;
            return (
              <div className="teamPerson" key={person.id}>
                <div className="contentTeamPerson">
                  <div
                    className="contentTeamPersonPhoto"
                    style={{ backgroundImage: `url(${bgUrl})` }}
                    aria-label={person.name}
                  />
                  <div className="contentTeamPersonInfo">
                    <div className="contentTeamPersonInfoName">{person.name}</div>
                    <div className="contentTeamPersonInfoTitle">{person.position}</div>
                    <div className="contentTeamPersonInfoStaz">{person.experience}</div>
                  </div>
                </div>
              </div>
            );
          })}

          {!loading && !error && team.length === 0 && (
            <div>Состав команды пока пуст.</div>
          )}
        </div>
      </div>
    </div>
  );
}
