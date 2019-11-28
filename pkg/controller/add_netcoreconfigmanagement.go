package controller

import (
	"github.com/selcukusta/cm-operator/pkg/controller/netcoreconfigmanagement"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, netcoreconfigmanagement.Add)
}
