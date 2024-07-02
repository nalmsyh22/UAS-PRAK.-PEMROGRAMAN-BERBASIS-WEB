package handlers

import (
    "database/sql"
    "html/template"
    "log"
    "net/http"
)

func Register(db *sql.DB, tpl *template.Template) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodPost {
            username := r.FormValue("username")
            password := r.FormValue("password")

            _, err := db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", username, password)
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }

            http.Redirect(w, r, "/login", http.StatusSeeOther)
            return
        }
        tpl.ExecuteTemplate(w, "register.html", nil)
    }
}

func Login(db *sql.DB, tpl *template.Template) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodPost {
            username := r.FormValue("username")
            password := r.FormValue("password")

            var id int
            err := db.QueryRow("SELECT id FROM users WHERE username = ? AND password = ?", username, password).Scan(&id)
            if err != nil {
                if err == sql.ErrNoRows {
                    http.Error(w, "Invalid credentials", http.StatusUnauthorized)
                } else {
                    http.Error(w, err.Error(), http.StatusInternalServerError)
                }
                return
            }

            log.Printf("User %s logged in with ID %d", username, id) // Tambahkan log ini untuk debugging
            http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
            return
        }
        tpl.ExecuteTemplate(w, "login.html", nil)
    }
}

func Logout(tpl *template.Template) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Implementasi logika logout di sini, seperti menghapus sesi pengguna
        http.Redirect(w, r, "/login", http.StatusSeeOther)
    }
}
