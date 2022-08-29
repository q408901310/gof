package cmd

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gfile"
)

type Build struct {
	g.Meta        `name:"build" config:"gfcli.build"`
	File          string `name:"FILE" arg:"true"    brief:"building file path"`
	Name          string `short:"n"  name:"name"    brief:"output binary name"`
	Version       string `short:"v"  name:"version" brief:"output binary version"`
	Arch          string `short:"a"  name:"arch"    brief:"output binary architecture, multiple arch separated with ','"`
	System        string `short:"s"  name:"system"  brief:"output binary system, multiple os separated with ','"`
	Output        string `short:"o"  name:"output"  brief:"output binary path, used when building single binary file"`
	Path          string `short:"p"  name:"path"    brief:"output binary directory path, default is './temp'" d:"./temp"`
	Extra         string `short:"e"  name:"extra"   brief:"extra custom \"go build\" options"`
	Mod           string `short:"m"  name:"mod"     brief:"like \"-mod\" option of \"go build\", use \"-m none\" to disable go module"`
	Cgo           bool   `short:"c"  name:"cgo"     brief:"enable or disable cgo feature, it's disabled in default" orphan:"true"`
	VarMap        g.Map  `short:"r"  name:"varMap"  brief:"custom built embedded variable into binary"`
	PackSrc       string `short:"ps" name:"packSrc" brief:"pack one or more folders into one go file before building"`
	PackDst       string `short:"pd" name:"packDst" brief:"temporary go file path for pack, this go file will be automatically removed after built" d:"internal/packed/build_pack_data.go"`
	ExitWhenError bool   `short:"ew" name:"exitWhenError" brief:"exit building when any error occurs, default is false" orphan:"true"`
}

var (
	Restart = gcmd.Command{
		Name:  "restart",
		Usage: "main restart",
		Brief: "restart http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			g.Dump(gfile.RealPath("."))
			return nil
		},
	}
)
