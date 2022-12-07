package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/gogf/gf/v2/text/gstr"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/util/gtag"
)

var (
	Tpl = cTpl{}
)

type cTpl struct {
	g.Meta `name:"tpl" brief:"{cTplBrief}" eg:"{cTplEg}"`
}

const (
	cTplController = `package controller

var (
	${name} = c${name}{}
)

type c${name} struct{}`
	cTplService = `package service

type (
	s${name} struct{}
)

var (
	ins${name} = s${name}{}
)

func ${name}() *s${name} {
	return &ins${name}
}`
	cTplModel = `package model`
	cTplBrief = `create controller or service template`
	cTplEg    = `
gf tpl name -c
gf tpl name -c -s
`
	cTplNameBrief = `
file name.
`
)

func init() {
	gtag.Sets(g.MapStrStr{
		`cTplBrief`:     cTplBrief,
		`cTplEg`:        cTplEg,
		`cTplNameBrief`: cTplNameBrief,
	})
}

type cTplInput struct {
	g.Meta     `name:"tpl"`
	Name       string `name:"NAME" arg:"true" brief:"{cTplNameBrief}"`
	All        bool   `name:"all" orphan:"true" short:"a" brief:"create all template" default:"false"`
	Controller bool   `name:"controller" orphan:"true" short:"c" brief:"create controller template" default:"false"`
	Service    bool   `name:"service" orphan:"true" short:"s" brief:"create service template" default:"false"`
	Model      bool   `name:"model" orphan:"true" short:"m" brief:"create model template" default:"false"`
}
type cTplOutput struct{}

func (c cTpl) Index(ctx context.Context, in cTplInput) (out *cTplOutput, err error) {
	if in.Name == "" {
		fmt.Println("The Name field is required")
	}
	if in.All {
		in.Controller = true
		in.Service = true
		in.Model = true
	}
	if !in.Controller && !in.Service && !in.Model {
		g.Dump("Enter at least one between CONTROLLER and SERVICE and MODEL")
		return
	}

	var (
		name     = gstr.CaseSnake(in.Name)
		rootPath = gfile.Pwd()
		cFile    = gfile.Join(rootPath, "internal/controller/", name+".go")
		sFile    = gfile.Join(rootPath, "internal/service/", name+".go")
		mFile    = gfile.Join(rootPath, "internal/model/", name+".go")
	)

	if in.Controller {
		g.Dump("start create controller file")
		if !gfile.IsEmpty(cFile) {
			s := gcmd.Scanf(`the controller folder "%s" is not empty, files might be overwrote, continue? [y/n]: `, name)
			if strings.EqualFold(s, "n") {
				goto SERVICE
			}
		}
		err := gfile.PutContents(cFile, gstr.Replace(cTplController, "${name}", gstr.UcFirst(name)))
		if err != nil {
			g.Dump(cFile + " create error")
			return nil, err
		}
		g.Dump(cFile + " ok")
	}

SERVICE:
	if in.Service {
		g.Dump("start create service file")
		if !gfile.IsEmpty(sFile) {
			s := gcmd.Scanf(`the service folder "%s" is not empty, files might be overwrote, continue? [y/n]: `, name)
			if strings.EqualFold(s, "n") {
				goto MODEL
			}
		}
		err := gfile.PutContents(sFile, gstr.Replace(cTplService, "${name}", gstr.UcFirst(name)))
		if err != nil {
			g.Dump(sFile + " create error")
			return nil, err
		}
		g.Dump(sFile + " ok")
	}
MODEL:
	if in.Model {
		g.Dump("start create model file")
		if !gfile.IsEmpty(mFile) {
			s := gcmd.Scanf(`the model folder "%s" is not empty, files might be overwrote, continue? [y/n]: `, name)
			if strings.EqualFold(s, "n") {
				return
			}
		}
		err := gfile.PutContents(mFile, gstr.Replace(cTplModel, "${name}", gstr.UcFirst(name)))
		if err != nil {
			g.Dump(mFile + " create error")
			return nil, err
		}
		g.Dump(mFile + " ok")
	}
	return
}
