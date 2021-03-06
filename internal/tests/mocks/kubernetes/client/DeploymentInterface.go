// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import autoscalingv1 "k8s.io/api/autoscaling/v1"
import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
import mock "github.com/stretchr/testify/mock"
import types "k8s.io/apimachinery/pkg/types"
import v1 "k8s.io/api/apps/v1"
import watch "k8s.io/apimachinery/pkg/watch"

// DeploymentInterface is an autogenerated mock type for the DeploymentInterface type
type DeploymentInterface struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0
func (_m *DeploymentInterface) Create(_a0 *v1.Deployment) (*v1.Deployment, error) {
	ret := _m.Called(_a0)

	var r0 *v1.Deployment
	if rf, ok := ret.Get(0).(func(*v1.Deployment) *v1.Deployment); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Deployment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*v1.Deployment) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: name, options
func (_m *DeploymentInterface) Delete(name string, options *metav1.DeleteOptions) error {
	ret := _m.Called(name, options)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *metav1.DeleteOptions) error); ok {
		r0 = rf(name, options)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteCollection provides a mock function with given fields: options, listOptions
func (_m *DeploymentInterface) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	ret := _m.Called(options, listOptions)

	var r0 error
	if rf, ok := ret.Get(0).(func(*metav1.DeleteOptions, metav1.ListOptions) error); ok {
		r0 = rf(options, listOptions)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: name, options
func (_m *DeploymentInterface) Get(name string, options metav1.GetOptions) (*v1.Deployment, error) {
	ret := _m.Called(name, options)

	var r0 *v1.Deployment
	if rf, ok := ret.Get(0).(func(string, metav1.GetOptions) *v1.Deployment); ok {
		r0 = rf(name, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Deployment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, metav1.GetOptions) error); ok {
		r1 = rf(name, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetScale provides a mock function with given fields: deploymentName, options
func (_m *DeploymentInterface) GetScale(deploymentName string, options metav1.GetOptions) (*autoscalingv1.Scale, error) {
	ret := _m.Called(deploymentName, options)

	var r0 *autoscalingv1.Scale
	if rf, ok := ret.Get(0).(func(string, metav1.GetOptions) *autoscalingv1.Scale); ok {
		r0 = rf(deploymentName, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*autoscalingv1.Scale)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, metav1.GetOptions) error); ok {
		r1 = rf(deploymentName, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: opts
func (_m *DeploymentInterface) List(opts metav1.ListOptions) (*v1.DeploymentList, error) {
	ret := _m.Called(opts)

	var r0 *v1.DeploymentList
	if rf, ok := ret.Get(0).(func(metav1.ListOptions) *v1.DeploymentList); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.DeploymentList)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(metav1.ListOptions) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Patch provides a mock function with given fields: name, pt, data, subresources
func (_m *DeploymentInterface) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (*v1.Deployment, error) {
	_va := make([]interface{}, len(subresources))
	for _i := range subresources {
		_va[_i] = subresources[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, name, pt, data)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *v1.Deployment
	if rf, ok := ret.Get(0).(func(string, types.PatchType, []byte, ...string) *v1.Deployment); ok {
		r0 = rf(name, pt, data, subresources...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Deployment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, types.PatchType, []byte, ...string) error); ok {
		r1 = rf(name, pt, data, subresources...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: _a0
func (_m *DeploymentInterface) Update(_a0 *v1.Deployment) (*v1.Deployment, error) {
	ret := _m.Called(_a0)

	var r0 *v1.Deployment
	if rf, ok := ret.Get(0).(func(*v1.Deployment) *v1.Deployment); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Deployment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*v1.Deployment) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateScale provides a mock function with given fields: deploymentName, scale
func (_m *DeploymentInterface) UpdateScale(deploymentName string, scale *autoscalingv1.Scale) (*autoscalingv1.Scale, error) {
	ret := _m.Called(deploymentName, scale)

	var r0 *autoscalingv1.Scale
	if rf, ok := ret.Get(0).(func(string, *autoscalingv1.Scale) *autoscalingv1.Scale); ok {
		r0 = rf(deploymentName, scale)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*autoscalingv1.Scale)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *autoscalingv1.Scale) error); ok {
		r1 = rf(deploymentName, scale)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateStatus provides a mock function with given fields: _a0
func (_m *DeploymentInterface) UpdateStatus(_a0 *v1.Deployment) (*v1.Deployment, error) {
	ret := _m.Called(_a0)

	var r0 *v1.Deployment
	if rf, ok := ret.Get(0).(func(*v1.Deployment) *v1.Deployment); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*v1.Deployment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*v1.Deployment) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Watch provides a mock function with given fields: opts
func (_m *DeploymentInterface) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	ret := _m.Called(opts)

	var r0 watch.Interface
	if rf, ok := ret.Get(0).(func(metav1.ListOptions) watch.Interface); ok {
		r0 = rf(opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(watch.Interface)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(metav1.ListOptions) error); ok {
		r1 = rf(opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
