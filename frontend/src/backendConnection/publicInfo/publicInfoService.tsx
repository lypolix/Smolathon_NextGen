import axios from "axios";
import type {
  Team,
  News,
  Services,
  Projects,
  Traffic,
} from "../../types";
import type { AxiosResponse } from "axios";

export const API_PUBLIC = `http://localhost:8080/api`;

const $api = axios.create({
  withCredentials: true,
  baseURL: API_PUBLIC,
});

// Контейнеры ответов от бэка
type NewsResp = { news: News[] };
type TeamResp = { team: Team[] };
type ServicesResp = { services: Services[] };
type ProjectsResp = { projects: Projects[] };
type TrafficResp = { traffic: Traffic }; // по логам — объект

// Точный тип под stats из бэкенда
export type StatsPayload = {
  collected_amount_total: number;
  evacuations_count: number;
  evacuators_count: number;
  fine_lot_income: number;
  fines_amount_total: number;
  orders_total: number;
  traffic_lights_active: number;
  trips_count: number;
  violations_total: number;
};
type StatsResp = { stats: StatsPayload } | { statistics: StatsPayload } | StatsPayload;

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

  static async getStatisticsInfo(): Promise<StatsPayload> {
    try {
      // Бэкенд присылает:
      // { "stats": { collected_amount_total: ..., ... } }
      // Поддержим также { "statistics": {...} } или прямой объект на случай изменений
      const response: AxiosResponse<StatsResp> = await $api.get<StatsResp>("/stats");
      console.log(response.data);
      const data: any = response.data;
      const payload: StatsPayload =
        data?.stats ?? data?.statistics ?? data;
      return payload as StatsPayload;
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
