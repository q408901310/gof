package main

import (
	_ "gof/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"gof/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
