import { Component, ViewChild } from '@angular/core';
import { NgForm } from '@angular/forms';
import { UserService } from './shared/user.service';
import { Router } from '@angular/router'


@Component({
  selector: 'app-login',
  templateUrl: 'login.component.html',
  styleUrls: ['login.component.css']  
})
export class LoginComponent{

    username: string = "";
    password: string = "";
     @ViewChild('loginForm') loginForm: NgForm;
	
	    constructor(private userService: UserService, private router:Router) {

    }

	onSubmit(loginForm: NgForm) {
       // if (this.loginForm.invalid) return;
        let user = { username: this.username, password: this.password };
        this.userService.loginUser(user)
            .then(() => {
                    this.loginForm.resetForm();
					this.router.navigate(['/threads']);
					
            });
    }
}