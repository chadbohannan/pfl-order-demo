import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppComponent } from './app.component';
import { ProductListComponent } from './product-list/product-list.component';
import { ArchwizardModule } from 'angular-archwizard';

@NgModule({
  declarations: [
    AppComponent,
    ProductListComponent
  ],
  imports: [
    ArchwizardModule,
    BrowserModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
