package model

import (
	"fmt"
	"net/http"
)

type String string
type routedata struct { //routeの情報を構造体に入れる
	Httpmethod  string
	Pash        string
	Packagename string
	Filemethod  string
}
type Y struct {
	Url         string
	html_data   map[string]string
	HttpStatece int
	Form        map[string]string
	Cookie			*http.Cookie
}
type Dynamicwebdata struct {
	Id       int
	Maindata string
	Side     string
	SideAll  string
}

func Model() {
	//reder input
	route_data := Readroute()
	//読み込んだ物の処理
	//var view_in controllers.T
	//var y Y
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public/"))))
	for i := 0; i < len(route_data); i++ {
		fmt.Printf("Pash:%s  ", route_data[i].Pash)
		Webmake(route_data[i]) //routea_dataをwebを立てるためにWebmakeに飛ばす
	}
	//webの動的処理
	dynamic_data := Dyanamic_data()
	DynamicWebmake(dynamic_data) //routea_dataをwebを立てるためにWebmakeに飛ばす

	var port_number string = ":9000"
	fmt.Printf("port number%s", port_number)
	http.ListenAndServe(port_number, nil)
}
