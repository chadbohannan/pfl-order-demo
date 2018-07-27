import { Component, OnInit, Output, EventEmitter} from '@angular/core';

@Component({
  selector: 'app-user-details',
  templateUrl: './user-details.component.html',
  styleUrls: ['../app.global.css', './user-details.component.css']
})
export class UserDetailsComponent implements OnInit {

  @Output()
  values = new EventEmitter();

  details = {};


  constructor() {
    [
      "firstName",
      "lastName",
      "companyName",
      "address1",
      "address2",
      "city",
      "state",
      "postalCode",
      "countryCode",
      "email",
      "phone"
    ].forEach( param => this.details[param] = "");
  }

  ngOnInit() {
  }

  emitValues() {
    this.values.emit(this.details);
  }

}
