import { Time } from "@angular/common";

// 用户请求返回数据
export interface User {
  errorCode: number;
  msg: string;
  data: Authdata;
}

// 用户认证信息
export interface Authdata {
  token: string;
  NickName: string;
  result: string;
}
