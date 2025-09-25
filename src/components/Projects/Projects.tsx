import { Header } from "../Header/Header";
import { useState, useEffect } from "react";
import PublicService from "../../backendConnection/publicInfo/publicInfoService";
import type { Projects } from "../../types";
export function Projects() {
  const [projects, setProjects] = useState<Projects[] | undefined>(undefined);
  useEffect(() => {
    const getAllProjects = async () => {
      const result = await PublicService.getProjectsInfo();
      setProjects(result);
      console.log(result);
    };
    getAllProjects();
  }, []);
  return (
    <>
      <Header />
      <h1>Here will be our projects</h1>
    </>
  );
}
