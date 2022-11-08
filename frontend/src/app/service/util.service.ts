import { Injectable } from '@angular/core';
import {Subject, throwError, Observable} from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class UtilService {

  constructor() { }

  private snackbarContentSource = new Subject<string>();
  snackbarContent$ = this.snackbarContentSource.asObservable();

  sendError(content: string) {
    this.snackbarContentSource.next(content);
  }

  private sidenavSource = new Subject<boolean>();
  sidenavSource$ = this.sidenavSource.asObservable();

  toogleSidenav(opened?: boolean) {
    this.sidenavSource.next(opened);
  }

}
