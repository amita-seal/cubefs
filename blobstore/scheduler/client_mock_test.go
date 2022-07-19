// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cubefs/cubefs/blobstore/scheduler/client (interfaces: ClusterMgrAPI,BlobnodeAPI,IVolumeUpdater,ProxyAPI)

// Package scheduler is a generated GoMock package.
package scheduler

import (
	context "context"
	reflect "reflect"

	clustermgr "github.com/cubefs/cubefs/blobstore/api/clustermgr"
	proto "github.com/cubefs/cubefs/blobstore/common/proto"
	client "github.com/cubefs/cubefs/blobstore/scheduler/client"
	gomock "github.com/golang/mock/gomock"
)

// MockClusterMgrAPI is a mock of ClusterMgrAPI interface.
type MockClusterMgrAPI struct {
	ctrl     *gomock.Controller
	recorder *MockClusterMgrAPIMockRecorder
}

// MockClusterMgrAPIMockRecorder is the mock recorder for MockClusterMgrAPI.
type MockClusterMgrAPIMockRecorder struct {
	mock *MockClusterMgrAPI
}

// NewMockClusterMgrAPI creates a new mock instance.
func NewMockClusterMgrAPI(ctrl *gomock.Controller) *MockClusterMgrAPI {
	mock := &MockClusterMgrAPI{ctrl: ctrl}
	mock.recorder = &MockClusterMgrAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClusterMgrAPI) EXPECT() *MockClusterMgrAPIMockRecorder {
	return m.recorder
}

// AddMigrateTask mocks base method.
func (m *MockClusterMgrAPI) AddMigrateTask(arg0 context.Context, arg1 *proto.MigrateTask) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddMigrateTask", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddMigrateTask indicates an expected call of AddMigrateTask.
func (mr *MockClusterMgrAPIMockRecorder) AddMigrateTask(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMigrateTask", reflect.TypeOf((*MockClusterMgrAPI)(nil).AddMigrateTask), arg0, arg1)
}

// AddMigratingDisk mocks base method.
func (m *MockClusterMgrAPI) AddMigratingDisk(arg0 context.Context, arg1 *client.MigratingDiskMeta) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddMigratingDisk", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddMigratingDisk indicates an expected call of AddMigratingDisk.
func (mr *MockClusterMgrAPIMockRecorder) AddMigratingDisk(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMigratingDisk", reflect.TypeOf((*MockClusterMgrAPI)(nil).AddMigratingDisk), arg0, arg1)
}

