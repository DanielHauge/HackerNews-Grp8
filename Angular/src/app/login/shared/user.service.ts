import { Injectable } from '@angular/core';
import { Http } from '@angular/http';

@Injectable()
export class UserService {
	private isUserLoggedIn: boolean;
	private username:string;
    constructor(private http: Http) {
		this.isUserLoggedIn = false;
    }

    loginUser(user: { username: string; password: string; }) {
		console.log(user);
		this.setUsername(user.username);
		this.setUserLoggedIn();
        return this.http.post(`/app/threads/`, user)
            .toPromise();
    }
	setUserLoggedIn(){
		this.isUserLoggedIn = true;
	}
	getUserLoggedIn():boolean{
		return this.isUserLoggedIn;
	}
	setUsername(name:string){
		this.username = name;
	}
	getUsername():string{
		return this.username;
	}
	getUsernameText():string{
		if(this.getUserLoggedIn()){
			return 'Log out';
		}
		else{
			return 'Log in';
		}
	}

   
}