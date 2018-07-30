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
  quantity = 0;

  priceAttributes =[
    ["envelopePrice", "Envelope Price"],
    ["mailingPrice", "Mailing Price"],
    ["promotionalDiscount", "Promotional Discount"],
    ["retailFulfillmentPrice", "Retail Fulfillment Price"],
    ["retailPrintPrice", " Retail Print Price"],
    ["retailReimbursementPrice", "Retail Reimbursement Price"],
    ["retailRushPrice", "Retail Rush Price"],
    ["retailShippingPrice", "Retail Shipping Price"],
    ["secondSheetPrice", "Second Sheet Price"],
    ["printingCostEach", "Printing Cost Each"],
    ["printPrice", "Print Price"],
    ["rushPrice", "Rush Price"],
    ["shipPrice", "Ship Price"],
    ["totalPrintingPrice", "Total Shipping Price"],
  ];

  printAttr(attr){
    console.log(attr);
  }

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
          that.quantity = obj.results.data.items[0].quantity;
          that.response = null;
        } else {
          that.response = response;
        }
      }, error => {
        that.response = "Error: " + error.text();
      });
  }
}
