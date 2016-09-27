package coolpay

import (
  "errors"
  "github.com/astaxie/beego/httplib"
)

type ApiClient struct {
  url string
}

type InputRecipient struct {
  name string
}

type InputPayment struct {
  amount       float64
  currency     string // 'GBP'
  recipient_id string // uuid
}

func NewApiClient(url) *ApiClient {
  c := new(ApiClient)
  c.url := url
  return c
}

//TODO: record ALL requests and responses using Request model

func (c *ApiClient) NewToken(username string, password string) (output string, err error) {
  req := httplib.Post(c.url + "/login")
  req.Param("username", username)
  req.Param("password", password)
  output, err := req.Response()
  if err != nil {
    return nil, err
  }

  return output, nil
}

func (c *ApiClient) NewRecipient(token string, name string) (output string, err error) {
  req := httplib.Post(c.url + "/recipients")
  recipient := new(InputRecipient)
  recipient.name := name
  req.Param("recipient", recipient)
  req.Header("Content-Type", "application/json")
  req.Header("Authorization", "Bearer " + token)
  output, err := req.Response()
  if err != nil {
    return nil, err
  }

  return output, nil
}

func (c *ApiClient) NewPayment(token string, currencyCode string, amount int, recipient_id string) (output string, err error) {
  req := httplib.Post(c.url + "/payments")
  payment := new(InputPayment)
  payment.currency     := currencyCode
  payment.amount       := amount / 100.0
  payment.recipient_id := recipient_id
  req.Param("payment", payment)
  req.Header("Content-Type", "application/json")
  req.Header("Authorization", "Bearer " + token)
  output, err := req.Response()
  if err != nil {
    return nil, err
  }

  return output, nil
}
