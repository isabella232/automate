package gen

import (
	"bytes"
	"fmt"

	"github.com/dave/jennifer/jen"
	"github.com/golang/protobuf/proto"
	pgs "github.com/lyft/protoc-gen-star"
	pgsgo "github.com/lyft/protoc-gen-star/lang/go"

	"github.com/chef/automate/components/automate-grpc/protoc-gen-a2-config/api/a2conf"
)

const a2ConfPkg = "github.com/chef/automate/components/automate-grpc/protoc-gen-a2-config/api/a2conf"

type A2ServiceConfigModule struct {
	*pgs.ModuleBase
	ctx pgsgo.Context
}

func NewA2ServiceConfigModule() pgs.Module {
	return &A2ServiceConfigModule{
		ModuleBase: &pgs.ModuleBase{},
	}
}

func (m *A2ServiceConfigModule) InitContext(c pgs.BuildContext) {
	m.ModuleBase.InitContext(c)
	m.ctx = pgsgo.InitContext(c.Parameters())
}

func (m *A2ServiceConfigModule) Name() string { return "a2_service_config" }

func (m *A2ServiceConfigModule) Execute(targets map[string]pgs.File, pkgs map[string]pgs.Package) []pgs.Artifact {

OUTER:
	for _, f := range targets {
		messages := f.AllMessages()
		acc := &PortAcc{}

		for _, msg := range messages {
			if !proto.HasExtension(msg.Descriptor().GetOptions(), a2conf.E_ServiceConfig) {
				continue
			}

			iext, err := proto.GetExtension(msg.Descriptor().GetOptions(), a2conf.E_ServiceConfig)
			m.CheckErr(err)
			serviceConfigInfo := iext.(*a2conf.ServiceConfig)

			m.visit(msg, acc)

			out := bytes.Buffer{}
			m.applyTemplate(&out, f.File().Name().String(), msg, serviceConfigInfo, acc)
			generatedFileName := m.ctx.OutputPath(f).SetExt(".a2svc.go").String()
			m.Debugf("%s", generatedFileName)
			m.AddGeneratorFile(generatedFileName, out.String())

			// We only support one config type per service
			break OUTER
		}
	}

	return m.Artifacts()
}

type PortAcc struct {
	portInfo   []PortInfo
	fieldStack []pgs.Field
}

func (acc *PortAcc) PushField(f pgs.Field) {
	acc.fieldStack = append(acc.fieldStack, f)
}

func (acc *PortAcc) PopField() {
	acc.fieldStack = acc.fieldStack[:len(acc.fieldStack)-1]
}

func (acc *PortAcc) SnapshotFields() []pgs.Field {
	fields := make([]pgs.Field, len(acc.fieldStack))
	copy(fields, acc.fieldStack)
	return fields
}

type PortInfo struct {
	Port      a2conf.Port
	SourceLoc string
	Path      []pgs.Field
}

func (m *A2ServiceConfigModule) visit(msg pgs.Message, acc *PortAcc) {
	if msg == nil {
		return
	}
	fields := msg.Fields()
	for _, field := range fields {
		m.Debugf("Push %s", field.Name())
		acc.PushField(field)

		file := field.File().InputPath()
		line := field.SourceCodeInfo().Location().Span[0] + 1
		sourceLoc := fmt.Sprintf("%s:%d", file, line)
		fieldType := field.Type()
		descriptor := field.Descriptor()

		if proto.HasExtension(descriptor.GetOptions(), a2conf.E_Port) {
			iext, err := proto.GetExtension(descriptor.GetOptions(), a2conf.E_Port)
			m.CheckErr(err)
			ext := iext.(*a2conf.Port)
			if ext.Name == "" {
				m.Fail("A name is required for the port")
			}

			if ext.Default <= 0 {
				// We don't automatically allocate ports yet, so
				// you must tell us what the default port to use is
				m.Fail("A valid default port must be provided")
			}

			fieldMsg := fieldType.Embed()

			m.Debugf("Field %s is a port", fieldMsg.FullyQualifiedName())
			switch fieldMsg.FullyQualifiedName() {
			case ".google.protobuf.Int32Value", ".google.protobuf.StringValue":
			default:
				m.Failf("%s is not a supported Message type for Port", fieldMsg.FullyQualifiedName())
			}

			for _, portInfo := range acc.portInfo {
				if portInfo.Port.Name == ext.Name {
					m.Failf("%s port name %q already in use by previous annotation at %s",
						sourceLoc, portInfo.Port.Name, portInfo.SourceLoc)
				}
			}
			acc.portInfo = append(acc.portInfo, PortInfo{
				Port:      *ext,
				SourceLoc: sourceLoc,
				Path:      acc.SnapshotFields(),
			})
		} else {
			next := fieldType.Embed()
			m.visit(next, acc)
		}
		m.Debugf("Pop %s", field.Name())
		acc.PopField()
	}
}

const commentFormat = `// Code generated by protoc-gen-a2config. DO NOT EDIT.
// source: %s
`

