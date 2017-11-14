/*export class UncaughtExceptionHandler implements ErrorHandler {
    handleError(error: any) {
        JL().fatalException('Uncaught Exception', error);
    }
}*/
import { RollbarModule, RollbarService } from 'angular-rollbar'

import { NgModule , ErrorHandler} from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { HttpModule } from '@angular/http';
import { FormsModule } from '@angular/forms';
//import { InMemoryWebApiModule } from 'angular-in-memory-web-api';<!-- used for dummy backend during development-->
//import { InMemoryThreadService } from './backend';
import { AppComponent } from './app.component';
import { appRouterModule } from "./app.routes";

import { UserService, LoginComponent, UserUpdateComponent, UserResetPasswordComponent } from './login';
import { ThreadListComponent, ThreadComponent, ThreadService, ThreadCommentFormComponent, ThreadDetailsComponent, ThreadSubmitComponent } from './threads';


import { HeaderComponent } from './header/header.component';
import { FooterComponent } from './footer/footer.component';
import { AuthguardGuard } from './authguard.guard';
// ... other imports ...
//import { JL } from 'jsnlog';
//JL().setOptions({ "defaultAjaxUrl":''});/*"level": JL.getWarnLevel() ,*/ 
//JL().setOptions({defaultAjaxUrl: 'any'});
/*JL.setOptions({
    "defaultAjaxUrl": "139.59.157.29:5000"
});*/

  

@NgModule({
    imports: [
        BrowserModule,
        HttpModule,
        FormsModule,
        appRouterModule,
        RollbarModule.forRoot({
            accessToken: '73823db57eb9461ca6e162596fe9aa67'
        })
        //InMemoryWebApiModule.forRoot(InMemoryThreadService)
    ],
    providers: [ ThreadService , UserService, AuthguardGuard,  { provide: ErrorHandler, useClass: RollbarService } ],
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