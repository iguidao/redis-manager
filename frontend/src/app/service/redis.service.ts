import { Injectable } from '@angular/core';
import {HttpClient, HttpErrorResponse} from '@angular/common/http';
import {Subject, throwError, Observable} from 'rxjs';
import {catchError, tap} from 'rxjs/operators';
import {environment} from '../../environments/environment';
import {
  User
} from './model';

@Injectable({
  providedIn: 'root'
})
export class RedisService {

  constructor(private http: HttpClient) { }

}



