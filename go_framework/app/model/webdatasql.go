package model

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

//sql書くよ
func Dyanamic_data() []Dynamicwebdata {
	db, err := sql.Open("mysql", "blog:@/Blog")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()                            // 関数がリターンする直前に呼び出される
	rows, err := db.Query("SELECT * FROM blog") //
	if err != nil {
		panic(err.Error())
	}
	var i int = 0
	var dw []Dynamicwebdata
	//var sidealldata string = "<p><a href=\"/teapot\">・紅茶でも飲みませんか？</a></p>"
	dw = make([]Dynamicwebdata, 0, 70)
	for rows.Next() {
		var id int
		var main string
		var title string
		if err := rows.Scan(&id, &title, &main); err != nil {
			panic(err.Error())
		}
		//	sidealldata = "<p><a href=\"/home" + strconv.Itoa(id) + "\">" + title + "</a></p>" + sidealldata
		dw = append(dw, Dynamicwebdata{id, main, title, ""})
		i++
	}
	return dw
}
