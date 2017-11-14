import { Component, OnInit } from '@angular/core';
import { ThreadService } from '../shared/thread.service';
import { ThreadDisplay } from '../shared/thread-display.model';
import { RollbarService } from 'angular-rollbar';
@Component({
    selector: 'app-thread-list',
    templateUrl: 'thread-list.component.html',
    styleUrls: ['thread-list.component.css']
})

export class ThreadListComponent implements OnInit {
    threads: ThreadDisplay[];
    counter:number;
    alertMsg = "";

    constructor(private threadService: ThreadService, private rollbar: RollbarService) {
        rollbar.info('ThreadListComponent: constructor called');
    }
    loadMore(){
        this.rollbar.info('ThreadListComponent: loadMore()');               
        this.counter +=100;
        this.threadService
        .getThreads(this.counter)
        .then(threads => this.threads = this.threads.concat(threads));
    }
    ngOnInit() {
        this.counter = 0;
        this.threadService
                .getThreads(this.counter)
                .then((threads) => {
                    this.threads = threads;
                    this.rollbar.info('ThreadListComponent: getThreads()');                    
                    
                },
                reason => {
                    this.alertMsg = "Sorry! Something went wrong. Could not contact the server! Please contact us.";
                    this.rollbar.error('ThreadListComponent: No connection to the database!');                    
                    console.error(reason);
                }

                );
    }
}