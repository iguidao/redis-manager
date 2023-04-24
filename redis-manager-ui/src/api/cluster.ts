import RequestHttp from '../utils/request'

// 新增codis平台
namespace RClusterupdate  {
    export interface ReqForm {
        name: string;
        nodes: string;
        password: string;
    }
    export interface ResData {
        data: {}
        errorCode: number;
        msg: string;
    }
}
// 删除
export const delClusterCfg = (params: RClusterupdate.ReqForm) => {
    return RequestHttp.delete<RClusterupdate.ResData>('/cluster/v1/del', {params: params});
}

//新增
export const addClusterCfg = (params: RClusterupdate.ReqForm) => {
    return RequestHttp.post<RClusterupdate.ResData>('/cluster/v1/add', params);
}


namespace RClusterlist {
    // 拿到配置数据
    export interface ResData {
        data: [{
            ID: number;
            CreatedAt: string;
            UpdatedAt: string;
            DeletedAt: string;
            Name: string;
            Nodes: string;
            Password: string;
        }]
        errorCode: number;
        msg: string;
    }
  }
// 配置获取
export const listCluster = () => {
    // 返回的数据格式可以和服务端约定
    return RequestHttp.get<RClusterlist.ResData>('/cluster/v1/list');
}
namespace RClusternodes {

    export interface ReqForm {
        cluster_id: string;
    }
    export interface Node {
        CreateTime: string;
        CluserId: string;
        NodeId: string;
        Address: string;
        Flags: string;
        LinkState: string;
        RunStatus: string;
        SlotRange: string;
        SlotNumber: string; 
        Children: Node[]; 
    }
    // 拿到配置数据
    export interface ResData {
        data: [{
            CreateTime: string;
            CluserId: string;
            NodeId: string;
            Address: string;
            Flags: string;
            LinkState: string;
            RunStatus: string;
            SlotRange: string;
            SlotNumber: string;
            Children: Node[];
        }]
        errorCode: number;
        msg: string;
    }
  }
// 配置获取
export const listClusterNodes = (params: RClusternodes.ReqForm) => {
    // 返回的数据格式可以和服务端约定
    return RequestHttp.get<RClusternodes.ResData>('/cluster/v1/nodes', { params: params});
}

namespace RClustermaster {
    export interface ReqForm {
        cluster_id: string;
    }
    export interface ResData {
        data: [{
            CreatedAt: string;
            CluserId: string;
            NodeId: string;
            Ip: string;
            Port: string;
            SlotRange: string;
        }]
        errorCode: number;
        msg: string;
    }
}
export const listClusterMaster = (params: RClusternodes.ReqForm) => {
    return RequestHttp.get<RClusternodes.ResData>('/cluster/v1/masters', { params: params});
}