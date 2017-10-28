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
        //this.user = {username: this.userService.getUsername,password: this.userService.getPassword, new_password:""};
        this.user = {username: ""+this.userService.getUsername(),password: "", new_password:""};
        
	}
	onSubmit(userUpdateForm: NgForm) {
       // if (this.userUpdateForm.invalid) return;
        //let user = { username: this.username, password: this.password };
        this.userService.updateUser(this.user)
            .then(() => {
                this.user.password="";
                this.user.new_password="";					
            },
            reason => {
                console.error(reason);
                this.alertMsg = "Whoops, something went wrong when updating password... Unexpected Error.";
            }        
        );
    }
}