// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ConfigDataDbDao is the data access object for table config_data_db.
type ConfigDataDbDao struct {
	table   string              // table is the underlying table name of the DAO.
	group   string              // group is the database configuration group name of current DAO.
	columns ConfigDataDbColumns // columns contains all the column names of Table for convenient usage.
}

// ConfigDataDbColumns defines and stores column names for table config_data_db.
type ConfigDataDbColumns struct {
	ZoneId     string // 区id(引用mst_zone中的zone_id,0表示所有的区)
	UserName   string // 数据库用户名
	Password   string // 数据库密码
	DbName     string // 数据库名称
	MasterHost string // 主库ip
	MasterPort string // 主库端口
	SlaveHost  string // 从库ip
	SlavePort  string // 从库端口
}

// configDataDbColumns holds the columns for table config_data_db.
var configDataDbColumns = ConfigDataDbColumns{
	ZoneId:     "zone_id",
	UserName:   "user_name",
	Password:   "password",
	DbName:     "db_name",
	MasterHost: "master_host",
	MasterPort: "master_port",
	SlaveHost:  "slave_host",
	SlavePort:  "slave_port",
}

// NewConfigDataDbDao creates and returns a new DAO object for table data access.
func NewConfigDataDbDao() *ConfigDataDbDao {
	return &ConfigDataDbDao{
		group:   "default",
		table:   "config_data_db",
		columns: configDataDbColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ConfigDataDbDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ConfigDataDbDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ConfigDataDbDao) Columns() ConfigDataDbColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ConfigDataDbDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ConfigDataDbDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the merror from function f if it returns non-nil merror.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ConfigDataDbDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
