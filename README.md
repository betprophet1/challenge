# Prophet Engineering Challenge
## Technical stack
| Technology | version | Description |
|------------|---------|-------------|
| gORM       | latest  | Object Relational Mapping to interact with relational database |
| gorilla/mux| latest  | To implement rest API for golang |
| golang     | 1.16    | golang language|
| docker     | latest  | it's required to build and run app |
| docker-comose| latest | it's required to build and run app |
# How to run
```
$docker-compose up
```
In case, Your machine has already setup Make file. Could you run command below
```
$make docker-compose-up
```
#Test APIs
1. Place wager
should you use both postman and curl to test. In case, I will provide you a curl command as below
###Request
```
$curl --location --request POST 'http://localhost:8080/wagers' \
--header 'Content-Type: application/json' \
--data-raw '{
    "total_wager_value": 50.25,
    "odds": 30,
    "selling_percentage": 20,
    "selling_price": 40.25
}'
```
### Response
#### HTTP 200 OK
```
{
    "id": 2,
    "total_wager_value": 50.25,
    "odds": 30,
    "selling_percentage": 20,
    "selling_price": 40.25,
    "current_selling_price": 40.25,
    "percentage_sold": 0,
    "amount_sold": 0,
    "placed_at": "2022-05-22T09:51:37.445Z"
}
```
#### Http 500 Internal server error
```
{
    "error": "Body must not be null or empty"
}
```
2. Buy
###Request
#### HTTP 200 Ok
```
curl --location --request POST 'http://localhost:8080/buy/2' \
--header 'Content-Type: application/json' \
--data-raw '{
    "buying_price": 30.05
}'
```
#### HTTP 500 INTERNAL SERVER ERROR
```
{
    "error": "Buying Price must be lesser or equal current selling price"
}
```
3. Get List Wager
### Request
```
curl --location --request GET 'http://localhost:8080/wagers?page=1&limit=200'
```
### Response
#### HTTP 200 OK
```
{
    "limit": 100,
    "page": 1,
    "sort": "Id desc",
    "total_rows": 2,
    "total_pages": 1,
    "rows": [
        {
            "id": 2,
            "total_wager_value": 50.25,
            "odds": 30,
            "selling_percentage": 20,
            "selling_price": 40.25,
            "current_selling_price": 30.05,
            "percentage_sold": 59.801,
            "amount_sold": 1,
            "placed_at": "2022-05-22T09:51:37.445Z"
        },
        {
            "id": 1,
            "total_wager_value": 50.25,
            "odds": 30,
            "selling_percentage": 20,
            "selling_price": 40.25,
            "current_selling_price": 40.25,
            "percentage_sold": 0,
            "amount_sold": 0,
            "placed_at": "2022-05-22T09:48:35.739Z"
        }
    ]
}
```
