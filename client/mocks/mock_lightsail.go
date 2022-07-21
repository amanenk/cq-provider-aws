// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudquery/cq-provider-aws/client (interfaces: LightsailClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	lightsail "github.com/aws/aws-sdk-go-v2/service/lightsail"
	gomock "github.com/golang/mock/gomock"
)

// MockLightsailClient is a mock of LightsailClient interface.
type MockLightsailClient struct {
	ctrl     *gomock.Controller
	recorder *MockLightsailClientMockRecorder
}

// MockLightsailClientMockRecorder is the mock recorder for MockLightsailClient.
type MockLightsailClientMockRecorder struct {
	mock *MockLightsailClient
}

// NewMockLightsailClient creates a new mock instance.
func NewMockLightsailClient(ctrl *gomock.Controller) *MockLightsailClient {
	mock := &MockLightsailClient{ctrl: ctrl}
	mock.recorder = &MockLightsailClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLightsailClient) EXPECT() *MockLightsailClientMockRecorder {
	return m.recorder
}

// GetBucketAccessKeys mocks base method.
func (m *MockLightsailClient) GetBucketAccessKeys(arg0 context.Context, arg1 *lightsail.GetBucketAccessKeysInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetBucketAccessKeysOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBucketAccessKeys", varargs...)
	ret0, _ := ret[0].(*lightsail.GetBucketAccessKeysOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBucketAccessKeys indicates an expected call of GetBucketAccessKeys.
func (mr *MockLightsailClientMockRecorder) GetBucketAccessKeys(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBucketAccessKeys", reflect.TypeOf((*MockLightsailClient)(nil).GetBucketAccessKeys), varargs...)
}

// GetBuckets mocks base method.
func (m *MockLightsailClient) GetBuckets(arg0 context.Context, arg1 *lightsail.GetBucketsInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetBucketsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetBuckets", varargs...)
	ret0, _ := ret[0].(*lightsail.GetBucketsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBuckets indicates an expected call of GetBuckets.
func (mr *MockLightsailClientMockRecorder) GetBuckets(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBuckets", reflect.TypeOf((*MockLightsailClient)(nil).GetBuckets), varargs...)
}

// GetInstances mocks base method.
func (m *MockLightsailClient) GetInstances(arg0 context.Context, arg1 *lightsail.GetInstancesInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetInstancesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetInstances", varargs...)
	ret0, _ := ret[0].(*lightsail.GetInstancesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetInstances indicates an expected call of GetInstances.
func (mr *MockLightsailClientMockRecorder) GetInstances(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetInstances", reflect.TypeOf((*MockLightsailClient)(nil).GetInstances), varargs...)
}

// GetRelationalDatabaseEvents mocks base method.
func (m *MockLightsailClient) GetRelationalDatabaseEvents(arg0 context.Context, arg1 *lightsail.GetRelationalDatabaseEventsInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetRelationalDatabaseEventsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRelationalDatabaseEvents", varargs...)
	ret0, _ := ret[0].(*lightsail.GetRelationalDatabaseEventsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRelationalDatabaseEvents indicates an expected call of GetRelationalDatabaseEvents.
func (mr *MockLightsailClientMockRecorder) GetRelationalDatabaseEvents(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRelationalDatabaseEvents", reflect.TypeOf((*MockLightsailClient)(nil).GetRelationalDatabaseEvents), varargs...)
}

// GetRelationalDatabaseLogEvents mocks base method.
func (m *MockLightsailClient) GetRelationalDatabaseLogEvents(arg0 context.Context, arg1 *lightsail.GetRelationalDatabaseLogEventsInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetRelationalDatabaseLogEventsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRelationalDatabaseLogEvents", varargs...)
	ret0, _ := ret[0].(*lightsail.GetRelationalDatabaseLogEventsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRelationalDatabaseLogEvents indicates an expected call of GetRelationalDatabaseLogEvents.
func (mr *MockLightsailClientMockRecorder) GetRelationalDatabaseLogEvents(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRelationalDatabaseLogEvents", reflect.TypeOf((*MockLightsailClient)(nil).GetRelationalDatabaseLogEvents), varargs...)
}

// GetRelationalDatabaseLogStreams mocks base method.
func (m *MockLightsailClient) GetRelationalDatabaseLogStreams(arg0 context.Context, arg1 *lightsail.GetRelationalDatabaseLogStreamsInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetRelationalDatabaseLogStreamsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRelationalDatabaseLogStreams", varargs...)
	ret0, _ := ret[0].(*lightsail.GetRelationalDatabaseLogStreamsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRelationalDatabaseLogStreams indicates an expected call of GetRelationalDatabaseLogStreams.
func (mr *MockLightsailClientMockRecorder) GetRelationalDatabaseLogStreams(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRelationalDatabaseLogStreams", reflect.TypeOf((*MockLightsailClient)(nil).GetRelationalDatabaseLogStreams), varargs...)
}

// GetRelationalDatabaseParameters mocks base method.
func (m *MockLightsailClient) GetRelationalDatabaseParameters(arg0 context.Context, arg1 *lightsail.GetRelationalDatabaseParametersInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetRelationalDatabaseParametersOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRelationalDatabaseParameters", varargs...)
	ret0, _ := ret[0].(*lightsail.GetRelationalDatabaseParametersOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRelationalDatabaseParameters indicates an expected call of GetRelationalDatabaseParameters.
func (mr *MockLightsailClientMockRecorder) GetRelationalDatabaseParameters(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRelationalDatabaseParameters", reflect.TypeOf((*MockLightsailClient)(nil).GetRelationalDatabaseParameters), varargs...)
}

// GetRelationalDatabaseSnapshots mocks base method.
func (m *MockLightsailClient) GetRelationalDatabaseSnapshots(arg0 context.Context, arg1 *lightsail.GetRelationalDatabaseSnapshotsInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetRelationalDatabaseSnapshotsOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRelationalDatabaseSnapshots", varargs...)
	ret0, _ := ret[0].(*lightsail.GetRelationalDatabaseSnapshotsOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRelationalDatabaseSnapshots indicates an expected call of GetRelationalDatabaseSnapshots.
func (mr *MockLightsailClientMockRecorder) GetRelationalDatabaseSnapshots(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRelationalDatabaseSnapshots", reflect.TypeOf((*MockLightsailClient)(nil).GetRelationalDatabaseSnapshots), varargs...)
}

// GetRelationalDatabases mocks base method.
func (m *MockLightsailClient) GetRelationalDatabases(arg0 context.Context, arg1 *lightsail.GetRelationalDatabasesInput, arg2 ...func(*lightsail.Options)) (*lightsail.GetRelationalDatabasesOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetRelationalDatabases", varargs...)
	ret0, _ := ret[0].(*lightsail.GetRelationalDatabasesOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRelationalDatabases indicates an expected call of GetRelationalDatabases.
func (mr *MockLightsailClientMockRecorder) GetRelationalDatabases(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRelationalDatabases", reflect.TypeOf((*MockLightsailClient)(nil).GetRelationalDatabases), varargs...)
}
