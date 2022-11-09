import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { LoginComponent } from './login/login.component';
import { PageComponent } from './page/page.component';
// import { CodisComponent } from './page/codis/codis.component'
const routes: Routes = [
  { path: 'login', component: LoginComponent },
  { path: 'page', component: PageComponent },
  // { path: 'page/codis', component: CodisComponent },
  { path: 'page/:appname', component: PageComponent },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }







