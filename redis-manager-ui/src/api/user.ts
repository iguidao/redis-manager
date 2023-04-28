import RequestHttp from '../utils/request'

namespace Ruser {
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
            usertype: string;
        }
        errorCode: number;
        msg: string;
    }
    export interface PasswordReqForm {
        old: string;
        new: string;
      }
    export interface PasswordResData {
        data: {
            result: string;
        }
        errorCode: number;
        msg: string;
    }
    export interface listResData {
        data: [{
            ID: number;
            CreatedAt: string;
            UpdatedAt: string;
            UserName: string;
            Email: string;
            UserType: string;
        }]
        errorCode: number;
        msg: string;
    }
    export interface typeResData {
        data: [{
            label: string;
            value: string;
        }]
        errorCode: number;
        msg: string;
    }
    export interface createReqForm {
        username: string;
        password: string;
        mail: string;
      }
    export interface createResData {
        data: {
            result: string;
        }
        errorCode: number;
        msg: string;
    }
    export interface changeReqForm {
        username: string;
        usertype: string;
      }
    export interface changeResData {
        data: {
            result: string;
        }
        errorCode: number;
        msg: string;
    }
    export interface delReqForm {
        username: string;
      }
    export interface delResData {
        data: {
            result: string;
        }
        errorCode: number;
        msg: string;
    }
}
  // 用户登录
export const login = (params: Ruser.LoginReqForm) => {
    // 返回的数据格式可以和服务端约定
    return RequestHttp.post<Ruser.LoginResData>('/auth/v1/sign-in', params);
}

export const userpassword = (params: Ruser.PasswordReqForm) => {
    // 返回的数据格式可以和服务端约定
    return RequestHttp.post<Ruser.PasswordResData>('/auth/v1/password', params);
}

export const userlist = () => {
    // 返回的数据格式可以和服务端约定
    return RequestHttp.get<Ruser.listResData>('/user/v1/list');
}

export const usertypelist = () => {
    // 返回的数据格式可以和服务端约定
    return RequestHttp.get<Ruser.typeResData>('/user/v1/utype');
}

export const useradd = (params: Ruser.createReqForm) => {
    // 返回的数据格式可以和服务端约定
    return RequestHttp.post<Ruser.createResData>('/user/v1/add', params);
}

export const userchange = (params: Ruser.changeReqForm) => {
    // 返回的数据格式可以和服务端约定
    return RequestHttp.post<Ruser.changeResData>('/user/v1/change', params);
}

export const userdel = (params: Ruser.delReqForm) => {
    // 返回的数据格式可以和服务端约定
    return RequestHttp.delete<Ruser.delResData>('/user/v1/del', {params: params});
}

