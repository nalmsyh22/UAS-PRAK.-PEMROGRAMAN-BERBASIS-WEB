package handlers

import (
    "database/sql"
    "html/template"
    "net/http"
    "log"
    "strconv"
    "train-ticket/models"
)

type Order struct {
    ID        int
    UserID    int
    TicketID  int
    OrderDate string
}

func OrderTicket(db *sql.DB, tpl *template.Template) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method == http.MethodPost {
            userID, err := strconv.Atoi(r.FormValue("user_id"))
            if err != nil {
                http.Error(w, err.Error(), http.StatusBadRequest)
                return
            }

            ticketID, err := strconv.Atoi(r.FormValue("ticket_id"))
            if err != nil {
                http.Error(w, err.Error(), http.StatusBadRequest)
                return
            }

            _, err = db.Exec("INSERT INTO orders (user_id, ticket_id) VALUES (?, ?)", userID, ticketID)
            if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
            }

            http.Redirect(w, r, "/orders", http.StatusSeeOther)
            return
        }

        tickets, err := getAllTickets(db)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        tpl.ExecuteTemplate(w, "order.html", tickets)
    }
}

func getAllTickets(db *sql.DB) ([]models.Ticket, error) {
    rows, err := db.Query("SELECT id, train, date, price FROM tickets")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var tickets []models.Ticket
    for rows.Next() {
        var ticket models.Ticket
        if err := rows.Scan(&ticket.ID, &ticket.Train, &ticket.Date, &ticket.Price); err != nil {
            return nil, err
        }
        tickets = append(tickets, ticket)
    }
    return tickets, nil
}

func ViewOrders(db *sql.DB, tpl *template.Template) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        rows, err := db.Query("SELECT orders.id, users.username, tickets.train, orders.order_date FROM orders JOIN users ON orders.user_id = users.id JOIN tickets ON orders.ticket_id = tickets.id")
        if err != nil {
            log.Fatal(err)
        }
        defer rows.Close()

        var orders []struct {
            OrderID   int
            Username  string
            Train     string
            OrderDate string
        }
        for rows.Next() {
            var order struct {
                OrderID   int
                Username  string
                Train     string
                OrderDate string
            }
            if err := rows.Scan(&order.OrderID, &order.Username, &order.Train, &order.OrderDate); err != nil {
                log.Fatal(err)
            }
            orders = append(orders, order)
        }

        tpl.ExecuteTemplate(w, "view_orders.html", orders)
    }
}
