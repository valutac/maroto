// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import consts "github.com/johnfercher/maroto/pkg/consts"

import mock "github.com/stretchr/testify/mock"
import props "github.com/johnfercher/maroto/pkg/props"

// Image is an autogenerated mock type for the Image type
type Image struct {
	mock.Mock
}

// AddFromBase64 provides a mock function with given fields: b64, marginTop, indexCol, qtdCols, colHeight, prop, extension
func (_m *Image) AddFromBase64(b64 string, marginTop float64, indexCol float64, qtdCols float64, colHeight float64, prop props.Rect, extension consts.Extension) error {
	ret := _m.Called(b64, marginTop, indexCol, qtdCols, colHeight, prop, extension)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, float64, float64, float64, float64, props.Rect, consts.Extension) error); ok {
		r0 = rf(b64, marginTop, indexCol, qtdCols, colHeight, prop, extension)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AddFromFile provides a mock function with given fields: path, marginTop, indexCol, qtdCols, colHeight, prop
func (_m *Image) AddFromFile(path string, marginTop float64, indexCol float64, qtdCols float64, colHeight float64, prop props.Rect) error {
	ret := _m.Called(path, marginTop, indexCol, qtdCols, colHeight, prop)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, float64, float64, float64, float64, props.Rect) error); ok {
		r0 = rf(path, marginTop, indexCol, qtdCols, colHeight, prop)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
