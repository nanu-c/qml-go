package qml

// #cgo CPPFLAGS: -I../cpp
// #cgo CXXFLAGS: -std=c++0x  -Wall -fno-strict-aliasing
// #cgo LDFLAGS: -lstdc++
// #cgo pkg-config: Qt5Core Qt5Widgets Qt5Quick
//
// #include <stdlib.h>
// #include "goitemmodel_api.h"
//
import "C"
import "unsafe"

type qModelIndex struct {
	ptr    unsafe.Pointer
	engine *Engine
}

func mkModelIndex(ptr uintptr, engine *Engine) ModelIndex {
	if ptr == 0 {
		return nil
	}
	return &qModelIndex{
		ptr:    unsafe.Pointer(ptr),
		engine: engine,
	}
}

type ModelIndex interface {
	// ModelIndex can only be created from a ItemModel
	internal_ModelIndex()
	Child(row, col int) ModelIndex
	Sibling(row, col int) ModelIndex
	Column() int
	Row() int
	Data(role Role) interface{}
	Flags() ItemFlags
	InternalId() uintptr
	InternalPointer() uintptr
	IsValid() bool
	Model() ItemModel
	Parent() ModelIndex
}

func (i *qModelIndex) internal_ModelIndex() {}

func (i *qModelIndex) Child(row, col int) ModelIndex {
	return mkModelIndex(uintptr(C.modelIndexChild(i.ptr, C.int(row), C.int(col))), i.engine)
}

func (i *qModelIndex) Sibling(row, col int) ModelIndex {
	return mkModelIndex(uintptr(C.modelIndexSibling(i.ptr, C.int(row), C.int(col))), i.engine)
}

func (i *qModelIndex) Column() int {
	return int(C.modelIndexColumn(i.ptr))
}

func (i *qModelIndex) Row() int {
	return int(C.modelIndexRow(i.ptr))
}

func (i *qModelIndex) Data(role Role) interface{} {
	var dvalue C.DataValue
	C.modelIndexData(i.ptr, C.int(role), &dvalue)
	return unpackDataValue(&dvalue, i.engine)
}

func (i *qModelIndex) Flags() ItemFlags {
	return ItemFlags(C.modelIndexFlags(i.ptr))
}

func (i *qModelIndex) InternalId() uintptr {
	return uintptr(C.modelIndexInternalId(i.ptr))
}

func (i *qModelIndex) InternalPointer() uintptr {
	return uintptr(C.modelIndexInternalPointer(i.ptr))
}

func (i *qModelIndex) IsValid() bool {
	if i == nil || i.ptr == nil {
		return false
	}
	return bool(C.modelIndexIsValid(i.ptr))
}

func (i *qModelIndex) Model() ItemModel {
	return itemModelFromCPP(uintptr(C.modelIndexModel(i.ptr)), i.engine)
}

func (i *qModelIndex) Parent() ModelIndex {
	return mkModelIndex(uintptr(C.modelIndexParent(i.ptr)), i.engine)
}
