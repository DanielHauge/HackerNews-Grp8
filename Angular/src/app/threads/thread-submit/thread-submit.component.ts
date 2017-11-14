import { Component, ViewChild, OnInit  } from '@angular/core';
import { NgForm } from '@angular/forms';
import { ThreadService } from '../shared/thread.service';
import { UserService } from '../../login/shared/user.service';
import { Router } from '@angular/router'
import { RollbarService } from 'angular-rollbar';

@Component({
    selector: 'app-thread-submit',
    templateUrl: 'thread-submit.component.html',
    styleUrls: ['thread-submit.component.css']
})

export class ThreadSubmitComponent implements OnInit {
    alertMsg = "";
    
    post_title: string = "";
    post_url: string = "";

    @ViewChild('submitThreadForm') submitThreadForm: NgForm;
	
        constructor(private threadService: ThreadService, 
            private userService: UserService,  
            private router:Router , 
            private rollbar: RollbarService) {

    }
	ngOnInit() {
		if(!this.userService.getUserLoggedIn()){
			this.router.navigate(['/login']);
            this.rollbar.info('ThreadSubmitComponent: ngOnInit() ');                    
            
		}
		
	}
	onSubmit(submitThreadForm: NgForm) {
        if (this.submitThreadForm.invalid) return;
        let username = this.userService.getUsername();
        let password = this.userService.getPassword();

        //Example: {"post_title": "NYC Developer Dilemma", "post_text": "", "hanesst_id": 4, "post_type": "story", "post_parent": -1, "username": "onebeerdave", "pwd_hash": "fwozXFe7g0", "post_url": "http://avc.blogs.com/a_vc/2006/10/the_nyc_develop.html"}
        let thread  = {post_title: this.post_title, post_text: '', hanesst_id: -1, post_type: "story", post_parent: -1, username: username, pwd_hash: password, post_url: this.post_url}
        
        this.threadService.submitThread(thread)
            .then((response) => {
                this.alertMsg = "Your thread was successfully submitted!";
                this.submitThreadForm.resetForm();
                console.log('submitThreadForm');
                console.log(response);
					
            },
             reason => {
                console.warn(reason);
                if(reason.status != 200){
                    this.alertMsg = "Whoops... something went wrong";
                    this.rollbar.error('ThreadSubmitComponent: if Whoops... something went wrong!');                    
                    
                }
                else{
                    this.alertMsg = "Whoops... something went wrong";
                    this.rollbar.error('ThreadSubmitComponent: else Whoops... something went wrong!');                    
                    
                }                
                
            })
            .catch(	response => { 			
                    console.error(response);
                    this.alertMsg = "Whoops... something went wrong";
                    this.rollbar.error('ThreadSubmitComponent: submitThread(thread)'+response);                    
                    
            });

    }
}
