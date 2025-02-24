// Code generated by go-enum DO NOT EDIT.
// Version: example
// Revision: example
// Build Date: example
// Built By: example

package example

import (
	"errors"
	"fmt"
)

const (
	// ForceLowerTypeDataSwap is a ForceLowerType of type DataSwap.
	ForceLowerTypeDataSwap ForceLowerType = iota
	// ForceLowerTypeBootNode is a ForceLowerType of type BootNode.
	ForceLowerTypeBootNode
)

var ErrInvalidForceLowerType = errors.New("not a valid ForceLowerType")

const _ForceLowerTypeName = "dataswapbootnode"

var _ForceLowerTypeMap = map[ForceLowerType]string{
	ForceLowerTypeDataSwap: _ForceLowerTypeName[0:8],
	ForceLowerTypeBootNode: _ForceLowerTypeName[8:16],
}

// String implements the Stringer interface.
func (x ForceLowerType) String() string {
	if str, ok := _ForceLowerTypeMap[x]; ok {
		return str
	}
	return fmt.Sprintf("ForceLowerType(%d)", x)
}

var _ForceLowerTypeValue = map[string]ForceLowerType{
	_ForceLowerTypeName[0:8]:  ForceLowerTypeDataSwap,
	_ForceLowerTypeName[8:16]: ForceLowerTypeBootNode,
}

// ParseForceLowerType attempts to convert a string to a ForceLowerType.
func ParseForceLowerType(name string) (ForceLowerType, error) {
	if x, ok := _ForceLowerTypeValue[name]; ok {
		return x, nil
	}
	return ForceLowerType(0), fmt.Errorf("%s is %w", name, ErrInvalidForceLowerType)
}
