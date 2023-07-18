# faker
Golang Library to fake personal information

## Installing
```bash
go get -u github.com/joram/faker/faker
```

## Functions

### Names
#### `GetMaleFirstName`
returns a random male first name based on the weights provided by stats Canada.
#### `GetFemaleFirstName`
returns a random female first name based on the weights provided by stats Canada.
#### `GetFirstName`
returns a random first name based on the weights provided by stats Canada.
#### `GetLastName`
returns a random last name based on the weights provided by stats Canada.

### Address
#### `GetAddress`
returns a random Canadian address based off of canadian postal code information.

### Age
#### `GetAge`
returns a random age from 0 to 100 based on stats canada historical data.
#### `GetBirthday`
returns a random birthday based on stats canada historical data.

### Credit Card
#### `GetCreditCard`
returns a random Visa or Mastercard

### Email
#### `GetEmail`
returns a random email address with a com,net,org,ca top level domain.

### IP Address
#### `GetIPAddress`
returns a random IP address

### Phone
#### `GetPhone`
returns a random phone number

## Data Sources
- stats canada
- https://github.com/djbelieny/geoinfo-dataset/blob/master/full_dataset_csv.zip