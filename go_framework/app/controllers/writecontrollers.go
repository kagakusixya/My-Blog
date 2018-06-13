package controllers

import (
	"database/sql"
	"fmt"
	"strconv"
)

func (t *T) Write() {
	t.Url = "write.html"
}
func (t *T) Preview_post() {
	t.Html_data = map[string]string{
		"main":  t.Form["main"],
		"title": "<h1>" + t.Form["side"] + "</h1>",
		"side":  "<p><a href=\"/teapot\">" + "・" + t.Form["side"] + "</a></p>",
	}
	t.Url = "burogu.html"
}

func (t *T) Send() {
	dbr, err := sql.Open("mysql", "blog:@/Blog")
	if err != nil {
		panic(err.Error())
	}
	defer dbr.Close()                             // 関数がリターンする直前に呼び出される
	rowss, err := dbr.Query("SELECT * FROM blog") //
	if err != nil {
		panic(err.Error())
	}
	var i int = 0
	//var sidealldata string = "<p><a href=\"/teapot\">・紅茶でも飲みませんか？</a></p>"

	for rowss.Next() {
		i++
	}
	db, err := sql.Open("mysql", "blogwrite:@/Blog")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close() // 関数がリターンする直前に呼び出される
	rows, err := db.Exec("insert into  Blog.blog (id,title,main) values(" + strconv.Itoa(i+1) + ",\"" + t.Form["main"] + "\",\"" + t.Form["side"] + " \")")
	if err != nil {
		panic(err.Error())
	}
	id, err := rows.LastInsertId()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(id)

	// 影響を与えた行数を返す
	n, err := rows.RowsAffected()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(n)
	t.Url = "teapot.html"
}
