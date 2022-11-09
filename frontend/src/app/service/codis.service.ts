import { Injectable } from '@angular/core';
import { HttpClient, HttpErrorResponse} from '@angular/common/http';
import {Subject, throwError, Observable} from 'rxjs';
import {catchError,shareReplay, tap} from 'rxjs/operators';
import { environment } from '../../environments/environment';
import { CodisInfo, CodisList } from './model';

@Injectable({
  providedIn: 'root'
})
export class CodisService {

  constructor(private http: HttpClient) { }

  private base = environment.ServerUrl;

  private snackbarContentSource = new Subject<string>();
  snackbarContent$ = this.snackbarContentSource.asObservable();
  sendError(content: string) {
    this.snackbarContentSource.next(content);
  }

  listCodis() {
    let url = this.base + '/redis-manager/codis/v1/list';
    return this.http.get<CodisInfo>(url)
      .pipe(catchError(err => this.handleError(err)));
  }

  private handleError(error: HttpErrorResponse) {
    let msg = '';
    if (error.error instanceof ErrorEvent) {
      // A client-side or network error occurred. Handle it accordingly.
      msg = `Cannot request: ${error.error.message}`;
    } else {
      // The backend returned an unsuccessful response code.
      // The response body may contain clues as to what went wrong,
      msg = `Status: ${error.status}, Body: ` + JSON.stringify(error.error);
    }
    console.error(msg);
    this.sendError(msg);
    // return an observable with a user-facing error message
    return throwError(msg);
  };
}


