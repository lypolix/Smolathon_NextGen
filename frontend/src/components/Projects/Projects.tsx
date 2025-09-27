import { Header } from "../Header/Header";
import { useState, useEffect } from "react";
import PublicService from "../../backendConnection/publicInfo/publicInfoService";
import type { ProjectsResponse } from "../../types";
import "./Projects.css";

export function Projects() {
  const [projects, setProjects] = useState<ProjectsResponse | undefined>(
    undefined
  );

  useEffect(() => {
    const getAllProjects = async () => {
      const result = await PublicService.getProjectsInfo();
      setProjects(result);
      console.log(result);
    };
    getAllProjects();
  }, []);

  // берём последние 4 проекта из ответа
  const lastProjects = projects?.projects.slice(-4) ?? [];

  return (
    <>
      <div className="serve-page">
        <div className="container">
          <Header
            menuBgColor="white"
            textColor="#203716"
            activeTextColor="white"
          />
          <h1 className="projectsHeading">Проекты</h1>

          <div className="project1">
            <h2 className="project1Heading">
              {lastProjects[0]?.title ?? "Название проекта"}
            </h2>
            <p className="project1Content">
              {lastProjects[0]?.description ??
                "Описание недоступно или проект не найден."}
            </p>
          </div>

          <div className="projectsInLine">
            <div className="project2">
              <h3 className="project2Heading">
                {lastProjects[1]?.title ?? "Название проекта"}
              </h3>
              <p className="project2Content">
                {lastProjects[1]?.description ??
                  "Описание недоступно или проект не найден."}
              </p>
            </div>
            <div className="project3">
              <h3 className="project2Heading">
                {lastProjects[2]?.title ?? "Название проекта"}
              </h3>
              <p className="project2Content">
                {lastProjects[2]?.description ??
                  "Описание недоступно или проект не найден."}
              </p>
            </div>
            <div className="project4">
              <h3 className="project2Heading">
                {lastProjects[3]?.title ?? "Название проекта"}
              </h3>
              <p className="project2Content">
                {lastProjects[3]?.description ??
                  "Описание недоступно или проект не найден."}
              </p>
            </div>
          </div>
        </div>
      </div>
    </>
  );
}
