import axios, { AxiosRequestConfig } from "axios";
import { API_BASE_URL } from "@/const";

const AUTH_KEY = "auth_token";

export const getAuthToken = () => localStorage.getItem(AUTH_KEY) || "";
export const setAuthToken = (token?: string) => {
  if (!token) {
    localStorage.removeItem(AUTH_KEY);
    return;
  }
  localStorage.setItem(AUTH_KEY, token);
};

export const apiClient = axios.create({
  baseURL: API_BASE_URL,
  timeout: 12000,
  headers: {
    "Content-Type": "application/json",
  },
});

apiClient.interceptors.request.use((config) => {
  const token = getAuthToken();
  if (token) {
    config.headers = config.headers || {};
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
});

export const fetcher = async <T>(url: string, params?: Record<string, unknown>): Promise<T> => {
  const response = await apiClient.get<T>(url, { params });
  return response.data;
};

export const apiGet = async <T>(url: string, config?: AxiosRequestConfig) => {
  const res = await apiClient.get<T>(url, config);
  return res.data;
};

export const apiPost = async <T>(url: string, data?: unknown, config?: AxiosRequestConfig) => {
  const res = await apiClient.post<T>(url, data, config);
  return res.data;
};

export const apiPut = async <T>(url: string, data?: unknown, config?: AxiosRequestConfig) => {
  const res = await apiClient.put<T>(url, data, config);
  return res.data;
};

export const apiDelete = async <T>(url: string, config?: AxiosRequestConfig) => {
  const res = await apiClient.delete<T>(url, config);
  return res.data;
};
