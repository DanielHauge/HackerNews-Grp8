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
    alertMsg = "";
    
    username: string = "";
    password: string = "";

    reg_username: string = "";
    reg_password: string = "";
    reg_email_addr: string = "";
    @ViewChild('loginForm') loginForm: NgForm;
	
	    constructor(private userService: UserService, private router:Router) {

    }
	ngOnInit() {
		if(this.userService.getUserLoggedIn()){
			//this.router.navigate(['/login']);
            this.userService.setUserLoggedOut()
		}
		
	}
	onSubmit(loginForm: NgForm) {
       // if (this.loginForm.invalid) return;
        let user = { username: this.username, password: this.password };
        this.userService.loginUser(user)
            .then((response) => {
                console.log(response);
                this.userService.setUsername(this.username);
                this.userService.setUserLoggedIn();
                    this.loginForm.resetForm();
					this.router.navigate(['/threads']);
					
            },
             reason => {
                console.warn(reason);
                if(reason){
                    console.log(reason);
                    this.alertMsg = "Whoops... could not log in.";
                    
                }
               
                this.loginForm.resetForm();
                
            })
            .catch(	response => { 			
                    console.error(response);
                    this.alertMsg = "Whoops... something went wrong";
                    this.loginForm.resetForm();
            });
    }
    onSubmit2(registerForm: NgForm) {
        console.log('register users');
         let user = { username: this.reg_username, password: this.reg_password , email_addr: this.reg_email_addr};
         this.userService.registerUser(user)
             .then((response) => {
                 this.userService.setUsername(this.reg_username);
                 this.userService.setUserLoggedIn();
                    // this.registerForm.resetForm();
                 this.router.navigate(['/threads']);
                     
             },
             reason => {
                console.warn(reason);
                if(reason.status){
                    this.alertMsg = "Please type a valid username or password";
                }
                else{
                    this.alertMsg = "Registration failed.";
                    
                }                
                this.loginForm.resetForm();
                
            })
            .catch(	response => { 			
                    console.error(response);
                    this.alertMsg = "Whoops... something went wrong";
                    this.loginForm.resetForm();
            });
     }
}