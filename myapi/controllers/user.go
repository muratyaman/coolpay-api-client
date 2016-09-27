package controllers

import (
  "myapi/models"
  "encoding/json"

  "github.com/astaxie/beego"
)

// Operations about Users
type UserController struct {
  beego.Controller
}

// @Title CreateUser
// @Description create users
// @Param  body    body   models.User  true    "body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (u *UserController) Post() {
  var user models.User
  json.Unmarshal(u.Ctx.Input.RequestBody, &user)
  uid := models.AddUser(user)
  u.Data["json"] = map[string]string{"uid": uid}
  u.ServeJSON()
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (u *UserController) GetAll() {
  users := models.GetAllUsers()
  u.Data["json"] = users
  u.ServeJSON()
}

// @Title Get
// @Description get user by id
// @Param   id    path   string  true    "The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :id is empty
// @router /:id [get]
func (u *UserController) Get() {
  id := u.GetString(":id")
  if id != "" {
    user, err := models.GetUser(id)
    if err != nil {
      u.Data["json"] = err.Error()
    } else {
      u.Data["json"] = user
    }
  }
  u.ServeJSON()
}

// @Title Update
// @Description update the user
// @Param  id    path   string  true    "The id you want to update"
// @Param  body    body   models.User  true    "body for user content"
// @Success 200 {object} models.User
// @Failure 403 :id is not int
// @router /:id [put]
func (u *UserController) Put() {
  id := u.GetString(":id")
  if id != "" {
    var user models.User
    json.Unmarshal(u.Ctx.Input.RequestBody, &user)
    uu, err := models.UpdateUser(id, &user)
    if err != nil {
      u.Data["json"] = err.Error()
    } else {
      u.Data["json"] = uu
    }
  }
  u.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param  id    path   string  true    "The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (u *UserController) Delete() {
  id := u.GetString(":id")
  models.DeleteUser(id)
  u.Data["json"] = "delete success!"
  u.ServeJSON()
}

// @Title Login
// @Description Logs user into the system
// @Param  username    query   string  true    "The username for login"
// @Param  password    query   string  true    "The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [get]
func (u *UserController) Login() {
  username := u.GetString("username")
  password := u.GetString("password")
  if models.UserLogin(username, password) {
    u.Data["json"] = "login success"
  } else {
    u.Data["json"] = "user not exist"
  }
  u.ServeJSON()
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *UserController) Logout() {
  u.Data["json"] = "logout success"
  u.ServeJSON()
}
