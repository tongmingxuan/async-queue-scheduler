// Package Service /*
package Service

import (
	"BaseFrameServer/app/Common/Utils"
	"BaseFrameServer/app/Dao"
	"BaseFrameServer/app/Model"
	"BaseFrameServer/app/Object"
	"encoding/json"
	"errors"
	"github.com/syyongx/php2go"
	"github.com/tongmingxuan/tmx-server/tmxServer"
	"gorm.io/gorm"
)

type RouteService struct {
	BaseService
}

func (service RouteService) checkParam(param map[string]interface{}, routeDao *Dao.RouteDao) Result {
	queueInfo := tmxServer.AntToString(param["queue_info"])

	if queueInfo == "" {
		return service.ServiceError("queue_info:获取信息异常", param["queue_info"], 500)
	}

	queueInfoStruct := Object.RouteQueueInfo{}

	if json.Unmarshal([]byte(queueInfo), &queueInfoStruct) != nil {
		return service.ServiceError("queue_info:json解析异常", param["queue_info"], 500)
	}

	//尝试获取链接
	redisConnection := Utils.GetRedisConnection(queueInfoStruct.Connection)

	res := redisConnection.Ping()

	_, err := res.Result()

	if err != nil {
		return service.ServiceError("queue_info:connection:获取链接异常:请检查链接配置是否正确"+err.Error(), param["queue_info"], 500)
	}

	if queueInfoStruct.QueueName == "" {
		return service.ServiceError("queue_info:queue_name:为空", param["queue_info"], 500)
	}

	upRouteId := tmxServer.StringToInt(param["up_route_id"].(string), "up_route_id类型转化异常")

	mainRouteId := 0
	if !php2go.Empty(param["main_route_id"]) {
		mainRouteId_ := param["main_route_id"].(string)

		mainRouteId = tmxServer.StringToInt(mainRouteId_, "main_route_id类型转化异常")
	}

	if mainRouteId > 0 {
		mainRouteInfo, _ := routeDao.FindInfoByWhere(map[string]interface{}{
			"id": mainRouteId,
		})

		if mainRouteInfo.Id == 0 {
			return service.ServiceError("main_route_id:对应路由不存在", param["main_route_id"], 500)
		}
	}

	if upRouteId > 0 {
		upInfo, _ := routeDao.FindInfoByWhere(map[string]interface{}{
			"id": upRouteId,
		})

		if upInfo.Id == 0 {
			return service.ServiceError("up_route_id:对应路由不存在", param["up_route_id"], 500)
		}
	}

	return service.ServiceSuccess("验证参数成功", map[string]int{
		"mainRouteId": mainRouteId,
		"upRouteId":   upRouteId,
	})
}

// CreateRoute
// @Description: 创建路由数据
// @receiver service
// @param param
// @return Result
func (service RouteService) CreateRoute(param map[string]interface{}) Result {

	queueInfo := tmxServer.AntToString(param["queue_info"])

	if queueInfo == "" {
		return service.ServiceError("queue_info:获取信息异常", param["queue_info"], 500)
	}

	//获取 gorm实例
	conn := service.GetGormConnection()

	//获取 routeDao
	routeDao := Dao.GetRouteDao(conn)

	//验证数据 且model不开启事务
	checkResult := service.checkParam(param, routeDao)

	if checkResult.Code != 200 {
		return service.ServiceError("checkResult:error:"+checkResult.Message, checkResult.Data, 500)
	}

	//判断 当前路由名称是否存在
	routeInfo, _ := routeDao.FindInfoByWhere(map[string]interface{}{
		"route_name": param["route_name"],
	})

	if routeInfo.Id > 0 {
		return service.ServiceError("checkResult:error:当前路由名称已经存在", routeInfo, 500)
	}

	mapByIdInfo := checkResult.Data.(map[string]int)

	mainRouteId, _ := mapByIdInfo["mainRouteId"]

	upRouteId, _ := mapByIdInfo["upRouteId"]

	param["main_route_id"] = mainRouteId

	//开启事务
	txErr := routeDao.GormDb.Transaction(func(tx *gorm.DB) error {
		//更换gorm对象
		routeDao.GormDb = tx

		routeCreate := Model.RouteModel{
			RouteName: param["route_name"].(string),
			Describe:  param["describe"].(string),
			QueueType: param["queue_type"].(string),
			QueueInfo: queueInfo,
		}

		if upRouteId > 0 {
			routeCreate.UpRouteId = upRouteId
			routeCreate.RouteLevelType = Dao.RouteLevelTypeByChild
		}

		if mainRouteId > 0 {
			routeCreate.MainRouteId = mainRouteId
		}

		if mainRouteId == 0 && upRouteId == 0 {
			routeCreate.RouteLevelType = Dao.RouteLevelTypeByMain
		}

		routeCreate.OpenStatus = Dao.RouteOpenStatusByClose

		insertInfo := routeDao.CreateInTx(routeCreate)

		insertId := insertInfo.Id

		if insertId == 0 {
			return errors.New("新增数据异常")
		}

		if mainRouteId == 0 && upRouteId == 0 {
			row := routeDao.UpdateInfoByWhere(map[string]interface{}{
				"id": insertId,
			}, map[string]interface{}{
				"main_route_id": insertId,
			})

			if row.RowsAffected < 1 {
				return errors.New("更新数据异常:响应行数小于1")
			}
		}

		return nil
	})

	if txErr != nil {
		return service.ServiceError("操作数据库异常:"+txErr.Error(), nil, 500)
	}

	return service.ServiceSuccess("success", nil)
}
