import RequestHttp from '../utils/request'

namespace Rruledata {
    export interface AddReqForm {
        identity: string;
        path: string;
        method: string;
    }
    export interface AddResData {
        data: {
            result: string;
        }
        errorCode: number;
        msg: string;
    }
    export interface ListResData {
        data: [{
            ID: number;
            identity: string;
            method: string;
            note: string;
            path: string;
        }]
        errorCode: number;
        msg: string;
    }
    export interface CfgResData {
        data: {
            method: [{
                label: string;
                value: string;
            }]
            url: [{
                label: string;
                value: string;   
            }]
        }
        errorCode: number;
        msg: string;
    }
}
export const rulelist = () => {
    // 返回的数据格式可以和服务端约定
    return RequestHttp.get<Rruledata.ListResData>('/rule/v1/list');
}

export const ruledel = (params: Rruledata.AddReqForm) => {
    return RequestHttp.delete<Rruledata.AddResData>('/rule/v1/del', {params: params});
}

export const rulecfg = () => {
    return RequestHttp.get<Rruledata.CfgResData>('/rule/v1/cfg');
}

export const ruleadd = (params: Rruledata.AddReqForm) => {
    return RequestHttp.post<Rruledata.AddResData>('/rule/v1/add', params);
}