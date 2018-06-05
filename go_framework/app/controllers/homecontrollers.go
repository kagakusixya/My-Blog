package controllers

/*
type T struct {
	Url       string
	Html_data map[string]string
}
*/
func (t *T) Home() {
	t.Html_data = map[string]string{
		"main": " どうも、Kagakusixya(科学者)です。このサイトはブログかつ作者が勉強した時にまとめてたものです。",
		//"side": "<p><a href=\"/teapot\">・紅茶でも飲みませんか？</a></p>",
		"side": "<p><a href=\"/teapot\">・紅茶でも飲みませんか？</a></p>",
	}
	t.Url = "burogu.html"
}
func (t *T) Teapot() {
	t.HttpStatece = 418
	t.Url = "teapot.html"
}

/*
func (t *T) Homehell() {
	t.Url = "tesut.html.tpl"
}
*/
