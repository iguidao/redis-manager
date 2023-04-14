import axios, { AxiosRequestConfig, AxiosRequestHeaders } from "axios";
import {ElMessage} from 'element-plus'


interface AdaptAxiosRequestConfig extends AxiosRequestConfig {
  headers: AxiosRequestHeaders
}

// config
const axiosInstance = axios.create({
    baseURL: "/api",
    timeout: 30000
});

// Interceptors
axiosInstance.interceptors.request.use(
    (config): AdaptAxiosRequestConfig => {
        let token = sessionStorage.getItem('token')
        if (token) {
          config.headers['Authorization'] = token
        }
        return config;
    },
    (error): any => {
        return Promise.reject(error);
    }
);

axiosInstance.interceptors.response.use(
    async (response): Promise<any> => {
        return response;
    },
    async (error): Promise<any> => {
        return Promise.reject(error);
    }
);

// 4.导出 axios 实例
export default axiosInstance;