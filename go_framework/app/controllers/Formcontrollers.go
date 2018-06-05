package controllers

func (t *T) Formcontrollers() {
	t.Html_data = map[string]string{
		"key1": t.Form["name"],
		"key2": t.Form["say"],
		//"ket3": say,
	}

	t.Url = "keiziban.html.tpl"
}
