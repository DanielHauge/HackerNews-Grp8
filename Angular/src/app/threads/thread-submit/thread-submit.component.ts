import { Component, ViewChild, OnInit  } from '@angular/core';
import { NgForm } from '@angular/forms';
import { ThreadService } from '../shared/thread.service';
import { UserService } from '../../login/shared/user.service';
import { Router } from '@angular/router'

@Component({
    selector: 'app-thread-submit',
    templateUrl: 'thread-submit.component.html',
    styleUrls: ['thread-submit.component.css']
})

export class ThreadSubmitComponent implements OnInit {
    post_title: string = "";
    post_url: string = "";
    post_text: string = "";
    

    @ViewChild('submitThreadForm') submitThreadForm: NgForm;
	
	    constructor(private threadService: ThreadService, private userService: UserService,  private router:Router) {

    }
	ngOnInit() {
		if(!this.userService.getUserLoggedIn()){
			this.router.navigate(['/login']);

		}
		
	}
	onSubmit(submitThreadForm: NgForm) {
        if (this.submitThreadForm.invalid) return;
        //let thread = { post_title: this.post_title, post_url: this.post_url, post_text: post_text };
        let thread = { name: this.post_title, comment: this.post_url };
        /*this.threadService.addThread(this.threadId, thread)
            .then(() => {
                    //this.submitThreadForm.resetForm();
					//TODO if everything goes well see all threads including submitted
					this.router.navigate(['/threads']);
					
            });*/
    }
}
