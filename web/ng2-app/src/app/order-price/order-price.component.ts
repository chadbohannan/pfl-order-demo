import { Component, Input, OnChanges, OnInit } from '@angular/core';
import { Http, Headers } from '@angular/http';

@Component({
  selector: 'app-order-price',
  templateUrl: './order-price.component.html',
  styleUrls: ['./order-price.component.css']
})
export class OrderPriceComponent implements OnChanges, OnInit {
  @Input()
  order: any;

  response: any;
  itemPrice: any;

  constructor(private http: Http) { }

  ngOnInit() { }

  ngOnChanges() {
    if (this.order) {
      this.postPriceQuery();
    }
  }

  orderText(): string {
    return JSON.stringify(this.order, null, 4);
  }

  responseText(): string {
    if (this.response &&
      this.response.results &&
      this.response.results.data &&
      this.response.results.data.items &&
      this.response.results.data.items.length > 0) {
      this.itemPrice = this.response.results.data.items[0].itemPrice;
      return JSON.stringify(this.itemPrice, null, 4);
    }
    return JSON.stringify(this.response, null, 4);
  }

  postPriceQuery() {
    const url = '/api/price';
    const headers = new Headers();
    headers.append('Content-Type', 'application/json');
    const body = JSON.stringify(this.order);
    const that = this;
    this.http.post(url, body, { headers: headers })
      .subscribe(response => {
        const obj = response.json();
        if (obj &&
          obj.results &&
          obj.results.data &&
          obj.results.data.items &&
          obj.results.data.items.length > 0) {
          that.itemPrice = obj.results.data.items[0].itemPrice;
          that.response = null;
        } else {
          that.response = response;
        }
      }, error => {
        that.response = "Error: " + error.text();
      });
  }
}
