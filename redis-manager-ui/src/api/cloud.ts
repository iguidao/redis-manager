import RequestHttp from '../utils/request'

namespace RRegionlist {
    export interface ListReqForm {
        cloud: string;
    }
    export interface ListResData {
        data: {
            region_list: {
                TotalCount: number;
                RequestId: string;
                RegionSet: [{
                    Region: string;
                    RegionName: string;
                    RegionState: string;
                }]
            }
        }
        errorCode: number;
        msg: string;
    } 
}
export const listCloudRegion = (params: RRegionlist.ListReqForm) => {
    // 返回的数据格式可以和服务端约定
    return RequestHttp.get<RRegionlist.ListResData>('/cloud/v1/region', {params: params});
}

namespace RRedislist {
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
}
export const listCloudRedis = (params: RRedislist.ListReqForm) => {
    // 返回的数据格式可以和服务端约定
    return RequestHttp.get<RRedislist.ListResData>('/cloud/v1/list', {params: params});
}

namespace RChangePassword {
    export interface ReqForm {
        cloud: string;
        instanceid: string;
        password: string;
    }
    export interface ResData {
        data: {}
        errorCode: number;
        msg: string;
    } 
}
export const changeCloudRedisPw = (params: RChangePassword.ReqForm) => {
    // 返回的数据格式可以和服务端约定
    return RequestHttp.post<RChangePassword.ResData>('/cloud/v1/password', params);
}
