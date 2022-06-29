# Prophet Engineering Challenge
## Technical stack
| Technology     | version | 
|----------------|---------|
| postgres       | latest  | 
| Laravel        | 9.x     |
| docker         | latest  |
| docker-compose | latest  |
# How to run

#### Run docker

```text
docker-compose up -d --build
```
#### Create database
```text
docker-compose postgres sh
su - postgres
psql
CREATE DATABASE laravel;
```
#### Run laravel
```
$./start.sh
```
#Test APIs
1. Place wager
   should you use both postman and curl to test. In case, I will provide you a curl command as below
   ###Request
```
$curl --location --request POST 'http://laravel.localhost:22080/wagers' \
--header 'Content-Type: application/json' \
--data-raw '{
    "total_wager_value": 50,
    "odds":     30,
    "selling_percentage": 10,
    "selling_price": 50
}'
```
### Response
#### HTTP 200 OK
```
{
    "id": 11,
    "total_wager_value": 50,
    "odds": 30,
    "selling_percentage": 10,
    "selling_price": 50,
    "current_selling_price": 0,
    "percentage_sold": 0,
    "amount_sold": 0,
    "placed_at": "2022-06-29T15:40:25.276613Z"
}
```

2. Buy
   ###Request
```
curl --location --request POST 'http://laravel.localhost:22080/buy/2' \
--header 'Content-Type: application/json' \
--data-raw '{
    "buying_price": 10
}'
```

### Response
#### HTTP 200 OK
```
{
    "id": 4,
    "wager_id": "12",
    "buying_price": 10,
    "bought_at": "2022-06-29T16:28:12.350366Z"
}
```
#### HTTP 500 INTERNAL SERVER ERROR
```
{
    "errors": {
        "buying_price": [
            "The buying price must be lesser or equal to current_selling_price of the wager_id."
        ]
    }
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
    "items": [
        {
            "id": 12,
            "total_wager_value": "50",
            "odds": "30",
            "selling_percentage": "10",
            "selling_price": "50",
            "current_selling_price": "10",
            "percentage_sold": "20",
            "amount_sold": "1",
            "placed_at": "2022-06-29 16:23:57"
        },
        {
            "id": 11,
            "total_wager_value": "50",
            "odds": "30",
            "selling_percentage": "10",
            "selling_price": "50",
            "current_selling_price": "0",
            "percentage_sold": "0",
            "amount_sold": "0",
            "placed_at": "2022-06-29 15:40:25"
        },
        {
            "id": 10,
            "total_wager_value": "50",
            "odds": "30",
            "selling_percentage": "10",
            "selling_price": "40",
            "current_selling_price": "80",
            "percentage_sold": "160",
            "amount_sold": "3",
            "placed_at": "2022-06-29 14:58:59"
        },
        {
            "id": 9,
            "total_wager_value": "50",
            "odds": "30",
            "selling_percentage": "10",
            "selling_price": "40.25",
            "current_selling_price": "0",
            "percentage_sold": "0",
            "amount_sold": "0",
            "placed_at": "2022-06-28 09:59:05"
        },
        {
            "id": 8,
            "total_wager_value": "50",
            "odds": "30",
            "selling_percentage": "10",
            "selling_price": "40.25",
            "current_selling_price": "0",
            "percentage_sold": "0",
            "amount_sold": "0",
            "placed_at": "2022-06-28 09:55:11"
        },
        {
            "id": 7,
            "total_wager_value": "50.25",
            "odds": "30",
            "selling_percentage": "20",
            "selling_price": "40.25",
            "current_selling_price": "0",
            "percentage_sold": "0",
            "amount_sold": "0",
            "placed_at": "2022-06-28 04:12:41"
        },
        {
            "id": 6,
            "total_wager_value": "50.25",
            "odds": "30",
            "selling_percentage": "20",
            "selling_price": "40.25",
            "current_selling_price": "0",
            "percentage_sold": "0",
            "amount_sold": "0",
            "placed_at": "2022-06-28 04:09:37"
        },
        {
            "id": 5,
            "total_wager_value": "50.25",
            "odds": "30",
            "selling_percentage": "20",
            "selling_price": "40.25",
            "current_selling_price": "0",
            "percentage_sold": "0",
            "amount_sold": "0",
            "placed_at": "2022-06-28 04:05:07"
        },
        {
            "id": 2,
            "total_wager_value": "50.25",
            "odds": "30",
            "selling_percentage": "20",
            "selling_price": "40.25",
            "current_selling_price": "0",
            "percentage_sold": "0",
            "amount_sold": "0",
            "placed_at": "2022-06-28 04:02:08"
        },
        {
            "id": 1,
            "total_wager_value": "50.25",
            "odds": "30",
            "selling_percentage": "20",
            "selling_price": "40.25",
            "current_selling_price": "0",
            "percentage_sold": "0",
            "amount_sold": "0",
            "placed_at": "2022-06-28 03:59:50"
        }
    ],
    "total": 10,
    "page": 1,
    "last_page": 1,
    "count": 10
}
```

# I guess your challenge wrong. Because current_selling_price is 0, but buying_price must be lesser or equal to current_selling_price of the wager_id so buying_price always is 0. Please check it.