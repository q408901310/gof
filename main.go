package main

import (
	_ "gof/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"gof/internal/cmd"
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
