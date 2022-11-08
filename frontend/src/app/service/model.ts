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

// 文章请求返回数据
export interface Article {
  errorCode: number;
  msg: string;
  data: ArticleData;
}

export interface ArticleData {
  Token: string;
  ArticleId: string;
}

// 文章列表请求返回数据
export interface ArticleList {
  errorCode: number;
  msg: string;
  data: AListData;
}

export interface AListData {
  lists: AListInfoList[];
  total: number;
}

export interface AListInfoList {
  Id: string;
  CreatedAt: Time;
  UpdatedAt: Time;
  DeletedAt: Time;
  ArticleTitle: string;
  ArticleContent: string;
  AuthorName: string;
  State: boolean;
}

// qiniukodoken请求返回数据
export interface QiniukodoToken {
  errorCode: number;
  msg: string;
  data: QiniuTokenData;
}

export interface QiniuTokenData {
  kodotoken: string;
}

// qiniu返回
export interface KodoResult {
  hash: string;
  key: string;
}

//小说列表
export interface NovelList {

}
//小说分类列表
export interface NovelTypesList {
  errorCode: number;
  msg: string;
  data: NovelTypesData;
}
export interface NovelTypesData {
  categorys: NovelTypeCategorys[];
}

export interface NovelTypeCategorys {
}

export interface NovelTypeCategoryinfo {
  id: number;
  pid: number;
  name: string;
  desc: string;
  child: NovelTypeCategoryChild[];
}
export interface NovelTypeCategoryChild {
  id: number;
  pid: number;
  name: string;
  desc: string;
  child: boolean;
}