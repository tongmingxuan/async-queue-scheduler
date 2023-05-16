// Package Route /*
package Route

import (
	"BaseFrameServer/app/Controller"
	"github.com/gin-gonic/gin"
)

func AdminRouteInit(g *gin.Engine) {

	adminController := Controller.AdminController{}
	routeController := Controller.RouteController{}

	admin := g.Group("/admin")
	admin.POST("/login", adminController.Login)

	admin.GET("/getUserInfo", adminController.GetUserInfo)

	admin.POST("createRoute", routeController.CreateRoute)
}
