package build

import "github.com/zjaco13/proto-test/go/codegen"

func Build() *codegen.EksBlueprint {
	return &codegen.EksBlueprint{Id: "test-stack"}
}
