package model

import (
	"html/template"
	"log"
	"net/http"
	"strings"

	"../controllers"
)

func (route_data routedata) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		//w.WriteHeader(418)
		funcMap := template.FuncMap{
			"safehtml": func(text string) template.HTML { return template.HTML(text) },
		}
		s := controllers.Routeontroller_package(route_data.Filemethod)
		w.WriteHeader(s.HttpStatece)
		t := template.Must(template.New("").Funcs(funcMap).ParseFiles("app/views/" + string(s.Url)))
		mozi := String(s.Url)
		var k string
		var su int = 0
		for i := 0; i < len(mozi); i++ {
			if mozi[i] == '/' {
				su = i
			}
		}
		if su != 0 {
			for i := su + 1; i < len(mozi); i++ {
				k = k + string(mozi[i])
			}
		} else {
			k = string(mozi)
		}
		if err := t.ExecuteTemplate(w, k, s.Html_data); err != nil {
			log.Fatal(err)
		}
		//	w.WriteHeader(s.HttpStatece)
		//	w.WriteHeader(418)
	}
	if r.Method == "POST" {
		r.ParseForm()
		a := make(map[string]string)
		for k, v := range r.Form {
			a[k] = strings.Join(v, "")
		}

		s := controllers.Routeontroller_package_post(route_data.Filemethod, a)
		t := template.Must(template.ParseFiles("app/views/" + string(s.Url)))
		mozi := String(s.Url)
		var k string
		var su int = 0
		for i := 0; i < len(mozi); i++ {
			if mozi[i] == '/' {
				su = i
			}
		}
		if su != 0 {
			for i := su + 1; i < len(mozi); i++ {
				k = k + string(mozi[i])
			}
		} else {
			k = string(mozi)
		}
		if err := t.ExecuteTemplate(w, k, s.Html_data); err != nil {
			log.Fatal(err)
		}
	}
}

func Webmake(route_data routedata) { //route_dataとviewのHTMLのURLを取得（配列にする予定）

	http.Handle(route_data.Pash, routedata(route_data))
	//route_data[i].Pashの部分がlocalhost:9000/aとかになる
	//http.ListenAndServe(":9000", nil)
}
