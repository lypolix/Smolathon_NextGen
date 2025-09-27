import $api from "../http";
import type { AxiosResponse } from "axios";
import type { AuthResponse } from "../model/AuthResponse";

export default class AuthService {
  static async loginAdmin(email: string, password: string): Promise<AxiosResponse<AuthResponse>> {
    return $api.post<AuthResponse>("/admin/login", { email, password });
  }

  static async loginEditor(email: string, password: string): Promise<AxiosResponse<AuthResponse>> {
    return $api.post<AuthResponse>("/editor/login", { email, password });
  }

  static async logout(): Promise<void> {
    return $api.post("/logout");
  }
}
