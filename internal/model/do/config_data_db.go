// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ConfigDataDb is the golang structure of table config_data_db for DAO operations like Where/Data.
type ConfigDataDb struct {
	g.Meta     `orm:"table:config_data_db, do:true"`
	ZoneId     interface{} // 区id(引用mst_zone中的zone_id,0表示所有的区)
	UserName   interface{} // 数据库用户名
	Password   interface{} // 数据库密码
	DbName     interface{} // 数据库名称
	MasterHost interface{} // 主库ip
	MasterPort interface{} // 主库端口
	SlaveHost  interface{} // 从库ip
	SlavePort  interface{} // 从库端口
}
