package controllers

func (t *T) Write() {
	t.Url = "write.html"
}
func (t *T) Write_post() {
	t.Html_data = map[string]string{
		"main": t.Form["title"],
		//"side": "<p><a href=\"/teapot\">・紅茶でも飲みませんか？</a></p>",
		"side": t.Form["write"],
	}
	t.Url = "burogu.html"
}
