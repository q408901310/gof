// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"gof/internal/dao/internal"
)

// internalConfigGroupChannelDao is internal type for wrapping internal DAO implements.
type internalConfigGroupChannelDao = *internal.ConfigGroupChannelDao

// configGroupChannelDao is the data access object for table config_group_channel.
// You can define custom methods on it to extend its functionality as you wish.
type configGroupChannelDao struct {
	internalConfigGroupChannelDao
}

var (
	// ConfigGroupChannel is globally public accessible object for table config_group_channel operations.
	ConfigGroupChannel = configGroupChannelDao{
		internal.NewConfigGroupChannelDao(),
	}
)

// Fill with you ideas below.
