package main

import (
	"crypto/sha256"
	"database/sql"
	"fmt"
)

func main() {
	db, err := sql.Open("mysql", "login:@/Blog")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close() // 関数がリターンする直前に呼び出される
	//	id := "karasawa"
	str := "pantu1024"
	solt := "iroiro"
	hashstr := sha256.Sum256([]byte(str + solt))
	fmt.Printf("%x", hashstr)
	//rows := "insert into  Blog.logintable (iduser,password,solt) values(\"" + id + "\"," + "\"" + *(*string)(unsafe.Pointer(&hashstr)) + "\"" + ",\"" + solt + "\")"
	//fmt.Printf(rows)
	//	rows, err := db.Exec("insert into  Blog.logintable (iduser,password,solt) values(\"" + id + "\"," + "\"" + *(*string)(unsafe.Pointer(&hashstr)) + "\"" + ",\"" + solt + "\")")
	/*
	     if err != nil {
	   		panic(err.Error())
	   	}
	   	ids, err := rows.LastInsertId()
	   	if err != nil {
	   		panic(err.Error())
	   	}
	   	fmt.Println(ids)

	   	// 影響を与えた行数を返す
	   	n, err := rows.RowsAffected()
	   	if err != nil {
	   		panic(err.Error())
	   	}
	   	fmt.Println(n)
	*/
}
