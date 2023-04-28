import RequestHttp from '../utils/request'

namespace Rconfig {
    // 拿到配置数据
    export interface ListResData {
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
    export interface ResData {
        data: [{
            label: string;
            value: string;
        }]
        errorCode: number;
        msg: string;
    }
    export interface ReqForm {
        key: string;
        value: string;
    }
}
// 已配置获取
export const listCfg = () => {
    // 返回的数据格式可以和服务端约定
    return RequestHttp.get<Rconfig.ListResData>('/cfg/v1/list');
}

// 默认配置获取
export const listDefaultCfg = () => {
    // 返回的数据格式可以和服务端约定
    return RequestHttp.get<Rconfig.ResData>('/cfg/v1/listdefault');
}

// 删除
export const delCfg = (params: Rconfig.ReqForm) => {
    return RequestHttp.delete<Rconfig.ResData>('/cfg/v1/del', {params: params});
}

//变更
export const updateCfg = (params: Rconfig.ReqForm) => {
    return RequestHttp.post<Rconfig.ResData>('/cfg/v1/update', params);
}