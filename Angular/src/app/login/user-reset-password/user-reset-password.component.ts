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


     @ViewChild('userResetPasswordForm') userUpdateForm: NgForm;
	
	    constructor(private userService: UserService, private router:Router) {

    }
	ngOnInit() {
		if(this.userService.getUserLoggedIn()){
            /*Should never happen*/
        }

	}
	onSubmit(hnForm: NgForm) {
      // if (hnForm.invalid) return;
       console.log(this.username);
        //let user = { username: this.username, password: this.password };
        this.userService.resetPassword(this.username)
            .then(() => {
                hnForm.resetForm();
                this.alertMsg = "Check your email!";
            });
    }
}