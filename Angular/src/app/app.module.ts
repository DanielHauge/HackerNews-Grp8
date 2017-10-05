import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpModule } from '@angular/http';
import { FormsModule } from '@angular/forms';
import { InMemoryWebApiModule } from 'angular-in-memory-web-api';
import { AppComponent } from './app.component';
import { appRouterModule } from "./app.routes";

import { LoginComponent } from './login';
import { ThreadListComponent, ThreadComponent, ThreadService, ThreadCommentFormComponent, ThreadDetailsComponent } from './threads';

import { InMemoryThreadService } from './backend';


@NgModule({
    imports: [
        BrowserModule,
        HttpModule,
        FormsModule,
        appRouterModule,
        InMemoryWebApiModule.forRoot(InMemoryThreadService)
    ],
    providers: [ ThreadService ],
    declarations: [
        AppComponent, 
        ThreadComponent,
        ThreadListComponent,
        ThreadCommentFormComponent,
        ThreadDetailsComponent,
        LoginComponent,
    ],
    bootstrap: [AppComponent]
})
export class AppModule {

}