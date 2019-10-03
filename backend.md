# Software Engineer Challenge (Backend)

## Introduction

In Prophet, we are receiving wagers for sale both day and night. As a software engineer in Prophet, you have to provide a reliable backend system to clients. Your task here is to implement three endpoints to list/place/buy wagers.

## Requirements

1. We value a **clean**, **simple**, working solution.
2. The application must be run in Docker, candidate must provide `docker-composer.yml` and `start.sh` bash script at the root of the project, which should setup all relevant services/applications.
3. We prefer PHP (Laravel/Lumen), but the solution can also be written in Python or Node.js.
4. Candidates must submit the project as a git repository (github.com, bitbucket.com, gitlab.com). The repository must avoid containing the words `prophet`, `betprophet` and `challenge`.
5. Having unit/integration tests is a strong bonus.
6. As we run automated tests on your project, you must comply to the API requirement as stipulated below. You can assume Docker is already installed in the test machine.
7. The solution must be production ready.

## Problem Statement

1. Must be a RESTful HTTP API listening to port `8080` (or you can use another port instead and describe in the README)
2. The API must implement 3 endpoints with path, method, request and response body as specified
    - One endpoint to create a wager
        - To create a wager you must supply `total_wager_value`, `odds`, `selling_percentage`, `selling_price`
        - The API responds with an object of the wager

    - One endpoint to accept buying a wager, or part of a wager (fractional)
        - A purchase can be made multiple times against a single wager while `current_selling_price` is still positive.

    - One endpoint to list wagers
3. The request input should be validated before processing. The server should return proper error responses in case validation fails.
4. A database must be used (at Prophet we typically use MySQL). Database installation and initialization must be done in `start.sh`
5. All responses must be in json format for success and failure responses.

## API Interface

You are expected to follow the API specification as follows. Your implementation should not have any deviations on the method, URI path, request and response body. Such alterations may cause our automated tests to fail.

#### Place Wager

- Method: `POST`
- URL path: `/wagers`
- Request body:
    ```
    {
        "total_wager_value": <total_wager_value>,
        "odds": <odds>,
        "selling_percentage": <selling_percentage>,
        "selling_price": <selling_price>,
    }
    ```

- Response:
    Header: `HTTP 201`
    Body:
    ```
    {
        "id": <wager_id>,
        "total_wager_value": <total_wager_value>,
        "odds": <odds>,
        "selling_percentage": <selling_percentage>,
        "selling_price": <selling_price>,
        "current_selling_price": <current_selling_price>,
        "percentage_sold": <percentage_sold>,
        "amount_sold": <amount_sold>,
        "placed_at": <placed_at>
    }
    ```
    or

    Header: `HTTP <HTTP_CODE>`
    Body:
    ```
    {
        "error": "ERROR_DESCRIPTION"
    }
    ```

- Requirements:

    - `total_wager_value` must be specified as a positive integer above 0
    - `odds` must be specified as a positive integer above 0
    - `selling_percentage` must be specified as an integer between 1 and 100
    - `selling_price` must be specified as a positive decimal value to two decimal places, it is a monetary value
    - `selling_price` must be greater than `total_wager_value` * (`selling_percentage` / 100)
    - `id` should be an auto increment field
    - `current_selling_price` should be the `selling_price` until a `Buy Wager` action is taken against this wager record
    - `percentage_sold` should be null until a `Buy Wager` action is taken against this wager record
    - `amount_sold` should be null until a `Buy Wager` action is taken against this wager record
    - `placed_at` should be a timestamp at the completion of the request


#### Buy wager

- Method: `POST`
- URL path: `/buy/:wager_id`
- Request body:
    ```
    {
        "buying_price": <buying_price>
    }
    ```

- Response:
    Header: `HTTP 201`
    Body:
    ```
    {
        "id": <purchase_id>,
        "wager_id": <wager_id>,
        "buying_price": <buying_price>,
        "bought_at": <bought_at>
    }
    ```
    or

    Header: `HTTP <HTTP_CODE>`
    Body:
    ```
    {
        "error": "ERROR_DESCRIPTION"
    }
    ```

- Requirements:
    - `buying_price` should be an positive decimal value
    - `buying_price` must be lesser or equal to `current_selling_price` of the `wager_id`
    - A successful purchase should update the wager fields `current_selling_price`, `percentage_sold`, `amount_sold`
    - `id` should be an auto increment field
    - `bought_at` should be a timestamp at completion of the request


#### Wager list

- Method: `GET`
- URL path: `/wagers?page=:page&limit=:limit`
- Response:
    Header: `HTTP 200`
    Body:
    ```
    [
        {
            "id": <wager_id>,
            "total_wager_value": <total_wager_value>,
            "odds": <odds>,
            "selling_percentage": <selling_percentage>,
            "selling_price": <selling_price>,
            "current_selling_price": <current_selling_price>,
            "percentage_sold": <percentage_sold>,
            "amount_sold": <amount_sold>,
            "placed_at": <placed_at>
        }
        ...
    ]
    ```

Questions? We love to answer: techchallenge@betprophet.co
