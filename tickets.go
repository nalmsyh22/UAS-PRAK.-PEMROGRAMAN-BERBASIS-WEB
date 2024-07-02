package handlers

import (
    "database/sql"
    "html/template"
    "net/http"
    "log"
    "train-ticket/models"
)

func Tickets(tpl *template.Template, db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        rows, err := db.Query("SELECT id, train, date, price FROM tickets")
        if err != nil {
            log.Fatal(err)
        }
        defer rows.Close()

        var tickets []models.Ticket
        for rows.Next() {
            var ticket models.Ticket
            if err := rows.Scan(&ticket.ID, &ticket.Train, &ticket.Date, &ticket.Price); err != nil {
                log.Fatal(err)
            }
            tickets = append(tickets, ticket)
        }

        if err := rows.Err(); err != nil {
            log.Fatal(err)
        }

        tpl.ExecuteTemplate(w, "tickets.html", tickets)
    }
}
