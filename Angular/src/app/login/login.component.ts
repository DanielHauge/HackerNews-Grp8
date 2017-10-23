import { Component, OnInit, OnDestroy } from '@angular/core';
import { ActivatedRoute } from "@angular/router";


@Component({
  selector: 'app-login',
  templateUrl: 'login.component.html',
  styleUrls: ['login.component.css']  
})
export class LoginComponent implements OnInit, OnDestroy {


  constructor(private route:ActivatedRoute) { }

  ngOnInit() {
    

  }

  ngOnDestroy(): void {
  }
}