package controllers

//Login.....system
func (t *T) Login() {
	t.Url = "login.html"
}
func (t *T) Login_post() {
	if t.Form["id"] == "id" {
		if t.Form["password"] == "password" {
			t.Html_data = map[string]string{
				"id":       t.Form["id"],
				"password": t.Form["password"],
			}
			t.Url = "loginok.html"
		}
	}
	if t.Url != "loginok.html" {
		t.Url = "err.html"
	}
}
