package controllers

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
	t.Html_data = map[string]string{
		"main":  t.Form["main"],
		"title": "<h1>" + "おはよう" + t.Form["side"] + "</h1>",
		"side":  "<p><a href=\"/teapot\">" + "・" + t.Form["side"] + "</a></p>",
	}
	t.Url = "burogu.html"
}
