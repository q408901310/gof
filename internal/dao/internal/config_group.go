// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ConfigGroupDao is the data access object for table config_group.
type ConfigGroupDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns ConfigGroupColumns // columns contains all the column names of Table for convenient usage.
}

// ConfigGroupColumns defines and stores column names for table config_group.
type ConfigGroupColumns struct {
	GroupId   string //
	GroupName string //
}

// configGroupColumns holds the columns for table config_group.
var configGroupColumns = ConfigGroupColumns{
	GroupId:   "group_id",
	GroupName: "group_name",
}

// NewConfigGroupDao creates and returns a new DAO object for table data access.
func NewConfigGroupDao() *ConfigGroupDao {
	return &ConfigGroupDao{
		group:   "default",
		table:   "config_group",
		columns: configGroupColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ConfigGroupDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ConfigGroupDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ConfigGroupDao) Columns() ConfigGroupColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ConfigGroupDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ConfigGroupDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the merror from function f if it returns non-nil merror.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ConfigGroupDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
