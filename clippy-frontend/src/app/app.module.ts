import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { NoopAnimationsModule } from '@angular/platform-browser/animations';
import { ItemsoverviewComponent } from './core/itemsoverview/itemsoverview.component';
import { ItemComponent } from './core/item/item.component';

@NgModule({
  declarations: [AppComponent, ItemsoverviewComponent, ItemComponent],
  imports: [BrowserModule, AppRoutingModule, NoopAnimationsModule],
  providers: [],
  bootstrap: [AppComponent],
})
export class AppModule {}
