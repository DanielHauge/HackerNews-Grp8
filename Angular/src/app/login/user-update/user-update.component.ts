import { Component, ViewChild } from '@angular/core';
import { NgForm } from '@angular/forms';
import { UserService } from '../shared/user.service';
import { Router } from '@angular/router'
import { User } from '../shared/user.model';


@Component({
  selector: 'app-user-update',
  templateUrl: 'user-update.component.html',
  styleUrls: ['user-update.component.css']  
})
export class UserUpdateComponent{
    alertMsg = "";
    user:User;
    username: string = "";
    password: string = "";


     @ViewChild('userUpdateForm') userUpdateForm: NgForm;
	
	    constructor(private userService: UserService, private router:Router) {

    }
	ngOnInit() {
		if(!this.userService.getUserLoggedIn()){
            /*Should never happen*/
			this.router.navigate(['/login']);
        }
         this.user = {username: ""+this.userService.getUsername(),password: "", new_password:""};
        
	}
	onSubmit(userUpdateForm: NgForm) {
        if (this.userUpdateForm.invalid) return;
         this.userService.updateUser(this.user)
            .then(() => {

                this.alertMsg = "Your password was successfully updated!";
                this.userUpdateForm.resetForm();
                this.user.password="";
                this.user.new_password="";
                this.userUpdateForm.resetForm();
                
                
            },
            reason => {
                console.error(reason);
                this.alertMsg = "Whoops, something went wrong when updating password... Unexpected Error.";
                this.userUpdateForm.resetForm();
            }        
        );
    }
}