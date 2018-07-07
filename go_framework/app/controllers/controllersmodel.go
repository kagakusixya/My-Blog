package controllers

import (
	"fmt"
	"reflect"
	"net/http"
)

type T struct {
	Url         string
	Html_data   map[string]string
	Form        map[string]string
	HttpStatece int
	Cookie			*http.Cookie
	Cookier         string
}
type Post struct{
	An map[string]string
	CookieData string
}
/*
func Ok(x string) string {
	//配列で送り,0番目の数値をURLの情報として取得する。
	return x

}
*/

//controllerを読み込むfuncを作らなければいけない
func Routeontroller_package(route_data string,poust Post) T {
	//for i := 0; i < len(route_data); i++ {
	var t T
	t.HttpStatece = 200
	x := route_data //メソッド名
	t.Cookier =poust.CookieData
	reflect.ValueOf(&t).MethodByName(x).Call([]reflect.Value{}) //xの名前のメソッドをcontorollersから実行するaは変数用
	fmt.Printf("%s of call:%s\n", x, t.Url)
	//↑どのようなメソッドをcallしたかの確認
	return t
}
func Routeontroller_package_post(route_data string, poust Post) T {
	//for i := 0; i < len(route_data); i++ {
	var t T
	x := route_data
	t.Form = poust.An
	t.Cookier =poust.CookieData
	reflect.ValueOf(&t).MethodByName(x).Call([]reflect.Value{}) //xの名前のメソッドをcontorollersから実行するaは変数用
	fmt.Printf("%s of call:%s\n", x, t.Url)
	//↑どのようなメソッドをcallしたかの確認
	return t
}
