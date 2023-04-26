import RequestHttp from '../utils/request'

namespace Rlogin {
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
  }
  // 用户登录
export const login = (params: Rlogin.LoginReqForm) => {
    // 返回的数据格式可以和服务端约定
    return RequestHttp.post<Rlogin.LoginResData>('/auth/v1/sign-in', params);
}
namespace RUserPassword {
    export interface ReqForm {
        old: string;
        new: string;
      }
    export interface ResData {
        data: {
            result: string;
        }
        errorCode: number;
        msg: string;
    }
}
export const userpassword = (params: RUserPassword.ReqForm) => {
    // 返回的数据格式可以和服务端约定
    return RequestHttp.post<RUserPassword.ResData>('/auth/v1/password', params);
}

namespace RUserList {
    export interface ResData {
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
}
export const userlist = () => {
    // 返回的数据格式可以和服务端约定
    return RequestHttp.get<RUserList.ResData>('/user/v1/list');
}

namespace RUserType {
    export interface ResData {
        data: [{
            label: string;
            value: string;
        }]
        errorCode: number;
        msg: string;
    }
}
export const usertypelist = () => {
    // 返回的数据格式可以和服务端约定
    return RequestHttp.get<RUserType.ResData>('/user/v1/utype');
}

namespace RUserCreate {
    export interface ReqForm {
        username: string;
        password: string;
        mail: string;
      }
    export interface ResData {
        data: {
            result: string;
        }
        errorCode: number;
        msg: string;
    }
}
export const useradd = (params: RUserCreate.ReqForm) => {
    // 返回的数据格式可以和服务端约定
    return RequestHttp.post<RUserCreate.ResData>('/user/v1/add', params);
}

namespace RUserChange {
    export interface ReqForm {
        username: string;
        usertype: string;
      }
    export interface ResData {
        data: {
            result: string;
        }
        errorCode: number;
        msg: string;
    }
}
export const userchange = (params: RUserChange.ReqForm) => {
    // 返回的数据格式可以和服务端约定
    return RequestHttp.post<RUserChange.ResData>('/user/v1/change', params);
}


namespace RUserChange {
    export interface ReqForm {
        username: string;
      }
    export interface ResData {
        data: {
            result: string;
        }
        errorCode: number;
        msg: string;
    }
}
export const userdel = (params: RUserChange.ReqForm) => {
    // 返回的数据格式可以和服务端约定
    return RequestHttp.delete<RUserChange.ResData>('/user/v1/del', {params: params});
}

