import { Component, ViewChild } from '@angular/core';
import { WizardComponent } from 'angular-archwizard';


@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.global.css', './app.component.css']
})
export class AppComponent {

  // TODO formalize Product model as a class

  selectedProduct = {};                       // wizard step 1
  productDetails = new Map<string, string>(); // wizard step 2

  @ViewChild("wizard")
  public wizard: WizardComponent;

  // selectedProduct is an input parameter to the Product Details step
  productSelected(selection) {
    console.log("productSelected:" + selection.name);
    this.selectedProduct = selection;
    this.goToNextStep();
  }

  goToNextStep() {
    this.wizard.navigation.goToNextStep();
  }

  goToPreviousStep() {
    this.wizard.navigation.goToPreviousStep();
  }
}
