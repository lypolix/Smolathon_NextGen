import axios from "axios";
import type { AuthResponse } from "../model/AuthResponse";
export const API_URL = `http://localhost:5000/auth`;
const $api=axios.create({
    withCredentials: true,
    baseURL: API_URL
});

$api.interceptors.request.use((config)=>{
    config.headers.Authorization = `Bearer ${localStorage.getItem("token")}`;
    return config;
})
$api.interceptors.response.use(
    (config)=>{
        return config;
    },
    async (error)=>{
        console.log("not authorized")
        const originalRequest= error.config;
        originalRequest._isRetry=true;
        if (error.response.status == 401 && error.config && !error.config._isRetry) {
      try {
        const response = await axios.get<AuthResponse>(`${API_URL}/refresh`, {
          withCredentials: true,
        });
        localStorage.setItem("token", response.data.accessToken);
        return $api.request(originalRequest);
      } catch (error) {
        console.log('НЕ АВТОРИЗОВАН')
      }
    }
    throw error
    }
)
export default $api;
