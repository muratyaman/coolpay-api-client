package controllers

import (
  "myapi/models"
  "encoding/json"

  "github.com/astaxie/beego"
)

// Operations about Requests
type CoolpayController struct {
  beego.Controller
}

// @Title CreateRequest
// @Description create requests
// @Param  body    body   models.Request  true    "body for request content"
// @Success 200 {str} models.Request.id
// @Failure 403 body is empty
// @router / [post]
func (u *CoolpayController) Post() {
  var request models.Request
  json.Unmarshal(u.Ctx.Input.RequestBody, &request)
  id := models.AddRequest(request)
  u.Data["json"] = map[string]string{"id": id}
  u.ServeJSON()
}

// @Title GetAll
// @Description get all Requests
// @Success 200 {object} models.Request
// @router / [get]
func (u *CoolpayController) GetAll() {
  requests := models.GetAllRequests()
  u.Data["json"] = requests
  u.ServeJSON()
}

// @Title Get
// @Description get request by uid
// @Param  uid    path   string  true    "The key for staticblock"
// @Success 200 {object} models.Request
// @Failure 403 :id is empty
// @router /:id [get]
func (u *CoolpayController) Get() {
  id := u.GetString(":id")
  if id != "" {
    request, err := models.GetRequest(id)
    if err != nil {
      u.Data["json"] = err.Error()
    } else {
      u.Data["json"] = request
    }
  }
  u.ServeJSON()
}

// @Title Update
// @Description update the request
// @Param  id     path   string  true    "The id you want to update"
// @Param  body   body   models.Request  true    "body for request content"
// @Success 200 {object} models.Request
// @Failure 403 :id is not int
// @router /:id [put]
func (u *CoolpayController) Put() {
  id := u.GetString(":id")
  if id != "" {
    var request models.Request
    json.Unmarshal(u.Ctx.Input.RequestBody, &request)
    uu, err := models.UpdateRequest(id, &request)
    if err != nil {
      u.Data["json"] = err.Error()
    } else {
      u.Data["json"] = uu
    }
  }
  u.ServeJSON()
}

// @Title Delete
// @Description delete the request
// @Param   id  path     string  true    "The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (u *CoolpayController) Delete() {
  id := u.GetString(":id")
  models.DeleteRequest(id)
  u.Data["json"] = "delete success!"
  u.ServeJSON()
}

// @Title Login
// @Description Logs request into the system
// @Success 200 {string} login success
// @Failure 403 request not exist
// @router /login [post]
func (u *CoolpayController) RequestLogin() {
  if models.RequestLogin() {
    u.Data["json"] = "login success"
  } else {
    u.Data["json"] = "request not exist"
  }
  u.ServeJSON()
}
