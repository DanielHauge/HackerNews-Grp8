import { Component, OnInit, OnDestroy, Input } from '@angular/core';
import { ActivatedRoute } from "@angular/router";
import { ThreadService } from '../shared/thread.service';
import { Thread } from '../shared/thread.model';
import { UserService } from '../../login/shared/user.service';
import { Router } from '@angular/router'



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
		private threadService:ThreadService, private userService: UserService, private router:Router) { }
	//commentVote(index:number){
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
		}
	}
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