import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';

import { AuthenticationService } from '../service/auth.service';
import { User } from '../service/model';

@Component({
  selector: 'login.component',
  templateUrl: 'login.component.html',
  styleUrls: ['login.component.css'],
})

export class LoginComponent implements OnInit {
  username: string = '';
  password: string = '';
  error!: string;

  user!: User;

  constructor(
    private router: Router,
    private authService: AuthenticationService,
  ) { }

  ngOnInit() {
    // this.router.navigate(['/page/dashboard'])
    this.authService.verify()
      .then(() => this.router.navigate(['/page/dashboard']))
  }

  login() {
    this.error = '';
    if (this.password == '' ) {
      this.error = '还没填好呢！'
      return
    };
    this.authService.login(this.username, this.password)
      .then(() => {
        this.authService.notify();
        this.router.navigate(['/page/dashboard']);
      })
      .catch(err => this.error = err)
  }
}