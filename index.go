package handlers

import (
    "html/template"
    "net/http"
)

func Index(tpl *template.Template) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        tpl.ExecuteTemplate(w, "index.html", nil)
    }
}
