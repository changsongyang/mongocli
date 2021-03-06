// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongocli/internal/store (interfaces: ProcessDiskMeasurementsLister)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	mongodbatlas "go.mongodb.org/atlas/mongodbatlas"
	reflect "reflect"
)

// MockProcessDiskMeasurementsLister is a mock of ProcessDiskMeasurementsLister interface
type MockProcessDiskMeasurementsLister struct {
	ctrl     *gomock.Controller
	recorder *MockProcessDiskMeasurementsListerMockRecorder
}

// MockProcessDiskMeasurementsListerMockRecorder is the mock recorder for MockProcessDiskMeasurementsLister
type MockProcessDiskMeasurementsListerMockRecorder struct {
	mock *MockProcessDiskMeasurementsLister
}

// NewMockProcessDiskMeasurementsLister creates a new mock instance
func NewMockProcessDiskMeasurementsLister(ctrl *gomock.Controller) *MockProcessDiskMeasurementsLister {
	mock := &MockProcessDiskMeasurementsLister{ctrl: ctrl}
	mock.recorder = &MockProcessDiskMeasurementsListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProcessDiskMeasurementsLister) EXPECT() *MockProcessDiskMeasurementsListerMockRecorder {
	return m.recorder
}

// ProcessDiskMeasurements mocks base method
func (m *MockProcessDiskMeasurementsLister) ProcessDiskMeasurements(arg0, arg1 string, arg2 int, arg3 string, arg4 *mongodbatlas.ProcessMeasurementListOptions) (*mongodbatlas.ProcessDiskMeasurements, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessDiskMeasurements", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(*mongodbatlas.ProcessDiskMeasurements)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProcessDiskMeasurements indicates an expected call of ProcessDiskMeasurements
func (mr *MockProcessDiskMeasurementsListerMockRecorder) ProcessDiskMeasurements(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessDiskMeasurements", reflect.TypeOf((*MockProcessDiskMeasurementsLister)(nil).ProcessDiskMeasurements), arg0, arg1, arg2, arg3, arg4)
}
