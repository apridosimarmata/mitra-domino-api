package Routes

import (
	"domino/domino/Controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	grp_user := r.Group("/api/user")
	{
		grp_user.POST("create", Controllers.CreateUser)

		//Authentication
		grp_user.POST("OTP/:phone", Controllers.CheckUserOtpByPhone)
		grp_user.GET("OTP/:phone", Controllers.GenerateUserOTPByPhone)
		grp_user.GET("reauthenticate", Controllers.GenerateNewJWTToken)

		//Authorization
		grp_user.GET("authorize", Controllers.AuthorizeUser)
	}

	grp_player := r.Group("/api/player")
	{
		grp_player.GET(":id", Controllers.GetPlayerByID)
		grp_player.GET("send/:id", Controllers.SendVoucher)
		grp_player.GET("login_topbos", Controllers.LoginTopBos)
		grp_player.GET("history", Controllers.GetHistory)
	}

	grp_agen := r.Group("/api/agen")
	{
		grp_agen.GET("orders", Controllers.GetOrdersByAgen)
	}
	return r
}
