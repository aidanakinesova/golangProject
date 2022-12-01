package myTickets

import (
	"database/sql"
	"fmt"
	"go/basic/g/modules"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostTicket(db *sql.DB) gin.HandlerFunc{
	return func(c *gin.Context){
		var ticket modules.Ticket
		var ctx = c.Request.Context()

		if e:=c.ShouldBindJSON(&ticket); e!=nil{
			c.JSON(http.StatusBadRequest,"error")
			return
		}

		// var row = db.QueryRowContext(ctx, "SELECT MAX(Id) FROM `tickets`")  
		var row = db.QueryRowContext(ctx, "SELECT `id` FROM `tickets` ORDER BY `id` DESC LIMIT 1)")  

		if e := row.Scan(&ticket.Id); e != nil{
			if e==sql.ErrNoRows{
				c.JSON(http.StatusNotFound,"ERROR")
				return
			}
			c.JSON(http.StatusInternalServerError, row.Scan(&ticket.Id))
		}
		c.JSON(http.StatusOK, ticket)
 
		if _, e := db.ExecContext(ctx, fmt.Sprintf("INSERT INTO `tickets` VALUES(row.Scan(&ticket.Id), '%s','%s','%s','%s','%s','%s',%d);", ticket.FromWhere, ticket.ToWhere, ticket.DepartureDate, ticket.DepartureTime, ticket.ArrivalTime, ticket.Duration, ticket.Price)); e != nil {
			c.JSON(http.StatusInternalServerError,"internal error")
			return
		}


		c.Writer.Header().Add("Location",fmt.Sprintf("/tickets/%d",ticket.Id+1))
		c.JSON(http.StatusCreated,"succesfully created")
	}
	
}