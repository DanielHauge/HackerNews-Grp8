import { Component, ViewChild } from '@angular/core';
import { NgForm } from '@angular/forms';
import { UserService } from '../shared/user.service';
import { Router } from '@angular/router'
import { User } from '../shared/user.model';
import { RollbarService } from 'angular-rollbar';


@Component({
  selector: 'app-user-reset-password',
  templateUrl: 'user-reset-password.component.html',
  styleUrls: ['user-reset-password.component.css']  
})
export class UserResetPasswordComponent{
    alertMsg = "";
    username: string = "";


     @ViewChild('hnForm') hnForm: NgForm;
	
        constructor(private userService: UserService, 
            private router:Router, 
            private rollbar: RollbarService) {

    }
	ngOnInit() {
		if(this.userService.getUserLoggedIn()){
            /*Should never happen*/
        }

	}
	onSubmit(hnForm: NgForm) {

       console.log('User submitted recover/reset password..');
       console.log(this.username);
        this.userService.resetPassword(this.username)
            .then(() => {
                this.alertMsg = "Check your email!";                
                this.hnForm.resetForm();
            },
            reason => {
                /*The correct behaviour ends up here*/
                console.error(reason);
                this.alertMsg = "Check your email!";                
                this.hnForm.resetForm();
            }  );
    }
}