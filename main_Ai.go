package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/mux"
)

type Tickets struct {
	Id, Price uint16
	FromWhere, ToWhere, Duration string
	// DepartureDate time.Date
	DepartureTime, ArrivalTime time.Time
}

var tkts = []Tickets{}
var showTicket = Tickets{}


func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/tickets.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/golangProject") // ?parseTime=true&loc=Asia%2FCalcutta
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// db := helper.DbConn()
	// res := time.Now().UTC()
	res, err := db.Query("SELECT `id`, `fromWhere`, `toWhere`, `departureTime`, `arrivalTime`, `duration` from `tickets`")
	if err != nil {
		panic(err.Error())
	}

	tkts = []Tickets{}

	for res.Next() {
		var tkt Tickets
		err = res.Scan(&tkt.Id, &tkt.FromWhere, &tkt.ToWhere, &tkt.Duration, &tkt.DepartureTime, &tkt.ArrivalTime)
		if err != nil {
			panic(err.Error())
		}

		tkts = append(tkts, tkt)
	}

	t.ExecuteTemplate(w, "tickets", tkts)

}

func show_ticket(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	t, err := template.ParseFiles("templates/ticketDetail.html")

	if err != nil {
		fmt.Fprintf(w, err.Error())
	}

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/golangProject")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// выборка данных
	res, err := db.Query(fmt.Sprintf("SELECT `id`, `fromWhere`, `toWhere`, `duration` FROM `tickets` WHERE `id` = '%s'", vars["id"]))
	if err != nil {
		panic(err.Error())
	}

	showTicket = Tickets{}

	for res.Next() { 
		var tkt Tickets
		err = res.Scan(&tkt.Id, &tkt.FromWhere, &tkt.ToWhere, &tkt.Duration) // метод который позволяет нам определитьь существует ли у нас какое то значение
		if err != nil {
			panic(err.Error())
		}

		showTicket = tkt 
	}

	t.ExecuteTemplate(w, "ticketDetail", showTicket)	
}

func handleFunc() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/", index).Methods("GET")
	rtr.HandleFunc("/post/{id:[0-9]+}", show_ticket).Methods("GET")

	http.Handle("/", rtr)
	http.ListenAndServe(":8080", nil)
}

func mainAi() {
	handleFunc()
}