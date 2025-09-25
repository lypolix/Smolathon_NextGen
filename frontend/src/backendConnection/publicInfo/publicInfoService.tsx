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

export default class PublicService {
  static async getTeamInfo(): Promise<Team[] | undefined> {
    try {
      const response: AxiosResponse<Team[]> = await $api.get<Team[]>("/team");
      console.log(response.data);
      return response.data;
    } catch (error) {
      console.log("ошибка при получении команды");
      throw error;
    }
  }

  static async getNewsInfo(): Promise<News[] | undefined> {
    try {
      const response: AxiosResponse<News[]> = await $api.get<News[]>("/news");
      console.log(response.data);
      return response.data;
    } catch (error) {
      console.log("ошибка при получении новостей");
      throw error;
    }
  }

  static async getServicesInfo(): Promise<Services[] | undefined> {
    try {
      const response: AxiosResponse<Services[]> = await $api.get<Services[]>(
        "/services"
      );
      console.log(response.data);
      return response.data;
    } catch (error) {
      console.log("ошибка при получении услуг");
      throw error;
    }
  }

  static async getTrafficInfo(): Promise<Traffic[] | undefined> {
    try {
      const response: AxiosResponse<Traffic[]> = await $api.get<Traffic[]>(
        "/traffic"
      );
      console.log(response.data);
      return response.data;
    } catch (error) {
      console.log("ошибка при получении трафика");
      throw error;
    }
  }

  static async getStatisticsInfo(): Promise<Statistics[] | undefined> {
    try {
      const response: AxiosResponse<Statistics[]> = await $api.get<
        Statistics[]
      >("/stats");
      console.log(response.data);
      return response.data;
    } catch (error) {
      console.log("ошибка при получении статистики");
      throw error;
    }
  }

  static async getProjectsInfo(): Promise<Projects[] | undefined> {
    try {
      const response: AxiosResponse<Projects[]> = await $api.get<Projects[]>(
        "/projects"
      );
      console.log(response.data);
      return response.data;
    } catch (error) {
      console.log("ошибка при получении проектов");
      throw error;
    }
  }
}
