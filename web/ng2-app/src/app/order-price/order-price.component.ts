import { Component, Input, OnChanges, OnInit } from '@angular/core';

@Component({
  selector: 'app-order-price',
  templateUrl: './order-price.component.html',
  styleUrls: ['./order-price.component.css']
})
export class OrderPriceComponent implements OnChanges, OnInit {
  @Input()
  order: any;

  response = "";

  constructor() { }

  ngOnInit() {
  }

  ngOnChanges() {
    // console.log('order changed, TODO lookup price')
  }

  orderText(): string {
    return JSON.stringify(this.order, null, 4);
  }
}
