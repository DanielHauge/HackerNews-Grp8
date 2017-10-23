import { Component, OnInit } from '@angular/core';
import { ThreadService } from '../shared/thread.service';
import { Thread } from '../shared/thread.model';

@Component({
    selector: 'app-thread-list',
    templateUrl: 'thread-list.component.html',
    styleUrls: ['thread-list.component.css']
})

export class ThreadListComponent implements OnInit {
    threads: Thread[];

    constructor(private threadService: ThreadService) {
        
    }

    ngOnInit() {
        this.threadService
                .getEntries()
                .then(threads => this.threads = threads);
    }
}