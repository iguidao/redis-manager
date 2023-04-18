import RequestHttp from '../utils/request'


namespace Rcodisupdate  {
    export interface UpdateReqForm {
        cname: string;
        curl: string;
    }
    export interface SettingResData {
        data: number
        errorCode: number;
        msg: string;
    }
}
// 删除
export const delCfg = (params: Rcodisupdate.UpdateReqForm) => {
    return RequestHttp.delete<Rcodisupdate.SettingResData>('/cfg/v1/del', {params: params});
}

//新增
export const addCfg = (params: Rcodisupdate.UpdateReqForm) => {
    return RequestHttp.post<Rcodisupdate.SettingResData>('/codis/v1/add', params);
}


namespace Rcodislist {
    // 拿到配置数据
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
  }
// 配置获取
export const listCodis = () => {
    // 返回的数据格式可以和服务端约定
    return RequestHttp.get<Rcodislist.ListResData>('/codis/v1/list');
}

namespace Rclusterlist {
    export interface ListReqForm {
        cname: string;
        curl: string;
    }
    export interface ListResData {
        data: [string]
        errorCode: number;
        msg: string;
    } 
}
export const listCluster = (params: Rclusterlist.ListReqForm) => {
    // 返回的数据格式可以和服务端约定
    return RequestHttp.get<Rclusterlist.ListResData>('/codis/v1/cluster', {params: params});
}