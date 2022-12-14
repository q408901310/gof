// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ConfigGroupChannelDao is the data access object for table config_group_channel.
type ConfigGroupChannelDao struct {
	table   string                    // table is the underlying table name of the DAO.
	group   string                    // group is the database configuration group name of current DAO.
	columns ConfigGroupChannelColumns // columns contains all the column names of Table for convenient usage.
}

// ConfigGroupChannelColumns defines and stores column names for table config_group_channel.
type ConfigGroupChannelColumns struct {
	Id        string //
	GroupId   string // 分组ID
	ChannelId string // 组名称
}

// configGroupChannelColumns holds the columns for table config_group_channel.
var configGroupChannelColumns = ConfigGroupChannelColumns{
	Id:        "id",
	GroupId:   "group_id",
	ChannelId: "channel_id",
}

// NewConfigGroupChannelDao creates and returns a new DAO object for table data access.
func NewConfigGroupChannelDao() *ConfigGroupChannelDao {
	return &ConfigGroupChannelDao{
		group:   "default",
		table:   "config_group_channel",
		columns: configGroupChannelColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ConfigGroupChannelDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ConfigGroupChannelDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ConfigGroupChannelDao) Columns() ConfigGroupChannelColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ConfigGroupChannelDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ConfigGroupChannelDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the merror from function f if it returns non-nil merror.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ConfigGroupChannelDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
