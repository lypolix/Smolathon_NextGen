import axios from "axios";

export const API_URL = `http://localhost:8080/api/auth`;

const $api = axios.create({
  withCredentials: true,
  baseURL: API_URL,
});

$api.interceptors.request.use((config) => {
  const token = localStorage.getItem("token");
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

$api.interceptors.response.use(
  (response) => response,
  (error) => {
    console.log("not authorized");
    return Promise.reject(error);
  }
);

export default $api;
