import { Component, EventEmitter, Input, Output } from '@angular/core';
import { User } from "../service/model";
import { UtilService } from '../service/util.service';

@Component({
  selector: 'application-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.css']
})
export class HeaderComponent {
  user!: User;
  @Input()
  set currentUser(user: User) {
    this.user = user;
  };

  @Output() onlogout = new EventEmitter<any>();

  constructor(
    private utilservice: UtilService,
  ) { }
  
  logout() {
    this.onlogout.emit();
  }

  toggle(opened?: boolean): void {
    this.utilservice.toogleSidenav(opened);
  }
}
