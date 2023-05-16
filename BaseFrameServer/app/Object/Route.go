// Package Object /*
package Object

// RouteQueueInfo
// @Description: 路由参数 queue_info解析
type RouteQueueInfo struct {
	//redis链接池名称
	Connection string `json:"connection"`
	//redis队列名称
	QueueName string `json:"queue_name"`
}

type RouteCreate struct {
	Id             int    `gorm:"primaryKey" json:"id"`
	MainRouteId    int    `json:"main_route_id"`
	UpRouteId      int    `json:"up_route_id"`
	RouteName      string `json:"route_name"`
	Describe       string `json:"describe"`
	QueueType      string `json:"queue_type"`
	QueueInfo      string `json:"queue_info"`
	CreatedAt      string `json:"created_at"`
	RouteLevelType string `json:"route_level_type"`
	OpenStatus     int    `json:"open_status"`
}
