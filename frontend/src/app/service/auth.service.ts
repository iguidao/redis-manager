import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { JwtHelperService } from "@auth0/angular-jwt";
import { Subject } from "rxjs";


import { User } from "./model";
import { environment } from '../../environments/environment';

@Injectable()
export class AuthenticationService {
  constructor(private http: HttpClient) { }

  private base = environment.ServerUrl;
  private jwtHelper: JwtHelperService = new JwtHelperService();

  private loginSource = new Subject<string>();
  loginAction$ = this.loginSource.asObservable();
  notify() {
    this.loginSource.next();
  }

  currentUser(): User {
    return JSON.parse(localStorage.getItem('currentUser') || '{ "errorCode": 400, "data": { "Token": "" } }') as User;
  }

  verify(): Promise<User> {
    let user = JSON.parse(localStorage.getItem('currentUser') || '{ "errorCode": 400, "data": { "Token": "" } }')
    if (user && user.data.Toke != "") {
      let expire = this.jwtHelper.getTokenExpirationDate(user.data.Token);
      if (expire != null) {
        console.log("expire: ",expire)
        let validin = expire.getTime() - (new Date()).getTime();  
        if (validin <= 10) {
          return Promise.reject("needa login")
        } else {
          return this.refresh()
        }
      } else {
        return this.refresh()
      }
    } else {
      return Promise.reject("needa login")
    }
  }

  refresh(): Promise<User> {
    return this.http.post(this.base + '/redis-manager/auth/v1/refresh', null)
      .toPromise()
      .then(response => {
        let user = response as User;
        localStorage.setItem('currentUser', JSON.stringify(user))
        return user
      })
      .catch(this.handleError)
  }
  login(uname: string, up: string): Promise<User> {
    return this.http.post(this.base + '/redis-manager/user/v1/sign-in', { 'username': uname, 'password': up })
      .toPromise()
        .then(response => {
          let user = response as User;
          // console.log("response: ", user.data.token)
          if (user && user.data.token) {
            localStorage.setItem('currentUser', JSON.stringify(user));
            // console.log("dengluchenggong: ",localStorage.getItem('currentUser') || '{ "errorCode": 400, "data": {} }')
            return user
          } else {
            return Promise.reject("登陆失败")
          }
      })
      .catch(err => this.handleError(err))
  }

  logout(): void {
    localStorage.removeItem('currentUser');
    this.http.post(this.base + '/logout', null)
  }

  private handleError(error: any): Promise<any> {
    if (error.error && error.error.error) {
      return Promise.reject(error.error.error)
    } else {
      return Promise.reject(error.message || error)
    }
  }
}