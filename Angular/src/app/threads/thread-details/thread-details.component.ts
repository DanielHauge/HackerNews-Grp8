import { Component, OnInit, OnDestroy, Input } from '@angular/core';
import { ActivatedRoute } from "@angular/router";
import { ThreadService } from '../shared/thread.service';
import { Thread } from '../shared/thread.model';
import { UserService } from '../../login/shared/user.service';
import { Router } from '@angular/router'
import { RollbarService } from 'angular-rollbar';



@Component({
	selector: 'app-thread-details',
	templateUrl: 'thread-details.component.html',
	styleUrls: ['thread-details.component.css']  
})


export class ThreadDetailsComponent implements OnInit, OnDestroy {
	@Input() thread: Thread;
	sub:any;
	alertMsg = "";
	isLiked:boolean = false;
	
	constructor(private route:ActivatedRoute,
		private threadService:ThreadService, 
		private userService: UserService, 
		private router:Router, 
		private rollbar: RollbarService
	) { }
	ngOnInit() {
		this.sub = this.route.params.subscribe(params => {
			let id = Number.parseInt(params['id']);
			this.threadService.getThread(id).then( data => this.thread = data );
		});
	}
	commentVote(comment:any){
		if(this.userService.getUserLoggedIn()){
			if(comment.vote == false){
				comment.vote = true ;
				comment.points += 1 ;
				
				let username = this.userService.getUsername();			
				this.threadService.upvotePost({thread_id: -1,comment_id: comment.id, username: username});
			}
		}
		else{
			this.router.navigate(['/login']);
			
			this.alertMsg = "You must be logged in to be able upvote.";
		}
	}
	threadVote(){
		if(this.userService.getUserLoggedIn()){
				
			if(this.isLiked == false ){

				this.isLiked = true;
				let username = this.userService.getUsername();			
				this.threadService.upvotePost({thread_id: this.thread.id, comment_id: -1, username: username});
				this.thread.points += 1;
			}
		}
		else{
			this.router.navigate(['/login']);
			
			this.alertMsg = "You must be logged in to be able upvote.";
			this.rollbar.warn('ThreadDetailsComponent: You must be logged in to be able upvote.');                    
			
		}
	}
	onCommentAdded( comment:any  ) {

			this.thread.commentamount += 1;
			comment = {username:comment.username,comment:comment.comment,vote:0,points:0,time:'Just now'}
			this.thread.comments.unshift(comment);
				
	}



	ngOnDestroy(): void {
		this.sub.unsubscribe();
	}
}