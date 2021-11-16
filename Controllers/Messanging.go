package Controllers

import (
	"domino/domino/Models"
	"domino/domino/Utils"
	"fmt"
	"os"
	"time"

	twilio "github.com/twilio/twilio-go"

	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

var loca, _ = time.LoadLocation("Asia/Jakarta")

func sendOTP(phone string) bool {
	var user Models.User
	err := Models.GetUserByPhone(&user, phone)

	otpCode := GenerateOTP()
	client := twilio.NewRestClient()

	params := &openapi.CreateMessageParams{}
	params.SetTo("+" + phone)
	params.SetFrom(os.Getenv("TWILIO_PHONE_NUMBER"))
	params.SetBody("Berikut adalah kode hutahita (domino) milikmu " + otpCode + "\nJangan berikan kode ini kepada siapapun.")

	_, err = client.ApiV2010.CreateMessage(params)
	if err != nil {
		return false
	} else {
		user.Otp = otpCode
		Models.UpdateUser(&user, Utils.IntToString(user.Id))
		fmt.Println("ok")
		return true
	}
}
