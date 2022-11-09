import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-codis',
  templateUrl: './codis.component.html',
  styleUrls: ['./codis.component.css']
})
export class CodisComponent implements OnInit {

  constructor() { }

  ngOnInit(): void {
    console.log("laile codis")
  }

}
