import { Component, OnInit } from '@angular/core';
import { ThreadService } from '../shared/thread.service';
import { ThreadDisplay } from '../shared/thread-display.model';

@Component({
    selector: 'app-thread-list',
    templateUrl: 'thread-list.component.html',
    styleUrls: ['thread-list.component.css']
})

export class ThreadListComponent implements OnInit {
    threads: ThreadDisplay[];
    counter:number;
    alertMsg = "";
    constructor(private threadService: ThreadService) {
        
    }
    loadMore(){
        this.counter +=100;
        this.threadService
        .getThreads(this.counter)
        .then(threads => this.threads = this.threads.concat(threads));
    }
    ngOnInit() {
        this.counter = 0;
        this.threadService
                .getThreads(this.counter)
                //.then(threads => this.threads = threads);
                .then((threads) => {
                    this.threads = threads;
                },
                reason => {
                    this.alertMsg = "Sorry! Something went wrong. Could not contact the server!";
                    console.error(reason);
                }

                );
    }
}