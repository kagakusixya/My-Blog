package model

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func (Dynamic_webdata Dynamicwebdata) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//w.WriteHeader(418)
	funcMap := template.FuncMap{
		"safehtml": func(text string) template.HTML { return template.HTML(text) },
	}
	var map_data map[string]string
	map_data = map[string]string{
		"main":  Dynamic_webdata.Maindata,
		"title": "<h1>" + Dynamic_webdata.Side + "</h1>",
		"side":  Dynamic_webdata.SideAll,
	}
	t := template.Must(template.New("").Funcs(funcMap).ParseFiles("app/views/burogu.html"))
	if err := t.ExecuteTemplate(w, "burogu.html", map_data); err != nil {
		log.Fatal(err)
	}
	//	w.WriteHeader(s.HttpStatece)
	//	w.WriteHeader(418)
}

func DynamicWebmake(Dynamic_webdata []Dynamicwebdata) { //route_dataとviewのHTMLのURLを取得（配列にする予定）
	//var id int = Dynamic_webdata.Id
	var sidealldata string = "<p><a href=\"/teapot\">・紅茶でも飲みませんか？</a></p>"
	for i := len(Dynamic_webdata) - 1; 0 <= i; i-- {
		if Dynamic_webdata[i].Id == 1 {
			sidealldata = "<p><a href=\"/home" + "\">" + "・" + Dynamic_webdata[i].Side + "</a></p>" + sidealldata
		} else {
			sidealldata = "<p><a href=\"/home" + strconv.Itoa(Dynamic_webdata[i].Id) + "\">" + "・" + Dynamic_webdata[i].Side + "</a></p>" + sidealldata
		}
	}
	for i := 0; i < len(Dynamic_webdata); i++ {
		Dynamic_webdata[i].SideAll = sidealldata
		var url string
		if Dynamic_webdata[i].Id == 1 {
			url = "/home"
		} else {
			url = "/home" + strconv.Itoa(Dynamic_webdata[i].Id)
		}
		http.Handle(url, (Dynamic_webdata[i]))
		fmt.Printf("\n%s\n", url)
	}
	//route_data[i].Pashの部分がlocalhost:9000/aとかになる
	//http.ListenAndServe(":9000", nil)
}