// AllocVolumeUnit mocks base method.
func (m *MockClusterMgrAPI) AllocVolumeUnit(arg0 context.Context, arg1 proto.Vuid) (*client.AllocVunitInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllocVolumeUnit", arg0, arg1)
	ret0, _ := ret[0].(*client.AllocVunitInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AllocVolumeUnit indicates an expected call of AllocVolumeUnit.
func (mr *MockClusterMgrAPIMockRecorder) AllocVolumeUnit(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllocVolumeUnit", reflect.TypeOf((*MockClusterMgrAPI)(nil).AllocVolumeUnit), arg0, arg1)
}

// DeleteMigrateTask mocks base method.
func (m *MockClusterMgrAPI) DeleteMigrateTask(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMigrateTask", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMigrateTask indicates an expected call of DeleteMigrateTask.
func (mr *MockClusterMgrAPIMockRecorder) DeleteMigrateTask(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMigrateTask", reflect.TypeOf((*MockClusterMgrAPI)(nil).DeleteMigrateTask), arg0, arg1)
}

// DeleteMigratingDisk mocks base method.
func (m *MockClusterMgrAPI) DeleteMigratingDisk(arg0 context.Context, arg1 proto.TaskType, arg2 proto.DiskID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMigratingDisk", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMigratingDisk indicates an expected call of DeleteMigratingDisk.
func (mr *MockClusterMgrAPIMockRecorder) DeleteMigratingDisk(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMigratingDisk", reflect.TypeOf((*MockClusterMgrAPI)(nil).DeleteMigratingDisk), arg0, arg1, arg2)
}

// GetConfig mocks base method.
func (m *MockClusterMgrAPI) GetConfig(arg0 context.Context, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConfig", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConfig indicates an expected call of GetConfig.
func (mr *MockClusterMgrAPIMockRecorder) GetConfig(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfig", reflect.TypeOf((*MockClusterMgrAPI)(nil).GetConfig), arg0, arg1)
}

// GetDiskInfo mocks base method.
func (m *MockClusterMgrAPI) GetDiskInfo(arg0 context.Context, arg1 proto.DiskID) (*client.DiskInfoSimple, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDiskInfo", arg0, arg1)
	ret0, _ := ret[0].(*client.DiskInfoSimple)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDiskInfo indicates an expected call of GetDiskInfo.
func (mr *MockClusterMgrAPIMockRecorder) GetDiskInfo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDiskInfo", reflect.TypeOf((*MockClusterMgrAPI)(nil).GetDiskInfo), arg0, arg1)
}

// GetMigrateTask mocks base method.
func (m *MockClusterMgrAPI) GetMigrateTask(arg0 context.Context, arg1 string) (*proto.MigrateTask, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMigrateTask", arg0, arg1)
	ret0, _ := ret[0].(*proto.MigrateTask)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMigrateTask indicates an expected call of GetMigrateTask.
func (mr *MockClusterMgrAPIMockRecorder) GetMigrateTask(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMigrateTask", reflect.TypeOf((*MockClusterMgrAPI)(nil).GetMigrateTask), arg0, arg1)
}

// GetMigratingDisk mocks base method.
func (m *MockClusterMgrAPI) GetMigratingDisk(arg0 context.Context, arg1 proto.TaskType, arg2 proto.DiskID) (*client.MigratingDiskMeta, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMigratingDisk", arg0, arg1, arg2)
	ret0, _ := ret[0].(*client.MigratingDiskMeta)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMigratingDisk indicates an expected call of GetMigratingDisk.
func (mr *MockClusterMgrAPIMockRecorder) GetMigratingDisk(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMigratingDisk", reflect.TypeOf((*MockClusterMgrAPI)(nil).GetMigratingDisk), arg0, arg1, arg2)
}

// GetService mocks base method.
func (m *MockClusterMgrAPI) GetService(arg0 context.Context, arg1 string, arg2 proto.ClusterID) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetService", arg0, arg1, arg2)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetService indicates an expected call of GetService.
func (mr *MockClusterMgrAPIMockRecorder) GetService(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetService", reflect.TypeOf((*MockClusterMgrAPI)(nil).GetService), arg0, arg1, arg2)
}

// GetVolumeInfo mocks base method.
func (m *MockClusterMgrAPI) GetVolumeInfo(arg0 context.Context, arg1 proto.Vid) (*client.VolumeInfoSimple, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVolumeInfo", arg0, arg1)
	ret0, _ := ret[0].(*client.VolumeInfoSimple)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVolumeInfo indicates an expected call of GetVolumeInfo.
func (mr *MockClusterMgrAPIMockRecorder) GetVolumeInfo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVolumeInfo", reflect.TypeOf((*MockClusterMgrAPI)(nil).GetVolumeInfo), arg0, arg1)
}

// GetVolumeInspectCheckPoint mocks base method.
func (m *MockClusterMgrAPI) GetVolumeInspectCheckPoint(arg0 context.Context) (*proto.VolumeInspectCheckPoint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVolumeInspectCheckPoint", arg0)
	ret0, _ := ret[0].(*proto.VolumeInspectCheckPoint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVolumeInspectCheckPoint indicates an expected call of GetVolumeInspectCheckPoint.
func (mr *MockClusterMgrAPIMockRecorder) GetVolumeInspectCheckPoint(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVolumeInspectCheckPoint", reflect.TypeOf((*MockClusterMgrAPI)(nil).GetVolumeInspectCheckPoint), arg0)
}

// ListAllMigrateTasks mocks base method.
func (m *MockClusterMgrAPI) ListAllMigrateTasks(arg0 context.Context, arg1 proto.TaskType) ([]*proto.MigrateTask, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllMigrateTasks", arg0, arg1)
	ret0, _ := ret[0].([]*proto.MigrateTask)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllMigrateTasks indicates an expected call of ListAllMigrateTasks.
func (mr *MockClusterMgrAPIMockRecorder) ListAllMigrateTasks(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllMigrateTasks", reflect.TypeOf((*MockClusterMgrAPI)(nil).ListAllMigrateTasks), arg0, arg1)
}

// ListAllMigrateTasksByDiskID mocks base method.
func (m *MockClusterMgrAPI) ListAllMigrateTasksByDiskID(arg0 context.Context, arg1 proto.TaskType, arg2 proto.DiskID) ([]*proto.MigrateTask, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllMigrateTasksByDiskID", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*proto.MigrateTask)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllMigrateTasksByDiskID indicates an expected call of ListAllMigrateTasksByDiskID.
func (mr *MockClusterMgrAPIMockRecorder) ListAllMigrateTasksByDiskID(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllMigrateTasksByDiskID", reflect.TypeOf((*MockClusterMgrAPI)(nil).ListAllMigrateTasksByDiskID), arg0, arg1, arg2)
}

// ListBrokenDisks mocks base method.
func (m *MockClusterMgrAPI) ListBrokenDisks(arg0 context.Context, arg1 int) ([]*client.DiskInfoSimple, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBrokenDisks", arg0, arg1)
	ret0, _ := ret[0].([]*client.DiskInfoSimple)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBrokenDisks indicates an expected call of ListBrokenDisks.
func (mr *MockClusterMgrAPIMockRecorder) ListBrokenDisks(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBrokenDisks", reflect.TypeOf((*MockClusterMgrAPI)(nil).ListBrokenDisks), arg0, arg1)
}

// ListClusterDisks mocks base method.
func (m *MockClusterMgrAPI) ListClusterDisks(arg0 context.Context) ([]*client.DiskInfoSimple, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListClusterDisks", arg0)
	ret0, _ := ret[0].([]*client.DiskInfoSimple)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListClusterDisks indicates an expected call of ListClusterDisks.
func (mr *MockClusterMgrAPIMockRecorder) ListClusterDisks(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListClusterDisks", reflect.TypeOf((*MockClusterMgrAPI)(nil).ListClusterDisks), arg0)
}

// ListDiskVolumeUnits mocks base method.
func (m *MockClusterMgrAPI) ListDiskVolumeUnits(arg0 context.Context, arg1 proto.DiskID) ([]*client.VunitInfoSimple, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListDiskVolumeUnits", arg0, arg1)
	ret0, _ := ret[0].([]*client.VunitInfoSimple)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDiskVolumeUnits indicates an expected call of ListDiskVolumeUnits.
func (mr *MockClusterMgrAPIMockRecorder) ListDiskVolumeUnits(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDiskVolumeUnits", reflect.TypeOf((*MockClusterMgrAPI)(nil).ListDiskVolumeUnits), arg0, arg1)
}

// ListDropDisks mocks base method.
func (m *MockClusterMgrAPI) ListDropDisks(arg0 context.Context) ([]*client.DiskInfoSimple, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListDropDisks", arg0)
	ret0, _ := ret[0].([]*client.DiskInfoSimple)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListDropDisks indicates an expected call of ListDropDisks.
func (mr *MockClusterMgrAPIMockRecorder) ListDropDisks(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListDropDisks", reflect.TypeOf((*MockClusterMgrAPI)(nil).ListDropDisks), arg0)
}

// ListMigrateTasks mocks base method.
func (m *MockClusterMgrAPI) ListMigrateTasks(arg0 context.Context, arg1 proto.TaskType, arg2 *clustermgr.ListKvOpts) ([]*proto.MigrateTask, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListMigrateTasks", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*proto.MigrateTask)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListMigrateTasks indicates an expected call of ListMigrateTasks.
func (mr *MockClusterMgrAPIMockRecorder) ListMigrateTasks(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMigrateTasks", reflect.TypeOf((*MockClusterMgrAPI)(nil).ListMigrateTasks), arg0, arg1, arg2)
}

// ListMigratingDisks mocks base method.
func (m *MockClusterMgrAPI) ListMigratingDisks(arg0 context.Context, arg1 proto.TaskType) ([]*client.MigratingDiskMeta, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListMigratingDisks", arg0, arg1)
	ret0, _ := ret[0].([]*client.MigratingDiskMeta)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListMigratingDisks indicates an expected call of ListMigratingDisks.
func (mr *MockClusterMgrAPIMockRecorder) ListMigratingDisks(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListMigratingDisks", reflect.TypeOf((*MockClusterMgrAPI)(nil).ListMigratingDisks), arg0, arg1)
}

// ListRepairingDisks mocks base method.
func (m *MockClusterMgrAPI) ListRepairingDisks(arg0 context.Context) ([]*client.DiskInfoSimple, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListRepairingDisks", arg0)
	ret0, _ := ret[0].([]*client.DiskInfoSimple)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListRepairingDisks indicates an expected call of ListRepairingDisks.
func (mr *MockClusterMgrAPIMockRecorder) ListRepairingDisks(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListRepairingDisks", reflect.TypeOf((*MockClusterMgrAPI)(nil).ListRepairingDisks), arg0)
}

// ListVolume mocks base method.
func (m *MockClusterMgrAPI) ListVolume(arg0 context.Context, arg1 proto.Vid, arg2 int) ([]*client.VolumeInfoSimple, proto.Vid, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListVolume", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*client.VolumeInfoSimple)
	ret1, _ := ret[1].(proto.Vid)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListVolume indicates an expected call of ListVolume.
func (mr *MockClusterMgrAPIMockRecorder) ListVolume(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListVolume", reflect.TypeOf((*MockClusterMgrAPI)(nil).ListVolume), arg0, arg1, arg2)
}

// LockVolume mocks base method.
func (m *MockClusterMgrAPI) LockVolume(arg0 context.Context, arg1 proto.Vid) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LockVolume", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// LockVolume indicates an expected call of LockVolume.
func (mr *MockClusterMgrAPIMockRecorder) LockVolume(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LockVolume", reflect.TypeOf((*MockClusterMgrAPI)(nil).LockVolume), arg0, arg1)
}

// Register mocks base method.
func (m *MockClusterMgrAPI) Register(arg0 context.Context, arg1 client.RegisterInfo) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register.
func (mr *MockClusterMgrAPIMockRecorder) Register(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockClusterMgrAPI)(nil).Register), arg0, arg1)
}

// ReleaseVolumeUnit mocks base method.
func (m *MockClusterMgrAPI) ReleaseVolumeUnit(arg0 context.Context, arg1 proto.Vuid, arg2 proto.DiskID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReleaseVolumeUnit", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReleaseVolumeUnit indicates an expected call of ReleaseVolumeUnit.
func (mr *MockClusterMgrAPIMockRecorder) ReleaseVolumeUnit(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReleaseVolumeUnit", reflect.TypeOf((*MockClusterMgrAPI)(nil).ReleaseVolumeUnit), arg0, arg1, arg2)
}

// SetDiskDropped mocks base method.
func (m *MockClusterMgrAPI) SetDiskDropped(arg0 context.Context, arg1 proto.DiskID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetDiskDropped", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetDiskDropped indicates an expected call of SetDiskDropped.
func (mr *MockClusterMgrAPIMockRecorder) SetDiskDropped(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDiskDropped", reflect.TypeOf((*MockClusterMgrAPI)(nil).SetDiskDropped), arg0, arg1)
}

// SetDiskRepaired mocks base method.
func (m *MockClusterMgrAPI) SetDiskRepaired(arg0 context.Context, arg1 proto.DiskID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetDiskRepaired", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetDiskRepaired indicates an expected call of SetDiskRepaired.
func (mr *MockClusterMgrAPIMockRecorder) SetDiskRepaired(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDiskRepaired", reflect.TypeOf((*MockClusterMgrAPI)(nil).SetDiskRepaired), arg0, arg1)
}

// SetDiskRepairing mocks base method.
func (m *MockClusterMgrAPI) SetDiskRepairing(arg0 context.Context, arg1 proto.DiskID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetDiskRepairing", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetDiskRepairing indicates an expected call of SetDiskRepairing.
func (mr *MockClusterMgrAPIMockRecorder) SetDiskRepairing(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDiskRepairing", reflect.TypeOf((*MockClusterMgrAPI)(nil).SetDiskRepairing), arg0, arg1)
}

// SetVolumeInspectCheckPoint mocks base method.
func (m *MockClusterMgrAPI) SetVolumeInspectCheckPoint(arg0 context.Context, arg1 proto.Vid) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetVolumeInspectCheckPoint", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetVolumeInspectCheckPoint indicates an expected call of SetVolumeInspectCheckPoint.
func (mr *MockClusterMgrAPIMockRecorder) SetVolumeInspectCheckPoint(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetVolumeInspectCheckPoint", reflect.TypeOf((*MockClusterMgrAPI)(nil).SetVolumeInspectCheckPoint), arg0, arg1)
}

// UnlockVolume mocks base method.
func (m *MockClusterMgrAPI) UnlockVolume(arg0 context.Context, arg1 proto.Vid) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnlockVolume", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnlockVolume indicates an expected call of UnlockVolume.
func (mr *MockClusterMgrAPIMockRecorder) UnlockVolume(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnlockVolume", reflect.TypeOf((*MockClusterMgrAPI)(nil).UnlockVolume), arg0, arg1)
}

// UpdateMigrateTask mocks base method.
func (m *MockClusterMgrAPI) UpdateMigrateTask(arg0 context.Context, arg1 *proto.MigrateTask) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMigrateTask", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMigrateTask indicates an expected call of UpdateMigrateTask.
func (mr *MockClusterMgrAPIMockRecorder) UpdateMigrateTask(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMigrateTask", reflect.TypeOf((*MockClusterMgrAPI)(nil).UpdateMigrateTask), arg0, arg1)
}

// UpdateVolume mocks base method.
func (m *MockClusterMgrAPI) UpdateVolume(arg0 context.Context, arg1, arg2 proto.Vuid, arg3 proto.DiskID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateVolume", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateVolume indicates an expected call of UpdateVolume.
func (mr *MockClusterMgrAPIMockRecorder) UpdateVolume(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateVolume", reflect.TypeOf((*MockClusterMgrAPI)(nil).UpdateVolume), arg0, arg1, arg2, arg3)
}

// MockBlobnodeAPI is a mock of BlobnodeAPI interface.
type MockBlobnodeAPI struct {
	ctrl     *gomock.Controller
	recorder *MockBlobnodeAPIMockRecorder
}

// MockBlobnodeAPIMockRecorder is the mock recorder for MockBlobnodeAPI.
type MockBlobnodeAPIMockRecorder struct {
	mock *MockBlobnodeAPI
}

// NewMockBlobnodeAPI creates a new mock instance.
func NewMockBlobnodeAPI(ctrl *gomock.Controller) *MockBlobnodeAPI {
	mock := &MockBlobnodeAPI{ctrl: ctrl}
	mock.recorder = &MockBlobnodeAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBlobnodeAPI) EXPECT() *MockBlobnodeAPIMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockBlobnodeAPI) Delete(arg0 context.Context, arg1 proto.VunitLocation, arg2 proto.BlobID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockBlobnodeAPIMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockBlobnodeAPI)(nil).Delete), arg0, arg1, arg2)
}

// MarkDelete mocks base method.
func (m *MockBlobnodeAPI) MarkDelete(arg0 context.Context, arg1 proto.VunitLocation, arg2 proto.BlobID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarkDelete", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// MarkDelete indicates an expected call of MarkDelete.
func (mr *MockBlobnodeAPIMockRecorder) MarkDelete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarkDelete", reflect.TypeOf((*MockBlobnodeAPI)(nil).MarkDelete), arg0, arg1, arg2)
}

// RepairShard mocks base method.
func (m *MockBlobnodeAPI) RepairShard(arg0 context.Context, arg1 string, arg2 proto.ShardRepairTask) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RepairShard", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// RepairShard indicates an expected call of RepairShard.
func (mr *MockBlobnodeAPIMockRecorder) RepairShard(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RepairShard", reflect.TypeOf((*MockBlobnodeAPI)(nil).RepairShard), arg0, arg1, arg2)
}

// MockVolumeUpdater is a mock of IVolumeUpdater interface.
type MockVolumeUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockVolumeUpdaterMockRecorder
}

// MockVolumeUpdaterMockRecorder is the mock recorder for MockVolumeUpdater.
type MockVolumeUpdaterMockRecorder struct {
	mock *MockVolumeUpdater
}

// NewMockVolumeUpdater creates a new mock instance.
func NewMockVolumeUpdater(ctrl *gomock.Controller) *MockVolumeUpdater {
	mock := &MockVolumeUpdater{ctrl: ctrl}
	mock.recorder = &MockVolumeUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVolumeUpdater) EXPECT() *MockVolumeUpdaterMockRecorder {
	return m.recorder
}

// UpdateFollowerVolumeCache mocks base method.
func (m *MockVolumeUpdater) UpdateFollowerVolumeCache(arg0 context.Context, arg1 string, arg2 proto.Vid) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateFollowerVolumeCache", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateFollowerVolumeCache indicates an expected call of UpdateFollowerVolumeCache.
func (mr *MockVolumeUpdaterMockRecorder) UpdateFollowerVolumeCache(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFollowerVolumeCache", reflect.TypeOf((*MockVolumeUpdater)(nil).UpdateFollowerVolumeCache), arg0, arg1, arg2)
}

// UpdateLeaderVolumeCache mocks base method.
func (m *MockVolumeUpdater) UpdateLeaderVolumeCache(arg0 context.Context, arg1 proto.Vid) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateLeaderVolumeCache", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateLeaderVolumeCache indicates an expected call of UpdateLeaderVolumeCache.
func (mr *MockVolumeUpdaterMockRecorder) UpdateLeaderVolumeCache(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateLeaderVolumeCache", reflect.TypeOf((*MockVolumeUpdater)(nil).UpdateLeaderVolumeCache), arg0, arg1)
}

// MockMqProxyAPI is a mock of ProxyAPI interface.
type MockMqProxyAPI struct {
	ctrl     *gomock.Controller
	recorder *MockMqProxyAPIMockRecorder
}

// MockMqProxyAPIMockRecorder is the mock recorder for MockMqProxyAPI.
type MockMqProxyAPIMockRecorder struct {
	mock *MockMqProxyAPI
}

// NewMockMqProxyAPI creates a new mock instance.
func NewMockMqProxyAPI(ctrl *gomock.Controller) *MockMqProxyAPI {
	mock := &MockMqProxyAPI{ctrl: ctrl}
	mock.recorder = &MockMqProxyAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMqProxyAPI) EXPECT() *MockMqProxyAPIMockRecorder {
	return m.recorder
}

// SendShardRepairMsg mocks base method.
func (m *MockMqProxyAPI) SendShardRepairMsg(arg0 context.Context, arg1 proto.Vid, arg2 proto.BlobID, arg3 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendShardRepairMsg", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendShardRepairMsg indicates an expected call of SendShardRepairMsg.
func (mr *MockMqProxyAPIMockRecorder) SendShardRepairMsg(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendShardRepairMsg", reflect.TypeOf((*MockMqProxyAPI)(nil).SendShardRepairMsg), arg0, arg1, arg2, arg3)
}
