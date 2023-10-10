package service

import (
	"context"
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
	groupId, _ := dao.ChannelGroup.Ctx(ctx).Value("group", "channel", user.Channel)
	if groupId == nil {
		return
	}
	var servers []*entity.Server
	err = dao.Server.Ctx(ctx).Where("group", groupId.Int()).Scan(&servers)
	if err != nil {
		return nil, err
	}
	sort.Slice(servers, func(i, j int) bool {
		return servers[i].Sort < servers[j].Sort
	})
	for _, server := range servers {
		out = append(out, &model.ServerListItem{
			Id:     int(server.Id),
			Url:    server.Socket,
			Name:   server.Name,
			Status: server.Status,
			Hot:    server.Hot,
		})
	}
	return
}
