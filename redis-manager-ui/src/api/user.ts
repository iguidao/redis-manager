import RequestHttp from '../utils/request'

namespace Login {
    // 用户登录表单
    export interface LoginReqForm {
      username: string;
      password: string;
    }
    // 登录成功后返回的token
    export interface LoginResData {
        data: {
            result: string;
            token: string;
            username: string;
        }
        errorCode: number;
        msg: string;
    }
  }
  // 用户登录
  export const login = (params: Login.LoginReqForm) => {
      // 返回的数据格式可以和服务端约定
      return RequestHttp.post<Login.LoginResData>('/user/v1/sign-in', params);
  }
  