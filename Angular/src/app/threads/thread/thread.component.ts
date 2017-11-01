import { Component, Input } from '@angular/core';
import { Thread } from '../shared/thread.model';

@Component({
    selector: 'app-thread',
    templateUrl: 'thread.component.html',
    styleUrls: ['thread.component.css']
})

export class ThreadComponent {
    @Input() thread: Thread;

    
}