package myTickets

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"go/basic/g/modules"

	"github.com/gin-gonic/gin"
)

func PutTicket(db *sql.DB) gin.HandlerFunc{
	return func(c *gin.Context){	
		var id, _ = strconv.Atoi(c.Param("id"))
		var ctx = c.Request.Context()
		var myTickets modules.MyTickets

		if e := c.ShouldBindJSON(&myTickets); e!=nil{
			c.JSON(http.StatusBadRequest,"please try again or enter the valid data")
			return
		}

		if _,e := db.ExecContext(ctx,fmt.Sprintf("UPDATE `myTickets` SET `userId`='%d',`ticketId`='%d' WHERE `id`=%d;", myTickets.UserId, myTickets.TicketId, id)); e!=nil {
			c.JSON(http.StatusInternalServerError, "internal server error")
			return
		}

		c.JSON(http.StatusOK,"succesfully updated!!!")

	}
}