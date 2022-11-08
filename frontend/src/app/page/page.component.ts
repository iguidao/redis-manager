import { Component, OnInit, ViewChild } from '@angular/core';
import { MatSidenav } from "@angular/material/sidenav";
import {
  BreakpointObserver,
  Breakpoints,
  BreakpointState,
} from '@angular/cdk/layout';

import { Router, ActivatedRoute, ParamMap } from '@angular/router';
import { map, switchMap } from 'rxjs/operators';
import { of, Observable } from 'rxjs';

import { User } from "../service/model";
import { UtilService } from '../service/util.service';
import { AuthenticationService } from '../service/auth.service';

@Component({
  selector: 'app-page',
  templateUrl: './page.component.html',
  styleUrls: ['./page.component.css']
})
export class PageComponent implements OnInit {
  user!: User;
  // app: AppInfo = {} as AppInfo;
  appscore: number = 0;
  select!: string;
  env!: string;
  appname!: string;
  envs!: string[];
  snackbar: any;


  @ViewChild(MatSidenav, { static: true })
  sidenav!: MatSidenav;

  constructor(
    private route: ActivatedRoute,
    private router: Router,
    private utilservice: UtilService,
    private authService: AuthenticationService,
    private breakpointObserver: BreakpointObserver,
  ) {

    utilservice.sidenavSource$.subscribe(o => {
      if (false === o) {
        this.sidenav.close();
      } else if (true === o) {
        this.sidenav.open();
      } else {
        this.sidenav.toggle();
      }
    })
  }

  public isHandset$: Observable<boolean> = this.breakpointObserver
    .observe(Breakpoints.Handset)
    .pipe(map((result: BreakpointState) => result.matches));
  
  ngOnInit() {
    this.route.paramMap.pipe(
      switchMap((params: ParamMap) => of(params.get('appname')))
    ).subscribe((appname: string | null) => {
      if (appname && appname != 'null') {
        this.appname = appname;
        localStorage.setItem('app', appname);
        this.router.navigate(['page', this.appname], {
          queryParamsHandling: 'preserve',
        })
        // this.refreshAppInfo();
      }else{
        this.router.navigate(['page', 'dashboard'], {
          queryParamsHandling: 'preserve',
        })     
      }
    })
  }
  selectApp(app: string) {
    this.router.navigate(['page', app], {
      queryParamsHandling: "preserve",
    })
  }
}