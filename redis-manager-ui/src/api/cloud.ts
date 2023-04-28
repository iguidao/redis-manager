import RequestHttp from '../utils/request'

namespace RCloud {
    export interface RegionReqForm {
        cloud: string;
    }
    export interface RegionResData {
        data: {
            region_list: [{
                Region: string;
                RegionName: string;
                RegionState: string;
                RegionId: string;
                RegionEndpoint: string;
                LocalName: string;
            }]
        }
        errorCode: number;
        msg: string;
    } 
    export interface ListReqForm {
        cloud: string;
        region: string;
    }
    export interface ListResData {
        data: {
            redis_list: [
                {
                    Cloud: string;
                    InstanceName: string;
                    InstanceId: string;        
                    PrivateIp: string;
                    Port: string;
                    Region: string;
                    Createtime: string;
                    Size: number;
                    InstanceStatus: string;
                    RedisShardSize: number;
                    RedisShardNum: number;
                    RedisReplicasNum: number;
                    NoAuth: Boolean;
                    PublicIp: string;
                }
            ]
        }
        errorCode: number;
        msg: string;
    } 
    export interface PasswordReqForm {
        cloud: string;
        instanceid: string;
        password: string;
    }
    export interface PasswordResData {
        data: {}
        errorCode: number;
        msg: string;
    } 
}
export const listCloudRegion = (params: RCloud.RegionReqForm) => {
    // 返回的数据格式可以和服务端约定
    return RequestHttp.get<RCloud.RegionResData>('/cloud/v1/region', {params: params});
}

export const listCloudRedis = (params: RCloud.ListReqForm) => {
    // 返回的数据格式可以和服务端约定
    return RequestHttp.get<RCloud.ListResData>('/cloud/v1/list', {params: params});
}


export const changeCloudRedisPw = (params: RCloud.PasswordReqForm) => {
    // 返回的数据格式可以和服务端约定
    return RequestHttp.post<RCloud.PasswordResData>('/cloud/v1/password', params);
}
