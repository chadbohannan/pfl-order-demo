# pfl-order-demo

[Live Instance](https://pfl-order-demo.appspot.com)

## Task
Using the PFL API create an application that calls out to our API to retrieve 
a list of products available displays the products in a user friendly manner, 
giving the user a way to select which product for which they would like to 
place an order (some products have a template consisting of fields that should
be input by the user) submits an order to our TEST API, displaying the returned
order number to the user after submission.

API Documentation: https://www.printingforless.com/partners/api/APIs/Getting_Started


## Breakdown

The scope of the task appears to be narrow enough that I can ignore the greater 
context of a coherient website and present an interface that can compose an
order that is accepted by the PFL API.

### Workflow
From reviewing the API documentation I've inferred a plausible workflow:

 1. List Products
   * [Get Products](#get-products) 
 2. Select Product
   * [Get Product Details](#get-product-details) 
 3. Customize Product
   * Custom product parameter tempates
   * Upload PDFs or use set of static files
 4. Input Recipient Data
   * Name
   * Address
   * Phone
   * Email
 5. Input Customer Data
   * Name
   * Address
   * Phone
   * Email
 6. Preview Order
   * [Post Price Lookup](#post-price-lookup)
 7. Place Order
   * [Post Order](#post-order)
 8. Review Order

To simplify the User|Account Managment features that would be associated with 
the greater context of a website the server could automatically generate a new
User entity with a 1-to-1 relationship with a broser session.

Extending the application to support multiple sessions for a User likely
involves new account setup workflows with password persistence and email 
verification and is assumed to be outside the scope of this task.

### Media Upload

The [Post Order](#post-order) interface expects a URI for the PDF content.
A custom upload interface is certainly reasonable but would consume time.
I could use some static PDFs instead.

## API

### Get Products

 ```
 GET https://testapi.pfl.com/products?apikey=123
 {
    "results": {
        "errors": [],
        "messages": [],
        "data": [
            {
                "id":                int,
                "productID":         int,
                "name":              string,
                "description":       string,
                "imageURL":          string,
                "hasTemplate":       boolean,
                "quantityDefault":   int,
                "quantityIncrement": int,
                "quantityMaximum":   int,
                "quantityMinimum":   int,
            }
        ]
    }
 }
 ```

### Get Product Details

```
GET https://testapi.pfl.com/products?apikey=123&id=123
{
    "results": {
        "errors": [],
        "messages": [],
        "data": [
            {
                "id":                int,
                "productID":         int,
                "name":              string,
                "description":       string,
                "imageURL":          string,
                "hasTemplate":       boolean,
                "quantityDefault":   int,
                "quantityIncrement": int,
                "quantityMaximum":   int,
                "quantityMinimum":   int,
                shippingMethodDefault": string,
                "deliveredPrices":[  
                    {  
                        "deliveryMethodCode": string,
                        "description": string,
                        "isDefault": boolean,
                        "locationType": string,
                        "price": float,
                        "country": null,
                        "countryCode": null,
                        "created": string,
                    },
                ],
                "templateFields": {
                    "fieldlist": {
                        "field": [
                            {
                                "required": string, bool (Y | N)
                                "visible": string,  bool (Y | N)
                                "type": string, enum (SINGLELINE | MULTILINE | GRAPHICUPLOAD)
                                "linelimit": string, integer
                                "fieldname": string,
                                "prompt": [
                                    {
                                        "language": string, e.g "en-US"
                                        "text": string
                                    }
                                ],
                                "default": string,
                                "orgvalue": string,
                                "htmlfieldname": string
                            }
                        ]
                    }
                }
            }
        ],
    }
}
```

### Post Order
[apidoc]: https://www.printingforless.com/partners/api/APIs/Orders_API/Create_Order_Examples/Partner_Provides_Template_Data
APIs: Orders API: Create Order Examples: (Partner Provides Template Data)[apidoc]

```
POST https://testapi.pfl.com/products?apikey=123&id=123
{
    "partnerOrderReference": "MyReferenceNumber",
    "orderCustomer": {  
        "firstName":   string,
        "lastName":    string,
        "companyName": string,
        "address1":    string,
        "address2":    string,
        "city":        string,
        "state":       string,
        "postalCode":  string,
        "countryCode": string,
        "email":       string,
        "phone":       string
    },
    "items": [
        {
            "itemSequenceNumber":   int,
            "productID":            int,
            "quantity":             int,
            "productionDays":       int,
            "partnerItemReference": string,
            "itemFile":             string e.g. "http://www.x.com/files/art.pdf"
        }
    ],
    "shipments": [
        {
            "shipmentSequenceNumber": int,
            "firstName":              string,
            "lastName":               string,
            "companyName":            string,
            "address1":               string,
            "address2":               string,
            "city":                   string,
            "state":                  string,
            "postalCode":             string,
            "countryCode":            string,
            "phone":                  string,  
            "shippingMethod":         string, // "FDXG"
            "IMBSerialNumber":        string, // "004543450"
        }  
    ],
    "payments":[
        {
            "paymentMethod": string, // "FDXG"
            "paymentID":     string,
            "paymentAmount": float   // 3.00
        }
    ],
    "billingVariables":[
        {
            "key":   string, // "BillingVariable1Name"
            "value": string  // "BillingVariable1Value"
        }
    ]
}
```

### Post Price Lookup
```
POST https://testapi.pfl.com/products?apikey=123&id=123
```
Same as #(post-order)


## Build dependencies
Google Cloud SDK
Go >= 1.8
npm
angular-cli

## Build

TODO

## Deploy

```./shipit.sh```

## Configure

TODO setting PFL credentials
