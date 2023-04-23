import RequestHttp from '../utils/request'

namespace Rboard {

    // 登录成功后返回的token
    export interface ResData {
        data: {
            aliredis: number;
            codis: number;
            txredis: number;
            cluster: number;
        }
        errorCode: number;
        msg: string;
    }
  }
  // 用户登录
  export const BoardDesc = () => {
      // 返回的数据格式可以和服务端约定
      return RequestHttp.get<Rboard.ResData>('/board/v1/desc');
  }
  