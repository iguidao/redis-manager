import RequestHttp from '../utils/request'

// 新增codis平台
namespace Rcodis  {
    export interface ReqForm {
        cname: string;
        curl: string;
    }
    export interface SettingResData {
        data: number
        errorCode: number;
        msg: string;
    }
    export interface ListResData {
        data: {
            lists: [{
                ID: number;
                CreatedAt: string;
                UpdatedAt: string;
                DeletedAt: string;
                Curl: string;
                Cname: string;
            }];
            total: number;
        }
        errorCode: number;
        msg: string;
    }
    export interface ListclusterReqForm {
        cname: string;
        curl: string;
    }
    export interface ListclusterResData {
        data: [string]
        errorCode: number;
        msg: string;
    } 
    export interface ListgroupReqForm {
        cluster_name: string;
        curl: string;
    }
    export interface ListgroupResData {
        data: [string]
        errorCode: number;
        msg: string;
    } 
    export interface opnodeReqForm {
        curl: string;
        cluster_name: string;
        add_proxy: string;
        add_server: string;
        del_proxy: number;
        del_group: number;
        op_type: string
    }
    export interface opnodeResData {
        data: string
        errorCode: number;
        msg: string;
    }
}
// 删除
export const delCodisCfg = (params: Rcodis.ReqForm) => {
    return RequestHttp.delete<Rcodis.SettingResData>('/cfg/v1/del', {params: params});
}

//新增
export const addCodisCfg = (params: Rcodis.ReqForm) => {
    return RequestHttp.post<Rcodis.SettingResData>('/codis/v1/add', params);
}

// 配置获取
export const listCodis = () => {
    // 返回的数据格式可以和服务端约定
    return RequestHttp.get<Rcodis.ListResData>('/codis/v1/list');
}
//获取codis平台集群列表
export const listCodisCluster = (params: Rcodis.ListclusterReqForm) => {
    // 返回的数据格式可以和服务端约定
    return RequestHttp.get<Rcodis.ListclusterResData>('/codis/v1/cluster', {params: params});
}
//获取codis集群的group列表
export const listCodisGroup = (params: Rcodis.ListgroupReqForm) => {
    // 返回的数据格式可以和服务端约定
    return RequestHttp.get<Rcodis.ListgroupResData>('/codis/v1/group', {params: params});
}
//针对codis进行扩容缩容
export const opCodisNode = (params: Rcodis.opnodeReqForm) => {
    return RequestHttp.post<Rcodis.opnodeResData>('/codis/v1/opnode', params);
}