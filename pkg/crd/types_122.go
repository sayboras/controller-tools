//go:build go1.22
// +build go1.22

package crd

import "go/types"

func convert(in types.Type) *types.TypeName {
	switch in.(type) {
	case *types.Named:
		return in.(*types.Named).Obj()
	case *types.Alias:
		return in.(*types.Alias).Obj()
	}
	return nil
}
