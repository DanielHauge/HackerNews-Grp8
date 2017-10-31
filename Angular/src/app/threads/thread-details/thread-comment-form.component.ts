import { Component, EventEmitter, Input, Output, ViewChild } from '@angular/core';
import { NgForm } from '@angular/forms';
import { ThreadService } from '../shared/thread.service';
import { UserService } from '../../login/shared/user.service';

@Component({
    selector: 'app-thread-comment-form',
    templateUrl: 'thread-comment-form.component.html'
})

export class ThreadCommentFormComponent {
    alertMsg = "";
    
    comment: string = "";
    canMakeComment:boolean;
    @Input() threadId: number;
    @Output() onCommentAdded = new EventEmitter<{ username: string; comment: string; password: string }>();
    @ViewChild('commentForm') commentForm: NgForm;
    
    constructor(private threadService: ThreadService, private userService: UserService) {

    }
    ngOnInit() {
        this.canMakeComment = this.userService.getUserLoggedIn();
        
	}
    onSubmit(commentForm: NgForm) {
        if (this.commentForm.invalid) return;
        let username = this.userService.getUsername();
        let password = this.userService.getPassword();
        let comment = { username: username, comment: this.comment, password: password };
        
        
        
        this.threadService.addComment(this.threadId, comment )
            /*.then(() => {
               // delete comment.password;
                this.onCommentAdded.emit(comment);
                    this.commentForm.resetForm();
            });*/

            .then((response) => {
                console.log(response);
                this.onCommentAdded.emit(comment);
                this.commentForm.resetForm();
					
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
}