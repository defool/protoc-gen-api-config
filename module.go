package main

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"
)

type mod struct {
	*pgs.ModuleBase
	pgsgo.Context
}

func newMod() pgs.Module {
	return &mod{ModuleBase: &pgs.ModuleBase{}}
}

func (m *mod) InitContext(c pgs.BuildContext) {
	m.ModuleBase.InitContext(c)
	m.Context = pgsgo.InitContext(c.Parameters())
}

func (mod) Name() string {
	return "api-config"
}

func (m mod) Execute(targets map[string]pgs.File, gpkgs map[string]pgs.Package) []pgs.Artifact {
	initLogger(enableLogger)

	methods := make([]MethodInfo, 0)
	svcKey := "Service"
	for _, f := range targets {
		pkgName := *f.Descriptor().Package
		tmp := strings.SplitN(pkgName, ".", 2)
		pkgNameX := tmp[0]
		version := "v1"
		if len(tmp) > 1 {
			version = tmp[1]
		}
		for _, svc := range f.Services() {
			svcName := svc.Name().String()
			svcNameX := strings.TrimSuffix(svcName, svcKey)
			svcNameX = strcase.ToKebab(svcNameX)
			for _, method := range svc.Methods() {
				methodName := method.Name().String()
				methodNameX := strcase.ToKebab(methodName)
				methods = append(methods, MethodInfo{
					Name: fmt.Sprintf("%s.%s.%s", pkgName, svcName, methodName),
					Path: fmt.Sprintf("/api/%s/%s/%s/%s", pkgNameX, version, svcNameX, methodNameX),
				})
			}
		}
	}

	info := &Config{
		Methods: methods,
	}
	filename := m.Context.Params().OutputPath() + "/api-config.yaml"
	m.genAPIConfig(info, filename)
	return m.Artifacts()
}

func (m mod) genAPIConfig(info *Config, filename string) {
	buf := bytes.NewBuffer(nil)
	tp := template.Must(template.New("").Parse(fieldTemplate))
	err := tp.Execute(buf, info)
	m.CheckErr(err)
	m.AddGeneratorFile(filename, buf.String())
}

func firstLowger(s string) string {
	if len(s) == 0 {
		return ""
	}
	return strings.ToLower(s[:1]) + s[1:]
}
