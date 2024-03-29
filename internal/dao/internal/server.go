// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ServerDao is the data access object for table server.
type ServerDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of current DAO.
	columns ServerColumns // columns contains all the column names of Table for convenient usage.
}

// ServerColumns defines and stores column names for table server.
type ServerColumns struct {
	Id        string //
	Group     string //
	Name      string //
	Ip        string //
	Http      string //
	Socket    string //
	MaxReg    string //
	MaxOnline string //
	Status    string //
	Hot       string //
	Sort      string //
	Open      string //
	CreatedAt string //
	UpdatedAt string //
	DeletedAt string //
}

// serverColumns holds the columns for table server.
var serverColumns = ServerColumns{
	Id:        "id",
	Group:     "group",
	Name:      "name",
	Ip:        "ip",
	Http:      "http",
	Socket:    "socket",
	MaxReg:    "max_reg",
	MaxOnline: "max_online",
	Status:    "status",
	Hot:       "hot",
	Sort:      "sort",
	Open:      "open",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewServerDao creates and returns a new DAO object for table data access.
func NewServerDao() *ServerDao {
	return &ServerDao{
		group:   "default",
		table:   "server",
		columns: serverColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *ServerDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *ServerDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *ServerDao) Columns() ServerColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *ServerDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *ServerDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *ServerDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
