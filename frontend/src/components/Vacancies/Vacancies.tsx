import { Header } from "../Header/Header";
import "./Vacancies.css";
import { useNavigate } from "react-router-dom";
export function Vacancies() {
  const navigate=useNavigate()
  
  return (
    <>
      <button onClick={()=>navigate("/team")} className="vacanciesBackButton">Назад</button>
      <div className="team-page">
        <div className="container">
          <Header />
          <h1 className="teamHeading">Вакансии</h1>
          <div className="teamPeople">
            <div className="teamWithButton">
              <div className="teamPerson">
                <div className="contentTeamPerson">
                  <div className="contentTeamPersonPhoto"></div>
                  <div className="contentTeamPersonInfo">
                    <div className="contentTeamPersonInfoName">
                      Менеджер по подбору персонала
                    </div>
                    <div className="contentTeamPersonInfoTitle">
                      Опыт работы: 1 год
                    </div>
                    <div className="contentTeamPersonInfoStaz">
                      ЗП: от 70 000
                    </div>
                  </div>
                </div>
              </div>
              <button className="buttonTeamPersonClick">Откликнуться</button>
            </div>
            <div className="teamWithButton">
              <div className="teamPerson">
                <div className="contentTeamPerson">
                  <div className="contentTeamPersonPhoto"></div>
                  <div className="contentTeamPersonInfo">
                    <div className="contentTeamPersonInfoName">
                      Менеджер по подбору персонала
                    </div>
                    <div className="contentTeamPersonInfoTitle">
                      Опыт работы: 1 год
                    </div>
                    <div className="contentTeamPersonInfoStaz">
                      ЗП: от 70 000
                    </div>
                  </div>
                </div>
              </div>
              <button className="buttonTeamPersonClick">Откликнуться</button>
            </div>
            <div className="teamWithButton">
              <div className="teamPerson">
                <div className="contentTeamPerson">
                  <div className="contentTeamPersonPhoto"></div>
                  <div className="contentTeamPersonInfo">
                    <div className="contentTeamPersonInfoName">
                      Менеджер по подбору персонала
                    </div>
                    <div className="contentTeamPersonInfoTitle">
                      Опыт работы: 1 год
                    </div>
                    <div className="contentTeamPersonInfoStaz">
                      ЗП: от 70 000
                    </div>
                  </div>
                </div>
              </div>
              <button className="buttonTeamPersonClick">Откликнуться</button>
            </div>
          </div>
        </div>
      </div>
    </>
  );
}
