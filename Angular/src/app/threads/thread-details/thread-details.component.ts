import { Component, OnInit, OnDestroy, Input } from '@angular/core';
import { ActivatedRoute } from "@angular/router";
import { ThreadService } from '../shared/thread.service';
import { Thread } from '../shared/thread.model';

@Component({
	selector: 'app-thread-details',
	templateUrl: 'thread-details.component.html',
	styleUrls: ['thread-details.component.css']  
})


export class ThreadDetailsComponent implements OnInit, OnDestroy {
	@Input() thread: Thread;
	sub:any;

	onCommentAdded(comment: {name: string; comment: string;}) {
		this.thread.comments.push(comment);
	}
	

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