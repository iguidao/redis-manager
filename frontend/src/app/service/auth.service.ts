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
    return JSON.parse(localStorage.getItem('currentUser') || '{ "errorCode": 400, "data": {} }') as User;
  }

  verify(): Promise<User> {
    let user = JSON.parse(localStorage.getItem('currentUser') || '{ "errorCode": 400, "data": {} }')
    if (user && user.errorCode == 400 && user.data.Token) {
      let expire = this.jwtHelper.getTokenExpirationDate(user.data.Token);
      if (expire != null) {
        let validin = expire.getTime() - (new Date()).getTime();  
        if (validin <= 10) {
          return Promise.reject("need login")
        } else {
          return this.refresh()
        }
      } else {
        return this.refresh()
      }
    } else {
      return Promise.reject("need login")
    }
  }

  refresh(): Promise<User> {
    return this.http.post(this.base + '/account-admin/daoguan/v1/refresh', null)
      .toPromise()
      .then(response => {
        let user = response as User;
        localStorage.setItem('currentUser', JSON.stringify(user))
        return user
      })
      .catch(this.handleError)
  }

  login(userphone: number, password: string): Promise<User> {

    return this.http.post(this.base + '/account-admin/auth/v1/login', { Phone: userphone, Password: password })
      .toPromise()
        .then(response => {

        let user = response as User;
        if (user && user.data.token) {
          localStorage.setItem('currentUser', JSON.stringify(user));
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