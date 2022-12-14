// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"gof/internal/dao/internal"
)

// internalConfigResDbDao is custom type for wrapping custom DAO implements.
type internalConfigResDbDao = *internal.ConfigResDbDao

// configResDbDao is the data access object for table config_res_db.
// You can define custom methods on it to extend its functionality as you wish.
type configResDbDao struct {
	internalConfigResDbDao
}

var (
	// ConfigResDb is globally public accessible object for table config_res_db operations.
	ConfigResDb = configResDbDao{
		internal.NewConfigResDbDao(),
	}
)

// Fill with you ideas below.
