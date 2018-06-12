package controllers

func (t *T) Teapot() {
	t.HttpStatece = 418
	t.Url = "teapot.html"
}
