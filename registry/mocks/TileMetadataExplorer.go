// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	models "github.com/monitoror/monitoror/models"
	mock "github.com/stretchr/testify/mock"

	registry "github.com/monitoror/monitoror/registry"

	versions "github.com/monitoror/monitoror/api/config/versions"
)

// TileMetadataExplorer is an autogenerated mock type for the TileMetadataExplorer type
type TileMetadataExplorer struct {
	mock.Mock
}

// GetMinimalVersion provides a mock function with given fields:
func (_m *TileMetadataExplorer) GetMinimalVersion() versions.RawVersion {
	ret := _m.Called()

	var r0 versions.RawVersion
	if rf, ok := ret.Get(0).(func() versions.RawVersion); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(versions.RawVersion)
	}

	return r0
}

// GetVariant provides a mock function with given fields: variantName
func (_m *TileMetadataExplorer) GetVariant(variantName models.VariantName) (registry.VariantMetadataExplorer, bool) {
	ret := _m.Called(variantName)

	var r0 registry.VariantMetadataExplorer
	if rf, ok := ret.Get(0).(func(models.VariantName) registry.VariantMetadataExplorer); ok {
		r0 = rf(variantName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(registry.VariantMetadataExplorer)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(models.VariantName) bool); ok {
		r1 = rf(variantName)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// GetVariantsNames provides a mock function with given fields:
func (_m *TileMetadataExplorer) GetVariantsNames() []models.VariantName {
	ret := _m.Called()

	var r0 []models.VariantName
	if rf, ok := ret.Get(0).(func() []models.VariantName); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.VariantName)
		}
	}

	return r0
}
