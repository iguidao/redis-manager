import axiosInstance from './index'

// 数据返回的接口
// 定义请求响应参数，不含data
interface Result {
    errorCode: number;
    msg: string
}

// 请求响应参数，包含data
interface ResultData<T = any> extends Result {
    data?: T;
}

class RequestHttp {
    get<T>(url: string, params?: object): Promise<ResultData<T>> {
        return axiosInstance.get(url, {params});
    }
    post<T>(url: string, params?: object): Promise<ResultData<T>> {
        return axiosInstance.post(url, params);
    }
    put<T>(url: string, params?: object): Promise<ResultData<T>> {
        return axiosInstance.put(url, params);
    }
    delete<T>(url: string, params?: object): Promise<ResultData<T>> {
        return axiosInstance.delete(url, {params});
    }
}
// 导出一个实例对象
export default RequestHttp