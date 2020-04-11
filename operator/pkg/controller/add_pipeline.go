package controller

import (
	"github.com/ClareChu/tigger/operator/pkg/controller/pipeline"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, pipeline.Add)
}
