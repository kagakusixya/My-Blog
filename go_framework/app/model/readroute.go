package model

import (
	"bufio"
	"os"
	"strings"
)

func Readroute() []routedata {
	var route_data []routedata            //構造体の宣言
	route_data = make([]routedata, 0, 70) //makeで長さ0で70のスライス構造を作成
	//ｘは何段までrouteを読み込むかが不変的であるためにappendで必要な分だけをスライスする
	fp, err := os.Open("conf/route") //routeファイルを開くもしなければエラー
	if err != nil {
		panic(err)
	}
	//ファイルを開く---------------------------------------------------
	s := bufio.NewScanner(fp) //スキャンするための変数を確保
	for s.Scan() {
		data_array := strings.Fields(s.Text())                                                                 //for文が回るたびに[]data配列にスペースごとのstringを読み込む
		route_data = append(route_data, routedata{data_array[0], data_array[1], data_array[2], data_array[3]}) //appendで配列を拡張するfor文が回るたびにmakeで０だった部分が+1
	} //!!!data_arrayが6以上の時またはそれ以下の時のエラー表示!!!
	fp.Close()
	return route_data
}
