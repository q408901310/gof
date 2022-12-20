// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"gof/internal/dao/internal"
)

// internalConfigChannelDao is custom type for wrapping custom DAO implements.
type internalConfigChannelDao = *internal.ConfigChannelDao

// configChannelDao is the data access object for table config_channel.
// You can define custom methods on it to extend its functionality as you wish.
type configChannelDao struct {
	internalConfigChannelDao
}

var (
	// ConfigChannel is globally public accessible object for table config_channel operations.
	ConfigChannel = configChannelDao{
		internal.NewConfigChannelDao(),
	}
)

// Fill with you ideas below.
