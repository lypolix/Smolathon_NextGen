import { Header } from "../Header/Header";
import "./Team.css";
import { useState, useEffect } from "react";
import PublicService from "../../backendConnection/publicInfo/publicInfoService";
import type { Team } from "../../types";
export function Team() {
  const [team, setTeam] = useState<Team[] | undefined>(undefined);
  useEffect(() => {
    const getAllTeam = async () => {
      const result = await PublicService.getTeamInfo();
      setTeam(result);
      console.log(result);
    };
    getAllTeam();
  }, []);
  return (
    <>
      <div className="team-page">
        <div className="container">
          <Header />
          <h1 className="teamHeading">Команда</h1>
          <p className="teamDescribe">
            Мы стараемся, чтобы на дорогах Смоленска Вам было приятно ездить
          </p>
          <div className="teamPeople">
            <div className="teamPerson">
              <div className="contentTeamPerson">
                <div className="contentTeamPersonPhoto"></div>
                <div className="contentTeamPersonInfo">
                  <div className="contentTeamPersonInfoName">Иван Иванов</div>
                  <div className="contentTeamPersonInfoTitle">Должность</div>
                  <div className="contentTeamPersonInfoStaz">Стаж</div>
                </div>
              </div>
            </div>
            <div className="teamPerson">
              <div className="contentTeamPerson">
                <div className="contentTeamPersonPhoto"></div>
                <div className="contentTeamPersonInfo">
                  <div className="contentTeamPersonInfoName">Иван Иванов</div>
                  <div className="contentTeamPersonInfoTitle">Должность</div>
                  <div className="contentTeamPersonInfoStaz">Стаж</div>
                </div>
              </div>
            </div>
            <div className="teamPerson">
              <div className="contentTeamPerson">
                <div className="contentTeamPersonPhoto"></div>
                <div className="contentTeamPersonInfo">
                  <div className="contentTeamPersonInfoName">Иван Иванов</div>
                  <div className="contentTeamPersonInfoTitle">Должность</div>
                  <div className="contentTeamPersonInfoStaz">Стаж</div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </>
  );
}
