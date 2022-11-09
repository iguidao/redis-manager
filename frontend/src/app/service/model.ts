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
  username: string;
  result: string;
}

// codis信息
export interface CodisInfo {
  errorCode: number;
  msg: string;
  data: CodisData;
}
export interface CodisData {
  lists: CodisList[];
  total: number;
}
export interface CodisList {
  ID: number;
  CreatedAt: Time;
  UpdatedAt: Time;
  DeletedAt: Time;
  Curl: string;
  Cname: string;
}