func (m *A2ServiceConfigModule) applyTemplate(
	buf *bytes.Buffer,
	protoFileName string,
	message pgs.Message,
	serviceConfigInfo *a2conf.ServiceConfig,
	acc *PortAcc) {

	importPath := m.ctx.ImportPath(message).String()
	pkgName := m.ctx.PackageName(message).String()

	f := jen.NewFilePathName(importPath, pkgName)
	f.HeaderComment(fmt.Sprintf(commentFormat, protoFileName))

	f.Comment("ServiceName returns the name of the service this config belongs to")
	f.Add(m.generateServiceNameFunc(message, serviceConfigInfo.Name))
	f.Comment("BindPort sets the port tagged with the given name")
	f.Add(m.generateBindPortMethod(message, acc))
	f.Comment("ListPorts lists all the ports exposed by the config")
	f.Add(m.generateListPorts(message, acc))
	f.Comment("GetPort gets the port tagged with the given name. If the value is not set, it returns 0.")
	f.Add(m.generateGetPortMethod(message, acc))
	m.CheckErr(f.Render(buf))
}

func (m *A2ServiceConfigModule) generateServiceNameFunc(message pgs.Message, name string) *jen.Statement {
	// func (m *ConfigRequest) ServiceName() string {
	f := jen.Func().Params(jen.Id("m").Id("*" + m.ctx.Name(message).String())).Id("ServiceName").Params().
		Params(jen.String()).
		BlockFunc(func(g *jen.Group) {
			g.Return(jen.Lit(name))
		})
	return f
}

func (m *A2ServiceConfigModule) generateListPorts(message pgs.Message, acc *PortAcc) *jen.Statement {
	// func (m *ConfigRequest) ListPorts() []api.PortInfo
	f := jen.Func().Params(
		jen.Id("m").Id("*" + m.ctx.Name(message).String()),
	).Id("ListPorts").Params().Params(jen.Index().Qual(a2ConfPkg, "PortInfo")).BlockFunc(func(g *jen.Group) {
		// Example:
		// return []api.PortInfo {
		//   api.Port{
		//	   Name: "service",
		//     Default: 10143,
		//     Protocol: "grpc",
		//   },
		// }
		ret := jen.Index().Qual(a2ConfPkg, "PortInfo").ValuesFunc(func(vg *jen.Group) {
			for _, p := range acc.portInfo {
				vg.Add(jen.Qual(a2ConfPkg, "PortInfo").Values(jen.Dict{
					jen.Id("Name"):     jen.Lit(p.Port.Name),
					jen.Id("Default"):  jen.Uint16().Call(jen.Lit(p.Port.Default)),
					jen.Id("Protocol"): jen.Lit(p.Port.Protocol),
				}))
			}
		})
		g.Return(ret)
	})
	return f
}

func (m *A2ServiceConfigModule) generateBindPortMethod(message pgs.Message, acc *PortAcc) *jen.Statement {
	// func (m *ConfigRequest) BindPort(name string, value uint16) error
	f := jen.Func().Params(
		jen.Id("m").Id("*"+m.ctx.Name(message).String()),
	).Id("BindPort").Params(
		jen.Id("name").String(),
		jen.Id("value").Uint16(),
	).Params(jen.Error()).BlockFunc(func(g *jen.Group) {
		// switch name {
		g.Switch(jen.Id("name")).BlockFunc(func(sg *jen.Group) {
			for _, portInfo := range acc.portInfo {
				// Example:
				// case "service":
				sg.Case(jen.Lit(portInfo.Port.Name)).BlockFunc(func(cg *jen.Group) {
					// Create first struct on the ConfigRequest object
					// Example:
					// v0 := &m.V1
					// if *v0 == nil {
					//   *v0 = &ConfigRequest_V1{}
					// }
					cg.Id("v0").Op(":=").Op("&").Id("m").Dot(m.ctx.Name(portInfo.Path[0]).String())
					cg.If(jen.Op("*").Id("v0").Op("==").Nil()).Block(
						jen.Op("*").Id("v0").Op("=").Op("&").Qual(
							m.ctx.ImportPath(portInfo.Path[0].Type().Embed()).String(),
							m.ctx.Name(portInfo.Path[0].Type().Embed()).String()).Values(),
					)

					// Continue creating structs where nil is found
					for i, path := range portInfo.Path {
						if i == 0 {
							continue
						}
						cur := fmt.Sprintf("v%d", i)
						prev := fmt.Sprintf("v%d", i-1)

						// Example:
						// v1 := &(*v0).Sys
						cg.Id(cur).Op(":=").Op("&").Parens(jen.Op("*").Id(prev)).Dot(m.ctx.Name(path).String())

						if i == len(portInfo.Path)-1 {
							// If we're at a port annotation, we need to create the correct wrapper type
							// (because nobody could decide how big a port is) and fill it in with the
							// value cast to the right size.
							switch path.Type().Embed().FullyQualifiedName() {
							case ".google.protobuf.Int32Value":
								// Example:
								// v3 := &(*v2).Port
								// *v3 = &wrappers.Int32Value{Value: int32(value)}
								cg.Op("*").Id(cur).Op("=").Op("&").
									Qual("github.com/golang/protobuf/ptypes/wrappers", "Int32Value").
									Values(
										jen.Dict{
											jen.Id("Value"): jen.Int32().Call(jen.Id("value")),
										},
									)
							case ".google.protobuf.StringValue":
								// Example:
								// valueStr := strconv.Itoa(int(value))
								// *v3 = &wrappers.StringValue{Value: valueStr}
								cg.Id("valueStr").Op(":=").Qual("strconv", "Itoa").Call(jen.Int().Call(jen.Id("value")))
								cg.Op("*").Id(cur).Op("=").Op("&").
									Qual("github.com/golang/protobuf/ptypes/wrappers", "StringValue").
									Values(
										jen.Dict{
											jen.Id("Value"): jen.Id("valueStr"),
										},
									)
							default:
								m.Failf("unsupported message type for port")
							}
						} else {
							// If we're not at a port annotation, we need to create the struct if it's missing
							// Example:
							// if *v1 == nil {
							//   *v1 = &ConfigRequest_V1_System{}
							// }
							cg.If(jen.Op("*").Id(cur).Op("==").Nil()).Block(
								jen.Op("*").Id(cur).Op("=").Op("&").Qual(
									m.ctx.ImportPath(path.Type().Embed()).String(),
									m.ctx.Name(path.Type().Embed()).String()).Values(),
							)
						}
					}
				})
			}
			sg.Default().Block(
				jen.Return(jen.Qual(a2ConfPkg, "ErrPortNotFound")),
			)
		})
		g.Return(jen.Nil())
	})

	return f
}

