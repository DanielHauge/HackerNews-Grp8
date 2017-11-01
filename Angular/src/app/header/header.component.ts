import { Component, OnInit } from '@angular/core';
import { UserService } from '../login/shared/user.service';

@Component({
  selector: 'app-header',
  templateUrl: './header.component.html',
  styleUrls: ['./header.component.css']
})
export class HeaderComponent implements OnInit {

  constructor(private userService: UserService) { }
  username:string;
  loginText:string;
  getUserName (){ return this.userService.getUsername()};
  getUserNameText (){ return this.userService.getUsernameText()};
  
  ngOnInit() {

    
  }

}



