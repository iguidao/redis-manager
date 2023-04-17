import RequestHttp from '../utils/request'

namespace Rhistory {
    // 拿到配置数据
    export interface HistoryResData {
        data: [
            {
                ID: number;
                CreatedAt: string;
                UpdatedAt: string;
                DeletedAt: string;
                UserId: number;
                OpInfo: string;
                OpParams: string;
            }
        ]
        errorCode: number;
        msg: string;
    }
  }
  // 配置获取
  export const list = () => {
      // 返回的数据格式可以和服务端约定
      return RequestHttp.get<Rhistory.HistoryResData>('/ophistory/v1/list');
  }
  