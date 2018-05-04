package main

import (
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"fmt"
	"time"

	"github.com/martini-contrib/sessionauth"
)

var (
	storeUser map[interface{}]*User
	PassWord  string
)

func init() {
	PassWord = "hello" //uuid.New()
	storeUser = make(map[interface{}]*User)
}

type User struct {
	time          int64
	SeatNo        string `form:"SeatNo"`
	CardNo        string `form:"CardNo"`
	DeviceId      string `form:"DeviceId"`
	authenticated bool
}

func GenerateAnonymousUser() sessionauth.User {
	return &User{}
}

func (u *User) Login() {
	u.authenticated = true
}

func (u *User) Logout() {
	u.authenticated = false
}

func (u *User) IsAuthenticated() bool {
	return u.authenticated
}

func (u *User) hash() interface{} {
	src := u.SeatNo + u.CardNo + u.DeviceId + fmt.Sprint(u.time)
	hasher := sha1.New()
	hasher.Write([]byte(src))
	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}

func (u *User) UniqueId() interface{} {
	if u.time == 0 {
		u.time = time.Now().Unix()
	}
	id := u.hash()
	storeUser[id] = u
	return id
}

func (u *User) GetById(id interface{}) error {
	fmt.Println(id)
	if v, ok := storeUser[id]; ok {
		if v.hash() != id {
			return errors.New("User is change")
		}
		u.CardNo = v.CardNo
		u.DeviceId = v.DeviceId
		u.SeatNo = v.SeatNo
		u.authenticated = v.authenticated
		u.time = v.time
		return nil
	}
	return errors.New("User not found id")
}
