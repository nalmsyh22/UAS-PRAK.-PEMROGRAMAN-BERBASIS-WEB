package main

import (
    "database/sql"
    "html/template"
    "log"
    "net/http"

    _ "github.com/go-sql-driver/mysql"
    "train-ticket/handlers"
)

var tpl *template.Template
var db *sql.DB

func init() {
    tpl = template.Must(template.ParseGlob("templates/*.html"))
    var err error
    db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/train_ticket")
    if err != nil {
        log.Fatal(err)
    }
}

func main() {
    http.HandleFunc("/", handlers.Index(tpl))
    http.HandleFunc("/register", handlers.Register(db, tpl))
    http.HandleFunc("/login", handlers.Login(db, tpl))
    http.HandleFunc("/logout", handlers.Logout(tpl))
    http.HandleFunc("/dashboard", handlers.Dashboard(tpl))
    http.HandleFunc("/tickets", handlers.Tickets(tpl, db))
    http.HandleFunc("/order", handlers.OrderTicket(db, tpl))
    http.HandleFunc("/orders", handlers.ViewOrders(db, tpl))

    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

    log.Println("Server started on: http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
