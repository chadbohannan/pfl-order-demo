import { Component, ViewChild } from '@angular/core';
import { WizardComponent } from 'angular-archwizard';


@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.global.css', './app.component.css']
})
export class AppComponent {

  // TODO formalize Product model as a class

  selectedProduct: any; // wizard step 1
  productDetails: any;  // wizard step 2
  quantity = 0;
  shippingMethod = "";
  recipient: any;       // wizard step 3
  order: any;           // wizard step 4

  @ViewChild("wizard")
  public wizard: WizardComponent;

  // selectedProduct is an input parameter to the Product Details step
  productSelected(selection) {
    this.selectedProduct = selection;
    this.quantity = 4;
    this.goToNextStep();
  }

  onProductConfig(configuration) {
    this.productDetails = configuration.details;
    this.quantity = configuration.quantity;
    this.shippingMethod = configuration.shippingMethod;
    this.goToNextStep();
  }

  onRecipientDetails(recipient) {
    this.recipient = recipient;

    // TODO compose an order object
    const templateData = [];
    Object.keys(this.productDetails).forEach(element => {
      console.log(element);
      templateData.push({
        templateDataName: "CompanyName",
        templateDataValue: "Colorful Baloons Ltd."
      });
    });
    

    this.order = {
      orderCustomer: {
        firstName: "Hard",
        lastName:  "Coded",
        address1: "12345 Main St",
        address2: "Suite 0",
        city: "Bozeman",
        state: "Mt",
        postalCode: "59715",
        countryCode: "US",
        phone: "1234567890",
        email: "a@b.com"
      },
      items: [
        {
          itemSequenceNumber: 1,
          productID: this.selectedProduct.productID,
          quantity: this.quantity,
          templateData: templateData
        },
      ],
      shipments: [
        Object.assign({
          shipmentSequenceNumber: 1,
          shippingMethod: this.shippingMethod
        }, this.recipient)
      ]
    };
    this.goToNextStep();
  }

  goToNextStep() {
    this.wizard.navigation.goToNextStep();
  }

  goToPreviousStep() {
    this.wizard.navigation.goToPreviousStep();
  }
}
