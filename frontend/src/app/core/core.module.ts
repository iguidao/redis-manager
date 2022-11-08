import { NgModule, SkipSelf, Optional } from '@angular/core';

import { AppRoutingModule } from '../app-routing.module';
import { ShareModule } from '../share/share.module';
import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';


@NgModule({
  declarations: [],
  imports: [
    BrowserModule,
    ShareModule,
    BrowserAnimationsModule,
    AppRoutingModule
  ],
  exports: [
    ShareModule,
    BrowserModule,
    BrowserAnimationsModule,
    AppRoutingModule
  ],
})
export class CoreModule { 
  constructor(@SkipSelf() @Optional() parentModule: CoreModule) {
    if (parentModule) {
      throw new Error('CoreModule 只能被appModule引入');
    }
  }
}
