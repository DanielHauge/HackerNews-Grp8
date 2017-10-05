import { Component, OnInit, OnDestroy } from '@angular/core';
import { ActivatedRoute } from "@angular/router";
import { ThreadService } from '../shared/thread.service';
import { Thread } from '../shared/thread.model';

@Component({
  selector: 'app-thread-details',
  templateUrl: 'thread-details.component.html',
  styleUrls: ['thread-details.component.css']  
})
export class ThreadDetailsComponent implements OnInit, OnDestroy {
  thread: Thread;
  sub:any;

  constructor(private route:ActivatedRoute,
              private threadService:ThreadService) { }

  ngOnInit() {
    
    this.sub = this.route.params.subscribe(params => {
      let id = Number.parseInt(params['id']);
      this.threadService.getThread(id).then( data => this.thread = data );
    });
  }

  ngOnDestroy(): void {
    this.sub.unsubscribe();
  }
}