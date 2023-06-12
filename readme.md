
# Request For Quotation (RFQ) using Cosmos SDK

![Request For Quotation](https://github.com/anuragpal/quotation-mvp/blob/master/rfq.png)

## Request For Quotation (RFQ)

A request for quotation (RFQ) is a document that a company or organization sends to potential suppliers to request quotes for products or services. RFQs are often used when a company or organization is looking to purchase a large quantity of goods or services, or when they need to compare prices from multiple suppliers.

### Built With
* [Golang (go1.20.4)](https://go.dev)
* [Cosmos SDK (v0.46.7)](https://cosmos.network)
* [Ignite Cli (v0.26.1)](https://ignite.com)

### Installation
```sh
git clone https://github.com/anuragpal/quotation-mvp.git
cd quotation-mvp
ignite chain serve
```

### Commands
**Request Quotation**
```sh
quotation-mvpd tx quotation request-quotation "{\"id\":\"ba78ec76-3980-425e-9053-95ce86d9ce52\",\"tx\":{\"data\":\"[{\\\"product_name\\\": \\\"Apple Mac Pro\\\",\\\"quantity\\\": 5},{\\\"product_name\\\": \\\"Hp Spectra\\\",\\\"quantity\\\": 12},{\\\"product_name\\\": \\\"Dell Gaming Laptop\\\",\\\"quantity\\\": 5},{\\\"product_name\\\": \\\"Dell Vostro\\\",\\\"quantity\\\": 8}]\",\"sender\":\"importer@kadmine.com\",\"receiver\":[\"developer.anurag+1@gmail.com\",\"developer.anurag+2@gmail.com\",\"developer.anurag+3@gmail.com\",\"developer.anurag+4@gmail.com\"]},\"envelope\":{\"event\":\"request-quotation\",\"isEncrypt\":false,\"timestamp\":\"\"},\"version\":\"1.0\"}" --from alice
```

**Send Porposal**
```sh
quotation-mvpd tx quotation send-proposal "{\"id\":\"eb348c5b-a018-4a6e-ac42-58fcd8c217b3\",\"tx\":{\"rqId\":\"ba78ec76-3980-425e-9053-95ce86d9ce52\",\"data\":\"[{\\\"product_name\\\":\\\"Apple Mac Pro\\\",\\\"quantity\\\":5,\\\"price\\\":875000},{\\\"product_name\\\":\\\"Hp Spectra\\\",\\\"quantity\\\":12,\\\"price\\\":600000},{\\\"product_name\\\":\\\"Dell Gaming Laptop\\\",\\\"quantity\\\":5,\\\"price\\\":360000},{\\\"product_name\\\":\\\"Dell Vostro\\\",\\\"quantity\\\":8,\\\"price\\\":560000}]\",\"sender\":\"importer@kadmine.com\",\"receiver\":\"developer.anurag+1@gmail.com\"},\"envelope\":{\"event\":\"send-proposal\",\"isEncrypt\":false,\"timestamp\":\"\"},\"version\":\"1.0\"}" --from alice
```

**Accept/Reject Proposal**
```sh
quotation-mvpd tx quotation accept-or-reject-proposal "{\"id\":\"eb348c5b-a018-4a6e-ac42-58fcd8c217b3\",\"tx\":{\"action\":true},\"envelope\":{\"event\":\"accept-or-reject-quotation\",\"isEncrypt\":false,\"timestamp\":\"\"},\"version\":\"1.0\"}" --from alice
```


## Authors
- [@anuragpal](https://www.github.com/anuragpal)


## Reference
- [RFQ](https://sievo.com/blog/the-simple-request-for-quotation-rfq-process-for-procurement)