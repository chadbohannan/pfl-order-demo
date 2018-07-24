import { Component, OnInit, Output, EventEmitter } from '@angular/core';

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
          "id": 1,
          "productID": 2,
          "name": "uploads!",
          "description": "glitter",
          "imageURL": "https://biglittleidea.org/api/v1/media/5367079314128896/file",
          "hasTemplate": true,
          "quantityDefault": 10,
          "quantityIncrement": 10,
          "quantityMaximum": 100,
          "quantityMinimum": 10,
        },
        {
          "id": 3,
          "productID": 4,
          "name": "art of war",
          "description": "video game",
          "imageURL": "https://biglittleidea.org/api/v1/media/5646620347596800/file",
          "hasTemplate": true,
          "quantityDefault": 10,
          "quantityIncrement": 10,
          "quantityMaximum": 100,
          "quantityMinimum": 10,
        }
      ]
    }
  };

  constructor() { }

  ngOnInit() { }


  onSelect(i: number) {
    console.log('selected:' + this.productData.results.data[i].name);
    this.selection.emit(this.productData.results.data[i]);
  }
}
