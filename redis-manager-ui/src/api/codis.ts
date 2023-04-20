import RequestHttp from '../utils/request'

// 新增codis平台
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
export const delCodisCfg = (params: Rcodisupdate.UpdateReqForm) => {
    return RequestHttp.delete<Rcodisupdate.SettingResData>('/cfg/v1/del', {params: params});
}

//新增
export const addCodisCfg = (params: Rcodisupdate.UpdateReqForm) => {
    return RequestHttp.post<Rcodisupdate.SettingResData>('/codis/v1/add', params);
}

// 获取codis平台
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
//获取codis平台集群列表
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
export const listCodisCluster = (params: Rclusterlist.ListReqForm) => {
    // 返回的数据格式可以和服务端约定
    return RequestHttp.get<Rclusterlist.ListResData>('/codis/v1/cluster', {params: params});
}
//获取codis集群的group列表
namespace Rgrouplist {
    export interface ListReqForm {
        cluster_name: string;
        curl: string;
    }
    export interface ListResData {
        data: [string]
        errorCode: number;
        msg: string;
    } 
}
export const listCodisGroup = (params: Rgrouplist.ListReqForm) => {
    // 返回的数据格式可以和服务端约定
    return RequestHttp.get<Rgrouplist.ListResData>('/codis/v1/group', {params: params});
}