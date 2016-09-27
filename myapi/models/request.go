package models

import (
  "errors"
  //"strconv"
  "time"
  "github.com/satori/go.uuid"
  "myapi/coolpay"
)

var (
  RequestList map[string]*Request
)

func init() {
  RequestList = make(map[string]*Request)
  u := Request{"db05d5f7-0710-44ce-b507-1c748b0c4120", "2016-0926 23:29:00", "test request", "2016-0926 23:30:00", "test response"}
  RequestList["user_11111"] = &u
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
  u.id = id.String()
  var ts = time.Now()
  u.request_ts = ts.Format("2006-01-02 15:04:05")
  RequestList[u.id] = &u
  return u.id
}

func GetRequest(uid string) (u *Request, err error) {
  if u, ok := RequestList[uid]; ok {
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


func DeleteRequest(id string) {
  delete(RequestList, id)
}


func RequestLogin() (token string, err error) {
  url      := beego.AppConfig.String("coolpay_url")
  username := beego.AppConfig.String("coolpay_user")
  password := beego.AppConfig.String("coolpay_key")

  apiClient := NewApiClient(url)
  output, err := apiClient.NewToken(username, password)
  if err == nil {
    return output, nil
  }
  return nil, errors.New("API failure")
}
