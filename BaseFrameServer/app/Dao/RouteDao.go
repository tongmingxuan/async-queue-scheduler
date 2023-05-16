// Package Dao /*
package Dao

import (
	"BaseFrameServer/app/Model"
	"github.com/tongmingxuan/tmx-server/tmxServer"
	"gorm.io/gorm"
	"time"
)

var (
	QueueTypeByRedis string = "redis"
	QueueTypeByMq    string = "mq"

	RouteLevelTypeByMain  string = "main"
	RouteLevelTypeByChild string = "child"

	// RouteOpenStatusByClose open_status打开关闭状态 2关闭 (禁用)
	RouteOpenStatusByClose int = 2
	// RouteOpenStatusByOpen open_status打开关闭状态  1打开
	RouteOpenStatusByOpen int = 1
)

func GetRouteDao(d *tmxServer.Dao) *RouteDao {
	return &RouteDao{
		Dao: d,
	}
}

type RouteDao struct {
	*tmxServer.Dao
	tmxServer.BaseDao
}

func (dao RouteDao) GetModel() tmxServer.InterfaceModel {
	return Model.RouteModel{}
}

func (dao RouteDao) Create(data map[string]interface{}) *gorm.DB {
	data["created_at"] = time.Now().Format("2006-01-02 15:04:05")

	return dao.GormDb.Model(dao.GetModel()).Create(data)
}

func (dao *RouteDao) FindInfoByWhere(where interface{}) (Model.RouteModel, error) {
	routeModel := Model.RouteModel{}

	model := dao.GormDb.Model(dao.GetModel())

	query, err := dao.BuildFindByWhere(model, where)

	if err != nil {
		return routeModel, err
	}

	return routeModel, query.First(&routeModel).Error
}

func (dao *RouteDao) CreateInTx(data Model.RouteModel) Model.RouteModel {
	data.CreatedAt = tmxServer.Now()

	dao.GormDb.Model(dao.GetModel()).Create(&data)

	return data
}

func (dao *RouteDao) UpdateInfoByWhere(where, data map[string]interface{}) *gorm.DB {
	data["updated_at"] = time.Now().Format("2006-01-02 15:04:05")

	return dao.GormDb.Model(dao.GetModel()).Where(where).Updates(data)
}
