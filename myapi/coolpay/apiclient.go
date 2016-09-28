package coolpay

import (
  "github.com/astaxie/beego/httplib"
  "encoding/json"
  "bytes"
)

type ApiClient struct {
  url string
}

func NewApiClient(url string) *ApiClient {
  a := new(ApiClient)
  a.url = url
  return a
}

type InputLogin struct {
  username string
  apikey   string
}

func NewInputLogin(username string, apikey string) InputLogin {
  var i InputLogin
  i.username = username
  i.apikey   = apikey
  return i
}

type OutputLogin struct {
  token string
}

type InputNewRecipientDetails struct {
  name string
}

type InputNewRecipient struct {
  recipient InputNewRecipientDetails
}

func NewInputNewRecipient(name string) InputNewRecipient {
  var r InputNewRecipientDetails
  r.name = name
  var i InputNewRecipient
  i.recipient = r
  return i
}

type OutputNewRecipientDetails struct {
  id   string
  name string
}

type OutputNewRecipient struct {
  recipient OutputNewRecipientDetails
}

type InputNewPaymentDetails struct {
  amount       float64
  currency     string // 'GBP'
  recipient_id string // uuid
}

type InputNewPayment struct {
  payment InputNewPaymentDetails
}

type OutputNewPaymentDetails struct {
  id           string
  amount       float64
  currency     string // 'GBP'
  recipient_id string // uuid
}

type OutputNewPayment struct {
  payment OutputNewPaymentDetails
}

func NewInputNewPayment(amount float64, currency string, recipient_id string) InputNewPayment {
  var p InputNewPaymentDetails
  p.amount = amount
  p.currency = currency
  p.recipient_id = recipient_id
  var i InputNewPayment
  i.payment = p
  return i
}

// TODO: record ALL requests and responses using Request model

func (c *ApiClient) Login(input InputLogin) (string, error) {
  req := httplib.Post(c.url + "/login")
  req.Param("username", input.username)
  req.Param("apikey", input.apikey)
  req.Header("Content-Type", ":application/json")
  res, err := req.Response()
  defer res.Body.Close()
  token := ""
  if err == nil {
    outputLogin := new(OutputLogin)
    json.NewDecoder(res.Body).Decode(&outputLogin)
    token = outputLogin.token
  }
  return token, err
}


func (c *ApiClient) NewRecipient(token string, recipient InputNewRecipient) (*OutputNewRecipient, error) {
  req := httplib.Post(c.url + "/recipients")
  b := new(bytes.Buffer)
  json.NewEncoder(b).Encode(recipient)
  req.Body(b)
  req.Header("Content-Type", "application/json")
  req.Header("Authorization", "Bearer " + token)
  res, err := req.Response()
  defer res.Body.Close()
  if err == nil {
    outputNewRecipient := new(OutputNewRecipient)
    json.NewDecoder(res.Body).Decode(&outputNewRecipient)
    return outputNewRecipient, err
  }

  return nil, err
}


func (c *ApiClient) NewPayment(token string, payment InputNewPayment) (*OutputNewPayment, error) {
  var req     = httplib.Post(c.url + "/payments")
  b := new(bytes.Buffer)
  json.NewEncoder(b).Encode(payment)
  req.Body(b)
  req.Header("Content-Type", "application/json")
  req.Header("Authorization", "Bearer " + token)
  res, err := req.Response()
  defer res.Body.Close()
  if err == nil {
    outputNewPayment := new(OutputNewPayment)
    json.NewDecoder(res.Body).Decode(&outputNewPayment)
    return outputNewPayment, err
  }

  return nil, err
}
