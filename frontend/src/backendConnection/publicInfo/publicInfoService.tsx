import axios from "axios";
import type {
  New1Response,
  NewsResponse,
  Service1Response,
  ServicesResponse,
  Team1Response,
  TeamResponse,
  Project1Response,
  ProjectsResponse,
  StatisticsResponse,
  TrafficLightsResponse,
  FinesResponse,
  EvacuationsResponse,
  EvacuationRoutesResponse,
} from "../../types";
import type { AxiosResponse } from "axios";

export const API_PUBLIC = `http://localhost:8080/api`;
const $api = axios.create({
  withCredentials: true,
  baseURL: API_PUBLIC,
});

export default class PublicService {
  static async getTeamInfo(): Promise<TeamResponse | undefined> {
    try {
      const response: AxiosResponse<TeamResponse> = await $api.get<TeamResponse>("/team");
      console.log(response.data);
      return response.data;
    } catch (error) {
      console.log("ошибка при получении команды");
      throw error;
    }
  }

  static async getNewsInfo(): Promise<NewsResponse | undefined> {
    try {
      const response: AxiosResponse<NewsResponse> = await $api.get<NewsResponse>("/news");
      console.log(response.data);
      return response.data;
    } catch (error) {
      console.log("ошибка при получении новостей");
      throw error;
    }
  }

  static async getServicesInfo(): Promise<ServicesResponse | undefined> {
    try {
      const response: AxiosResponse<ServicesResponse> = await $api.get<ServicesResponse>(
        "/services"
      );
      console.log(response.data);
      return response.data;
    } catch (error) {
      console.log("ошибка при получении услуг");
      throw error;
    }
  }

  static async getTrafficInfo(): Promise<TrafficLightsResponse | undefined> {
    try {
      const response: AxiosResponse<TrafficLightsResponse> = await $api.get<TrafficLightsResponse>(
        "/traffic"
      );
      console.log(response.data);
      return response.data;
    } catch (error) {
      console.log("ошибка при получении трафика");
      throw error;
    }
  }

  static async getStatisticsInfo(): Promise<StatisticsResponse | undefined> {
    try {
      const response: AxiosResponse<StatisticsResponse> = await $api.get<
        StatisticsResponse
      >("/stats");
      console.log(response.data);
      return response.data;
    } catch (error) {
      console.log("ошибка при получении статистики");
      throw error;
    }
  }

  static async getProjectsInfo(): Promise<ProjectsResponse | undefined> {
    try {
      const response: AxiosResponse<ProjectsResponse> = await $api.get<ProjectsResponse>(
        "/projects"
      );
      console.log(response.data);
      return response.data;
    } catch (error) {
      console.log("ошибка при получении проектов");
      throw error;
    }
  }
  static async getNewInfo(itemId: number): Promise<New1Response | undefined> {
    try {
      const response: AxiosResponse<New1Response> = await $api.get<New1Response>(
        `/news/${itemId}`
      );
      console.log(response.data);
      return response.data;
    } catch (error) {
      console.log("ошибка при получении новости");
      throw error;
    }
  }
  static async getServiceInfo(itemId: number): Promise<Service1Response | undefined> {
    try {
      const response: AxiosResponse<Service1Response> = await $api.get<Service1Response>(
        `/services/${itemId}`
      );
      console.log(response.data);
      return response.data;
    } catch (error) {
      console.log("ошибка при получении услуги");
      throw error;
    }
  }
  static async getTeam1Info(itemId: number): Promise<Team1Response | undefined> {
    try {
      const response: AxiosResponse<Team1Response> = await $api.get<Team1Response>(
        `/team/${itemId}`
      );
      console.log(response.data);
      return response.data;
    } catch (error) {
      console.log("ошибка при получении одного человека");
      throw error;
    }
  }
  static async getProject1Info(itemId: number): Promise<Project1Response | undefined> {
    try {
      const response: AxiosResponse<Project1Response> = await $api.get<Project1Response>(
        `/projects/${itemId}`
      );
      console.log(response.data);
      return response.data;
    } catch (error) {
      console.log("ошибка при получении одного проекта");
      throw error;
    }
  }
  static async getTrafficLightsInfo(): Promise<TrafficLightsResponse | undefined> {
    try {
      const response: AxiosResponse<TrafficLightsResponse> = await $api.get<
        TrafficLightsResponse
      >("/traffic-lights");
      console.log(response.data);
      return response.data;
    } catch (error) {
      console.log("ошибка при получении светофоров");
      throw error;
    }
  }
  static async getFinesInfo(): Promise<FinesResponse | undefined> {
    try {
      const response: AxiosResponse<FinesResponse> = await $api.get<
        FinesResponse
      >("/fines");
      console.log(response.data);
      return response.data;
    } catch (error) {
      console.log("ошибка при получении штрафов");
      throw error;
    }
  }
  static async getEvacuationsInfo(): Promise<EvacuationsResponse | undefined> {
    try {
      const response: AxiosResponse<EvacuationsResponse> = await $api.get<
        EvacuationsResponse
      >("/evacuations");
      console.log(response.data);
      return response.data;
    } catch (error) {
      console.log("ошибка при получении эвакуаций");
      throw error;
    }
  }
  static async getEvacuationRoutesInfo(): Promise<EvacuationRoutesResponse | undefined> {
    try {
      const response: AxiosResponse<EvacuationRoutesResponse> = await $api.get<
        EvacuationRoutesResponse
      >("/evacuations-routes");
      console.log(response.data);
      return response.data;
    } catch (error) {
      console.log("ошибка при получении эвакуационных маршрутов");
      throw error;
    }
  }
}
