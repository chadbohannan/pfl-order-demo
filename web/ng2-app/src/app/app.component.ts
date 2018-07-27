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
  recipient: any;       // wizard step 3

  @ViewChild("wizard")
  public wizard: WizardComponent;

  // selectedProduct is an input parameter to the Product Details step
  productSelected(selection) {
    this.selectedProduct = selection;
    this.goToNextStep();
  }

  onProductConfig(configuration) {
    this.productDetails = configuration;
    this.goToNextStep();
  }

  onRecipientDetails(recipient) {
    this.recipient = recipient;
    this.goToNextStep();
  }

  goToNextStep() {
    this.wizard.navigation.goToNextStep();
  }

  goToPreviousStep() {
    this.wizard.navigation.goToPreviousStep();
  }
}
