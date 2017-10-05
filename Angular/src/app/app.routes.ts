import { Routes, RouterModule } from '@angular/router';
import { LoginComponent } from './login';
import { ThreadListComponent, ThreadComponent, ThreadService, ThreadCommentFormComponent , ThreadDetailsComponent } from './threads';

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
    path: 'login',
    component: LoginComponent,
  },
  
  // map '/' to '/persons' as our default route
  {
    path: '',
    redirectTo: '/threads',
    pathMatch: 'full'
  },
];

export const appRouterModule = RouterModule.forRoot(routes);
