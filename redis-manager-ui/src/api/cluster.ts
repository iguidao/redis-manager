import RequestHttp from '../utils/request'

// 新增codis平台
namespace RCluster  {
    export interface updateReqForm {
        name: string;
        nodes: string;
        password: string;
    }
    export interface updateResData {
        data: {}
        errorCode: number;
        msg: string;
    }
    export interface listResData {
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
    export interface nodeReqForm {
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
    export interface nodeResData {
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
    export interface masterReqForm {
        cluster_id: string;
    }
    export interface masterResData {
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
// 删除
export const delClusterCfg = (params: RCluster.updateReqForm) => {
    return RequestHttp.delete<RCluster.updateResData>('/cluster/v1/del', {params: params});
}

//新增
export const addClusterCfg = (params: RCluster.updateReqForm) => {
    return RequestHttp.post<RCluster.updateResData>('/cluster/v1/add', params);
}

// 配置获取
export const listCluster = () => {
    // 返回的数据格式可以和服务端约定
    return RequestHttp.get<RCluster.listResData>('/cluster/v1/list');
}

// 配置获取
export const listClusterNodes = (params: RCluster.nodeReqForm) => {
    // 返回的数据格式可以和服务端约定
    return RequestHttp.get<RCluster.nodeResData>('/cluster/v1/nodes', { params: params});
}

//获取master信息
export const listClusterMaster = (params: RCluster.masterReqForm) => {
    return RequestHttp.get<RCluster.masterResData>('/cluster/v1/masters', { params: params});
}