package Dao

import (
	"BaseFrameServer/app/Model"
	"github.com/tongmingxuan/tmx-server/tmxServer"
	"gorm.io/gorm"
	"time"
)

func GetRouteUndoDao(d *tmxServer.Dao) *RouteUndoDao {
	return &RouteUndoDao{
		Dao: d,
	}
}

type RouteUndoDao struct {
	*tmxServer.Dao
	tmxServer.BaseDao
}

func (dao RouteUndoDao) GetModel() tmxServer.InterfaceModel {
	return Model.RouteUndoModel{}
}

func (dao RouteUndoDao) Create(data map[string]interface{}) *gorm.DB {
	data["created_at"] = time.Now().Format("2006-01-02 15:04:05")
	data["updated_at"] = time.Now().Format("2006-01-02 15:04:05")
	
	return dao.GormDb.Model(dao.GetModel()).Create(data)
}
