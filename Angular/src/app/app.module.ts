import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpModule } from '@angular/http';
import { FormsModule } from '@angular/forms';
//import { InMemoryWebApiModule } from 'angular-in-memory-web-api';
import { AppComponent } from './app.component';
import { appRouterModule } from "./app.routes";

import { UserService, LoginComponent, UserUpdateComponent, UserResetPasswordComponent } from './login';
import { ThreadListComponent, ThreadComponent, ThreadService, ThreadCommentFormComponent, ThreadDetailsComponent, ThreadSubmitComponent } from './threads';

import { InMemoryThreadService } from './backend';
import { HeaderComponent } from './header/header.component';
import { FooterComponent } from './footer/footer.component';
import { AuthguardGuard } from './authguard.guard';


@NgModule({
    imports: [
        BrowserModule,
        HttpModule,
        FormsModule,
        appRouterModule,
        //InMemoryWebApiModule.forRoot(InMemoryThreadService)
    ],
    providers: [ ThreadService , UserService, AuthguardGuard],
    declarations: [
        AppComponent, 
        ThreadComponent,
        ThreadListComponent,
        ThreadCommentFormComponent,
        ThreadDetailsComponent,
        ThreadSubmitComponent,
        UserUpdateComponent, 
        UserResetPasswordComponent, 
        
        LoginComponent, HeaderComponent, FooterComponent,
    ],
    bootstrap: [AppComponent]
})
export class AppModule {

}