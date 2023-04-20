import RequestHttp from '../utils/request'

namespace RTxRegionlist {
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
export const listTxRegion = (params: RTxRegionlist.ListReqForm) => {
    // 返回的数据格式可以和服务端约定
    return RequestHttp.get<RTxRegionlist.ListResData>('/cloud/v1/region', {params: params});
}

namespace RTxRedislist {
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
export const listTxRedis = (params: RTxRedislist.ListReqForm) => {
    // 返回的数据格式可以和服务端约定
    return RequestHttp.get<RTxRedislist.ListResData>('/cloud/v1/list', {params: params});
}