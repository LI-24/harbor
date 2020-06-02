// Code generated by mockery v1.0.0. DO NOT EDIT.

package blob

import (
	context "context"

	blob "github.com/goharbor/harbor/src/controller/blob"

	distribution "github.com/docker/distribution"

	mock "github.com/stretchr/testify/mock"

	models "github.com/goharbor/harbor/src/pkg/blob/models"
)

// Controller is an autogenerated mock type for the Controller type
type Controller struct {
	mock.Mock
}

// AssociateWithArtifact provides a mock function with given fields: ctx, blobDigests, artifactDigest
func (_m *Controller) AssociateWithArtifact(ctx context.Context, blobDigests []string, artifactDigest string) error {
	ret := _m.Called(ctx, blobDigests, artifactDigest)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []string, string) error); ok {
		r0 = rf(ctx, blobDigests, artifactDigest)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AssociateWithProjectByDigest provides a mock function with given fields: ctx, blobDigest, projectID
func (_m *Controller) AssociateWithProjectByDigest(ctx context.Context, blobDigest string, projectID int64) error {
	ret := _m.Called(ctx, blobDigest, projectID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, int64) error); ok {
		r0 = rf(ctx, blobDigest, projectID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// AssociateWithProjectByID provides a mock function with given fields: ctx, blobID, projectID
func (_m *Controller) AssociateWithProjectByID(ctx context.Context, blobID int64, projectID int64) error {
	ret := _m.Called(ctx, blobID, projectID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64, int64) error); ok {
		r0 = rf(ctx, blobID, projectID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CalculateTotalSizeByProject provides a mock function with given fields: ctx, projectID, excludeForeign
func (_m *Controller) CalculateTotalSizeByProject(ctx context.Context, projectID int64, excludeForeign bool) (int64, error) {
	ret := _m.Called(ctx, projectID, excludeForeign)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, int64, bool) int64); ok {
		r0 = rf(ctx, projectID, excludeForeign)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64, bool) error); ok {
		r1 = rf(ctx, projectID, excludeForeign)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *Controller) Delete(ctx context.Context, id int64) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Ensure provides a mock function with given fields: ctx, digest, contentType, size
func (_m *Controller) Ensure(ctx context.Context, digest string, contentType string, size int64) (int64, error) {
	ret := _m.Called(ctx, digest, contentType, size)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, string, string, int64) int64); ok {
		r0 = rf(ctx, digest, contentType, size)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, int64) error); ok {
		r1 = rf(ctx, digest, contentType, size)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Exist provides a mock function with given fields: ctx, digest, options
func (_m *Controller) Exist(ctx context.Context, digest string, options ...blob.Option) (bool, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, digest)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string, ...blob.Option) bool); ok {
		r0 = rf(ctx, digest, options...)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, ...blob.Option) error); ok {
		r1 = rf(ctx, digest, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindMissingAssociationsForProject provides a mock function with given fields: ctx, projectID, blobs
func (_m *Controller) FindMissingAssociationsForProject(ctx context.Context, projectID int64, blobs []*models.Blob) ([]*models.Blob, error) {
	ret := _m.Called(ctx, projectID, blobs)

	var r0 []*models.Blob
	if rf, ok := ret.Get(0).(func(context.Context, int64, []*models.Blob) []*models.Blob); ok {
		r0 = rf(ctx, projectID, blobs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Blob)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64, []*models.Blob) error); ok {
		r1 = rf(ctx, projectID, blobs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: ctx, digest, options
func (_m *Controller) Get(ctx context.Context, digest string, options ...blob.Option) (*models.Blob, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, digest)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *models.Blob
	if rf, ok := ret.Get(0).(func(context.Context, string, ...blob.Option) *models.Blob); ok {
		r0 = rf(ctx, digest, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Blob)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, ...blob.Option) error); ok {
		r1 = rf(ctx, digest, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAcceptedBlobSize provides a mock function with given fields: sessionID
func (_m *Controller) GetAcceptedBlobSize(sessionID string) (int64, error) {
	ret := _m.Called(sessionID)

	var r0 int64
	if rf, ok := ret.Get(0).(func(string) int64); ok {
		r0 = rf(sessionID)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(sessionID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: ctx, params
func (_m *Controller) List(ctx context.Context, params models.ListParams) ([]*models.Blob, error) {
	ret := _m.Called(ctx, params)

	var r0 []*models.Blob
	if rf, ok := ret.Get(0).(func(context.Context, models.ListParams) []*models.Blob); ok {
		r0 = rf(ctx, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Blob)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.ListParams) error); ok {
		r1 = rf(ctx, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetAcceptedBlobSize provides a mock function with given fields: sessionID, size
func (_m *Controller) SetAcceptedBlobSize(sessionID string, size int64) error {
	ret := _m.Called(sessionID, size)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, int64) error); ok {
		r0 = rf(sessionID, size)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Sync provides a mock function with given fields: ctx, references
func (_m *Controller) Sync(ctx context.Context, references []distribution.Descriptor) error {
	ret := _m.Called(ctx, references)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []distribution.Descriptor) error); ok {
		r0 = rf(ctx, references)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Touch provides a mock function with given fields: ctx, _a1
func (_m *Controller) Touch(ctx context.Context, _a1 *models.Blob) (int64, error) {
	ret := _m.Called(ctx, _a1)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, *models.Blob) int64); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *models.Blob) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, _a1
func (_m *Controller) Update(ctx context.Context, _a1 *models.Blob) error {
	ret := _m.Called(ctx, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Blob) error); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
