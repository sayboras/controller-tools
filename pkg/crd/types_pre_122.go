//go:build !go1.22
// +build !go1.22

package crd

import "go/types"

func convert(in types.Type) *types.TypeName {
	if r, ok := in.(*types.Named); ok {
		return r.Obj()
	}
	return nil
}
