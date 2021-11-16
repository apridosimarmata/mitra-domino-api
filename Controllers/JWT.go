package Controllers

import (
	"domino/domino/Models"
	"fmt"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var mySigningKey = []byte("unicorns")

func GenerateJWTToken(phone string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["phone"] = phone
	claims["exp"] = time.Now().Add(time.Second * 1800).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GenerateJWTRefreshToken(phone string) (string, error) {
	refresh_token := jwt.New(jwt.SigningMethodHS256)

	claims := refresh_token.Claims.(jwt.MapClaims)

	claims["phone"] = phone
	claims["exp"] = time.Now().Add(time.Hour * 24 * 14).Unix()

	tokenString, err := refresh_token.SignedString(mySigningKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateJWTToken(tokenString string, phone *string) int {

	var code = 200

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			code = 400
		}
		checkExpired := token.Claims.(jwt.MapClaims).VerifyExpiresAt(time.Now().Unix(), false)
		if !checkExpired {
			code = 410
		}

		return mySigningKey, nil
	})

	if err != nil {
		fmt.Println(code)
		if code != 410 {
			code = 400
		}
	}

	*phone, _ = token.Claims.(jwt.MapClaims)["phone"].(string)

	if token.Valid {
		return code
	}

	return code
}

func GenerateNewJWTToken(c *gin.Context) {

	var authenticationResponse Models.AuthenticationResponse

	if c.Request.Header["Refresh-Token"] == nil {
		authenticationResponse.Code = 400
		c.JSON(http.StatusBadRequest, authenticationResponse)
	} else {
		token, err := jwt.Parse(c.Request.Header["Refresh-Token"][0], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				authenticationResponse.Code = 400
				c.JSON(http.StatusBadRequest, authenticationResponse)
			}
			checkExpired := token.Claims.(jwt.MapClaims).VerifyExpiresAt(time.Now().Unix(), false)
			if !checkExpired {
				authenticationResponse.Code = 410
				c.JSON(http.StatusGone, authenticationResponse)
			}

			return mySigningKey, nil
		})

		if err != nil {
			authenticationResponse.Code = 400
			c.JSON(http.StatusBadRequest, authenticationResponse)
		}

		if token.Valid {
			var phone string
			phone, _ = token.Claims.(jwt.MapClaims)["phone"].(string)
			authenticationResponse.Code = 200
			authenticationResponse.Token, _ = GenerateJWTToken(phone)
			authenticationResponse.RefreshToken, _ = GenerateJWTRefreshToken(phone)
			c.JSON(http.StatusOK, authenticationResponse)
		}

	}
}
