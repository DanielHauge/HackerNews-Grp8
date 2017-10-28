import { Injectable } from '@angular/core';
//import { Http } from '@angular/http';
import {Http, Headers, RequestOptions,Response} from '@angular/http';
import { User } from './user.model';

@Injectable()
export class UserService {
	private isUserLoggedIn: boolean;
	private username:string;
	private password:string;
    constructor(private http: Http) {
		this.isUserLoggedIn = false;
    }

    loginUser(user: { username: string; password: string; }) {
		console.log(user);
		console.log('http');
		let headers = new Headers();
		headers.append('Content-Type', 'text/plain');
		let options = new RequestOptions({ headers: headers });
		
        return this.http.post('http://165.227.151.217:9191/login', user, options)
			.toPromise()
			.then( response => { 
				console.error(response);
				console.error(response);
				
				 return response.json();
			});
			/*.catch(	response => { 			console.error(response);
			
			 return response.json();
			}
			);*/
			

	}
	registerUser(user: { username: string; password: string; }) {
		console.log(user);
		console.log('http');
		let headers = new Headers();
		headers.append('Content-Type', 'text/plain');
		let options = new RequestOptions({ headers: headers });
		
        return this.http.post('http://165.227.151.217:9191/create', user, options)
            .toPromise();
	}
	updateUser(user:User) {
		console.log('updateUser..');
		let headers = new Headers();
		headers.append('Content-Type', 'text/plain');
		let options = new RequestOptions({ headers: headers });
		
        return this.http.post('http://165.227.151.217:9191/update', user, options)
            .toPromise();
    }
	setUserLoggedIn(){
		this.isUserLoggedIn = true;
	}
	setUserLoggedOut(){
		this.isUserLoggedIn = false;
		this.username = '';
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
	setPassword(password:string){
		this.password = password;
	}
	getPassword():string{
		return this.password;
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