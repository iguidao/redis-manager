import { Injectable } from '@angular/core';
import {
  HttpRequest,
  HttpHandler,
  HttpEvent,
  HttpInterceptor
} from '@angular/common/http';

import { Observable } from 'rxjs';

@Injectable()
export class JwtInterceptor implements HttpInterceptor {

  constructor() {}

  intercept(request: HttpRequest<any>, next: HttpHandler): Observable<HttpEvent<any>> {
    let user = JSON.parse(localStorage.getItem('currentUser') || '{ "errorCode": 400, "data": { "token": "" } }')
    if (user && user.data.token != "") {
        request = request.clone({
            setHeaders: { Authorization: `${user.data.token}`}
        });
    }
    return next.handle(request);
  }
}

