package main

import (
	_ "gof/internal/packed"

	"gof/internal/cmd"

	"github.com/gogf/gf/v2/os/gctx"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
)

func main() {
	err := cmd.Main.AddCommand(&cmd.Restart)
	if err != nil {
		panic(err)
	}
	err = cmd.Main.AddObject(
		cmd.Tpl,
	)
	if err != nil {
		panic(err)
	}
	cmd.Main.Run(gctx.New())
}
