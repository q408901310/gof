// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// ConfigResDb is the golang structure for table config_res_db.
type ConfigResDb struct {
	ZoneId   int    `json:"zoneId"   ` // 区id(引用mst_zone中的zone_id,0表示所有的区)
	UserName string `json:"userName" ` // 资源库用户名
	Password string `json:"password" ` // 资源库密码
	Url      string `json:"url"      ` // 连接字符串
	Driver   string `json:"driver"   ` // DB驱动程序
}