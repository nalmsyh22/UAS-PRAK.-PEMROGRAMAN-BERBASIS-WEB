package handlers

import (
    "html/template"
    "net/http"
)

func Dashboard(tpl *template.Template) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        tpl.ExecuteTemplate(w, "dashboard.html", nil)
    }
}
