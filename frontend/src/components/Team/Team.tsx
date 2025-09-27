import { Header } from "../Header/Header";
import "./Team.css";
import { useState, useEffect } from "react";
import PublicService from "../../backendConnection/publicInfo/publicInfoService";
import type { TeamResponse, Team } from "../../types";
import { useNavigate } from "react-router-dom";

export function Team() {
  const [team, setTeam] = useState<TeamResponse | undefined>(undefined);
  const navigate = useNavigate();

  useEffect(() => {
    const getAllTeam = async () => {
      const result = await PublicService.getTeamInfo();
      setTeam(result);
      console.log(result);
    };
    getAllTeam();
  }, []);

  // Берём последние 3 человека
  const lastThreeTeamMembers: Team[] = team?.team.slice(-3).reverse() ?? [];

  return (
    <>
      <button className="TeamAllPeopleWorkingButton">Все сотрудники</button>
      <button onClick={() => navigate("/vacancies")} className="TeamAllPeopleWorkingButton2">
        Вакансии
      </button>
      <div className="team-page">
        <div className="container">
          <Header />
          <div className="teamHeading1">Команда</div>
          <p className="teamDescribe">
            Мы стараемся, чтобы на дорогах Смоленска Вам было приятно ездить
          </p>
          <div className="teamPeople">
            {lastThreeTeamMembers.length > 0 ? (
              lastThreeTeamMembers.map((member) => (
                <div className="teamPerson" key={member.id}>
                  <div className="contentTeamPerson">
                    <div
                      className="contentTeamPersonPhoto"
                      style={{
                        backgroundImage: `url(${member.photo_url})`,
                        backgroundSize: "cover",
                        backgroundPosition: "center",
                      }}
                    ></div>
                    <div className="contentTeamPersonInfo">
                      <div className="contentTeamPersonInfoName">{member.name}</div>
                      <div className="contentTeamPersonInfoTitle">{member.position}</div>
                      <div className="contentTeamPersonInfoStaz">{member.experience}</div>
                    </div>
                  </div>
                </div>
              ))
            ) : (
              <div>Команда пока не загружена</div>
            )}
          </div>
        </div>
      </div>
    </>
  );
}
