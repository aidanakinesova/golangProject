package ticket

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
 
		if _, e := db.ExecContext(ctx, fmt.Sprintf("INSERT INTO `tickets` VALUES(%d,'%s','%s',%d,'%s','%s');",ticket.Id, ticket.FromWhere, ticket.ToWhere, ticket.DepartureDate, ticket.DepartureTime, ticket.ArrivalTime, ticket.Duration, ticket.Price)); e != nil {
			c.JSON(http.StatusInternalServerError,"internal error")
			return
		}
		// чтобы проверить существуют в базе данных сохранились или нет введенные данные
		// var res modules.Customer
		// var row = db.QueryRowContext(ctx, fmt.Sprintf("SELECT * FROM `Customer` WHERE `Id`=%d",customer.Id))
		// if e := row.Scan(&res.Id, &res.Name, &res.Surname, &res.Age, &res.Sex, &res.CardNumber); e != nil{
		// 	if e==sql.ErrNoRows{
		// 		c.JSON(http.StatusNotFound,"there is now row with such id")
		// 		return
		// 	}

		// 	c.JSON(http.StatusInternalServerError,"error")
		// 	return
		// }
		// c.JSON(http.StatusOK, res)

		c.Writer.Header().Add("Location",fmt.Sprintf("/tickets/%d",ticket.Id))
		c.JSON(http.StatusCreated,"succesfully created")
	}
	
}