import { Component, Input, OnChanges, OnInit } from '@angular/core';
import { Http, Headers } from '@angular/http';

@Component({
  selector: 'app-product-details',
  templateUrl: './product-details.component.html',
  styleUrls: ['../app.global.css', './product-details.component.css']
})
export class ProductDetailsComponent implements OnInit {

  @Input()
  product: any;  // TODO formalize Product as a class
  details: any; // should be a superset of product
  quantity = 0;
  fieldMap = {};
  fieldList = [];

  constructor(private http: Http) { }

  ngOnInit() {
  }

  ngOnChanges() {
    if (this.product.productID) {
      this.getProductDetails(this.product.productID);
      this.fieldMap = {};
      this.fieldList = [];
      this.quantity = this.product.quantityMinimum;
    }
  }

  // getProductList doesn't need any parameters
  getProductDetails(productID: string) {
    const url = '/api/products/' + productID;
    const headers = new Headers();
    headers.append('Content-Type', 'application/json');
    this.http.get(url, { headers: headers })
      .subscribe(response => {
        const details = response.json();
        if (details && 
          details.results &&
          details.results.data &&
          details.results.data.templateFields &&
          details.results.data.templateFields.fieldlist &&
          details.results.data.templateFields.fieldlist.field) {
          this.fieldList = details.results.data.templateFields.fieldlist.field;
        }
        console.log(details);
        this.details = details;
      }, error => {
        console.log('GET product' + productID + ' err:' + error.json());
      });
  }

  canIncrementQuantity() {
    const max = this.product.quantityMaximum || 0;
    if (max == 0) {
      return true;
    }
    
    const inc = this.product.quantityIncrement || 1;
    if ((this.quantity + inc) <= this.product.quantityMaximum) {
        return true;
    }
    return false;
  } 

  canDecrementQuantity() {
    const inc = this.product.quantityIncrement || 1;
    if ( (this.quantity - inc) >= this.product.quantityMinimum) {
        return true;
    }
    return false;
  }

  onIncrementQuantity() {
    const inc = this.product.quantityIncrement || 1;
    if (this.product.quantityIncrement) {
      this.quantity += inc;
    }
  }

  onDecrementQuantity() {
    const inc = this.product.quantityIncrement || 1;
    if (this.product.quantityIncrement) {
      this.quantity -= inc;
    }
  }

  setDefaultFieldValue(field) {
    if (!this.fieldMap[field.fieldname] || 
      this.fieldMap[field.fieldname].length == 0 ) {
      this.fieldMap[field.fieldname] = field.orgvalue;
    }
  }
}
