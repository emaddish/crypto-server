# Crypto Server

Server to get currecy information.

Used api: https://api.hitbtc.com/ to get price information.
currenlty it supports only BTCUSD and ETHBTC

## Usage
#### `GET /currency/all`

**Output:** JSON
Response:
```json
  [
    {
        "id": "BTC",
        "fullName": "Bitcoin",
        "ask": "35397.33",
        "bid": "35390.09",
        "last": "35390.20",
        "open": "36565.75",
        "low": "34813.77",
        "high": "37900.00",
        "feeCurrency": "USD"
    },
    {
        "id": "ETH",
        "fullName": "Ethereum",
        "ask": "0.073367",
        "bid": "0.073357",
        "last": "0.073372",
        "open": "0.072600",
        "low": "0.072591",
        "high": "0.075059",
        "feeCurrency": "BTC"
    }
]
```

#### `GET /currency/ETC`

**Output:** JSON
Response:
```json
  {
    "id": "ETH",
    "fullName": "Ethereum",
    "ask": "0.073428",
    "bid": "0.073426",
    "last": "0.073423",
    "open": "0.073197",
    "low": "0.072591",
    "high": "0.075059",
    "feeCurrency": "BTC"
}
```
#### `GET /currency/BTC`

**Output:** JSON
Response:
```json
 {
    "id": "BTC",
    "fullName": "Bitcoin",
    "ask": "35357.50",
    "bid": "35349.24",
    "last": "35354.27",
    "open": "36565.75",
    "low": "34813.77",
    "high": "37900.00",
    "feeCurrency": "USD"
}
```


## Run Locally
I have used go mod to maintain dependencies.
```sh
$ go build
$ ./crypto-server

use below GET http call's using any Rest client:

http://localhost:10000/currency/all
http://localhost:10000/currency/BTC
http://localhost:10000/currency/ETH

```

## License
