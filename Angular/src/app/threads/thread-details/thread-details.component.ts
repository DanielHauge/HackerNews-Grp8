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
	alertMsg = "";
	

	onCommentAdded(threadId:number) {
		//this.thread.comments.push(comment);
		this.threadService.getComments(threadId)
		
		.then((response) => {
			console.log(response);
			this.thread.comments = response;
				
		},
		 reason => {
			console.warn(reason);
			if(reason.status){
				this.alertMsg = "Wrong username or password";
			}
			else{
				this.alertMsg = "Login failed.";
				
			}                
			
		})
		.catch(	response => { 			
				console.error(response);
				this.alertMsg = "Whoops... something went wrong";
		});

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