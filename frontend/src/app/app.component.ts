import { Component } from '@angular/core';
import { ActivatedRoute, ParamMap, Router } from '@angular/router';
import { switchMap } from 'rxjs/operators';
import { of } from "rxjs";
import { MatDialog } from '@angular/material/dialog';
import { MatSnackBar } from '@angular/material/snack-bar';
import { User } from "./service/model";
import { AuthenticationService } from "./service/auth.service";
import { UtilService } from './service/util.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'Redis Manager';
  theme!: string;
  user!: User;

  constructor(
    public dialog: MatDialog,
    private snackbar: MatSnackBar,
    private route: ActivatedRoute,
    private router: Router,
    private authService: AuthenticationService,
    private utilservice: UtilService,
  ) {
    this.utilservice.snackbarContent$.subscribe(content => {
      this.showError(content);
    });
  }

  ngOnInit() {

    this.authService.verify()
      .then(user => this.user = user)
      .catch(() => this.router.navigate(['login']))


    this.route.queryParamMap.pipe(
      switchMap((params: ParamMap) => of(params.get('env')))
    ).subscribe(env => this.theme = env + "-theme")

    this.authService.loginAction$.subscribe(
      () => this.user = this.authService.currentUser()
    );
  }

  logout() {
    this.authService.logout();
    this.user = {} as User;
    this.router.navigate(['login']);
  }

  showError(c: string) {
    this.snackbar.open(c, 'OK', {
      duration: 3000,
    });
  }
}
