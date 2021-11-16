package Controllers

import (
	"domino/domino/Models"
	"domino/domino/Utils"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetPlayerByID(c *gin.Context) {

	body, err := ioutil.ReadFile("./cookie.txt")
	if err != nil {
		fmt.Println(err)
	}

	client := &http.Client{}
	PostData := strings.NewReader("buyerId=" + c.Params.ByName("id"))

	req, err := http.NewRequest("POST", "https://trade.topbos.com/trade/queryBuyer", PostData)
	req.Header.Set("Cookie", "DAY_ONCE_SHOW_24=1; aliyungf_tc=2ed0bf0f79e48e2633335651020a1d53e7d54a9c5e7477bcd662d0bd55c497d8; trade-cookie-uid=887207; trade-cookie-token="+string(body))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	resp, err := client.Do(req)

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	}
	reg, err := regexp.Compile("[^a-zA-Z0-9 ]+")
	if err != nil {
		fmt.Println(err)
	}

	processedString := reg.ReplaceAllString(strings.Split((strings.Split(string(data), ":")[3]), ",")[0], "")
	fmt.Println(processedString)
	c.JSON(http.StatusOK, processedString)
}

func LoginTopBos(c *gin.Context) {
	client := &http.Client{}
	PostData := strings.NewReader("partnerId=887207&pwd=Kuda1234!")
	req, err := http.NewRequest("POST", "https://trade.topbos.com/trade/pwdLogin", PostData)

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	resp, err := client.Do(req)

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
	}

	f, err := os.Create("./cookie.txt")

	_, err2 := f.WriteString(strings.Split((strings.Split((resp.Header["Set-Cookie"][2]), ";")[0]), "=")[1])

	if err2 != nil {
		fmt.Println(err2)
	}
	//fmt.Println(strings.Split(resp.Header, ))
	c.JSON(http.StatusOK, string(data))

}

func SendVoucher(c *gin.Context) {
	body, err := ioutil.ReadFile("./cookie.txt")
	if err != nil {
		fmt.Println(err)
	}
	var phone string
	var user Models.User

	if c.Request.Header["Token"] != nil {
		switch ValidateJWTToken(c.Request.Header["Token"][0], &phone) {
		case 200:
			err := Models.GetUserByPhone(&user, phone)
			if err != nil {
				c.AbortWithStatus(http.StatusNotFound)
			}
			client := &http.Client{}
			PostData := strings.NewReader("itemId=5&buyerId=" + c.Params.ByName("id"))
			req, err := http.NewRequest("POST", "https://trade.topbos.com/trade/sellCard", PostData)
			req.Header.Set("Cookie", "DAY_ONCE_SHOW_24=1; aliyungf_tc=2ed0bf0f79e48e2633335651020a1d53e7d54a9c5e7477bcd662d0bd55c497d8; trade-cookie-uid=887207; trade-cookie-token="+string(body))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

			resp, err := client.Do(req)

			data, err := ioutil.ReadAll(resp.Body)

			if err != nil {
				fmt.Println(err)
			}
			if len(string(data)) == 31 {
				var order Models.Order
				order.BuyerId = c.Params.ByName("id")
				order.CreatedAt = Utils.Now()
				order.UserId = user.Id
				err = Models.CreateOrder(&order)
				fmt.Println("test")
				c.JSON(http.StatusOK, string(data))
			} else {
				c.AbortWithStatus(http.StatusBadRequest)
			}
		case 410:
			c.AbortWithStatus(http.StatusGone)
		default:
			c.AbortWithStatus(http.StatusBadRequest)
		}
	}

}

func GetHistory(c *gin.Context) {
	var phone string
	var user Models.User

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
				c.AbortWithStatus(http.StatusNotFound)
			}
			c.JSON(http.StatusOK, orders)
		case 410:
			c.AbortWithStatus(http.StatusGone)
		default:
			c.AbortWithStatus(http.StatusBadRequest)
		}
	}
}
