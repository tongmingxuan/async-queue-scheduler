// Package Dao /*
package Dao

import (
	"fmt"
	"github.com/tongmingxuan/tmx-server/tmxServer"
)

func GetMessageDao(d *tmxServer.Dao) *MessageDao {
	return &MessageDao{
		Dao: d,
	}
}

type MessageDao struct {
	*tmxServer.Dao
	tmxServer.BaseDao
}

// CreateMessageTable
// @Description: 主路由创建 同步创建对应路由消息表
// @receiver dao
// @param tableName
func (dao *MessageDao) CreateMessageTable(tableName string) {
	db := dao.GormDb

	db.Exec("CREATE TABLE `" + tableName + "` (  " +
		"`id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,  " +
		"`message_id` char(50) NOT NULL COMMENT '消息ID',  " +
		"`main_message_id` char(50) NOT NULL COMMENT '主消息ID',  " +
		"`up_message_id` char(50) DEFAULT NULL COMMENT '上级消息ID',  " +
		"`route_id` int(11) NOT NULL COMMENT '路由ID',  " +
		"`route_info` text NOT NULL COMMENT '当前路由数据',  " +
		"`trigger_status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '是否被触发   1未被触发  2已触发状态',  " +
		"`message_body` text NOT NULL COMMENT '接收的消息体',  " +
		"`message_queue_body` text NOT NULL COMMENT '投递的消息体',  " +
		"`message_info` text COMMENT '消息运行信息',  " +
		"`message_status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '消息状态 1初始状态(待调度) 2confirm成功 3调用finish接口(已调用finish) 4运行异常 5作废  6 创建子消息成功 7.判定运行成功',  " +
		"`message_type` char(10) NOT NULL DEFAULT 'main' COMMENT '消息类型  main主消息  child子消息',  " +
		"`route_queue_type` char(50) NOT NULL COMMENT '路由类型',  " +
		"`route_queue_info` text NOT NULL COMMENT '路由队列信息',  " +
		"`push_queue_at` datetime DEFAULT NULL COMMENT '推送队列时间',  " +
		"`finish_at` datetime DEFAULT NULL COMMENT '调用finish时间',  " +
		"`created_at` datetime NOT NULL COMMENT '创建时间',  " +
		"`updated_at` datetime DEFAULT NULL COMMENT '更新时间',  " +
		"PRIMARY KEY (`id`)) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;")

	res := db.Exec("SELECT`TABLE_NAME`FROM`INFORMATION_SCHEMA`.`TABLES`WHERE `TABLE_NAME` = '" + tableName + "';")

	fmt.Println(res)
}
