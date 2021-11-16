package Controllers

import (
	"domino/domino/Models"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var loc, _ = time.LoadLocation("Asia/Jakarta")

func CreateUser(c *gin.Context) {
	var user Models.User
	c.BindJSON(&user)
	user.Stock = 10
	err := Models.CreateUser(&user)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		sendOTP(user.Phone)
		c.Data(http.StatusOK, "application/json", []byte(`{"result":"Success"}`))
	}
}

func CheckUserOtpByPhone(c *gin.Context) {
	var user Models.User
	phone := c.Params.ByName("phone")
	err := Models.GetUserByPhone(&user, phone)

	var otp Models.OtpCheck
	c.BindJSON(&otp)

	var authenticationResponse Models.AuthenticationResponse

	if err != nil {
		authenticationResponse.Code = 404
		c.JSON(http.StatusNotFound, authenticationResponse)
	} else {
		if otp.Otp == user.Otp {
			authenticationResponse.Code = 200
			authenticationResponse.Token, _ = GenerateJWTToken(phone)
			authenticationResponse.RefreshToken, _ = GenerateJWTRefreshToken(phone)
			c.JSON(http.StatusOK, authenticationResponse)
		} else {
			authenticationResponse.Code = 406
			c.JSON(http.StatusNotAcceptable, authenticationResponse)
		}
	}

}

func GenerateUserOTPByPhone(c *gin.Context) {
	var user Models.User
	phone := c.Params.ByName("phone")
	err := Models.GetUserByPhone(&user, phone)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		sendOTP(user.Phone)
		c.Data(http.StatusOK, "application/json", []byte(`{"result":"Success"}`))
	}
}

func AuthorizeUser(c *gin.Context) {
	var phone string
	if c.Request.Header["Token"] != nil {
		switch ValidateJWTToken(c.Request.Header["Token"][0], &phone) {
		case 200:
			c.Data(http.StatusOK, "application/json", []byte(`{"result":"Authorized"}`))
		case 410:
			c.AbortWithStatus(http.StatusGone)
		default:
			c.AbortWithStatus(http.StatusBadRequest)
		}
	}
}

func GetUserByToken(c *gin.Context) {
	var phone string
	var user Models.User

	if c.Request.Header["Token"] != nil {
		fmt.Println("OK")
		switch ValidateJWTToken(c.Request.Header["Token"][0], &phone) {
		case 200:
			err := Models.GetUserByPhone(&user, phone)
			var profile Models.User

			profile.Phone = user.Phone
			if err != nil {
				c.AbortWithStatus(http.StatusNotFound)
			}
			c.JSON(http.StatusOK, profile)
		case 410:
			c.AbortWithStatus(http.StatusGone)
		default:
			c.AbortWithStatus(http.StatusBadRequest)
		}
	}
}
