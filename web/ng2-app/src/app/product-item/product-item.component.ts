import { Component, Input, OnInit } from '@angular/core';

@Component({
  selector: 'app-product-item',
  templateUrl: './product-item.component.html',
  styleUrls: ['../app.global.css', './product-item.component.css']
})
export class ProductItemComponent implements OnInit {

  @Input()
  product: any;  // TODO formalize Product as a class

  constructor() { }

  ngOnInit() { }

}
