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
		"main": Dynamic_webdata.Maindata,
		"side": Dynamic_webdata.Side,
	}
	t := template.Must(template.New("").Funcs(funcMap).ParseFiles("app/views/burogu.html"))
	if err := t.ExecuteTemplate(w, "burogu.html", map_data); err != nil {
		log.Fatal(err)
	}
	//	w.WriteHeader(s.HttpStatece)
	//	w.WriteHeader(418)
}

func DynamicWebmake(Dynamic_webdata Dynamicwebdata) { //route_dataとviewのHTMLのURLを取得（配列にする予定）
	//var id int = Dynamic_webdata.Id
	url := "/home" + strconv.Itoa(Dynamic_webdata.Id)
	http.Handle(url, (Dynamic_webdata))
	fmt.Printf("\n%s\n", url)
	//route_data[i].Pashの部分がlocalhost:9000/aとかになる
	//http.ListenAndServe(":9000", nil)
}
