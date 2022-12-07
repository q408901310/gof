package service

import (
	"context"
	"fmt"
	"gof/internal/dao"
	"gof/internal/model"
	"gof/internal/model/entity"
	"sort"
)

type (
	sServer struct{}
)

var (
	insServer = sServer{}
)

func Server() *sServer {
	return &insServer
}

func (s sServer) List(ctx context.Context) (out []*model.ServerListItem, err error) {
	user := User().GetUser(ctx)
	groupId, _ := dao.ConfigGroupChannel.Ctx(ctx).Value("group_id", "channel_id", user.Channel)
	if groupId == nil {
		return
	}
	var zones []*entity.ConfigZone
	err = dao.ConfigZone.Ctx(ctx).Where("group_id", groupId.Int()).Scan(&zones)
	if err != nil {
		return nil, err
	}
	sort.Slice(zones, func(i, j int) bool {
		return zones[i].DisplayOrder < zones[j].DisplayOrder
	})
	for _, zone := range zones {
		out = append(out, &model.ServerListItem{
			ZoneId: zone.ZoneId,
			Url:    fmt.Sprintf("ws://%s:%d/ws", zone.Host, zone.WebsocketPort),
			Name:   zone.ZoneName,
			Status: zone.Status,
		})
	}
	return
}
