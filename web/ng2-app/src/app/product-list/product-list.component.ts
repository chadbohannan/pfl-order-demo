import { Component, OnInit, Output, EventEmitter } from '@angular/core';
import { Http, Headers } from '@angular/http';

@Component({
  selector: 'app-product-list',
  templateUrl: './product-list.component.html',
  styleUrls: ['../app.global.css', './product-list.component.css']
})
export class ProductListComponent implements OnInit {
  @Output()
  selection = new EventEmitter();

  productData = {
    "results": {
      "errors": [],
      "messages": [],
      "data": [
        {
          "id": 0,
          "productID": 0,
          "name": "Loading....",
          "description": "Please be patient",
          "imageURL": "/assets/loading.gif",
          "hasTemplate": true,
          "quantityDefault": 10,
          "quantityIncrement": 10,
          "quantityMaximum": 100,
          "quantityMinimum": 10,
        }
      ]
    }
  };

  constructor(private http: Http) { }

  ngOnInit() {
    const url = '/api/products'
    const headers = new Headers();
    headers.append('Content-Type', 'application/json');
    this.http.get(url, { headers: headers })
      .subscribe(response => {
        this.productData = response.json();
      }, error => {
        console.log('GET products err:' + error.json());
      });
  }

  onSelect(i: number) {
    console.log('selected:' + this.productData.results.data[i].name);
    this.selection.emit(this.productData.results.data[i]);
  }


}
