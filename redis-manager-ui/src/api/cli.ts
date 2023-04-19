import RequestHttp from '../utils/request'

namespace Rcli {
    // 用户登录表单
    export interface CliReqForm {
        cache_type: string;
        cache_op: string;
        cluster_name: string;
        key_name: string;
        codis_url: string;
        group_name: string;

    }
    // 登录成功后返回的token
    export interface CliResData {
        data: {}
        errorCode: number;
        msg: string;
    }
  }
  // 用户登录
  export const cliRedisOpkey = (params: Rcli.CliReqForm) => {
      // 返回的数据格式可以和服务端约定
      return RequestHttp.post<Rcli.CliResData>('/cli/v1/opkey', params);
  }
  