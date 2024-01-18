package file

import (
	"os"
	"text/template"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ReadFileOrPanic(path string) []byte {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err.Error())
	}
	return data
}

func ReadFileOrPanicPulumi(path string) pulumi.StringInput {
	return pulumi.String(ReadFileOrPanic(path))
}

var funcMap = template.FuncMap{
	"sub": func(a int, b int) int64 {
		return int64(a - b)
	},
}

var Template = template.New("").Funcs(funcMap)
