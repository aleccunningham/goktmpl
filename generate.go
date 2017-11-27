package generate

import (
	"html/template"
	"net/http"
)

func parseTemplate(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseGlob("tmpl/*.yaml")
	t.Execute(w, err)
}
