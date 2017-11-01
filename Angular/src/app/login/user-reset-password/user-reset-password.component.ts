import { Component, ViewChild } from '@angular/core';
import { NgForm } from '@angular/forms';
import { UserService } from '../shared/user.service';
import { Router } from '@angular/router'
import { User } from '../shared/user.model';


@Component({
  selector: 'app-user-reset-password',
  templateUrl: 'user-reset-password.component.html',
  styleUrls: ['user-reset-password.component.css']  
})
export class UserResetPasswordComponent{
    alertMsg = "";
    username: string = "";


     @ViewChild('hnForm') hnForm: NgForm;
	
	    constructor(private userService: UserService, private router:Router) {

    }
	ngOnInit() {
		if(this.userService.getUserLoggedIn()){
            /*Should never happen*/
        }

	}
	onSubmit(hnForm: NgForm) {
       if (hnForm.invalid) return;
       console.log('User submitted recover/reset password..');
       console.log(this.username);
        this.userService.resetPassword(this.username)
            .then(() => {
                this.alertMsg = "Check your email!";                
                this.hnForm.resetForm();
            },
            reason => {
                console.error(reason);
                this.alertMsg = "Whoops, something went wrong. Make sure you put the correct username.";
            }  );
    }
}