import { Injectable } from '@angular/core';
import { HttpClient, HttpErrorResponse} from '@angular/common/http';
import {Subject, throwError, Observable} from 'rxjs';
import {catchError,shareReplay, tap} from 'rxjs/operators';
import { environment } from '../../environments/environment';
import { Article, ArticleList, QiniukodoToken } from './model';

@Injectable({
  providedIn: 'root'
})
export class GuidaoadminService {
  constructor(private http: HttpClient) { }

  private base = environment.ServerUrl;
  private snackbarContentSource = new Subject<string>();
  snackbarContent$ = this.snackbarContentSource.asObservable();

  sendError(content: string) {
    this.snackbarContentSource.next(content);
  }

  private userActionSource = new Subject<string>();
  userAction$ = this.userActionSource.asObservable();

  notify() {
    this.userActionSource.next();
  }
  currentArticle(): Article {
    return JSON.parse(localStorage.getItem('currentAtticle') || '') as Article;
  }

  AddArticle(articletitle: string, articlecontent: string){

    return this.http.post(this.base + '/article-admin/daoarticle/v1/add', { ArticleTitle: articletitle, ArticleContent: articlecontent })
  }

  VerifyArticle(articleid: string, articlestate: boolean){

    return this.http.put(this.base + '/article-admin/daoarticle/v1/verify', { ArticleId: articleid, ArticleState: articlestate})
  }

  DelArticle(articleid: string){

    return this.http.delete(this.base + '/article-admin/daoarticle/v1/del/' + articleid)
  }

  listArticle(status?: boolean) {
    let url = this.base + '/article-admin/daoarticle/v1/articles'+ `?status=${status}`;
    // if (status) {
    //   url = url + `?status=${status}`;
    // }

    return this.http.get<ArticleList>(url)
      .pipe(catchError(err => this.handleError(err)));
  }
  // 七牛
  GetKodoToken(){
    let url = this.base + '/guidao-admin/gdqn/v1/getkodotoken';
    return this.http.get<QiniukodoToken>(url)
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
