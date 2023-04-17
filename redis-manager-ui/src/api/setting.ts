import RequestHttp from '../utils/request'

namespace Rconfiglist {
    // 拿到配置数据
    export interface SettingResData {
        data: {
            lists: [{
                ID: number;
                CreatedAt: string;
                UpdatedAt: string;
                DeletedAt: string;
                Key: string;
                Name: string;
                Value: string;
            }];
            total: number;
        }
        errorCode: number;
        msg: string;
    }
  }
// 配置获取
export const listCfg = () => {
    // 返回的数据格式可以和服务端约定
    return RequestHttp.get<Rconfiglist.SettingResData>('/cfg/v1/list');
}

namespace Rconfiglistdefault {
    // 拿到配置数据
    export interface SettingResData {
        data: [{
            label: string;
            value: string;
        }]
        errorCode: number;
        msg: string;
    }
  }
// 配置获取
export const listDefaultCfg = () => {
    // 返回的数据格式可以和服务端约定
    return RequestHttp.get<Rconfiglistdefault.SettingResData>('/cfg/v1/listdefault');
}

namespace Rconfigupdate  {
    export interface UpdateReqForm {
        key: string;
        value: string;
    }
    export interface SettingResData {
        data: [{
            label: string;
            value: string;
        }]
        errorCode: number;
        msg: string;
    }
}
// 删除
export const delCfg = (params: Rconfigupdate.UpdateReqForm) => {
    return RequestHttp.delete<Rconfigupdate.SettingResData>('/cfg/v1/del', {params: params});
}

//变更
export const updateCfg = (params: Rconfigupdate.UpdateReqForm) => {
    return RequestHttp.post<Rconfigupdate.SettingResData>('/cfg/v1/update', params);
}