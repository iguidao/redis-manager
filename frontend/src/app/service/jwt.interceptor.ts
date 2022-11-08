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
    let user = JSON.parse(localStorage.getItem('currentUser') || '')
    if (user && user.data.Token) {
        request = request.clone({
            setHeaders: { Authorization: `${user.data.Token}`}
        });
    }
    return next.handle(request);
  }
}