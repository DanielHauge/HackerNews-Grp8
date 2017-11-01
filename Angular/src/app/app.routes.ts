import { Routes, RouterModule } from '@angular/router';
import { LoginComponent, UserUpdateComponent, UserResetPasswordComponent } from './login';
import { AuthguardGuard  } from './authguard.guard';

import { ThreadListComponent, ThreadComponent, ThreadService, ThreadCommentFormComponent , ThreadDetailsComponent, ThreadSubmitComponent } from './threads';

// Route config let's you map routes to components
const routes: Routes = [
  // map '/threads' to the threads-list component
  {
    path: 'threads',
    component: ThreadListComponent,
  },
  // HERE:route for ThreadDetailsComponent
  // map '/threads/:id' to thread-details component
  {
    path: 'threads/:id',
    component: ThreadDetailsComponent,
  },
  // HERE:route for UserUpdateComponent
  // map '/threads/:id' to user-update component
  {
    path: 'user/update',
    component: UserUpdateComponent,
  },
  // HERE:route for UserResetPasswordComponent
  // map '/threads/:id' to user-reset-password component  
  {
    path: 'user/resetpassword',
    component: UserResetPasswordComponent,
  },
  // HERE:route for LoginComponent
  // map '/threads/:id' to login component 
  {
    path: 'login',
    component: LoginComponent,
  }, 
  // HERE:route for ThreadSubmitComponent
  // map '/threads/:id' to thread-submit component    
  {
    path: 'submit',
  	//canActivate: [AuthguardGuard],
    component: ThreadSubmitComponent,
  },
  
  // map '/' to '/threads' as our default route
  {
    path: '',
    redirectTo: '/threads',
    pathMatch: 'full'
  },
];

export const appRouterModule = RouterModule.forRoot(routes);