func (m *A2ServiceConfigModule) generateGetPortMethod(message pgs.Message, acc *PortAcc) *jen.Statement {
	// func (m *ConfigRequest) GetPort(name string) (uint16, error)
	f := jen.Func().Params(
		jen.Id("m").Id("*"+m.ctx.Name(message).String()),
	).Id("GetPort").Params(
		jen.Id("name").String(),
		jen.Id("value").Uint16(),
	).Params(jen.Uint16(), jen.Error()).BlockFunc(func(g *jen.Group) {
		// switch name {
		g.Switch(jen.Id("name")).BlockFunc(func(sg *jen.Group) {
			for _, portInfo := range acc.portInfo {
				// Example:
				// case "service":
				sg.Case(jen.Lit(portInfo.Port.Name)).BlockFunc(func(cg *jen.Group) {
					// Example:
					// v0 := m.V1
					// if v0 == nil {
					//   return 0, nil
					// }
					cg.Id("v0").Op(":=").Id("m").Dot(m.ctx.Name(portInfo.Path[0]).String())
					cg.If(jen.Id("v0").Op("==").Nil()).Block(
						jen.Return(jen.Lit(0), jen.Nil()),
					)

					for i, path := range portInfo.Path {
						if i == 0 {
							continue
						}
						cur := fmt.Sprintf("v%d", i)
						prev := fmt.Sprintf("v%d", i-1)

						// Example:
						// v1 := v0.Sys
						cg.Id(cur).Op(":=").Id(prev).Dot(m.ctx.Name(path).String())

						if i == len(portInfo.Path)-1 {
							switch path.Type().Embed().FullyQualifiedName() {
							case ".google.protobuf.Int32Value":
								// Example:
								// return v3.Value, nil
								cg.Return(jen.Uint16().Call(jen.Id(cur).Dot("Value")), jen.Nil())
							case ".google.protobuf.StringValue":
								// Example:
								// parts := strings.Split(v3.Value, "-")
								// valueInt, err := strconv.ParseUint(parts[0], 10, 16)
								// if err != nil {
								//   return 0, err
								// }
								// return uint16(valueInt), nil
								cg.Id("parts").Op(":=").Qual("strings", "Split").
									Call(jen.Id(cur).Dot("Value"), jen.Lit("-"))
								cg.List(jen.Id("valueInt"), jen.Id("err")).Op(":=").
									Qual("strconv", "ParseUint").
									Call(jen.Id("parts").Index(jen.Lit(0)), jen.Lit(10), jen.Lit(16))
								cg.If(jen.Id("err")).Op("!=").Nil().Block(jen.Return(jen.Lit(0), jen.Id("err")))
								cg.Return(jen.Uint16().Call(jen.Id("valueInt")), jen.Nil())
							default:
								m.Failf("unsupported message type for port")
							}
						} else {
							// Example:
							// if v1 == nil {
							//   return 0, nil
							// }
							cg.If(jen.Id(cur).Op("==").Nil()).Block(
								jen.Return(jen.Lit(0), jen.Nil()),
							)
						}
					}
				})
			}
			sg.Default().Block(
				jen.Return(jen.Lit(0), jen.Qual(a2ConfPkg, "ErrPortNotFound")),
			)
		})
	})

	return f
}
