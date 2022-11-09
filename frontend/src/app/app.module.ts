import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { AppComponent } from './app.component';
import { CoreModule } from './core/core.module';
import { HeaderComponent } from './header/header.component';
import { LoginComponent } from "./login/login.component";
import { AppRoutingModule } from "./app-routing.module";
import { HttpClientModule, HTTP_INTERCEPTORS } from "@angular/common/http";
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

import { UtilService } from './service/util.service';
import { AuthenticationService } from "./service/auth.service";
import { JwtInterceptor } from './service/jwt.interceptor';
import { UserService } from './service/user.service';
import { PageComponent } from './page/page.component';
import { SelectComponent } from './page/select/select.component';
import { DashboardComponent } from './page/dashboard/dashboard.component';
import { UserComponent } from './page/user/user.component';
import { CodisComponent } from './page/codis/codis.component';
import { RedisComponent } from './page/redis/redis.component';



@NgModule({
  declarations: [
    AppComponent,
    HeaderComponent,
    LoginComponent,
    PageComponent,
    SelectComponent,
    DashboardComponent,
    UserComponent,
    CodisComponent,
    RedisComponent
  ],
  imports: [
    CoreModule,
    AppRoutingModule,
    HttpClientModule,
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule
  ],
  providers: [
    { provide: HTTP_INTERCEPTORS, useClass: JwtInterceptor, multi: true },
    UtilService,
    UserService,
    AuthenticationService
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
