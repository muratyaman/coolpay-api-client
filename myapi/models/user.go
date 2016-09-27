package models

import (
  "errors"
  "strconv"
  "time"
)

var (
  UserList map[string]*User
)

func init() {
  UserList = make(map[string]*User)
  u := User{"user_11111", "astaxie", "11111"}
  UserList["user_11111"] = &u
}

type User struct {
  id       string
  username string
  password string
}

func AddUser(u User) string {
  u.id = "user_" + strconv.FormatInt(time.Now().UnixNano(), 10)
  UserList[u.id] = &u
  return u.id
}

func GetUser(id string) (u *User, err error) {
  if u, ok := UserList[id]; ok {
    return u, nil
  }
  return nil, errors.New("User not exists")
}

func GetAllUsers() map[string]*User {
  return UserList
}

func UpdateUser(id string, uu *User) (a *User, err error) {
  if u, ok := UserList[id]; ok {
    if uu.username != "" {
      u.username = uu.username
    }
    if uu.password != "" {
      u.password = uu.password
    }
    return u, nil
  }
  return nil, errors.New("User Not Exist")
}

func UserLogin(username, password string) bool {
  for _, u := range UserList {
    if u.username == username && u.password == password {
      return true
    }
  }
  return false
}

func DeleteUser(id string) {
  delete(UserList, id)
}
