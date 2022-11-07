package main

import (
	"fmt"
	"google.golang.org/protobuf/compiler/protogen"
)

func main() {
	protogen.Options{}.Run(func(gen *protogen.Plugin) error {
		for _, f := range gen.Files {
			if !f.Generate {
				continue
			}
			generateFile(gen, f)
		}
		return nil
	})
}

// 生成.struct.go文件，参数为 输出插件gen，以及读取的文件file
func generateFile(gen *protogen.Plugin, file *protogen.File) {
	filename := file.GeneratedFilenamePrefix + ".struct.go"
	g := gen.NewGeneratedFile(filename, file.GoImportPath)
	// 输出 package packageName
	g.P("package ", file.GoPackageName)
	g.P() // 换行

	for _, m := range file.Messages {
		// 输出 type m.GoIdent struct {
		g.P("type ", m.GoIdent, " struct {")
		for _, field := range m.Fields {
			leadingComment := field.Comments.Leading.String()
			trailingComment := field.Comments.Trailing.String()

			line := fmt.Sprintf("%s %s `json:\"%s\"` %s", field.GoName, field.Desc.Kind(), field.Desc.JSONName(), trailingComment)
			// 输出 行首注释
			g.P(leadingComment)
			// 输出 行内容
			g.P(line)
		}
		// 输出 }
		g.P("}")
	}
	g.P() // 换行
}
