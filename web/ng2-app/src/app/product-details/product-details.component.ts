import { Component, Input, OnInit } from '@angular/core';

@Component({
  selector: 'app-product-details',
  templateUrl: './product-details.component.html',
  styleUrls: ['../app.global.css', './product-details.component.css']
})
export class ProductDetailsComponent implements OnInit {

  @Input()
  product: any;  // TODO formalize Product as a class

  constructor() { }

  ngOnInit() {
  }

}
