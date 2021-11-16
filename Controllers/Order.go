package Controllers

import (
	"domino/domino/Models"
	"domino/domino/Utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOrdersByAgen(c *gin.Context) {
	var user Models.User
	var phone string
	if c.Request.Header["Token"] != nil {
		switch ValidateJWTToken(c.Request.Header["Token"][0], &phone) {
		case 200:
			err := Models.GetUserByPhone(&user, phone)
			if err != nil {
				c.AbortWithStatus(http.StatusNotFound)
			}

			var orders []Models.Order
			err = Models.GetOrderByAgen(&orders, Utils.IntToString(user.Id))
			if err != nil {
				c.AbortWithStatus(http.StatusInternalServerError)
			}

			c.JSON(http.StatusOK, orders)
		case 410:
			c.AbortWithStatus(http.StatusGone)
		default:
			c.AbortWithStatus(http.StatusBadRequest)
		}
	}
}
