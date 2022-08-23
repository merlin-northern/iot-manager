// Copyright 2022 Northern.tech AS
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"

	app "github.com/mendersoftware/iot-manager/app"

	iotcore "github.com/mendersoftware/iot-manager/client/iotcore"

	iothub "github.com/mendersoftware/iot-manager/client/iothub"

	mock "github.com/stretchr/testify/mock"

	model "github.com/mendersoftware/iot-manager/model"

	uuid "github.com/google/uuid"
)

// App is an autogenerated mock type for the App type
type App struct {
	mock.Mock
}

// CreateIntegration provides a mock function with given fields: _a0, _a1
func (_m *App) CreateIntegration(_a0 context.Context, _a1 model.Integration) (*model.Integration, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *model.Integration
	if rf, ok := ret.Get(0).(func(context.Context, model.Integration) *model.Integration); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Integration)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.Integration) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DecommissionDevice provides a mock function with given fields: _a0, _a1
func (_m *App) DecommissionDevice(_a0 context.Context, _a1 string) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetDevice provides a mock function with given fields: _a0, _a1
func (_m *App) GetDevice(_a0 context.Context, _a1 string) (*model.Device, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *model.Device
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.Device); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Device)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDeviceIntegrations provides a mock function with given fields: _a0, _a1
func (_m *App) GetDeviceIntegrations(_a0 context.Context, _a1 string) ([]model.Integration, error) {
	ret := _m.Called(_a0, _a1)

	var r0 []model.Integration
	if rf, ok := ret.Get(0).(func(context.Context, string) []model.Integration); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Integration)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDeviceStateIntegration provides a mock function with given fields: _a0, _a1, _a2
func (_m *App) GetDeviceStateIntegration(_a0 context.Context, _a1 string, _a2 uuid.UUID) (*model.DeviceState, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *model.DeviceState
	if rf, ok := ret.Get(0).(func(context.Context, string, uuid.UUID) *model.DeviceState); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.DeviceState)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, uuid.UUID) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDeviceStateIoTCore provides a mock function with given fields: _a0, _a1, _a2
func (_m *App) GetDeviceStateIoTCore(_a0 context.Context, _a1 string, _a2 *model.Integration) (*model.DeviceState, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *model.DeviceState
	if rf, ok := ret.Get(0).(func(context.Context, string, *model.Integration) *model.DeviceState); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.DeviceState)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, *model.Integration) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDeviceStateIoTHub provides a mock function with given fields: _a0, _a1, _a2
func (_m *App) GetDeviceStateIoTHub(_a0 context.Context, _a1 string, _a2 *model.Integration) (*model.DeviceState, error) {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 *model.DeviceState
	if rf, ok := ret.Get(0).(func(context.Context, string, *model.Integration) *model.DeviceState); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.DeviceState)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, *model.Integration) error); ok {
		r1 = rf(_a0, _a1, _a2)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetEvents provides a mock function with given fields: ctx, filter
func (_m *App) GetEvents(ctx context.Context, filter model.EventsFilter) ([]model.Event, error) {
	ret := _m.Called(ctx, filter)

	var r0 []model.Event
	if rf, ok := ret.Get(0).(func(context.Context, model.EventsFilter) []model.Event); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Event)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.EventsFilter) error); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetIntegrationById provides a mock function with given fields: _a0, _a1
func (_m *App) GetIntegrationById(_a0 context.Context, _a1 uuid.UUID) (*model.Integration, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *model.Integration
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *model.Integration); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Integration)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetIntegrations provides a mock function with given fields: _a0
func (_m *App) GetIntegrations(_a0 context.Context) ([]model.Integration, error) {
	ret := _m.Called(_a0)

	var r0 []model.Integration
	if rf, ok := ret.Get(0).(func(context.Context) []model.Integration); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.Integration)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// HealthCheck provides a mock function with given fields: _a0
func (_m *App) HealthCheck(_a0 context.Context) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ProvisionDevice provides a mock function with given fields: _a0, _a1
func (_m *App) ProvisionDevice(_a0 context.Context, _a1 model.DeviceEvent) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, model.DeviceEvent) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveIntegration provides a mock function with given fields: _a0, _a1
func (_m *App) RemoveIntegration(_a0 context.Context, _a1 uuid.UUID) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetDeviceStateIntegration provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *App) SetDeviceStateIntegration(_a0 context.Context, _a1 string, _a2 uuid.UUID, _a3 *model.DeviceState) (*model.DeviceState, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 *model.DeviceState
	if rf, ok := ret.Get(0).(func(context.Context, string, uuid.UUID, *model.DeviceState) *model.DeviceState); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.DeviceState)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, uuid.UUID, *model.DeviceState) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetDeviceStateIoTCore provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *App) SetDeviceStateIoTCore(_a0 context.Context, _a1 string, _a2 *model.Integration, _a3 *model.DeviceState) (*model.DeviceState, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 *model.DeviceState
	if rf, ok := ret.Get(0).(func(context.Context, string, *model.Integration, *model.DeviceState) *model.DeviceState); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.DeviceState)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, *model.Integration, *model.DeviceState) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetDeviceStateIoTHub provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *App) SetDeviceStateIoTHub(_a0 context.Context, _a1 string, _a2 *model.Integration, _a3 *model.DeviceState) (*model.DeviceState, error) {
	ret := _m.Called(_a0, _a1, _a2, _a3)

	var r0 *model.DeviceState
	if rf, ok := ret.Get(0).(func(context.Context, string, *model.Integration, *model.DeviceState) *model.DeviceState); ok {
		r0 = rf(_a0, _a1, _a2, _a3)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.DeviceState)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, *model.Integration, *model.DeviceState) error); ok {
		r1 = rf(_a0, _a1, _a2, _a3)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetDeviceStatus provides a mock function with given fields: _a0, _a1, _a2
func (_m *App) SetDeviceStatus(_a0 context.Context, _a1 string, _a2 model.Status) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, model.Status) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetIntegrationCredentials provides a mock function with given fields: _a0, _a1, _a2
func (_m *App) SetIntegrationCredentials(_a0 context.Context, _a1 uuid.UUID, _a2 model.Credentials) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, model.Credentials) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SyncDevices provides a mock function with given fields: _a0, _a1, _a2
func (_m *App) SyncDevices(_a0 context.Context, _a1 int, _a2 bool) error {
	ret := _m.Called(_a0, _a1, _a2)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int, bool) error); ok {
		r0 = rf(_a0, _a1, _a2)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WithIoTCore provides a mock function with given fields: client
func (_m *App) WithIoTCore(client iotcore.Client) app.App {
	ret := _m.Called(client)

	var r0 app.App
	if rf, ok := ret.Get(0).(func(iotcore.Client) app.App); ok {
		r0 = rf(client)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(app.App)
		}
	}

	return r0
}

// WithIoTHub provides a mock function with given fields: client
func (_m *App) WithIoTHub(client iothub.Client) app.App {
	ret := _m.Called(client)

	var r0 app.App
	if rf, ok := ret.Get(0).(func(iothub.Client) app.App); ok {
		r0 = rf(client)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(app.App)
		}
	}

	return r0
}
