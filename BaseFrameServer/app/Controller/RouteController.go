// Package Controller /*
package Controller

import (
	"BaseFrameServer/app/Service"
	"fmt"
	"github.com/gin-gonic/gin"
)

type RouteController struct {
	BaseController
}

func (controller RouteController) CreateRoute(c *gin.Context) {

	param := controller.FilterData(controller.Params(c), map[string]string{
		"route_name": "路由名称为空",
		"describe":   "路由描述为空",
		"queue_type": "队列类型为空",
		"queue_info": "队列信息为空",
	})

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)

			controller.SetControllerError(c, r)
			return
		}
	}()

	result := Service.RouteService{}.CreateRoute(param)

	if result.Code == 200 {
		c.String(200, controller.JsonSuccess(result.Message, nil))
	} else {
		c.String(200, controller.JsonError(result.Message, nil, result.Code))
	}
}
