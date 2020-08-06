<p align="center"><img src="https://github.com/aliworkshop/samanpayment/blob/master/example/saman.jpg" height="150px"></p>

Saman payment library
=====================
Saman payment library for Go Language.
Which includes saman web api methods for connect to bank and verify transaction and reverse transaction.


## Installation

Installation is done using `go get`.
```
go get -u github.com/aliworkshop/samanpayment
```


#### These are following bank methods which is supported by this package
- [x]  **GetTokenRequest**
- [x]  **VerifyTransactionRequest**
- [x]  **ReverseTransactionRequest**

## GetTokenRequest
this method use for get token from bank webservice for connect to the bank payment gateway

this method receives 5 following parameter for get token
Parameter | Type
:-:|:-:
`resnum` | String
`amount` | Integer
`payer_phone` | String
`callback_url` | String
`transaction_key` | Integer

```go
saman := samanpayment.SamanConfig{
    TerminalId: 123456789,
}
result, err := saman.GetTokenRequest(resnum, transaction_key, amount, callback_url, mobile)
```

## VerifyTransactionRequest
this method use for verify transaction and update database if payment was successful

this method receive 2 parameters included transaction_key that you set it before and refnum that you get it from bank response in callback

```go
saman := samanpayment.SamanConfig{
    TerminalId: 123456789,
}
res, err := saman.VerifyTransactionRequest(transaction_key, resp.RefNum)
```

## ReverseTransactionRequest
this method use for reverse transaction that was unsuccessful

this method receive 2 parameters included transaction_key that you set it before and refnum that you get it from bank response in callback

```go
saman := samanpayment.SamanConfig{
    TerminalId: 123456789,
}
res, err := saman.VerifyTransactionRequest(transaction_key, resp.RefNum)
```

## Contributor
[Ali Torabi](https://github.com/aliworkshop)
If this library helps you in anyway, show your love :heart: by putting a :star: on this project :v:

#### More details and learn how to work visit [MiLearn](https://milearn.ir)
