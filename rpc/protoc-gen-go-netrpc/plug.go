package main

import (
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
	"github.com/golang/protobuf/protoc-gen-go/generator"
)

func init() {
	generator.RegisterPlugin(new(netrpcPlugin))
}

type netrpcPlugin struct {
	*generator.Generator
}

func (p *netrpcPlugin) Name() string {
	return "netrpc"
}

func (p *netrpcPlugin) Init(g *generator.Generator) {
	p.Generator = g
}

// GenerateImports方法调用自定义的genImportCode函数生成导入代码
func (p *netrpcPlugin) GenerateImports(file *generator.FileDescriptor) {
	if len(file.Service) > 0 {
		p.genImportCode(file)
	}
}

//Generate方法调用自定义的genServiceCode方法生成每个服务的代码
func (p *netrpcPlugin) Generate(file *generator.FileDescriptor) {
	for _, svc := range file.Service {
		p.genServiceCode(svc)
	}
}

func (p *netrpcPlugin) genImportCode(file *generator.FileDescriptor) {
	p.P("// TODO: import code")
}

func (p *netrpcPlugin) genServiceCode(svc *descriptor.ServiceDescriptorProto) {
	p.P("// TODO: service code, Name = " + svc.GetName())
}
