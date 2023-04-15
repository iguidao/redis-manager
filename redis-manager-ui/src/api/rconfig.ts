import RequestHttp from '../utils/request'

namespace Rconfig {
    // 拿到配置数据
    export interface LoginResData {
        data: {
            lists: [{
                ID: number;
                CreatedAt: string;
                UpdatedAt: string;
                DeletedAt: string;
                Name: string;
                Value: string;
                Note: string;
            }];
            total: number;
        }
        errorCode: number;
        msg: string;
    }
  }
  // 配置获取
  export const list = () => {
      // 返回的数据格式可以和服务端约定
      return RequestHttp.get<Rconfig.LoginResData>('/cfg/v1/list');
  }
  