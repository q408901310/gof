// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"gof/internal/dao/internal"
)

// internalConfigFightDao is internal type for wrapping internal DAO implements.
type internalConfigFightDao = *internal.ConfigFightDao

// configFightDao is the data access object for table config_fight.
// You can define custom methods on it to extend its functionality as you wish.
type configFightDao struct {
	internalConfigFightDao
}

var (
	// ConfigFight is globally public accessible object for table config_fight operations.
	ConfigFight = configFightDao{
		internal.NewConfigFightDao(),
	}
)

// Fill with you ideas below.