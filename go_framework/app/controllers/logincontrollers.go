package controllers

import (
	"net/http"
	"fmt"
)

//Login.....system
func (t *T) Login() {
	t.Url = "login/login.html"
}
func (t *T) Login_post() {
	fmt.Printf("id:%s\npassword:%s\n",t.Form["id"],t.Form["password"])
	if t.Form["id"] == "i" {
		if t.Form["password"] == "p" {
			t.Html_data = map[string]string{
				"id":       t.Form["id"],
				"password": t.Form["password"],
			}
			/*
			rand.Seed(time.Now().UnixNano())
			rand.Intn(1000000)
			*/
			cookie := &http.Cookie{
				Name: "hoge",
				Value: "Hello", // ここにcookieの名前を記述
			}
			// 2
			t.Cookie = cookie

			t.Url = "login/loginok.html"
		}
	}
	if t.Url != "login/loginok.html" {
		t.Url = "login/err.html"
	}
}
