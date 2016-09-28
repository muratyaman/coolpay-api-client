package models

import (
  "errors"
  //"strconv"
  "time"
  "github.com/satori/go.uuid"
  "github.com/astaxie/beego"

  "myapi/coolpay"
)

var (
  RequestList map[string]*Request
)

func init() {
  RequestList = make(map[string]*Request)
  r := Request{"db05d5f7-0710-44ce-b507-1c748b0c4120", "2016-09-26 23:29:00", "test request", "2016-09-26 23:30:00", "test response"}
  RequestList["db05d5f7-0710-44ce-b507-1c748b0c4120"] = &r
}

type Request struct {
  id            string
  request_ts    string
  request_data  string
  response_ts   string
  response_data string
}


func AddRequest(u Request) string {
  var id = uuid.NewV4()
  u.id   = id.String()
  var ts = time.Now()
  u.request_ts = ts.Format("2006-01-02 15:04:05")
  RequestList[u.id] = &u
  return u.id
}

func GetRequest(id string) (u *Request, err error) {
  if u, ok := RequestList[id]; ok {
    return u, nil
  }
  return nil, errors.New("Request not found")
}

func GetAllRequests() map[string]*Request {
  return RequestList
}

func UpdateRequest(id string, uu *Request) (a *Request, err error) {
  if u, ok := RequestList[id]; ok {
    if uu.request_ts != "" {
      u.request_ts = uu.request_ts
    }
    if uu.request_data != "" {
      u.request_data = uu.request_data
    }
    if uu.response_ts != "" {
      u.response_ts = uu.response_ts
    }
    if uu.response_data != "" {
      u.response_data = uu.response_data
    }
    return u, nil
  }
  return nil, errors.New("Request not found")
}


func RequestLogin() (string, error) {
  url      := beego.AppConfig.String("coolpay_url")
  username := beego.AppConfig.String("coolpay_user")
  apikey   := beego.AppConfig.String("coolpay_key")

  apiClient := coolpay.NewApiClient(url)
  inputLogin := coolpay.NewInputLogin(username, apikey)
  token, err := apiClient.Login(inputLogin)
  if err == nil {
    return token, err
  }
  return token, errors.New("API failure")
}
