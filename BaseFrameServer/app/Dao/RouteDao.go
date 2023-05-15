// Package Dao /*
package Dao

import (
	"BaseFrameServer/app/Model"
	"github.com/tongmingxuan/tmx-server/tmxServer"
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
