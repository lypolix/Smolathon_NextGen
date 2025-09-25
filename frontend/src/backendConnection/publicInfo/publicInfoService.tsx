import axios, { type AxiosResponse } from "axios";
import type {
  Team,
  News,
  Services,
  Projects,
  Statistics,
  Traffic
} from "../../types";

export const API_PUBLIC = "http://localhost:8080/api";

const $api = axios.create({
  baseURL: API_PUBLIC,
  withCredentials: false
});

// Точные типы обёрток ответов бэкенда
type NewsResp = { news: News[] };
type TeamResp = { team: Team[] };
type ServicesResp = { services: Services[] };
type ProjectsResp = { projects: Projects[] };
type TrafficResp = { traffic: Traffic };            // объект
type StatsResp = { stats: Statistics };             // объект

export default class PublicService {
  static async getTeamInfo(): Promise<Team[]> {
    const res: AxiosResponse<TeamResp> = await $api.get("/team");
    return res.data.team;
  }

  static async getNewsInfo(): Promise<News[]> {
    const res: AxiosResponse<NewsResp> = await $api.get("/news");
    return res.data.news;
  }

  static async getServicesInfo(): Promise<Services[]> {
    const res: AxiosResponse<ServicesResp> = await $api.get("/services");
    return res.data.services;
  }

  static async getProjectsInfo(): Promise<Projects[]> {
    const res: AxiosResponse<ProjectsResp> = await $api.get("/projects");
    return res.data.projects;
  }

  static async getTrafficInfo(): Promise<Traffic> {
    const res: AxiosResponse<TrafficResp> = await $api.get("/traffic");
    return res.data.traffic;
  }

  static async getStatisticsInfo(): Promise<Statistics> {
    const res: AxiosResponse<StatsResp> = await $api.get("/stats");
    return res.data.stats;
  }
}
