import { Routes, RouterModule } from '@angular/router';
import { LoginComponent, UserUpdateComponent } from './login';
import { AuthguardGuard  } from './authguard.guard';

import { ThreadListComponent, ThreadComponent, ThreadService, ThreadCommentFormComponent , ThreadDetailsComponent, ThreadSubmitComponent } from './threads';

// Route config let's you map routes to components
const routes: Routes = [
  // map '/stories' to the people list component
  {
    path: 'threads',
    component: ThreadListComponent,
  },
    // HERE: new route for StoryDetailsComponent
  // map '/persons/:id' to person details component
  {
    path: 'threads/:id',
    component: ThreadDetailsComponent,
  },
  {
    path: 'user/update',
    component: UserUpdateComponent,
  },
  {
    path: 'login',
    component: LoginComponent,
  },  
  {
    path: 'submit',
	canActivate: [AuthguardGuard],
    component: ThreadSubmitComponent,
  },
  
  // map '/' to '/persons' as our default route
  {
    path: '',
    redirectTo: '/threads',
    pathMatch: 'full'
  },
];

export const appRouterModule = RouterModule.forRoot(routes);
