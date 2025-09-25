import axios from "axios";
import type {
  Team,
  News,
  Services,
  Projects,
  Statistics,
  Traffic,
} from "../../types";
import type { AxiosResponse } from "axios";

export const API_PUBLIC = `http://localhost:8080/api`;

const $api = axios.create({
  withCredentials: true,
  baseURL: API_PUBLIC,
});

// Описываем ожидаемые “контейнеры” ответов от бэка
type NewsResp = { news: News[] };
type TeamResp = { team: Team[] };
type ServicesResp = { services: Services[] };
type ProjectsResp = { projects: Projects[] };
type TrafficResp = { traffic: Traffic }; // судя по логам, объект, не массив
type StatsResp = { stats: Statistics[] } | { statistics: Statistics[] }; // на случай разных ключей

export default class PublicService {
  static async getTeamInfo(): Promise<Team[]> {
    try {
      const response: AxiosResponse<TeamResp> = await $api.get<TeamResp>("/team");
      console.log(response.data);
      return response.data.team ?? [];
    } catch (error) {
      console.log("ошибка при получении команды");
      throw error;
    }
  }

  static async getNewsInfo(): Promise<News[]> {
    try {
      // Бэкенд возвращает { news: [...] }, пример: es={"news":[{...}, ...]}
      const response: AxiosResponse<NewsResp> = await $api.get<NewsResp>("/news");
      console.log(response.data);
      return response.data.news ?? [];
    } catch (error) {
      console.log("ошибка при получении новостей");
      throw error;
    }
  }

  static async getServicesInfo(): Promise<Services[]> {
    try {
      const response: AxiosResponse<ServicesResp> = await $api.get<ServicesResp>("/services");
      console.log(response.data);
      return response.data.services ?? [];
    } catch (error) {
      console.log("ошибка при получении услуг");
      throw error;
    }
  }

  static async getTrafficInfo(): Promise<Traffic> {
    try {
      const response: AxiosResponse<TrafficResp> = await $api.get<TrafficResp>("/traffic");
      console.log(response.data);
      return response.data.traffic;
    } catch (error) {
      console.log("ошибка при получении трафика");
      throw error;
    }
  }

  static async getStatisticsInfo(): Promise<Statistics[]> {
    try {
      // Поддержим оба варианта ключа (stats | statistics), чтобы не упасть при расхождении
      const response: AxiosResponse<StatsResp> = await $api.get<StatsResp>("/stats");
      console.log(response.data);
      const anyData = response.data as any;
      return anyData.stats ?? anyData.statistics ?? [];
    } catch (error) {
      console.log("ошибка при получении статистики");
      throw error;
    }
  }

  static async getProjectsInfo(): Promise<Projects[]> {
    try {
      const response: AxiosResponse<ProjectsResp> = await $api.get<ProjectsResp>("/projects");
      console.log(response.data);
      return response.data.projects ?? [];
    } catch (error) {
      console.log("ошибка при получении проектов");
      throw error;
    }
  }
}
